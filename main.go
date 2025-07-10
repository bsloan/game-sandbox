package main

import (
	"flag"
	"fmt"
	"github.com/bsloan/game-sandbox/asset"
	"github.com/bsloan/game-sandbox/boards"
	"github.com/bsloan/game-sandbox/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
)

const (
	ticksPerSecond            = 120
	screenWidth               = 320
	screenHeight              = 240
	tileSize                  = 16
	screenWidthTiles          = screenWidth / tileSize
	screenHeightTiles         = screenHeight / tileSize
	midgroundScrollMultiplier = 0.5
)

type viewport struct {
	viewX     float64
	viewY     float64
	maxViewX  float64
	maxViewY  float64
	view      *ebiten.Image
	midground *ebiten.Image
	midX      float64
	midY      float64
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
		p.midground = ebiten.NewImage(int(float64(boards.GameBoardPixelWidth)), int(float64(boards.GameBoardPixelHeight)))
		ht := float64(asset.HillsMidground.Bounds().Dy()) * 1.75
		midY := float64(boards.GameBoardPixelHeight) - ht
		for midX := 0; midX < boards.GameBoardPixelWidth; midX += 160 {
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(midX), midY)
			p.midground.DrawImage(asset.HillsMidground, &op)
		}
	}

	p.view.Clear()

	// render the tiles. calculate the top-left origin of the viewport in units of tiles (tileX,tileY).
	// then calculate the offset in pixels (ox,oy) that we begin drawing from the top-left tile from. the
	// pixel offset allows for smooth scrolling in smallelr units (pixel) instead of large (tile).
	tileX, tileY := p.TilePosition()
	ox, oy := float64(int(p.viewX)%tileSize), float64(int(p.viewY)%tileSize)
	yTileCount := 0
	for ty := tileY; ty <= (tileY+screenHeightTiles) && ty < len(boards.GameBoard); ty++ {
		xTileCount := 0
		for tx := tileX; tx <= (tileX+screenWidthTiles) && tx < len(boards.GameBoard[0]); tx++ {
			tile := boards.GameBoard[ty][tx]
			if tile != 0 {
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(xTileCount*tileSize)-ox, float64(yTileCount*tileSize)-oy)
				p.view.DrawImage(asset.TileImages[tile], &op)
			}
			xTileCount++
		}
		yTileCount++
	}

	// render sprites
	// TODO: refactor to a receiver method on Entity, GetDrawableEntities or similar
	for _, entity := range g.registry.Entities {
		if entity.Animations != nil {
			if entity.XPos >= p.viewX && entity.XPos < p.viewX+screenWidth && entity.YPos >= p.viewY && entity.YPos < p.viewY+screenHeight {
				x, y := entity.XPos-p.viewX, entity.YPos-p.viewY
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(x, y)
				p.view.DrawImage(entity.Image(), &op)
			}
		}
	}
}

type Game struct {
	vp       viewport
	board    [][]int
	debug    bool
	ticks    uint64
	registry entity.Registry
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.registry.Player().XPos += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.registry.Player().XPos -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.registry.Player().YPos += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.registry.Player().YPos -= 1
	}

	// render the image of the current viewport, centered on player
	g.vp.Center(g.registry.Player().XPos, g.registry.Player().YPos)
	g.vp.Draw(g)

	// animate sprites
	// TODO: refactor
	for _, entity := range g.registry.Entities {
		if entity.Animations != nil {
			entity.Animations[entity.State].Animate()
		}
	}

	// update ticks counter
	g.ticks++

	// return any errors
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x, y := g.vp.Position()
	noOp := ebiten.DrawImageOptions{}

	// render the static background image
	screen.DrawImage(asset.SkyBackground, &noOp)

	// render the middle layer
	screen.DrawImage(g.vp.midground.SubImage(image.Rect(int(g.vp.midX), int(g.vp.midY), int(g.vp.midX+screenWidth), int(g.vp.midY+screenHeight))).(*ebiten.Image), &noOp)

	// render the front layer (tiles and sprints) from the viewport
	screen.DrawImage(g.vp.view, &noOp)

	if g.debug {
		tx, ty := g.vp.TilePosition()
		debugMsg := fmt.Sprintf("TPS: %0.2f Origin X,Y: (%v, %v) Tile X,Y: (%v, %v)\nPlayer X,Y (%v, %v)", ebiten.ActualTPS(), x, y, tx, ty, g.registry.Player().XPos, g.registry.Player().YPos)
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

	player := entity.InitializePlayer(250, 250)
	r := entity.Registry{}
	r.AddEntity(*player)

	g := Game{
		debug: *debugMode,
		vp: viewport{
			maxViewX: float64(boards.GameBoardPixelWidth - screenWidth),
			maxViewY: float64(boards.GameBoardPixelHeight - screenHeight),
		},
		board:    boards.GameBoard,
		registry: r,
	}

	// set the initial position of the viewport
	g.vp.Center(player.XPos, player.YPos)

	// run the main loop
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
