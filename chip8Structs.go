package main

// translate chip 8 specs to instructions

type Chip8 struct {
	memory   uint8     // 4kiB
	register uint8     // v0 to vF
	stack    [16]uint8 // 16 bit stack

	I         uint16 // store memory address
	PC        uint16
	delay_reg uint8
	sound_reg uint8

	framebuffer [64 * 32]bool //pixels with state on or off
}
