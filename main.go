package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 320
	rectangleW   = 10
	rectangleH   = 10

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	// define the x and y

	screen.Fill(color.White)
	// coordinates
	for i := 0; i < 64; i++ {
		rectangleX := (screenWidth - rectangleW)

		for j := 0; j < 32; j++ {
			rectangleY := (screenHeight - rectangleH)

			vector.DrawFilledRect(screen, float32(rectangleX), float32(rectangleY), rectangleW, rectangleH, color.Black, false)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Decode an image from the image file's byte slice.

	// chip8 := Chip8{
	// 	4096,
	// 	16,
	// }
	// framebuffer := [64*32]bool

	// img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
