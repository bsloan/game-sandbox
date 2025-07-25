package main

import (
	"flag"
	"fmt"
	"github.com/bsloan/game-sandbox/asset"
	"github.com/bsloan/game-sandbox/boards"
	"github.com/bsloan/game-sandbox/entity"
	"github.com/ebitenui/ebitenui/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
	"image"
	"log"
)

// TODO: move to a global settings package
const (
	ticksPerSecond             = 60
	screenWidth                = 320
	screenHeight               = 240
	tileSize                   = 16
	screenWidthTiles           = screenWidth / tileSize
	screenHeightTiles          = screenHeight / tileSize
	midgroundScrollMultiplier  = 0.5
	Gravity                    = 450.0
	PlayerMaxVelocityX         = 100
	PlayerMaxVelocityY         = 300
	PlayerJumpVelocityLimit    = 100
	PlayerAccelerationStepX    = 1000
	PlayerJumpBoostHeight      = 30
	PlayerJumpInitialVelocity  = 6000
	PlayerJumpContinueVelocity = 400
)

type viewport struct {
	viewX     float64
	viewY     float64
	maxViewX  float64
	maxViewY  float64
	midX      float64
	midY      float64
	view      *ebiten.Image
	midground *ebiten.Image
}

// Move moves the viewport to pixel coordinates x,y in the game board.
// If coordinates are out-of-bounds of the map, they are adjusted to be
// within bounds. The middle layer for parallax scrolling is also updated
// according to the provided coordinates using the scroll multiplier.
func (p *viewport) Move(x, y float64) {
	if x < 0 {
		x = 0
	} else if x > p.maxViewX {
		x = p.maxViewX
	}
	if y < 0 {
		y = 0
	} else if y > p.maxViewY {
		y = p.maxViewY
	}
	p.viewX = x
	p.viewY = y
	p.midX = p.viewX * midgroundScrollMultiplier
	p.midY = p.viewY * midgroundScrollMultiplier
}

// Center centers the viewport around (x,y) in the game board.
func (p *viewport) Center(x, y float64) {
	p.Move(x-screenWidth/2, y-screenHeight/2)
}

// Position returns the pixel X, Y coordinates in the game board of the top-left
// pixel of the viewport.
func (p *viewport) Position() (float64, float64) {
	return p.viewX, p.viewY
}

// TilePosition returns the tile coordinates of the top-leftmost visible tile.
func (p *viewport) TilePosition() (int, int) {
	tx := int(p.viewX / tileSize)
	ty := int(p.viewY / tileSize)
	return tx, ty
}

func (p *viewport) InView(e *entity.Entity) bool {
	x, y := e.Position()
	return x >= p.viewX-float64(e.Image().Bounds().Dx()) &&
		x < p.viewX+screenWidth &&
		y >= p.viewY-float64(e.Image().Bounds().Dy()) &&
		y < p.viewY+screenHeight
}

