package hello

// translate chip 8 specs to instructions

type Chip8 struct {
	Framebuffer [64 * 32]bool //pixels with state on or off
}
