package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math"

	monitor "github.com/Minh-ctrl/go-CHIP18.git/monitor"
	chip8 "github.com/Minh-ctrl/go-CHIP18.git/struct"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = monitor.Columns * monitor.Scale
	screenHeight = monitor.Rows * monitor.Scale
	rectangleW   = 1 * monitor.Scale
	rectangleH   = 1 * monitor.Scale
	frameOX      = 0
	frameOY      = 0
	frameWidth   = monitor.Columns
	frameHeight  = monitor.Rows
	frameCount   = 8
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++
	return nil
}

var frame chip8.Chip8

func setPixel(x int, y int) {
	//
	if x > monitor.Columns {
		x -= monitor.Columns
	} else if x < 0 {
		x += monitor.Columns
	}
	if y > monitor.Rows {
		y -= monitor.Rows
	} else if y < 0 {
		y += monitor.Rows
	}
	var displayIndex = x + (y * monitor.Columns)
	// flip the value
	frame.Framebuffer[displayIndex] = !frame.Framebuffer[displayIndex]
	// cant return? need to go through golang again
}

//	func testRender() {
//		setPixel(12, 32)
//		setPixel(2, 1)
//	}
func paint(screen *ebiten.Image) {
	for i := 0; i < monitor.Columns*monitor.Rows; i++ {
		var x = (i % monitor.Columns) * monitor.Scale
		var y = math.Floor(float64(i)/monitor.Columns) * monitor.Scale

		if frame.Framebuffer[i] {
			vector.DrawFilledRect(screen, float32(x), float32(y), rectangleW, rectangleH, color.White, false)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(frameWidth)*monitor.Scale, float64(frameHeight)*monitor.Scale)
	op.GeoM.Translate(screenHeight, screenWidth)
	// define the x and y

	screen.Fill(color.Black)
	// draw is being run over and over again
	setPixel(64, 2)
	setPixel(64, 3)

	paint(screen)
}

// func clear(frameBuffer[]bool) {
// 	// nah this doesnt work
// 	var newFrameBuffer = frame.Framebuffer
// 	// reset state
// 	// frameBuffer = newFrameBuffer
// }

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
