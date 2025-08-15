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

	GRASS_SLOPE_STEEP_R_BASE
	GRASS_SLOPE_STEEP_R_START
	GRASS_SLOPE_STEEP_R_MAGIC_ROOT
	GRASS_SLOPE_STEEP_L_BASE
	GRASS_SLOPE_STEEP_L_START
	GRASS_SLOPE_STEEP_L_MAGIC_ROOT

	GRASS_FLOAT_LEFT
	GRASS_FLOAT_MIDDLE
	GRASS_FLOAT_RIGHT

	PLANK_LEFT_BASE
	PLANK_MIDDLE
	PLANK_LEFT_END
	PLANK_RIGHT_BASE
	PLANK_RIGHT_END

	GRASS_FOREGROUND_1
	GRASS_FOREGROUND_2
	BRANCH_FOREGROUND_1
	BRANCH_FOREGROUND_2

	CAVE_BACKGROUND_L_1
	CAVE_BACKGROUND_L_2
	CAVE_BACKGROUND_L_3
	CAVE_BACKGROUND_L_4
	CAVE_BACKGROUND_L_5
	CAVE_BACKGROUND_L_6
	CAVE_BACKGROUND_L_7
	CAVE_BACKGROUND_R_1
	CAVE_BACKGROUND_R_2
	CAVE_BACKGROUND_R_3
	CAVE_BACKGROUND_R_4
	CAVE_BACKGROUND_R_5
	CAVE_BACKGROUND_R_6
	CAVE_BACKGROUND_R_7
	CAVE_BACKGROUND_ROCKS
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

	//go:embed environment/grass_float_l.png
	grassFloatLeftPng []byte

	//go:embed environment/grass_float_m.png
	grassFloatMiddlePng []byte

	//go:embed environment/grass_float_r.png
	grassFloatRightPng []byte

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

	//go:embed environment/grass_slope_steep_r1.png
	grassSlopeSteepR1Png []byte

	//go:embed environment/grass_slope_steep_r2.png
	grassSlopeSteepR2Png []byte

	//go:embed environment/grass_slope_steep_l1.png
	grassSlopeSteepL1Png []byte

	//go:embed environment/grass_slope_steep_l2.png
	grassSlopeSteepL2Png []byte

	//go:embed environment/plank_base_l.png
	plankBaseLPng []byte

	//go:embed environment/plank_base_r.png
	plankBaseRPng []byte

	//go:embed environment/plank_middle.png
	plankMiddlePng []byte

	//go:embed environment/plank_end_l.png
	plankEndLPng []byte

	//go:embed environment/plank_end_r.png
	plankEndRPng []byte

	//go:embed environment/grass_fore_1.png
	grassForeground1Png []byte

	//go:embed environment/grass_fore_2.png
	grassForeground2Png []byte

	//go:embed environment/branch_fore_1.png
	branchForeground1Png []byte

	//go:embed environment/branch_fore_2.png
	branchForeground2Png []byte

	//go:embed environment/cave_back_l1.png
	caveBackgroundL1Png []byte

	//go:embed environment/cave_back_l2.png
	caveBackgroundL2Png []byte

	//go:embed environment/cave_back_l3.png
	caveBackgroundL3Png []byte

	//go:embed environment/cave_back_l4.png
	caveBackgroundL4Png []byte

	//go:embed environment/cave_back_l5.png
	caveBackgroundL5Png []byte

	//go:embed environment/cave_back_l6.png
	caveBackgroundL6Png []byte

	//go:embed environment/cave_back_l7.png
	caveBackgroundL7Png []byte

	//go:embed environment/cave_rocks.png
	caveRocksPng []byte

	//go:embed environment/grass_float_m_cave.png
	grassFloatMiddleCavePng []byte

	//go:embed environment/cave_ceil_1.png
	caveCeil1Png []byte

	//go:embed environment/back.png
	skyBackgroundPng []byte

	//go:embed environment/middle.png
	hillsMidgroundPng []byte

	GrassLeft           *ebiten.Image
	GrassMiddle         *ebiten.Image
	GrassRight          *ebiten.Image
	GrassFloatLeft      *ebiten.Image
	GrassFloatMiddle    *ebiten.Image
	GrassFloatRight     *ebiten.Image
	DirtMiddle          *ebiten.Image
	DirtLeft1           *ebiten.Image
	DirtLeft2           *ebiten.Image
	DirtRight1          *ebiten.Image
	DirtRight2          *ebiten.Image
	DirtRocks1          *ebiten.Image
	DirtRocks2          *ebiten.Image
	DirtBottom1         *ebiten.Image
	DirtBottom2         *ebiten.Image
	GrassSlopeR1        *ebiten.Image
	GrassSlopeR2        *ebiten.Image
	GrassSlopeR3        *ebiten.Image
	GrassSlopeR4        *ebiten.Image
	GrassSlopeL1        *ebiten.Image
	GrassSlopeL2        *ebiten.Image
	GrassSlopeL3        *ebiten.Image
	GrassSlopeL4        *ebiten.Image
	GrassSlopeSteepR1   *ebiten.Image
	GrassSlopeSteepR2   *ebiten.Image
	GrassSlopeSteepL1   *ebiten.Image
	GrassSlopeSteepL2   *ebiten.Image
	PlankBaseL          *ebiten.Image
	PlankBaseR          *ebiten.Image
	PlankEndL           *ebiten.Image
	PlankEndR           *ebiten.Image
	PlankMiddle         *ebiten.Image
	CaveBackgroundL1    *ebiten.Image
	CaveBackgroundL2    *ebiten.Image
	CaveBackgroundL3    *ebiten.Image
	CaveBackgroundL4    *ebiten.Image
	CaveBackgroundL5    *ebiten.Image
	CaveBackgroundL6    *ebiten.Image
	CaveBackgroundL7    *ebiten.Image
	CaveBackgroundR1    *ebiten.Image
	CaveBackgroundR2    *ebiten.Image
	CaveBackgroundR3    *ebiten.Image
	CaveBackgroundR4    *ebiten.Image
	CaveBackgroundR5    *ebiten.Image
	CaveBackgroundR6    *ebiten.Image
	CaveBackgroundR7    *ebiten.Image
	CaveBackgroundRocks *ebiten.Image

	GrassForeground1  *ebiten.Image
	GrassForeground2  *ebiten.Image
	BranchForeground1 *ebiten.Image
	BranchForeground2 *ebiten.Image

	SkyBackground  *ebiten.Image
	HillsMidground *ebiten.Image
	TileImages     = make(map[int]*ebiten.Image)
)

