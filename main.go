package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math"

	chip8struct "github.com/Minh-ctrl/go-CHIP18.git/struct"

	monitor "github.com/Minh-ctrl/go-CHIP18.git/monitor"
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

var chip8 chip8struct.Chip8

// stack push and pop implementation

func push(value uint16) {
	chip8.Stack = append(chip8.Stack, value)

}

func pop() {
	l := len(chip8.Stack)
	// chip8.PC = chip8.Stack[l-1] //assign location
	// chip8.Stack = chip8.Stack[:l-1] // pop
	chip8.PC, chip8.Stack = chip8.Stack[l-1], chip8.Stack[:l-1]
}

// instructions implementation

func intepret(instruction uint16) {
	chip8.PC += 2
	// bit shift to get x and y values

	kk := instruction & 0xFF
	x := (instruction & 0x0F00) >> 8
	y := (instruction & 0x00F0) >> 4

	switch line := instruction & 0xF000; line {
	case 0x0000:
		switch instruction {
		case 0x00E0:
			// CLS clear display
			clearFrame()
		case 0x00EE:
			// program counter to the address at the top of the stack, subtracts 1 from the stack pointer
			pop()
		}
	case 0x1000:
		//  JP addr
		// set program counter jump to nnn
		chip8.PC = instruction & 0xFFF
	case 0x2000:
		// CALL addr
		// call function (subroutine)
		// increment stack pointer then put PC on top of stack, then set pc to nnn
		push(chip8.PC)
		chip8.PC = instruction & 0xFFF
	case 0x3000:
		// 3xkk
		// SE Vx, byte
		if chip8.Vx[x] == kk {
			chip8.PC += 2
		}
	case 0x4000:
		// 4xkk
		// Vx, byte
		if chip8.Vx[x] != kk {
			chip8.PC += 2
		}
	case 0x5000:
		if chip8.Vx[x] == chip8.Vx[y] {
			chip8.PC += 2
		}
	case 0x6000:
		chip8.Vx[x] = kk
	case 0x7000:
		chip8.Vx[x] += chip8.Vx[y]
	case 0x8000:

		switch instruction & 0xF {
		case 0x0:
			chip8.Vx[x] = chip8.Vx[y]
		case 0x1:
			chip8.Vx[x] = chip8.Vx[x] | chip8.Vx[y]
		case 0x2:
			chip8.Vx[x] = chip8.Vx[x] & chip8.Vx[y]
		case 0x3:
			chip8.Vx[x] = chip8.Vx[x] ^ chip8.Vx[y]
		case 0x4:
			// chip8.Vx[x] = chip8.Vx[x] + chip8.Vx[y]
			sum := chip8.Vx[x] + chip8.Vx[y]
			if sum > 0xFF {
				chip8.Vx[0xF] = 1
				break
			}
			chip8.Vx[0xF] = 0

		case 0x5:
			// SUB Vx, Vy
			chip8.Vx[0xF] = 0
			if chip8.Vx[x] > chip8.Vx[y] {
				chip8.Vx[0xF] = 1
			}
			// Vy is subtracted from Vx, and the results stored in Vx.
			chip8.Vx[x] -= chip8.Vx[y]
		case 0x6:
			//  SHR Vx {, Vy}
			chip8.Vx[0xF] = chip8.Vx[x]
		}

	}

}

// functions for displaying monitor
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
	chip8.Framebuffer[displayIndex] = !chip8.Framebuffer[displayIndex]
}

func clearFrame() {
	// because i'm dumb
	for i := range chip8.Framebuffer {
		chip8.Framebuffer[i] = false
	}
}
func paint(screen *ebiten.Image) {
	for i := 0; i < monitor.Columns*monitor.Rows; i++ {
		var x = (i % monitor.Columns) * monitor.Scale
		var y = math.Floor(float64(i)/monitor.Columns) * monitor.Scale

		if chip8.Framebuffer[i] {
			vector.DrawFilledRect(screen, float32(x), float32(y), rectangleW, rectangleH, color.White, false)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(frameWidth)*monitor.Scale, float64(frameHeight)*monitor.Scale)
	op.GeoM.Translate(screenHeight, screenWidth)

	screen.Fill(color.Black)
	// draw is being run over and over again
	paint(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	chip8.PC = 0x200
	chip8.Stack = make([]uint16, 16)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
