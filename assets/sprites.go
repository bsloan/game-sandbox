package assets

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

var (
	//go:embed sprites/player_climb_1.png
	playerClimb1 []byte

	//go:embed sprites/player_climb_2.png
	playerClimb2 []byte

	//go:embed sprites/player_climb_3.png
	playerClimb3 []byte

	//go:embed sprites/player_crouch_1.png
	playerCrouch1 []byte

	//go:embed sprites/player_crouch_2.png
	playerCrouch2 []byte

	//go:embed sprites/player_hurt_1.png
	playerHurt1 []byte

	//go:embed sprites/player_hurt_2.png
	playerHurt2 []byte

	//go:embed sprites/player_idle_1.png
	playerIdle1 []byte

	//go:embed sprites/player_idle_2.png
	playerIdle2 []byte

	//go:embed sprites/player_idle_3.png
	playerIdle3 []byte

	//go:embed sprites/player_idle_4.png
	playerIdle4 []byte

	//go:embed sprites/player_jump_1.png
	playerJump1 []byte

	//go:embed sprites/player_jump_2.png
	playerJump2 []byte

	//go:embed sprites/player_run_1.png
	playerRun1 []byte

	//go:embed sprites/player_run_2.png
	playerRun2 []byte

	//go:embed sprites/player_run_3.png
	playerRun3 []byte

	//go:embed sprites/player_run_4.png
	playerRun4 []byte

	//go:embed sprites/player_run_5.png
	playerRun5 []byte

	//go:embed sprites/player_run_6.png
	playerRun6 []byte
)

var (
	PlayerRun1  *ebiten.Image
	PlayerRun2  *ebiten.Image
	PlayerRun3  *ebiten.Image
	PlayerRun4  *ebiten.Image
	PlayerRun5  *ebiten.Image
	PlayerRun6  *ebiten.Image
	PlayerJump1 *ebiten.Image
	PlayerJump2 *ebiten.Image
)

func LoadSprites() {
	GrassLeft = imageFromBytes(grassLeftPng)
	GrassMiddle = imageFromBytes(grassMiddlePng)
	GrassRight = imageFromBytes(grassRightPng)
	DirtMiddle = imageFromBytes(dirtMiddlePng)
	SkyBackground = imageFromBytes(skyBackgroundPng)
	HillsMidground = imageFromBytes(hillsMidgroundPng)
}
