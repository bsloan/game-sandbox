package main

import (
	"flag"
	"log"

	"github.com/bsloan/game-sandbox/asset"
	"github.com/bsloan/game-sandbox/game"
	"github.com/bsloan/game-sandbox/settings"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	debugMode := flag.Bool("debug", false, "Enable debug features")
	flag.Parse()

	asset.LoadTiles()
	asset.LoadSprites()

	ebiten.SetWindowSize(settings.ScreenWidth*2, settings.ScreenHeight*2)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Ninja Dōru ドール")
	ebiten.SetTPS(settings.TicksPerSecond)

	// initialize a new game
	g := game.NewGame(
		*debugMode,
	)

	// initialize the behavior of all the entities in the game
	game.InitializeEntityBehavior(g)

	// run the main loop
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
