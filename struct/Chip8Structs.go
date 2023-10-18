package chip18struct

// translate chip 8 specs to instructions

type Chip8 struct {
	memory   uint8     // 4096 bytes of memory
	register uint8     // v0 to vF
	u16Stack [16]uint8 // 16 bit stack

	I           uint16        // store memory address, quasi address pointer
	PC          uint16        // program counter, this is where we get the data?
	delay_reg   uint8         // delay timer
	sound_reg   uint8         // sound timer
	Framebuffer [64 * 32]bool //pixels with state on or off
}
