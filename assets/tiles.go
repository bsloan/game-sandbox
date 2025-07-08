package assets

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

	//go:embed environment/back.png
	skyBackgroundPng []byte

	//go:embed environment/middle.png
	hillsMidgroundPng []byte

	GrassLeft      *ebiten.Image
	GrassMiddle    *ebiten.Image
	GrassRight     *ebiten.Image
	DirtMiddle     *ebiten.Image
	SkyBackground  *ebiten.Image
	HillsMidground *ebiten.Image
	TileImages     = make(map[int]*ebiten.Image)
)

func LoadTiles() {
	GrassLeft = imageFromBytes(grassLeftPng)
	GrassMiddle = imageFromBytes(grassMiddlePng)
	GrassRight = imageFromBytes(grassRightPng)
	DirtMiddle = imageFromBytes(dirtMiddlePng)
	SkyBackground = imageFromBytes(skyBackgroundPng)
	HillsMidground = imageFromBytes(hillsMidgroundPng)
	TileImages[1] = GrassLeft
	TileImages[2] = GrassMiddle
	TileImages[3] = GrassRight
	TileImages[4] = DirtMiddle
}
