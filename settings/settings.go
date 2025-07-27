package settings

const (
	TicksPerSecond             = 60
	ScreenWidth                = 320
	ScreenHeight               = 240
	TileSize                   = 16
	ScreenWidthTiles           = ScreenWidth / TileSize
	ScreenHeightTiles          = ScreenHeight / TileSize
	MidgroundScrollMultiplier  = 0.5
	Gravity                    = 450.0
	PlayerMaxVelocityX         = 100
	PlayerMaxVelocityY         = 300
	PlayerJumpVelocityLimit    = 100
	PlayerAccelerationStepX    = 1000
	PlayerJumpBoostHeight      = 30
	PlayerJumpInitialVelocity  = 6000
	PlayerJumpContinueVelocity = 400
)
