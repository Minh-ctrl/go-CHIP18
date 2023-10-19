package monitor

import (
	chip8 "github.com/Minh-ctrl/go-CHIP18.git/struct"
)

const (
	Columns = 64
	Rows    = 32
	Scale   = 15
)

var frame chip8.Chip8

func setPixel(x int, y int, frameBuffer [64 * 32]bool) {
	//
	if x > Columns {
		x -= Columns
	} else if x < 0 {
		x += Columns
	}
	if y > Rows {
		y -= Rows
	} else if y < 0 {
		y += Rows
	}
	var displayIndex = x + (y * Columns)
	// flip the value
	frameBuffer[displayIndex] = !frameBuffer[displayIndex]
	// cant return? need to go through golang again
}

func clear(frameBuffer [64 * 32]bool) {
	var newFrameBuffer = frame.Framebuffer
	// reset state
	frameBuffer = newFrameBuffer
}