// Draw renders the foreground layer (tiles and sprites) within the currently visible section
// of the game board. When the frame is ready, this is later copied to the screen on top
// of any background and middle layers for a parallax scrolling affect.
func (p *viewport) Draw(g *Game) {
	// ebiten performance: avoid allocating a new image on every Update, use Clear instead
	if p.view == nil {
		p.view = ebiten.NewImage(screenWidth, screenHeight)
	}
	if p.midground == nil {
		// iterate horizontally in chunks of 160 pixels (width of midground image)
		// across the entire game board. draw any instances of the image that happen
		// to be visible within the viewport boundaries. skip any that are not visible.
		// draw them at constant height
		// TODO: refactor to generate the midground image elsewhere
		p.midground = ebiten.NewImage(int(float64(g.board.PixelWidth)), int(float64(g.board.PixelHeight)))
		ht := float64(asset.HillsMidground.Bounds().Dy())
		midY := float64(g.board.PixelHeight) - (ht + 100)
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
	p.view.DrawImage(g.vp.midground.SubImage(image.Rect(int(g.vp.midX), int(g.vp.midY), int(g.vp.midX+screenWidth), int(g.vp.midY+screenHeight))).(*ebiten.Image), &ebiten.DrawImageOptions{})

	// render the tiles. calculate the top-left origin of the viewport in units of tiles (tileX,tileY).
	// then calculate the offset in pixels (ox,oy) that we begin drawing from the top-left tile from. the
	// pixel offset allows for smooth scrolling in smaller units (pixel) instead of large (tile).
	tileX, tileY := p.TilePosition()
	ox, oy := float64(int(p.viewX)%tileSize), float64(int(p.viewY)%tileSize)
	yTileCount := 0
	for ty := tileY; ty <= (tileY+screenHeightTiles+1) && ty < g.board.TileHeight; ty++ {
		xTileCount := 0
		for tx := tileX; tx <= (tileX+screenWidthTiles+1) && tx < g.board.TileWidth; tx++ {
			tile := g.board.Map[ty][tx]
			if tile != 0 {
				// adjust the x,y position where the tile is rendered from to be its top-left corner.
				// the tile's x,y coordinates in chipmunk 2d space are at it's center, so need to pull it
				// back and to the left by tileSize/2, and adjust by (ox,oy) to get the correct pixel offset
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(xTileCount*tileSize)-ox-(tileSize/2), float64(yTileCount*tileSize)-oy-(tileSize/2))
				p.view.DrawImage(asset.TileImages[tile], &op)
			}
			xTileCount++
		}
		yTileCount++
	}

	// render sprites
	for _, entity := range g.registry.DrawableEntities() {
		if p.InView(entity) {
			// The entity's position in chipmunk 2d space is its center of gravity.
			// Use these coordinates as starting point for translating to view coordinates (in pixels)
			x, y := entity.Position()

			// Find the top-left corner of the entity's shape and pull its position back to top-left corner.
			x -= (entity.Shape.BB().R - entity.Shape.BB().L) / 2
			y -= (entity.Shape.BB().T - entity.Shape.BB().B) / 2

			// Adjust to get the correct pixel offset within the view
			x -= p.viewX
			y -= p.viewY

			// draw the entity
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)
			p.view.DrawImage(entity.Image(), &op)
		}
	}
}

type Game struct {
	vp       viewport
	board    boards.Gameboard
	debug    bool
	registry entity.Registry
	space    *cp.Space
}

func (g *Game) animateSprites() {
	for _, entity := range g.registry.Entities {
		if entity.Animations != nil {
			entity.Animations[entity.State].Animate()
		}
	}
}

