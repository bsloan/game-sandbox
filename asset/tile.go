package asset

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

const (
	EMPTY = iota
	GRASS_LEFT_EDGE
	GRASS_MIDDLE
	GRASS_RIGHT_EDGE
	DIRT_CENTER
	GRASS_SLOPE_R_START
	GRASS_SLOPE_R_BASE_1
	GRASS_SLOPE_R_MIDDLE
	GRASS_SLOPE_R_BASE_2
	GRASS_SLOPE_R_MAGIC_ROOT
	GRASS_SLOPE_L_START
	GRASS_SLOPE_L_BASE_1
	GRASS_SLOPE_L_MIDDLE
	GRASS_SLOPE_L_BASE_2
	GRASS_SLOPE_L_MAGIC_ROOT
	DIRT_CENTER_ROCKS_1
	DIRT_CENTER_ROCKS_2
	DIRT_L_1
	DIRT_L_2
	DIRT_R_1
	DIRT_R_2
	DIRT_BOTTOM_1
	DIRT_BOTTOM_2
)

var (
	//go:embed environment/grass_l.png
	grassLeftPng []byte

	//go:embed environment/grass_m.png
	grassMiddlePng []byte

	//go:embed environment/grass_r.png
	grassRightPng []byte

	//go:embed environment/dirt_m.png
	dirtMiddlePng []byte

	//go:embed environment/dirt_left_1.png
	dirtLeft1Png []byte

	//go:embed environment/dirt_left_2.png
	dirtLeft2Png []byte

	//go:embed environment/dirt_right_1.png
	dirtRight1Png []byte

	//go:embed environment/dirt_right_2.png
	dirtRight2Png []byte

	//go:embed environment/dirt_m_rocks_1.png
	dirtMiddleRocks1Png []byte

	//go:embed environment/dirt_m_rocks_2.png
	dirtMiddleRocks2Png []byte

	//go:embed environment/dirt_bot_1.png
	dirtBottom1Png []byte

	//go:embed environment/dirt_bot_2.png
	dirtBottom2Png []byte

	//go:embed environment/grass_slope_r1.png
	grassSlopeR1Png []byte

	//go:embed environment/grass_slope_r2.png
	grassSlopeR2Png []byte

	//go:embed environment/grass_slope_r3.png
	grassSlopeR3Png []byte

	//go:embed environment/grass_slope_r4.png
	grassSlopeR4Png []byte

	//go:embed environment/back.png
	skyBackgroundPng []byte

	//go:embed environment/middle.png
	hillsMidgroundPng []byte

	GrassLeft    *ebiten.Image
	GrassMiddle  *ebiten.Image
	GrassRight   *ebiten.Image
	DirtMiddle   *ebiten.Image
	DirtLeft1    *ebiten.Image
	DirtLeft2    *ebiten.Image
	DirtRight1   *ebiten.Image
	DirtRight2   *ebiten.Image
	DirtRocks1   *ebiten.Image
	DirtRocks2   *ebiten.Image
	DirtBottom1  *ebiten.Image
	DirtBottom2  *ebiten.Image
	GrassSlopeR1 *ebiten.Image
	GrassSlopeR2 *ebiten.Image
	GrassSlopeR3 *ebiten.Image
	GrassSlopeR4 *ebiten.Image
	GrassSlopeL1 *ebiten.Image
	GrassSlopeL2 *ebiten.Image
	GrassSlopeL3 *ebiten.Image
	GrassSlopeL4 *ebiten.Image

	SkyBackground  *ebiten.Image
	HillsMidground *ebiten.Image
	TileImages     = make(map[int]*ebiten.Image)
)

func LoadTiles() {
	GrassLeft = imageFromBytes(grassLeftPng)
	GrassMiddle = imageFromBytes(grassMiddlePng)
	GrassRight = imageFromBytes(grassRightPng)
	DirtMiddle = imageFromBytes(dirtMiddlePng)
	DirtLeft1 = imageFromBytes(dirtLeft1Png)
	DirtLeft2 = imageFromBytes(dirtLeft2Png)
	DirtRight1 = imageFromBytes(dirtRight1Png)
	DirtRight2 = imageFromBytes(dirtRight2Png)
	DirtBottom1 = imageFromBytes(dirtBottom1Png)
	DirtBottom2 = imageFromBytes(dirtBottom2Png)
	DirtRocks1 = imageFromBytes(dirtMiddleRocks1Png)
	DirtRocks2 = imageFromBytes(dirtMiddleRocks2Png)
	GrassSlopeR1 = imageFromBytes(grassSlopeR1Png)
	GrassSlopeR2 = imageFromBytes(grassSlopeR2Png)
	GrassSlopeR3 = imageFromBytes(grassSlopeR3Png)
	GrassSlopeR4 = imageFromBytes(grassSlopeR4Png)
	GrassSlopeL1 = flipImageXAxis(GrassSlopeR1)
	GrassSlopeL2 = flipImageXAxis(GrassSlopeR2)
	GrassSlopeL3 = flipImageXAxis(GrassSlopeR3)
	GrassSlopeL4 = flipImageXAxis(GrassSlopeR4)
	SkyBackground = imageFromBytes(skyBackgroundPng)
	HillsMidground = imageFromBytes(hillsMidgroundPng)

	// assign tile images to values
	TileImages[1] = GrassLeft
	TileImages[2] = GrassMiddle
	TileImages[3] = GrassRight
	TileImages[4] = DirtMiddle

	TileImages[5] = GrassSlopeR1
	TileImages[6] = GrassSlopeR2
	TileImages[7] = GrassSlopeR3
	TileImages[8] = GrassSlopeR4
	TileImages[9] = GrassSlopeR1 // Magic root tile for ~26 degree positive slope (grass)

	TileImages[10] = GrassSlopeL1
	TileImages[11] = GrassSlopeL2
	TileImages[12] = GrassSlopeL3
	TileImages[13] = GrassSlopeL4
	TileImages[14] = GrassSlopeL1 // Magic root tile for ~26 degree negative slope (grass)

	TileImages[DIRT_L_1] = DirtLeft1
	TileImages[DIRT_L_2] = DirtLeft2
	TileImages[DIRT_R_1] = DirtRight1
	TileImages[DIRT_R_2] = DirtRight2
	TileImages[DIRT_BOTTOM_1] = DirtBottom1
	TileImages[DIRT_BOTTOM_2] = DirtBottom2
	TileImages[DIRT_CENTER_ROCKS_1] = DirtRocks1
	TileImages[DIRT_CENTER_ROCKS_2] = DirtRocks2
}
