package asset

import (
	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
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

	//go:embed sprite/sword_dog_downslash1.png
	swordDogDownSlashRight1Png []byte

	//go:embed sprite/sword_dog_downslash2.png
	swordDogDownSlashRight2Png []byte

	//go:embed sprite/sword_dog_downslash3.png
	swordDogDownSlashRight3Png []byte

	//go:embed sprite/sword_dog_downslash4.png
	swordDogDownSlashRight4Png []byte

	//go:embed sprite/sword_dog_downslash5.png
	swordDogDownSlashRight5Png []byte

	//go:embed sprite/alligator_idle1.png
	alligatorIdle1Png []byte

	//go:embed sprite/alligator_idle2.png
	alligatorIdle2Png []byte

	//go:embed sprite/alligator_idle3.png
	alligatorIdle3Png []byte

	//go:embed sprite/alligator_idle4.png
	alligatorIdle4Png []byte

	//go:embed sprite/alligator_run1.png
	alligatorRun1Png []byte

	//go:embed sprite/alligator_run2.png
	alligatorRun2Png []byte

	//go:embed sprite/alligator_run3.png
	alligatorRun3Png []byte

	//go:embed sprite/alligator_run4.png
	alligatorRun4Png []byte

	//go:embed sprite/alligator_run5.png
	alligatorRun5Png []byte

	//go:embed sprite/alligator_run6.png
	alligatorRun6Png []byte

	//go:embed sprite/alligator_run7.png
	alligatorRun7Png []byte

	//go:embed sprite/alligator_run8.png
	alligatorRun8Png []byte

	//go:embed sprite/alligator_run9.png
	alligatorRun9Png []byte

	//go:embed sprite/alligator_slash1.png
	alligatorSlash1Png []byte

	//go:embed sprite/alligator_slash2.png
	alligatorSlash2Png []byte

	//go:embed sprite/alligator_slash3.png
	alligatorSlash3Png []byte

	//go:embed sprite/alligator_slash4.png
	alligatorSlash4Png []byte

	//go:embed sprite/alligator_slash5.png
	alligatorSlash5Png []byte

	//go:embed sprite/alligator_slash6.png
	alligatorSlash6Png []byte

	//go:embed sprite/alligator_slash7.png
	alligatorSlash7Png []byte

	//go:embed sprite/alligator_slash8.png
	alligatorSlash8Png []byte

	//go:embed sprite/alligator_slash9.png
	alligatorSlash9Png []byte

	//go:embed sprite/alligator_slash10.png
	alligatorSlash10Png []byte

	//go:embed sprite/alligator_slash11.png
	alligatorSlash11Png []byte

	//go:embed sprite/frog_idle_1.png
	frogIdle1Png []byte

	//go:embed sprite/frog_idle_2.png
	frogIdle2Png []byte

	//go:embed sprite/frog_idle_3.png
	frogIdle3Png []byte

	//go:embed sprite/frog_idle_4.png
	frogIdle4Png []byte

	//go:embed sprite/frog_jump_1.png
	frogJump1Png []byte

	//go:embed sprite/frog_jump_2.png
	frogJump2Png []byte

	//go:embed sprite/enemy_death_1.png
	enemyDeath1Png []byte

	//go:embed sprite/enemy_death_2.png
	enemyDeath2Png []byte

	//go:embed sprite/enemy_death_3.png
	enemyDeath3Png []byte

	//go:embed sprite/enemy_death_4.png
	enemyDeath4Png []byte

	//go:embed sprite/enemy_death_5.png
	enemyDeath5Png []byte

	//go:embed sprite/enemy_death_6.png
	enemyDeath6Png []byte

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

	SwordDogDownSlashRight1 *ebiten.Image
	SwordDogDownSlashRight2 *ebiten.Image
	SwordDogDownSlashRight3 *ebiten.Image
	SwordDogDownSlashRight4 *ebiten.Image
	SwordDogDownSlashRight5 *ebiten.Image

	SwordDogDownSlashLeft1 *ebiten.Image
	SwordDogDownSlashLeft2 *ebiten.Image
	SwordDogDownSlashLeft3 *ebiten.Image
	SwordDogDownSlashLeft4 *ebiten.Image
	SwordDogDownSlashLeft5 *ebiten.Image

	AlligatorIdleRight1 *ebiten.Image
	AlligatorIdleRight2 *ebiten.Image
	AlligatorIdleRight3 *ebiten.Image
	AlligatorIdleRight4 *ebiten.Image

	AlligatorIdleLeft1 *ebiten.Image
	AlligatorIdleLeft2 *ebiten.Image
	AlligatorIdleLeft3 *ebiten.Image
	AlligatorIdleLeft4 *ebiten.Image

	AlligatorRunRight1 *ebiten.Image
	AlligatorRunRight2 *ebiten.Image
	AlligatorRunRight3 *ebiten.Image
	AlligatorRunRight4 *ebiten.Image
	AlligatorRunRight5 *ebiten.Image
	AlligatorRunRight6 *ebiten.Image
	AlligatorRunRight7 *ebiten.Image
	AlligatorRunRight8 *ebiten.Image
	AlligatorRunRight9 *ebiten.Image

	AlligatorRunLeft1 *ebiten.Image
	AlligatorRunLeft2 *ebiten.Image
	AlligatorRunLeft3 *ebiten.Image
	AlligatorRunLeft4 *ebiten.Image
	AlligatorRunLeft5 *ebiten.Image
	AlligatorRunLeft6 *ebiten.Image
	AlligatorRunLeft7 *ebiten.Image
	AlligatorRunLeft8 *ebiten.Image
	AlligatorRunLeft9 *ebiten.Image

	AlligatorSlashRight1  *ebiten.Image
	AlligatorSlashRight2  *ebiten.Image
	AlligatorSlashRight3  *ebiten.Image
	AlligatorSlashRight4  *ebiten.Image
	AlligatorSlashRight5  *ebiten.Image
	AlligatorSlashRight6  *ebiten.Image
	AlligatorSlashRight7  *ebiten.Image
	AlligatorSlashRight8  *ebiten.Image
	AlligatorSlashRight9  *ebiten.Image
	AlligatorSlashRight10 *ebiten.Image
	AlligatorSlashRight11 *ebiten.Image

	AlligatorSlashLeft1  *ebiten.Image
	AlligatorSlashLeft2  *ebiten.Image
	AlligatorSlashLeft3  *ebiten.Image
	AlligatorSlashLeft4  *ebiten.Image
	AlligatorSlashLeft5  *ebiten.Image
	AlligatorSlashLeft6  *ebiten.Image
	AlligatorSlashLeft7  *ebiten.Image
	AlligatorSlashLeft8  *ebiten.Image
	AlligatorSlashLeft9  *ebiten.Image
	AlligatorSlashLeft10 *ebiten.Image
	AlligatorSlashLeft11 *ebiten.Image

	FrogIdleLeft1 *ebiten.Image
	FrogIdleLeft2 *ebiten.Image
	FrogIdleLeft3 *ebiten.Image
	FrogIdleLeft4 *ebiten.Image

	FrogIdleRight1 *ebiten.Image
	FrogIdleRight2 *ebiten.Image
	FrogIdleRight3 *ebiten.Image
	FrogIdleRight4 *ebiten.Image

	FrogJumpLeft1 *ebiten.Image
	FrogJumpLeft2 *ebiten.Image

	FrogJumpRight1 *ebiten.Image
	FrogJumpRight2 *ebiten.Image

	EnemyDeath1 *ebiten.Image
	EnemyDeath2 *ebiten.Image
	EnemyDeath3 *ebiten.Image
	EnemyDeath4 *ebiten.Image
	EnemyDeath5 *ebiten.Image
	EnemyDeath6 *ebiten.Image
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

	SwordDogDownSlashRight1 = imageFromBytes(swordDogDownSlashRight1Png)
	SwordDogDownSlashRight2 = imageFromBytes(swordDogDownSlashRight2Png)
	SwordDogDownSlashRight3 = imageFromBytes(swordDogDownSlashRight3Png)
	SwordDogDownSlashRight4 = imageFromBytes(swordDogDownSlashRight4Png)
	SwordDogDownSlashRight5 = imageFromBytes(swordDogDownSlashRight5Png)

	SwordDogDownSlashLeft1 = flipImageXAxis(SwordDogDownSlashRight1)
	SwordDogDownSlashLeft2 = flipImageXAxis(SwordDogDownSlashRight2)
	SwordDogDownSlashLeft3 = flipImageXAxis(SwordDogDownSlashRight3)
	SwordDogDownSlashLeft4 = flipImageXAxis(SwordDogDownSlashRight4)
	SwordDogDownSlashLeft5 = flipImageXAxis(SwordDogDownSlashRight5)

	AlligatorIdleRight1 = imageFromBytes(alligatorIdle1Png)
	AlligatorIdleRight2 = imageFromBytes(alligatorIdle2Png)
	AlligatorIdleRight3 = imageFromBytes(alligatorIdle3Png)
	AlligatorIdleRight4 = imageFromBytes(alligatorIdle4Png)

	AlligatorIdleLeft1 = flipImageXAxis(AlligatorIdleRight1)
	AlligatorIdleLeft2 = flipImageXAxis(AlligatorIdleRight2)
	AlligatorIdleLeft3 = flipImageXAxis(AlligatorIdleRight3)
	AlligatorIdleLeft4 = flipImageXAxis(AlligatorIdleRight4)

	AlligatorRunRight1 = imageFromBytes(alligatorRun1Png)
	AlligatorRunRight2 = imageFromBytes(alligatorRun2Png)
	AlligatorRunRight3 = imageFromBytes(alligatorRun3Png)
	AlligatorRunRight4 = imageFromBytes(alligatorRun4Png)
	AlligatorRunRight5 = imageFromBytes(alligatorRun5Png)
	AlligatorRunRight6 = imageFromBytes(alligatorRun6Png)
	AlligatorRunRight7 = imageFromBytes(alligatorRun7Png)
	AlligatorRunRight8 = imageFromBytes(alligatorRun8Png)
	AlligatorRunRight9 = imageFromBytes(alligatorRun9Png)

	AlligatorRunLeft1 = flipImageXAxis(AlligatorRunRight1)
	AlligatorRunLeft2 = flipImageXAxis(AlligatorRunRight2)
	AlligatorRunLeft3 = flipImageXAxis(AlligatorRunRight3)
	AlligatorRunLeft4 = flipImageXAxis(AlligatorRunRight4)
	AlligatorRunLeft5 = flipImageXAxis(AlligatorRunRight5)
	AlligatorRunLeft6 = flipImageXAxis(AlligatorRunRight6)
	AlligatorRunLeft7 = flipImageXAxis(AlligatorRunRight7)
	AlligatorRunLeft8 = flipImageXAxis(AlligatorRunRight8)
	AlligatorRunLeft9 = flipImageXAxis(AlligatorRunRight9)

	AlligatorSlashRight1 = imageFromBytes(alligatorSlash1Png)
	AlligatorSlashRight2 = imageFromBytes(alligatorSlash2Png)
	AlligatorSlashRight3 = imageFromBytes(alligatorSlash3Png)
	AlligatorSlashRight4 = imageFromBytes(alligatorSlash4Png)
	AlligatorSlashRight5 = imageFromBytes(alligatorSlash5Png)
	AlligatorSlashRight6 = imageFromBytes(alligatorSlash6Png)
	AlligatorSlashRight7 = imageFromBytes(alligatorSlash7Png)
	AlligatorSlashRight8 = imageFromBytes(alligatorSlash8Png)
	AlligatorSlashRight9 = imageFromBytes(alligatorSlash9Png)
	AlligatorSlashRight10 = imageFromBytes(alligatorSlash10Png)
	AlligatorSlashRight11 = imageFromBytes(alligatorSlash11Png)

	AlligatorSlashLeft1 = flipImageXAxis(AlligatorSlashRight1)
	AlligatorSlashLeft2 = flipImageXAxis(AlligatorSlashRight2)
	AlligatorSlashLeft3 = flipImageXAxis(AlligatorSlashRight3)
	AlligatorSlashLeft4 = flipImageXAxis(AlligatorSlashRight4)
	AlligatorSlashLeft5 = flipImageXAxis(AlligatorSlashRight5)
	AlligatorSlashLeft6 = flipImageXAxis(AlligatorSlashRight6)
	AlligatorSlashLeft7 = flipImageXAxis(AlligatorSlashRight7)
	AlligatorSlashLeft8 = flipImageXAxis(AlligatorSlashRight8)
	AlligatorSlashLeft9 = flipImageXAxis(AlligatorSlashRight9)
	AlligatorSlashLeft10 = flipImageXAxis(AlligatorSlashRight10)
	AlligatorSlashLeft11 = flipImageXAxis(AlligatorSlashRight11)

	FrogIdleLeft1 = imageFromBytes(frogIdle1Png)
	FrogIdleLeft2 = imageFromBytes(frogIdle2Png)
	FrogIdleLeft3 = imageFromBytes(frogIdle3Png)
	FrogIdleLeft4 = imageFromBytes(frogIdle4Png)

	FrogIdleRight1 = flipImageXAxis(FrogIdleLeft1)
	FrogIdleRight2 = flipImageXAxis(FrogIdleLeft2)
	FrogIdleRight3 = flipImageXAxis(FrogIdleLeft3)
	FrogIdleRight4 = flipImageXAxis(FrogIdleLeft4)

	FrogJumpLeft1 = imageFromBytes(frogJump1Png)
	FrogJumpLeft2 = imageFromBytes(frogJump2Png)

	FrogJumpRight1 = flipImageXAxis(FrogJumpLeft1)
	FrogJumpRight2 = flipImageXAxis(FrogJumpLeft2)

	EnemyDeath1 = imageFromBytes(enemyDeath1Png)
	EnemyDeath2 = imageFromBytes(enemyDeath2Png)
	EnemyDeath3 = imageFromBytes(enemyDeath3Png)
	EnemyDeath4 = imageFromBytes(enemyDeath4Png)
	EnemyDeath5 = imageFromBytes(enemyDeath5Png)
	EnemyDeath6 = imageFromBytes(enemyDeath6Png)
}
