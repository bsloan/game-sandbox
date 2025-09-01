package asset

import (
	"bytes"
	_ "embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	//go:embed general/title_screen_take1.png
	titleScreenPng []byte

	//go:embed general/BoldPixels1.4.ttf
	boldPixelsTtf []byte

	//go:embed general/map_take1.png
	levelSelectMapPng []byte

	//go:embed general/cursor_green.png
	levelSelectCursorGreenPng []byte

	//go:embed general/cursor_white.png
	levelSelectCursorWhitePng []byte

	//go:embed general/cursor_red.png
	levelSelectCursorRedPng []byte

	TitleScreen            *ebiten.Image
	LevelSelectMap         *ebiten.Image
	LevelSelectCursorGreen *ebiten.Image
	LevelSelectCursorWhite *ebiten.Image
	LevelSelectCursorRed   *ebiten.Image

	BoldPixelsFS *text.GoTextFaceSource
)

func LoadGeneralAssets() {
	// load static assets for title screen, level select map, etc
	TitleScreen = imageFromBytes(titleScreenPng)
	LevelSelectMap = imageFromBytes(levelSelectMapPng)
	LevelSelectCursorGreen = imageFromBytes(levelSelectCursorGreenPng)
	LevelSelectCursorWhite = imageFromBytes(levelSelectCursorWhitePng)
	LevelSelectCursorRed = imageFromBytes(levelSelectCursorRedPng)

	// load the font
	source, err := text.NewGoTextFaceSource(bytes.NewReader(boldPixelsTtf))
	if err != nil {
		log.Fatal(err)
	}
	BoldPixelsFS = source
}
