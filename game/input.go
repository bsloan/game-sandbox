package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	NonstandardGamepadButtonLeft  = 18
	NonstandardGamepadButtonRight = 16
	NonstandardGamepadButtonUp    = 15
	NonstandardGamepadButtonDown  = 17
	NonstandardGamepadButtonX     = 3
	NonstandardGamepadButtonA     = 0
)

func (g *Game) gamepadAvailable() bool {
	return len(g.gamepadIds) > 0
}

func (g *Game) inputAny() bool {
	return g.inputJump() || g.inputLeft() || g.inputRight() || g.inputUp() || g.inputDown() || g.inputAttack()
}

func (g *Game) inputJump() bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace) || (g.gamepadAvailable() && ebiten.IsGamepadButtonPressed(g.gamepadIds[0], NonstandardGamepadButtonA))
}

func (g *Game) inputLeft() bool {
	return ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) || (g.gamepadAvailable() && ebiten.IsGamepadButtonPressed(g.gamepadIds[0], NonstandardGamepadButtonLeft))
}

func (g *Game) inputRight() bool {
	return ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) || (g.gamepadAvailable() && ebiten.IsGamepadButtonPressed(g.gamepadIds[0], NonstandardGamepadButtonRight))
}

func (g *Game) inputUp() bool {
	return ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) || (g.gamepadAvailable() && ebiten.IsGamepadButtonPressed(g.gamepadIds[0], NonstandardGamepadButtonUp))
}

func (g *Game) inputDown() bool {
	return ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) || (g.gamepadAvailable() && ebiten.IsGamepadButtonPressed(g.gamepadIds[0], NonstandardGamepadButtonDown))
}

func (g *Game) inputAttack() bool {
	return ebiten.IsKeyPressed(ebiten.KeyK) || ebiten.IsKeyPressed(ebiten.KeyAlt) || (g.gamepadAvailable() && ebiten.IsGamepadButtonPressed(g.gamepadIds[0], NonstandardGamepadButtonX))
}
