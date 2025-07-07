package main

import (
	"flag"
	"fmt"
	"github.com/bsloan/game-sandbox/assets"
	"github.com/bsloan/game-sandbox/boards"
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

var (
	TileImages = make(map[int]*ebiten.Image)
)

type viewport struct {
	viewX     int
	viewY     int
	maxViewX  int
	maxViewY  int
	view      *ebiten.Image
	midground *ebiten.Image
	midX      int
	midY      int
}

func (p *viewport) Move(x, y int) {
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
	p.midX = int(float64(p.viewX) * midgroundScrollMultiplier)
	p.midY = int(float64(p.viewY) * midgroundScrollMultiplier)
}

// Position returns the pixel X, Y coordinates in the game board of the top-left
// pixel of the viewport.
func (p *viewport) Position() (int, int) {
	return p.viewX, p.viewY
}

// TilePosition returns the tile coordinates of the top-leftmost visible tile.
func (p *viewport) TilePosition() (int, int) {
	tx := p.viewX / tileSize
	ty := p.viewY / tileSize
	return tx, ty
}

// Draw renders the foreground (tiles and sprites) within the currently visible section
// of the game board. When the frame is ready, this is later copied to the screen on top
// of any background or middle layers for a parallax scrolling affect.
func (p *viewport) Draw() {
	// ebiten performance: avoid allocating a new image on every Update, use Clear instead
	if p.view == nil {
		p.view = ebiten.NewImage(screenWidth+tileSize, screenHeight+tileSize)
	}
	if p.midground == nil {
		// iterate horizontally in chunks of 160 pixels (width of midground image)
		// across the entire game board. draw any instances of the image that happen
		// to be visible within the viewport boundaries. skip any that are not visible.
		// draw them at constant height
		p.midground = ebiten.NewImage(int(float64(boards.GameBoardPixelWidth)), int(float64(boards.GameBoardPixelHeight)))
		midY := float64(300)
		for midX := 0; midX < boards.GameBoardPixelWidth; midX += 160 {
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(midX), midY)
			p.midground.DrawImage(assets.HillsMidground, &op)
		}
	}

	p.view.Clear()

	// render the tiles
	tileX, tileY := p.TilePosition()
	yTileCount := 0
	for ty := tileY; ty <= (tileY+screenHeightTiles+1) && ty < len(boards.GameBoard); ty++ {
		xTileCount := 0
		for tx := tileX; tx <= (tileX+screenWidthTiles+1) && tx < len(boards.GameBoard[0]); tx++ {
			tile := boards.GameBoard[ty][tx]
			if tile != 0 {
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(xTileCount*tileSize), float64(yTileCount*tileSize))
				p.view.DrawImage(TileImages[tile], &op)
			}
			xTileCount++
		}
		yTileCount++
	}

	// TODO: render sprites
}

type Game struct {
	vp    viewport
	board [][]int
	debug bool
	ticks uint64
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.vp.Move(g.vp.viewX+1, g.vp.viewY)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.vp.Move(g.vp.viewX-1, g.vp.viewY)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.vp.Move(g.vp.viewX, g.vp.viewY+1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.vp.Move(g.vp.viewX, g.vp.viewY-1)
	}

	// render the image of the current viewport
	g.vp.Draw()

	// update ticks counter
	g.ticks++

	// return any errors
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x, y := g.vp.Position()
	noOp := ebiten.DrawImageOptions{}

	// render the static background image
	screen.DrawImage(assets.SkyBackground, &noOp)

	// render the middle layer
	screen.DrawImage(g.vp.midground.SubImage(image.Rect(g.vp.midX, g.vp.midY, g.vp.midX+screenWidth, g.vp.midY+screenHeight)).(*ebiten.Image), &noOp)

	// render the front layer (tiles and sprints) from the viewport. calculate the
	// offset coords (ox, oy) from where we begin the copy from viewport to the screen.
	ox, oy := x%tileSize, y%tileSize
	screen.DrawImage(g.vp.view.SubImage(image.Rect(ox, oy, ox+screenWidth, oy+screenHeight)).(*ebiten.Image), &noOp)

	if g.debug {
		tx, ty := g.vp.TilePosition()
		debugMsg :=
			fmt.Sprintf(
				"TPS: %0.2f Origin X,Y: (%v, %v) Tile X,Y: (%v, %v)\nOffset X,Y (%v, %v)",
				ebiten.ActualTPS(), x, y, tx, ty, ox, oy)
		ebitenutil.DebugPrint(screen, debugMsg)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	debugMode := flag.Bool("debug", false, "Enable debug features")
	flag.Parse()

	assets.Initialize()

	// TODO: refactor
	TileImages[1] = assets.GrassLeft
	TileImages[2] = assets.GrassMiddle
	TileImages[3] = assets.GrassRight
	TileImages[4] = assets.DirtMiddle

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(ticksPerSecond)

	g := Game{
		debug: *debugMode,
		vp: viewport{
			viewX:    0,
			viewY:    0,
			maxViewX: boards.GameBoardPixelWidth - screenWidth,
			maxViewY: boards.GameBoardPixelHeight - screenHeight,
		},
		board: boards.GameBoard,
	}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
