package game

import (
	"fmt"
	"github.com/bsloan/game-sandbox/asset"
	"github.com/bsloan/game-sandbox/boards"
	"github.com/bsloan/game-sandbox/entity"
	"github.com/bsloan/game-sandbox/settings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp"
	"image"
	"image/color"
	"slices"
)

type Viewport struct {
	viewX     float64
	viewY     float64
	MaxViewX  float64
	MaxViewY  float64
	midX      float64
	midY      float64
	view      *ebiten.Image
	midground *ebiten.Image
}

// Move moves the viewport to pixel coordinates x,y in the game board.
// If coordinates are out-of-bounds of the map, they are adjusted to be
// within bounds. The middle layer for parallax scrolling is also updated
// according to the provided coordinates using the scroll multiplier.
func (p *Viewport) Move(x, y float64) {
	if x < 0 {
		x = 0
	} else if x > p.MaxViewX {
		x = p.MaxViewX
	}
	if y < 0 {
		y = 0
	} else if y > p.MaxViewY {
		y = p.MaxViewY
	}
	p.viewX = x
	p.viewY = y
	p.midX = p.viewX * settings.MidgroundScrollMultiplier
	p.midY = p.viewY * settings.MidgroundScrollMultiplier
}

// Center centers the viewport around (x,y) in the game board.
func (p *Viewport) Center(x, y float64) {
	p.Move(x-settings.ScreenWidth/2, y-settings.ScreenHeight/2)
}

// Position returns the pixel X, Y coordinates in the game board of the top-left
// pixel of the viewport.
func (p *Viewport) Position() (float64, float64) {
	return p.viewX, p.viewY
}

// TilePosition returns the tile coordinates of the top-leftmost visible tile.
func (p *Viewport) TilePosition() (int, int) {
	tx := int(p.viewX / settings.TileSize)
	ty := int(p.viewY / settings.TileSize)
	return tx, ty
}

func (p *Viewport) InView(e *entity.Entity) bool {
	x, y := e.Position()
	return x >= p.viewX-float64(e.Image().Bounds().Dx()) &&
		x < p.viewX+settings.ScreenWidth &&
		y >= p.viewY-float64(e.Image().Bounds().Dy()) &&
		y < p.viewY+settings.ScreenHeight
}

func translateTileCoordsToScreen(tileX, tileY int, offsetX, offsetY float64) (float64, float64) {
	return float64(tileX*settings.TileSize) - offsetX - (settings.TileSize / 2), float64(tileY*settings.TileSize) - offsetY - (settings.TileSize / 2)
}

// Draw renders the foreground layer (tiles and sprites) within the currently visible section
// of the game board. When the frame is ready, this is later copied to the screen on top
// of any background and middle layers for a parallax scrolling affect.
func (p *Viewport) Draw(g *Game) {
	// ebiten performance: avoid allocating a new image on every Update, use Clear instead
	if p.view == nil {
		p.view = ebiten.NewImage(settings.ScreenWidth, settings.ScreenHeight)
	}
	if p.midground == nil {
		// iterate horizontally in chunks of 160 pixels (width of midground image)
		// across the entire game board. draw any instances of the image that happen
		// to be visible within the viewport boundaries. skip any that are not visible.
		// draw them at constant height
		// TODO: refactor to generate the midground image elsewhere
		p.midground = ebiten.NewImage(int(float64(g.board.PixelWidth)), int(float64(g.board.PixelHeight)))
		ht := float64(asset.HillsMidground.Bounds().Dy())
		midY := float64(g.board.PixelHeight) - (ht + 55)
		for midX := 0; midX < g.board.PixelWidth; midX += 160 {
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(midX), midY)
			p.midground.DrawImage(asset.HillsMidground, &op)
		}
	}

	p.view.Clear()

	// render the static background image
	p.view.DrawImage(asset.SkyBackground, &ebiten.DrawImageOptions{})

	// render the middle layer
	p.view.DrawImage(g.vp.midground.SubImage(image.Rect(int(g.vp.midX), int(g.vp.midY), int(g.vp.midX+settings.ScreenWidth), int(g.vp.midY+settings.ScreenHeight))).(*ebiten.Image), &ebiten.DrawImageOptions{})

	// render the tiles. calculate the top-left origin of the viewport in units of tiles (tileX,tileY).
	// then calculate the offset in pixels (ox,oy) that we begin drawing from the top-left tile from. the
	// pixel offset allows for smooth scrolling in smaller units (pixel) instead of large (tile).
	tileX, tileY := p.TilePosition()
	ox, oy := float64(int(p.viewX)%settings.TileSize), float64(int(p.viewY)%settings.TileSize)
	yTileCount := 0
	// create a structure to hold any foreground tiles we find, so we can draw them later after sprites
	type tileCoords struct {
		X, Y, Tile int
	}
	var foregroundTiles []tileCoords

	for ty := tileY; ty <= (tileY+settings.ScreenHeightTiles+1) && ty < g.board.TileHeight; ty++ {
		xTileCount := 0
		for tx := tileX; tx <= (tileX+settings.ScreenWidthTiles+1) && tx < g.board.TileWidth; tx++ {
			tile := g.board.Map[ty][tx]
			if slices.Contains(boards.ForegroundTiles, tile) {
				// save foreground tiles for rendering later
				foregroundTiles = append(foregroundTiles, tileCoords{X: xTileCount, Y: yTileCount, Tile: tile})
			} else if tile != 0 {
				// adjust the x,y position where the tile is rendered from to be its top-left corner.
				// the tile's x,y coordinates in chipmunk 2d space are at it's center, so need to pull it
				// back and to the left by tileSize/2, and adjust by (ox,oy) to get the correct pixel offset
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(translateTileCoordsToScreen(xTileCount, yTileCount, ox, oy))
				p.view.DrawImage(asset.TileImages[tile], &op)
			}
			xTileCount++
		}
		yTileCount++
	}

	// render sprites
	for _, e := range g.registry.DrawableEntities() {
		if p.InView(e) && e.Shape != nil {
			// The e's position in chipmunk 2d space is its center of gravity.
			// Use these coordinates as starting point for translating to vp coordinates (in pixels)
			x, y := e.Position()

			// Find the top-left corner of the e's shape and pull its position back to top-left corner.
			x -= (e.Shape.BB().R - e.Shape.BB().L) / 2
			y -= (e.Shape.BB().T - e.Shape.BB().B) / 2

			// Adjust to get the correct pixel offset within the vp
			x -= p.viewX
			y -= p.viewY

			// draw the e
			entityImage := e.Image()
			if entityImage != nil {
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(x, y)
				if e.Damaged > 0 {
					// if this entity has taken damage, render with a solid red fill
					op.ColorScale.Scale(255.0, 0.0, 0.0, 255.0)
				}
				p.view.DrawImage(entityImage, &op)
			}
		}
	}

	// render foreground tiles
	for _, tileCoord := range foregroundTiles {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(translateTileCoordsToScreen(tileCoord.X, tileCoord.Y, ox, oy))
		p.view.DrawImage(asset.TileImages[tileCoord.Tile], &op)
	}
}

