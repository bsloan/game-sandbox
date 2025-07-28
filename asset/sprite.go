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

	//go:embed sprite/sword_dog_idle1.png
	swordDogIdleRight1Png []byte

	//go:embed sprite/sword_dog_idle2.png
	swordDogIdleRight2Png []byte

	//go:embed sprite/sword_dog_idle3.png
	swordDogIdleRight3Png []byte

	//go:embed sprite/sword_dog_idle4.png
	swordDogIdleRight4Png []byte

	//go:embed sprite/sword_dog_idle5.png
	swordDogIdleRight5Png []byte

	//go:embed sprite/sword_dog_idle6.png
	swordDogIdleRight6Png []byte

	//go:embed sprite/sword_dog_run1.png
	swordDogRunRight1Png []byte

	//go:embed sprite/sword_dog_run2.png
	swordDogRunRight2Png []byte

	//go:embed sprite/sword_dog_run3.png
	swordDogRunRight3Png []byte

	//go:embed sprite/sword_dog_run4.png
	swordDogRunRight4Png []byte

	//go:embed sprite/sword_dog_run5.png
	swordDogRunRight5Png []byte

	//go:embed sprite/sword_dog_run6.png
	swordDogRunRight6Png []byte

	//go:embed sprite/sword_dog_slash1.png
	swordDogBigSlashRight1Png []byte

	//go:embed sprite/sword_dog_slash2.png
	swordDogBigSlashRight2Png []byte

	//go:embed sprite/sword_dog_slash3.png
	swordDogBigSlashRight3Png []byte

	//go:embed sprite/sword_dog_slash4.png
	swordDogBigSlashRight4Png []byte

	//go:embed sprite/sword_dog_slash5.png
	swordDogBigSlashRight5Png []byte

	//go:embed sprite/sword_dog_slash6.png
	swordDogBigSlashRight6Png []byte

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

	SwordDogIdleRight1 *ebiten.Image
	SwordDogIdleRight2 *ebiten.Image
	SwordDogIdleRight3 *ebiten.Image
	SwordDogIdleRight4 *ebiten.Image
	SwordDogIdleRight5 *ebiten.Image
	SwordDogIdleRight6 *ebiten.Image

	SwordDogIdleLeft1 *ebiten.Image
	SwordDogIdleLeft2 *ebiten.Image
	SwordDogIdleLeft3 *ebiten.Image
	SwordDogIdleLeft4 *ebiten.Image
	SwordDogIdleLeft5 *ebiten.Image
	SwordDogIdleLeft6 *ebiten.Image

	SwordDogRunRight1 *ebiten.Image
	SwordDogRunRight2 *ebiten.Image
	SwordDogRunRight3 *ebiten.Image
	SwordDogRunRight4 *ebiten.Image
	SwordDogRunRight5 *ebiten.Image
	SwordDogRunRight6 *ebiten.Image

	SwordDogRunLeft1 *ebiten.Image
	SwordDogRunLeft2 *ebiten.Image
	SwordDogRunLeft3 *ebiten.Image
	SwordDogRunLeft4 *ebiten.Image
	SwordDogRunLeft5 *ebiten.Image
	SwordDogRunLeft6 *ebiten.Image

	SwordDogBigSlashRight1 *ebiten.Image
	SwordDogBigSlashRight2 *ebiten.Image
	SwordDogBigSlashRight3 *ebiten.Image
	SwordDogBigSlashRight4 *ebiten.Image
	SwordDogBigSlashRight5 *ebiten.Image
	SwordDogBigSlashRight6 *ebiten.Image

	SwordDogBigSlashLeft1 *ebiten.Image
	SwordDogBigSlashLeft2 *ebiten.Image
	SwordDogBigSlashLeft3 *ebiten.Image
	SwordDogBigSlashLeft4 *ebiten.Image
	SwordDogBigSlashLeft5 *ebiten.Image
	SwordDogBigSlashLeft6 *ebiten.Image
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

	SwordDogIdleRight1 = imageFromBytes(swordDogIdleRight1Png)
	SwordDogIdleRight2 = imageFromBytes(swordDogIdleRight2Png)
	SwordDogIdleRight3 = imageFromBytes(swordDogIdleRight3Png)
	SwordDogIdleRight4 = imageFromBytes(swordDogIdleRight4Png)
	SwordDogIdleRight5 = imageFromBytes(swordDogIdleRight5Png)
	SwordDogIdleRight6 = imageFromBytes(swordDogIdleRight6Png)

	SwordDogIdleLeft1 = flipImageXAxis(SwordDogIdleRight1)
	SwordDogIdleLeft2 = flipImageXAxis(SwordDogIdleRight2)
	SwordDogIdleLeft3 = flipImageXAxis(SwordDogIdleRight3)
	SwordDogIdleLeft4 = flipImageXAxis(SwordDogIdleRight4)
	SwordDogIdleLeft5 = flipImageXAxis(SwordDogIdleRight5)
	SwordDogIdleLeft6 = flipImageXAxis(SwordDogIdleRight6)

	SwordDogRunRight1 = imageFromBytes(swordDogRunRight6Png)
	SwordDogRunRight2 = imageFromBytes(swordDogRunRight5Png)
	SwordDogRunRight3 = imageFromBytes(swordDogRunRight4Png)
	SwordDogRunRight4 = imageFromBytes(swordDogRunRight3Png)
	SwordDogRunRight5 = imageFromBytes(swordDogRunRight2Png)
	SwordDogRunRight6 = imageFromBytes(swordDogRunRight1Png)

	SwordDogRunLeft1 = flipImageXAxis(SwordDogRunRight1)
	SwordDogRunLeft2 = flipImageXAxis(SwordDogRunRight2)
	SwordDogRunLeft3 = flipImageXAxis(SwordDogRunRight3)
	SwordDogRunLeft4 = flipImageXAxis(SwordDogRunRight4)
	SwordDogRunLeft5 = flipImageXAxis(SwordDogRunRight5)
	SwordDogRunLeft6 = flipImageXAxis(SwordDogRunRight6)

	SwordDogBigSlashRight1 = imageFromBytes(swordDogBigSlashRight1Png)
	SwordDogBigSlashRight2 = imageFromBytes(swordDogBigSlashRight2Png)
	SwordDogBigSlashRight3 = imageFromBytes(swordDogBigSlashRight3Png)
	SwordDogBigSlashRight4 = imageFromBytes(swordDogBigSlashRight4Png)
	SwordDogBigSlashRight5 = imageFromBytes(swordDogBigSlashRight5Png)
	SwordDogBigSlashRight6 = imageFromBytes(swordDogBigSlashRight6Png)

	SwordDogBigSlashLeft1 = flipImageXAxis(SwordDogBigSlashRight1)
	SwordDogBigSlashLeft2 = flipImageXAxis(SwordDogBigSlashRight2)
	SwordDogBigSlashLeft3 = flipImageXAxis(SwordDogBigSlashRight3)
	SwordDogBigSlashLeft4 = flipImageXAxis(SwordDogBigSlashRight4)
	SwordDogBigSlashLeft5 = flipImageXAxis(SwordDogBigSlashRight5)
	SwordDogBigSlashLeft6 = flipImageXAxis(SwordDogBigSlashRight6)
}
