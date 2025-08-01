package main

import (
	"flag"
	"github.com/bsloan/game-sandbox/asset"
	"github.com/bsloan/game-sandbox/boards"
	"github.com/bsloan/game-sandbox/entity"
	"github.com/bsloan/game-sandbox/game"
	"github.com/bsloan/game-sandbox/settings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"log"
)

func main() {
	debugMode := flag.Bool("debug", false, "Enable debug features")
	flag.Parse()

	asset.LoadTiles()
	asset.LoadSprites()

	ebiten.SetWindowSize(settings.ScreenWidth*2, settings.ScreenHeight*2)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(settings.TicksPerSecond)

	space := cp.NewSpace()
	space.SetGravity(cp.Vector{0, settings.Gravity})
	// allow no overlap between shapes in the space, to reduce prevalence of tile overlap/collision bug
	//space.SetCollisionSlop(0.00)

	player := entity.InitializePlayer(space, 35, 494)
	r := entity.Registry{}
	r.AddEntity(player)

	swordDog := entity.InitializeSwordDog(space, 290, 382)
	r.AddEntity(swordDog)

	sword := entity.InitializePlayerSword(space, 0, 0)
	r.AddEntity(sword)

	entity.InitializeCollisionHandlers(space)

	gameboard := boards.Gameboard{}
	gameboard.LoadGameboard(boards.Level1Map, space)

	// initialize a new game
	g := game.NewGame(
		game.Viewport{
			MaxViewX: float64(gameboard.PixelWidth - settings.ScreenWidth - settings.TileSize),
			MaxViewY: float64(gameboard.PixelHeight - settings.ScreenHeight - settings.TileSize),
		},
		gameboard,
		*debugMode,
		r,
		space,
	)

	// set the initial position of the viewport
	g.CenterViewport(player.Position())

	// run the main loop
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
