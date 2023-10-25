package chip8struct

// read cowgod's for all info
// Memory Map:
// +---------------+= 0xFFF (4095) End of Chip-8 RAM
// |               |
// |               |
// |               |
// |               |
// |               |
// | 0x200 to 0xFFF|
// |     Chip-8    |
// | Program / Data|
// |     Space     |
// |               |
// |               |
// |               |
// +- - - - - - - -+= 0x600 (1536) Start of ETI 660 Chip-8 programs
// |               |
// |               |
// |               |
// +---------------+= 0x200 (512) Start of most Chip-8 programs
// | 0x000 to 0x1FF|
// | Reserved for  |
// |  interpreter  |
// +---------------+= 0x000 (0) Start of Chip-8 RAM
type Chip8 struct {
	Memory   uint8   // 4096 bytes of memory
	Register []uint8 // v0 to vF
	PC       uint16  // program counter, this is where data starts. chip 8 is 0x200

	Stack         []uint16 // array of 16 16-bit values
	IndexRegister uint8

	Delay_timer uint8         // delay timer
	Sound_timer uint8         // sound timer
	Framebuffer [64 * 32]bool //pixels with state on or off, represented in Array

	Speed_cycle int

	Pause bool
}
