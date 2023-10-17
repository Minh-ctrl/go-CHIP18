package main

import (
	"image/color"
	_ "image/png"
	"log"

	chip8 "github.com/Minh-ctrl/go-CHIP18.git/struct"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 320
	rectangleW   = 20
	rectangleH   = 20

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
	var frame chip8.Chip8
	// log.Println(frame.Framebuffer[2*3*64])
	// populate
	for i := 0; i < 64; i++ {

		for j := 0; j < 32; j++ {
			if i%2 == 0 && j%2 == 0 {
				frame.Framebuffer[i*j] = true
			}
			if frame.Framebuffer[j*i] {
				vector.DrawFilledRect(screen, float32(i)*20, float32(j)*20, rectangleW, rectangleH, color.Black, false)
			}
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
