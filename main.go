package main

import (
	"image/color"
	_ "image/png"
	"log"

	monitor "github.com/Minh-ctrl/go-CHIP18.git/monitor"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = monitor.Columns * monitor.Scale
	screenHeight = monitor.Rows * monitor.Scale
	rectangleW   = 10 * monitor.Scale
	rectangleH   = 10 * monitor.Scale
	frameOX      = 0
	frameOY      = 0
	frameWidth   = 64
	frameHeight  = 32
	frameCount   = 8
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
	op.GeoM.Translate(float64(frameWidth), float64(frameHeight))
	op.GeoM.Translate(screenWidth, screenHeight)
	// define the x and y

	screen.Fill(color.Black)
	// coordinates
	// populate
	for i := 0; i < monitor.Rows; i++ {

		for j := 0; j < monitor.Columns; j++ {
			if i%2 == 0 {
				vector.DrawFilledRect(screen, float32(i)*monitor.Scale, float32(j)*monitor.Scale, rectangleW, rectangleH, color.White, false)

			}

			// if frame.Framebuffer[i] {
			// 	log.Println("placeholder")
			// }
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
