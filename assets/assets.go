package assets

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
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
)

var (
	GrassLeft   *ebiten.Image
	GrassMiddle *ebiten.Image
	GrassRight  *ebiten.Image
	DirtMiddle  *ebiten.Image
)

func imageFromBytes(pngBytes []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(pngBytes))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func Initialize() {
	GrassLeft = imageFromBytes(grassLeftPng)
	GrassMiddle = imageFromBytes(grassMiddlePng)
	GrassRight = imageFromBytes(grassRightPng)
	DirtMiddle = imageFromBytes(dirtMiddlePng)
}
