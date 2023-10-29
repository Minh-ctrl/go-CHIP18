package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"os"

	keyboard "github.com/Minh-ctrl/go-CHIP18.git/keyboard"
	"github.com/Minh-ctrl/go-CHIP18.git/monitor"
	chip8struct "github.com/Minh-ctrl/go-CHIP18.git/struct"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = monitor.Columns * 20
	screenHeight = monitor.Rows * 20
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
	// g.keys = inpututil.AppendPressedKeys(g.keys[:0]) //only call this in update function

	opcode := (uint16(chip8.Memory[chip8.PC]) << 8) | uint16(chip8.Memory[chip8.PC+1])
	intepret(opcode)

	return nil
}

var chip8 chip8struct.Chip8

// stack push and pop implementation

func init() {
	// init values
	ebiten.SetWindowSize(screenWidth, screenHeight)
	chip8.PC = 0x200
	chip8.Stack = make([]uint16, 16)
	chip8.IndexRegister = 0
	dat, err := os.ReadFile("IBM Logo.ch8")
	check(err)
	var uintData = []uint8(dat)
	// load program
	for index, value := range uintData {
		chip8.Memory[0x200+index] = value
	}

}
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
	// all values
	nnn := instruction & 0xFFF
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
			// RET
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
		// SE Vx, Vy
		if chip8.Vx[x] == chip8.Vx[y] {
			chip8.PC += 2
		}
	case 0x6000:
		// LD Vx, byte
		chip8.Vx[x] = (instruction & 0xFF)
	case 0x7000:
		// ADD Vx, byte
		chip8.Vx[x] = chip8.Vx[x] + (instruction & 0xFF)
	case 0x8000:
		switch instruction & 0xF {

		case 0x0:
			// LD Vx, Vy
			chip8.Vx[x] = chip8.Vx[y]
		case 0x1:
			// OR Vx, Vy
			chip8.Vx[x] = chip8.Vx[x] | chip8.Vx[y]
		case 0x2:
			// AND Vx, Vy
			chip8.Vx[x] = chip8.Vx[x] & chip8.Vx[y]
		case 0x3:
			// XOR Vx, Vy
			chip8.Vx[x] = chip8.Vx[x] ^ chip8.Vx[y]
		case 0x4:
			// ADD Vx, Vy
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
			//  SHR Vx {, Vy} this one is interesting for different implementation :think:
			chip8.Vx[0xF] = chip8.Vx[x] & 0x1
			chip8.Vx[x] >>= 1
		case 0x7:
			// SUBN Vx, Vy
			chip8.Vx[0xF] = 0
			if chip8.Vx[y] > chip8.Vx[x] {
				chip8.Vx[0xF] = 1

			}
			chip8.Vx[x] = chip8.Vx[y] - chip8.Vx[x]
		case 0xE:
			//  SHL Vx {, Vy}
			chip8.Vx[x] <<= 1
		default:
			// throw error
		}
	case 0x9000:
		if chip8.Vx[x] != chip8.Vx[y] {
			chip8.PC += 2
		}

	case 0xA000:
		// LD I, addr
		chip8.IndexRegister = instruction & 0xFFF

	case 0xB000:
		// JP V0, addr
		chip8.PC = nnn + chip8.Vx[0]

	case 0xC000:
		// RND Vx, byte
		randomValue := uint16(rand.Intn(256))
		chip8.Vx[x] = randomValue & kk

	case 0xD000:
		//  DRW Vx, Vy, nibble

		width := uint16(8)
		height := (instruction & 0xF)
		chip8.Vx[0xF] = 0

		for row := uint16(0); row < height; row++ {
			sprite := chip8.Memory[chip8.IndexRegister+row]
			for col := uint16(0); col < width; col++ {
				if (sprite & 0x80) > 0 {
					if setPixel(int(chip8.Vx[x]+col), int(chip8.Vx[y]+row)) {
						chip8.Vx[0xF] = 1
					}
				}
				sprite <<= 1
			}
		}

	case 0xE000:
		// keyboard
		switch instruction & kk {
		case 0x9E:
			if keyboard.KeyListener(chip8.Vx[x]) {
				chip8.PC += 2
			}
		case 0xA1:
			if !keyboard.KeyListener(chip8.Vx[x]) {
				chip8.PC += 2
			}
		}

	case 0xF000:
		switch instruction & kk {
		case 0x07:
			// - LD Vx, DT
			chip8.Vx[x] = uint16(chip8.Delay_timer)
		case 0x0A:
			// - LD Vx, K
			// Wait for a key press, store the value of the key in Vx.

			// All execution stops until a key is pressed, then the value of that key is stored in Vx.

			chip8.Pause = true
			// bit hacky, hopefully this works
			for key := range keyboard.KeyBoardMaps {
				if inpututil.IsKeyJustPressed(keyboard.KeyBoardMaps[key]) {
					chip8.Pause = false
					chip8.Vx[x] = key
				}
			}

			// if
		case 0x15:
			// - LD DT, Vx
			chip8.Delay_timer = chip8.Vx[x]
		case 0x18:
			// - LD ST, Vx
			chip8.Sound_timer = chip8.Vx[x]
		case 0x1E:
			// - ADD I, Vx
			chip8.IndexRegister += chip8.Vx[x]
		case 0x29:
			// - LD F, Vx
			chip8.IndexRegister = chip8.Vx[x]
		case 0x33:
			// -/LD B, Vx
			chip8.Memory[chip8.IndexRegister] = uint8(chip8.Vx[x] / 100)
			chip8.Memory[chip8.IndexRegister+1] = uint8((chip8.Vx[x] % 100) / 10)
			chip8.Memory[chip8.IndexRegister+2] = uint8(chip8.Vx[x] % 10)

		case 0x55:
			// - LD [I], Vx
			for i := uint8(0); i < uint8(x); i++ {
				chip8.Memory[uint8(chip8.IndexRegister)+i] = uint8(chip8.Vx[i])
			}

		case 0x65:
			// - LD Vx, [I]
			for i := uint16(0); i < uint16(x); i++ {
				chip8.Vx[i] = uint16(chip8.Memory[chip8.IndexRegister+i])
			}
		}
	}

}

// functions for displaying monitor
func setPixel(x int, y int) (result bool) {
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
	chip8.Framebuffer[displayIndex] ^= 1
	return chip8.Framebuffer[displayIndex] != 1
}

func clearFrame() {
	// because i'm dumb
	for i := range chip8.Framebuffer {
		chip8.Framebuffer[i] = 0
	}
}
func paint(screen *ebiten.Image) {
	for i := 0; i < monitor.Columns*monitor.Rows; i++ {
		var x = (i % monitor.Columns) * monitor.Scale
		var y = math.Floor(float64(i)/monitor.Columns) * monitor.Scale

		if chip8.Framebuffer[i] == 1 {
			vector.DrawFilledRect(screen, float32(x), float32(y), rectangleW, rectangleH, color.White, false)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(frameWidth)*monitor.Scale, float64(frameHeight)*monitor.Scale)
	op.GeoM.Translate(screenHeight, screenWidth)

	screen.Fill(color.Black)

	paint(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