type Game struct {
	vp         Viewport
	board      boards.Gameboard
	debug      bool
	registry   entity.Registry
	space      *cp.Space
	gamepadIds []ebiten.GamepadID
}

func NewGame(viewport Viewport, gameboard boards.Gameboard, debug bool, registry entity.Registry, space *cp.Space) *Game {
	game := Game{
		vp:         viewport,
		board:      gameboard,
		debug:      debug,
		registry:   registry,
		space:      space,
		gamepadIds: []ebiten.GamepadID{},
	}
	return &game
}

func (g *Game) CenterViewport(x, y float64) {
	g.vp.Center(x, y)
}

func (g *Game) animateSprites() {
	for _, e := range g.registry.Entities {
		if e != nil {
			if e.Animations != nil && e.Animations[e.State] != nil {
				newEntityState := e.Animations[e.State].Animate()
				if newEntityState != entity.Default {
					e.State = newEntityState
				}
			}
			if e.Damaged > 0 {
				e.Damaged-- // decrement number of ticks remaining to show damage color scale
			}
		}
	}
}

func (g *Game) drawPlayerHealth(screen *ebiten.Image) {
	player := g.registry.Player()
	barWidth := player.MaxHealth + 1
	var barHeight float32 = 5
	var x, y float32 = float32(settings.ScreenWidth-barWidth) - 2, 2

	// draw health bar background
	vector.DrawFilledRect(screen, x, y, float32(barWidth), barHeight, color.RGBA{R: 211, G: 211, B: 211, A: 255}, false)

	// draw filled portion representing player's remaining health
	filledWidth := max(player.Health, 0)
	vector.DrawFilledRect(screen, x+1, y+1, float32(filledWidth), barHeight-1, color.RGBA{R: 255, G: 0, B: 0, A: 255}, false)
}

func (g *Game) Cleanup() {
	g.registry.RemoveDead(g.space)
}

func (g *Game) Update() error {
	// get user input and move the player entity
	g.MovePlayer()

	// TODO: iterate all entities in the registry, get the relevant Move* function
	//  based on entity type, and run the move function (e.g., g.MoveEntity(entity) )
	//  For now, just query for sword directly.
	swordDog := g.registry.Query(entity.SwordDog)
	g.MoveSwordDog(swordDog)

	// render the image of the current viewport, centered on player
	g.vp.Center(g.registry.Player().Body.Position().X, g.registry.Player().Body.Position().Y)
	g.vp.Draw(g)

	// animate sprites
	g.animateSprites()

	// update physics space
	g.space.Step(1.0 / float64(ebiten.TPS()))

	// remove Dead entities from the space
	g.Cleanup()

	// return any errors
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.vp.view, &ebiten.DrawImageOptions{})

	g.drawPlayerHealth(screen)

	if g.debug {
		tx, ty := g.vp.TilePosition()
		px, py := g.registry.Player().Body.Position().X, g.registry.Player().Body.Position().Y
		vx, vy := g.registry.Player().Body.Velocity().X, g.registry.Player().Body.Velocity().Y
		x, y := g.vp.Position()
		debugMsg := fmt.Sprintf(
			"TPS: %0.2f Origin X,Y: (%0.2f, %0.2f) Tile X,Y: (%v, %v)\nPlayer X,Y (%0.2f, %0.2f)\nVelocity X,Y (%0.2f, %0.2f)\nGrounded: %v\nSlope: %v\nBoost: %v",
			ebiten.ActualTPS(), x, y, tx, ty, px, py, vx, vy, g.registry.Player().Grounded, g.registry.Player().OnSlope, g.registry.Player().Boost)
		ebitenutil.DebugPrint(screen, debugMsg)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return settings.ScreenWidth, settings.ScreenHeight
}
