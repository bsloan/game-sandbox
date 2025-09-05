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

var fadeValue float32 = 0.0
var fadeDesc = false

var levelSelectPoints = []LevelSelectPoint{
	{x: 28, y: 100, state: Playable},
	{x: 120, y: 34, state: Playable},
	{x: 86, y: 66, state: Playable},
	{x: 75, y: 140, state: Playable},
	{x: 175, y: 128, state: Playable},
	{x: 268, y: 56, state: Locked},
	{x: 260, y: 96, state: Locked},
}

func (g *Game) levelSelect() error {
	if g.vp.view == nil {
		// init viewport for the title screen if it's not set up yet
		g.vp.view = ebiten.NewImage(settings.ScreenWidth, settings.ScreenHeight)
	}

	g.vp.view.Clear()
	g.vp.view.DrawImage(asset.LevelSelectMap, &ebiten.DrawImageOptions{})

	// fade the cursors in and out
	cs := ebiten.ColorScale{}
	if fadeDesc {
		fadeValue -= 0.03
	} else {
		fadeValue += 0.03
	}
	if fadeValue <= 0.0 {
		fadeValue = 0.0
		fadeDesc = false
	} else if fadeValue >= 1.0 {
		fadeValue = 1.0
		fadeDesc = true
	}
	cs.SetR(fadeValue)
	cs.SetG(fadeValue)
	cs.SetB(fadeValue)
	cs.SetA(fadeValue)

	// draw the level cursors on the map
	for _, point := range levelSelectPoints {
		var cursorImage *ebiten.Image
		op := &ebiten.DrawImageOptions{}
		if point.state == Playable {
			cursorImage = asset.LevelSelectCursorWhite
			op.ColorScale = cs
		} else if point.state == Completed {
			cursorImage = asset.LevelSelectCursorGreen
		} else {
			// locked
			cursorImage = asset.LevelSelectCursorRed
		}
		op.GeoM.Translate(point.x, point.y)
		g.vp.view.DrawImage(cursorImage, op)
	}

	// get player input
	if (g.inputLeft() || g.inputUp()) && g.inputAvailable {

	} else if (g.inputRight() || g.inputDown()) && g.inputAvailable {

	} else if g.inputAttack() && g.inputAvailable {
		g.inputAvailable = false
		g.gameMode = InitializingGameplayMode
	} else if !g.inputAny() {
		g.inputAvailable = true
	}

	return nil
}
