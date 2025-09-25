package settings

const (
	TicksPerSecond                = 60
	ScreenWidth                   = 320
	ScreenHeight                  = 240
	TileSize                      = 16
	ScreenWidthTiles              = ScreenWidth / TileSize
	ScreenHeightTiles             = ScreenHeight / TileSize
	MidgroundScrollMultiplier     = 0.5
	Gravity                       = 450.0
	PlayerMaxVelocityX            = 100
	PlayerMaxRunningVelocityX     = 150
	PlayerMaxVelocityY            = 300
	PlayerJumpVelocityLimit       = 100
	PlayerAccelerationStep        = 1000
	PlayerRunningAccelerationStep = 1000
	PlayerJumpBoostHeight         = 30
	PlayerJumpInitialVelocity     = 6000
	PlayerJumpContinueVelocity    = 400
	PlayerClimbVelocity           = 600
	SwordDogAccelerationStep      = 500
	SwordDogMaxVelocityX          = 75
)
