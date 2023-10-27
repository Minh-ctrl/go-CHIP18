package keyboard

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	KeyBoardMaps = map[uint16]ebiten.Key{
		0x1: ebiten.Key1, // 1 1
		0x2: ebiten.Key2, // 2 2
		0x3: ebiten.Key3, // 3 3
		0xC: ebiten.Key4, // 4 C
		0x4: ebiten.KeyQ, // Q 4
		0x5: ebiten.KeyW, // W 5
		0x6: ebiten.KeyE, // E 6
		0xD: ebiten.KeyR, // R D
		0x7: ebiten.KeyA, // A 7
		0x8: ebiten.KeyS, // S 8
		0x9: ebiten.KeyD, // D 9
		0xE: ebiten.KeyF, // F E
		0xA: ebiten.KeyZ, // Z A
		0x0: ebiten.KeyX, // X 0
		0xB: ebiten.KeyC, // C B
		0xF: ebiten.KeyV, // V F
	}
)

func KeyListener(test uint16) (result bool) {
	return inpututil.IsKeyJustPressed(KeyBoardMaps[test]) //listener here
}
