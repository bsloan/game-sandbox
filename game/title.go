package game

import (
	"image/color"

	"github.com/bsloan/game-sandbox/asset"
	"github.com/bsloan/game-sandbox/settings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var titleOptions = map[string]GameMode{
	"New Game": InitializingGameplayMode,
	"About":    TitleMode,
	"Exit":     ExitingMode,
}

func (g *Game) titleScreen() error {
	if g.vp.view == nil {
		// init viewport for the title screen if it's not set up yet
		g.vp.view = ebiten.NewImage(settings.ScreenWidth, settings.ScreenHeight)
	}

	g.vp.view.Clear()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(40, 0)
	g.vp.view.DrawImage(asset.TitleScreen, op)

	// show title options - highlight the selected option in red
	options := []string{"New Game", "About", "Exit"}
	for i, option := range options {
		textOp := &text.DrawOptions{}
		textOp.GeoM.Translate(130, 160+float64(i*16))
		if g.titleSelection == i {
			textOp.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 0, B: 0, A: 255})
		} else {
			textOp.ColorScale.ScaleWithColor(color.White)
		}
		text.Draw(g.vp.view, option, &text.GoTextFace{
			Source: asset.BoldPixelsFS,
			Size:   16,
		}, textOp)
	}

	// make the selection
	if (g.inputLeft() || g.inputUp()) && g.inputAvailable {
		g.inputAvailable = false
		if g.titleSelection > 0 {
			g.titleSelection--
		}
	} else if (g.inputRight() || g.inputDown()) && g.inputAvailable {
		g.inputAvailable = false
		if g.titleSelection < len(options)-1 {
			g.titleSelection++
		}
	} else if g.inputAttack() {
		g.gameMode = titleOptions[options[g.titleSelection]]
	} else if !g.inputAny() {
		g.inputAvailable = true
	}

	return nil
}