func (g *Game) MovePlayer() {
	// TODO: maybe refactor movement to an interface on the Entity struct

	var p = g.registry.Player()

	if !input.AnyKeyPressed() && p.Grounded {
		if p.Facing == entity.Right {
			p.State = entity.IdleRight
		} else if p.Facing == entity.Left {
			p.State = entity.IdleLeft
		}
	}

	if !ebiten.IsKeyPressed(ebiten.KeySpace) {
		if p.Grounded {
			p.Boost = PlayerJumpBoostHeight
		} else {
			p.Boost = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Facing = entity.Right
		if p.Grounded {
			p.State = entity.MovingRight
		} else if p.State == entity.JumpingLeft {
			p.State = entity.JumpingRight
		}
		vx, vy := p.Body.Velocity().X, 0.0
		vx += PlayerAccelerationStepX
		p.Body.ApplyForceAtWorldPoint(cp.Vector{vx, vy}, p.Body.Position())
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Facing = entity.Left
		if p.Grounded {
			p.State = entity.MovingLeft
		} else if p.State == entity.JumpingRight {
			p.State = entity.JumpingLeft
		}
		vx, vy := p.Body.Velocity().X, 0.0
		vx -= PlayerAccelerationStepX
		p.Body.ApplyForceAtWorldPoint(cp.Vector{vx, vy}, p.Body.Position())
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		// TODO: crouch
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.Boost > 0 {
		if p.State == entity.JumpingRight || p.State == entity.JumpingLeft {
			// player is already in a jump, diminish boost
			p.Boost--
			p.Body.ApplyForceAtWorldPoint(cp.Vector{0, -PlayerJumpContinueVelocity}, p.Body.Position())
		} else {
			// player is in some other state, so must be initiating the jump
			if p.Facing == entity.Left {
				p.State = entity.JumpingLeft
			} else {
				p.State = entity.JumpingRight
			}
			p.Body.ApplyForceAtWorldPoint(cp.Vector{0, -PlayerJumpInitialVelocity}, p.Body.Position())
			p.Grounded = false
			p.Shape.SetFriction(0)
		}
		if p.Boost <= 0 {
			p.Boost = 0
			if p.Facing == entity.Left {
				p.State = entity.FallingLeft
			} else {
				p.State = entity.FallingRight
			}
		}
	}

	// determine if player is falling, change friction and sprite animation accordingly
	if p.Body.Velocity().Y > 70 && !p.OnSlope {
		// player has steady downward velocity and is falling
		if p.Facing == entity.Right {
			p.State = entity.FallingRight
		}
		if p.Facing == entity.Left {
			p.State = entity.FallingLeft
		}
		p.Grounded = false
		p.Shape.SetFriction(0)
	} else if p.Body.Velocity().Y >= 0.01 && !p.OnSlope {
		// player has a little bit of downward velocity but may not be falling
		// this helps prevent player from becoming "grounded" a.k.a. stuck on
		// vertically stacked tiles
		p.Grounded = false
		p.Shape.SetFriction(0)
	}
	if p.Grounded && !p.OnSlope {
		// reset the player to normal friction and y velocity for being on solid ground
		p.Shape.SetFriction(0.75)
		p.Body.SetVelocity(p.Body.Velocity().X, 0)
	}
	if p.Grounded && p.OnSlope && p.Body.Velocity().Y > 2 {
		// if the player is sliding down a slope, add extra friction to control it
		p.Shape.SetFriction(4.0)
	}

	// enforce maximum velocity in each direction
	if p.Body.Velocity().X > PlayerMaxVelocityX {
		p.Body.SetVelocity(PlayerMaxVelocityX, p.Body.Velocity().Y)
	}
	if p.Body.Velocity().X < -PlayerMaxVelocityX {
		p.Body.SetVelocity(-PlayerMaxVelocityX, p.Body.Velocity().Y)
	}
	if p.Body.Velocity().Y > PlayerMaxVelocityY {
		p.Body.SetVelocity(p.Body.Velocity().X, PlayerMaxVelocityY)
	}
	if p.Body.Velocity().Y < -PlayerJumpVelocityLimit {
		p.Body.SetVelocity(p.Body.Velocity().X, -PlayerJumpVelocityLimit)
	}
}

func (g *Game) Update() error {
	// get user input and move the player entity
	g.MovePlayer()

	// render the image of the current viewport, centered on player
	g.vp.Center(g.registry.Player().Body.Position().X, g.registry.Player().Body.Position().Y)
	g.vp.Draw(g)

	// animate sprites
	g.animateSprites()

	// update physics space
	g.space.Step(1.0 / float64(ebiten.TPS()))

	// return any errors
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.vp.view, &ebiten.DrawImageOptions{})
	if g.debug {
		tx, ty := g.vp.TilePosition()
		px, py := g.registry.Player().Body.Position().X, g.registry.Player().Body.Position().Y
		vx, vy := g.registry.Player().Body.Velocity().X, g.registry.Player().Body.Velocity().Y
		x, y := g.vp.Position()
		debugMsg := fmt.Sprintf(
			"TPS: %0.2f Origin X,Y: (%0.2f, %0.2f) Tile X,Y: (%v, %v)\nPlayer X,Y (%0.2f, %0.2f)\nVelocity X,Y (%0.2f, %0.2f)\nGrounded: %v Boost: %v",
			ebiten.ActualTPS(), x, y, tx, ty, px, py, vx, vy, g.registry.Player().Grounded, g.registry.Player().Boost)
		ebitenutil.DebugPrint(screen, debugMsg)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	debugMode := flag.Bool("debug", false, "Enable debug features")
	flag.Parse()

	asset.LoadTiles()
	asset.LoadSprites()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(ticksPerSecond)

	space := cp.NewSpace()
	space.SetGravity(cp.Vector{0, Gravity})
	// allow no overlap between shapes in the space, to reduce prevalence of tile overlap/collision bug
	//space.SetCollisionSlop(0.00)

	player := entity.InitializePlayer(space, 0, 0)
	r := entity.Registry{}
	r.AddEntity(player)

	gameboard := boards.Gameboard{}
	gameboard.LoadGameboard(boards.Level1Map, space)

	g := Game{
		debug: *debugMode,
		vp: viewport{
			maxViewX: float64(gameboard.PixelWidth - screenWidth - tileSize),
			maxViewY: float64(gameboard.PixelHeight - screenHeight - tileSize),
		},
		board:    gameboard,
		registry: r,
		space:    space,
	}

	// initialize all the tiles in this board into the game physics space
	//boards.InitializeTiles(space, boards.GameBoard)

	// set the initial position of the viewport
	g.vp.Center(player.Position())

	// run the main loop
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
