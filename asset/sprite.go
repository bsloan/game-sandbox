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
	playerIdleRight1Png []byte

	//go:embed sprite/player_idle_2.png
	playerIdleRight2Png []byte

	//go:embed sprite/player_idle_3.png
	playerIdleRight3Png []byte

	//go:embed sprite/player_idle_4.png
	playerIdleRight4Png []byte

	//go:embed sprite/player_jump_1.png
	playerJump1Png []byte

	//go:embed sprite/player_jump_2.png
	playerJump2Png []byte

	//go:embed sprite/player_run_1.png
	playerMoveRight1Png []byte

	//go:embed sprite/player_run_2.png
	playerMoveRight2Png []byte

	//go:embed sprite/player_run_3.png
	playerMoveRight3Png []byte

	//go:embed sprite/player_run_4.png
	playerMoveRight4Png []byte

	//go:embed sprite/player_run_5.png
	playerMoveRight5Png []byte

	//go:embed sprite/player_run_6.png
	playerMoveRight6Png []byte

	PlayerIdleRight1 *ebiten.Image
	PlayerIdleRight2 *ebiten.Image
	PlayerIdleRight3 *ebiten.Image
	PlayerIdleRight4 *ebiten.Image
	PlayerIdleLeft1  *ebiten.Image
	PlayerIdleLeft2  *ebiten.Image
	PlayerIdleLeft3  *ebiten.Image
	PlayerIdleLeft4  *ebiten.Image

	PlayerMoveRight1 *ebiten.Image
	PlayerMoveRight2 *ebiten.Image
	PlayerMoveRight3 *ebiten.Image
	PlayerMoveRight4 *ebiten.Image
	PlayerMoveRight5 *ebiten.Image
	PlayerMoveRight6 *ebiten.Image
	PlayerMoveLeft1  *ebiten.Image
	PlayerMoveLeft2  *ebiten.Image
	PlayerMoveLeft3  *ebiten.Image
	PlayerMoveLeft4  *ebiten.Image
	PlayerMoveLeft5  *ebiten.Image
	PlayerMoveLeft6  *ebiten.Image

	PlayerJumpRight1 *ebiten.Image
	PlayerJumpRight2 *ebiten.Image
)

func LoadSprites() {

	PlayerIdleRight1 = imageFromBytes(playerIdleRight1Png)
	PlayerIdleRight2 = imageFromBytes(playerIdleRight2Png)
	PlayerIdleRight3 = imageFromBytes(playerIdleRight3Png)
	PlayerIdleRight4 = imageFromBytes(playerIdleRight4Png)

	PlayerIdleLeft1 = flipImageXAxis(PlayerIdleRight1)
	PlayerIdleLeft2 = flipImageXAxis(PlayerIdleRight2)
	PlayerIdleLeft3 = flipImageXAxis(PlayerIdleRight3)
	PlayerIdleLeft4 = flipImageXAxis(PlayerIdleRight4)

	PlayerMoveRight1 = imageFromBytes(playerMoveRight1Png)
	PlayerMoveRight2 = imageFromBytes(playerMoveRight2Png)
	PlayerMoveRight3 = imageFromBytes(playerMoveRight3Png)
	PlayerMoveRight4 = imageFromBytes(playerMoveRight4Png)
	PlayerMoveRight5 = imageFromBytes(playerMoveRight5Png)
	PlayerMoveRight6 = imageFromBytes(playerMoveRight6Png)

	PlayerMoveLeft1 = flipImageXAxis(PlayerMoveRight1)
	PlayerMoveLeft2 = flipImageXAxis(PlayerMoveRight2)
	PlayerMoveLeft3 = flipImageXAxis(PlayerMoveRight3)
	PlayerMoveLeft4 = flipImageXAxis(PlayerMoveRight4)
	PlayerMoveLeft5 = flipImageXAxis(PlayerMoveRight5)
	PlayerMoveLeft6 = flipImageXAxis(PlayerMoveRight6)

}
