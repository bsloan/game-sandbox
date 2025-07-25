package asset

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
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

}
