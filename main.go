package main

import (
	"flag"
	"fmt"
	"github.com/bsloan/game-sandbox/asset"
	"github.com/bsloan/game-sandbox/boards"
	"github.com/bsloan/game-sandbox/entities"
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
	viewX      int
	viewY      int
	viewWidth  int
	viewHeight int
	maxViewX   int
	maxViewY   int
	view       *ebiten.Image
	midground  *ebiten.Image
	midX       int
	midY       int
}

// Move moves the viewport to pixel coordinates x,y in the game board.
// If coordinates are out-of-bounds of the map, they are adjusted to be
// within bounds. The middle layer for parallax scrolling is also updated
// according to the provided coordinates using the scroll multiplier.
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

// Draw renders the foreground layer (tiles and sprites) within the currently visible section
// of the game board. When the frame is ready, this is later copied to the screen on top
// of any background and middle layers for a parallax scrolling affect.
func (p *viewport) Draw(g *Game) {
	// ebiten performance: avoid allocating a new image on every Update, use Clear instead
	if p.view == nil {
		p.view = ebiten.NewImage(p.viewWidth, p.viewHeight)
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
				p.view.DrawImage(asset.TileImages[tile], &op)
			}
			xTileCount++
		}
		yTileCount++
	}

	// render sprites
	// TODO: refactor to a receiver method on Entity, GetDrawableEntities or similar
	for _, entity := range g.registry.Entities {
		if entity.ActiveImage != nil {
			if int(entity.XPos) >= p.viewX && int(entity.XPos) < p.viewX+p.viewWidth && int(entity.YPos) >= p.viewY && int(entity.YPos) < p.viewY+p.viewHeight {
				x, y := entity.XPos-float64(p.viewX), entity.YPos-float64(p.viewY)
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(x, y)
				p.view.DrawImage(entity.ActiveImage, &op)
			}
		}
	}
}

type Game struct {
	vp       viewport
	board    [][]int
	debug    bool
	ticks    uint64
	registry entities.Registry
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
	g.vp.Draw(g)

	// animate sprites
	// TODO: refactor
	for _, entity := range g.registry.Entities {
		if entity.Animations != nil {
			entity.ActiveImage = entity.Animations[entity.State].Animate()
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

	asset.LoadTiles()
	asset.LoadSprites()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(ticksPerSecond)

	r := entities.Registry{}
	player := entities.InitializePlayer(250, 250)
	r.AddEntity(*player)

	g := Game{
		debug: *debugMode,
		vp: viewport{
			viewWidth:  screenWidth + tileSize,
			viewHeight: screenHeight + tileSize,
			maxViewX:   boards.GameBoardPixelWidth - screenWidth,
			maxViewY:   boards.GameBoardPixelHeight - screenHeight,
		},
		board:    boards.GameBoard,
		registry: r,
	}

	// set the initial position of the viewport
	// TODO: refactor to a receiver function on the viewport
	g.vp.Move(int(player.XPos-screenWidth/2), int(player.YPos-screenHeight/2))

	// run the main loop
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
