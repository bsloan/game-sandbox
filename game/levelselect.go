package game

import (
	"github.com/bsloan/game-sandbox/asset"
	"github.com/bsloan/game-sandbox/settings"
	"github.com/hajimehoshi/ebiten/v2"
)

type LevelSelectState int

const (
	Playable LevelSelectState = iota
	Locked
	Completed
)

type LevelSelectPoint struct {
	x, y  float64
	state LevelSelectState
}

var levelSelectPoints = []LevelSelectPoint{
	{x: 75, y: 140, state: Playable},
}

func (g *Game) levelSelect() error {
	if g.vp.view == nil {
		// init viewport for the title screen if it's not set up yet
		g.vp.view = ebiten.NewImage(settings.ScreenWidth, settings.ScreenHeight)
	}

	g.vp.view.Clear()
	g.vp.view.DrawImage(asset.LevelSelectMap, &ebiten.DrawImageOptions{})

	// draw the level cursors on the map
	for _, point := range levelSelectPoints {
		var cursorImage *ebiten.Image
		if point.state == Playable {
			cursorImage = asset.LevelSelectCursorWhite
		} else if point.state == Completed {
			cursorImage = asset.LevelSelectCursorGreen
		} else {
			// locked
			cursorImage = asset.LevelSelectCursorRed
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(point.x, point.y)
		g.vp.view.DrawImage(cursorImage, op)
	}

	// get player input
	if (g.inputLeft() || g.inputUp()) && g.inputAvailable {

	} else if (g.inputRight() || g.inputDown()) && g.inputAvailable {

	} else if g.inputAttack() {

	} else if !g.inputAny() {
		g.inputAvailable = true
	}

	return nil
}
