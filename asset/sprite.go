package asset

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

var (
	//go:embed sprite/player_climb_1.png
	playerClimb1Png []byte

	//go:embed sprite/player_climb_2.png
	playerClimb2Png []byte

	//go:embed sprite/player_climb_3.png
	playerClimb3Png []byte

	//go:embed sprite/player_crouch_1.png
	playerCrouch1Png []byte

	//go:embed sprite/player_crouch_2.png
	playerCrouch2Png []byte

	//go:embed sprite/player_hurt_1.png
	playerHurt1Png []byte

	//go:embed sprite/player_hurt_2.png
	playerHurt2Png []byte

	//go:embed sprite/player_idle_1.png
	playerIdle1Png []byte

	//go:embed sprite/player_idle_2.png
	playerIdle2Png []byte

	//go:embed sprite/player_idle_3.png
	playerIdle3Png []byte

	//go:embed sprite/player_idle_4.png
	playerIdle4Png []byte

	//go:embed sprite/player_jump_1.png
	playerJump1Png []byte

	//go:embed sprite/player_jump_2.png
	playerJump2Png []byte

	//go:embed sprite/player_run_1.png
	playerRun1Png []byte

	//go:embed sprite/player_run_2.png
	playerRun2Png []byte

	//go:embed sprite/player_run_3.png
	playerRun3Png []byte

	//go:embed sprite/player_run_4.png
	playerRun4Png []byte

	//go:embed sprite/player_run_5.png
	playerRun5Png []byte

	//go:embed sprite/player_run_6.png
	playerRun6Png []byte

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
	PlayerRun1 = imageFromBytes(playerRun1Png)
	PlayerRun2 = imageFromBytes(playerRun2Png)
	PlayerRun3 = imageFromBytes(playerRun3Png)
	PlayerRun4 = imageFromBytes(playerRun4Png)
	PlayerRun5 = imageFromBytes(playerRun5Png)
	PlayerRun6 = imageFromBytes(playerRun6Png)
}