func LoadTiles() {
	GrassLeft = imageFromBytes(grassLeftPng)
	GrassMiddle = imageFromBytes(grassMiddlePng)
	GrassRight = imageFromBytes(grassRightPng)
	GrassFloatLeft = imageFromBytes(grassFloatLeftPng)
	GrassFloatMiddle = imageFromBytes(grassFloatMiddlePng)
	GrassFloatRight = imageFromBytes(grassFloatRightPng)
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
	GrassSlopeSteepR1 = imageFromBytes(grassSlopeSteepR1Png)
	GrassSlopeSteepR2 = imageFromBytes(grassSlopeSteepR2Png)
	GrassSlopeSteepL1 = imageFromBytes(grassSlopeSteepL1Png)
	GrassSlopeSteepL2 = imageFromBytes(grassSlopeSteepL2Png)
	PlankBaseL = imageFromBytes(plankBaseLPng)
	PlankBaseR = imageFromBytes(plankBaseRPng)
	PlankEndL = imageFromBytes(plankEndLPng)
	PlankEndR = imageFromBytes(plankEndRPng)
	PlankMiddle = imageFromBytes(plankMiddlePng)

	GrassForeground1 = imageFromBytes(grassForeground1Png)
	GrassForeground2 = imageFromBytes(grassForeground2Png)
	BranchForeground1 = imageFromBytes(branchForeground1Png)
	BranchForeground2 = imageFromBytes(branchForeground2Png)

	SkyBackground = imageFromBytes(skyBackgroundPng)
	HillsMidground = imageFromBytes(hillsMidgroundPng)

	CaveBackgroundL1 = imageFromBytes(caveBackgroundL1Png)
	CaveBackgroundL2 = imageFromBytes(caveBackgroundL2Png)
	CaveBackgroundL3 = imageFromBytes(caveBackgroundL3Png)
	CaveBackgroundL4 = imageFromBytes(caveBackgroundL4Png)
	CaveBackgroundL5 = imageFromBytes(caveBackgroundL5Png)
	CaveBackgroundL6 = imageFromBytes(caveBackgroundL6Png)
	CaveBackgroundL7 = imageFromBytes(caveBackgroundL7Png)
	CaveBackgroundR1 = flipImageXAxis(CaveBackgroundL1)
	CaveBackgroundR2 = flipImageXAxis(CaveBackgroundL2)
	CaveBackgroundR3 = flipImageXAxis(CaveBackgroundL3)
	CaveBackgroundR4 = flipImageXAxis(CaveBackgroundL4)
	CaveBackgroundR5 = flipImageXAxis(CaveBackgroundL5)
	CaveBackgroundR6 = flipImageXAxis(CaveBackgroundL6)
	CaveBackgroundR7 = flipImageXAxis(CaveBackgroundL7)
	CaveBackgroundRocks = imageFromBytes(caveRocksPng)

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

	TileImages[15] = DirtRocks1
	TileImages[16] = DirtRocks2

	TileImages[17] = DirtLeft1
	TileImages[18] = DirtLeft2
	TileImages[19] = DirtRight1
	TileImages[20] = DirtRight2
	TileImages[21] = DirtBottom1
	TileImages[22] = DirtBottom2

	TileImages[23] = GrassSlopeSteepR1
	TileImages[24] = GrassSlopeSteepR2
	TileImages[25] = GrassSlopeSteepR2 // Magic root tile for steep grass positive slope
	TileImages[26] = GrassSlopeSteepL1
	TileImages[27] = GrassSlopeSteepL2
	TileImages[28] = GrassSlopeSteepL2 // Magic root tile for steep grass negative slope

	TileImages[29] = GrassFloatLeft
	TileImages[30] = GrassFloatMiddle
	TileImages[31] = GrassFloatRight

	TileImages[32] = PlankBaseL
	TileImages[33] = PlankMiddle
	TileImages[34] = PlankEndL
	TileImages[35] = PlankBaseR
	TileImages[36] = PlankEndR

	TileImages[37] = GrassForeground1
	TileImages[38] = GrassForeground2
	TileImages[39] = BranchForeground1
	TileImages[40] = BranchForeground2

	TileImages[41] = CaveBackgroundL1
	TileImages[42] = CaveBackgroundL2
	TileImages[43] = CaveBackgroundL3
	TileImages[44] = CaveBackgroundL4
	TileImages[45] = CaveBackgroundL5
	TileImages[46] = CaveBackgroundL6
	TileImages[47] = CaveBackgroundL7

	TileImages[48] = CaveBackgroundR1
	TileImages[49] = CaveBackgroundR2
	TileImages[50] = CaveBackgroundR3
	TileImages[51] = CaveBackgroundR4
	TileImages[52] = CaveBackgroundR5
	TileImages[53] = CaveBackgroundR6
	TileImages[54] = CaveBackgroundR7

	TileImages[55] = CaveBackgroundRocks
}
