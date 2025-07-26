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
	playerJumpRight1Png []byte

	//go:embed sprite/player_fall_1.png
	playerFallRight1Png []byte

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

	//go:embed sprite/white_slash1.png
	whiteSlashRight1Png []byte

	//go:embed sprite/white_slash2.png
	whiteSlashRight2Png []byte

	//go:embed sprite/white_slash3.png
	whiteSlashRight3Png []byte

	//go:embed sprite/white_slash4.png
	whiteSlashRight4Png []byte

	//go:embed sprite/white_slash5.png
	whiteSlashRight5Png []byte

	//go:embed sprite/white_slash6.png
	whiteSlashRight6Png []byte

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
	PlayerJumpLeft1  *ebiten.Image
	PlayerFallRight1 *ebiten.Image
	PlayerFallLeft1  *ebiten.Image

	WhiteSlashRight1 *ebiten.Image
	WhiteSlashRight2 *ebiten.Image
	WhiteSlashRight3 *ebiten.Image
	WhiteSlashRight4 *ebiten.Image
	WhiteSlashRight5 *ebiten.Image
	WhiteSlashRight6 *ebiten.Image

	WhiteSlashLeft1 *ebiten.Image
	WhiteSlashLeft2 *ebiten.Image
	WhiteSlashLeft3 *ebiten.Image
	WhiteSlashLeft4 *ebiten.Image
	WhiteSlashLeft5 *ebiten.Image
	WhiteSlashLeft6 *ebiten.Image
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

	PlayerJumpRight1 = imageFromBytes(playerJumpRight1Png)

	PlayerJumpLeft1 = flipImageXAxis(PlayerJumpRight1)

	PlayerFallRight1 = imageFromBytes(playerFallRight1Png)

	PlayerFallLeft1 = flipImageXAxis(PlayerFallRight1)

	WhiteSlashRight1 = imageFromBytes(whiteSlashRight1Png)
	WhiteSlashRight2 = imageFromBytes(whiteSlashRight2Png)
	WhiteSlashRight3 = imageFromBytes(whiteSlashRight3Png)
	WhiteSlashRight4 = imageFromBytes(whiteSlashRight4Png)
	WhiteSlashRight5 = imageFromBytes(whiteSlashRight5Png)
	WhiteSlashRight6 = imageFromBytes(whiteSlashRight6Png)

	WhiteSlashLeft1 = flipImageXAxis(WhiteSlashRight1)
	WhiteSlashLeft2 = flipImageXAxis(WhiteSlashRight2)
	WhiteSlashLeft3 = flipImageXAxis(WhiteSlashRight3)
	WhiteSlashLeft4 = flipImageXAxis(WhiteSlashRight4)
	WhiteSlashLeft5 = flipImageXAxis(WhiteSlashRight5)
	WhiteSlashLeft6 = flipImageXAxis(WhiteSlashRight6)
}
