package keyboard

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	KeyBoardMaps = map[ebiten.Key]byte{
		ebiten.Key1: 0x1, // 1 1
		ebiten.Key2: 0x2, // 2 2
		ebiten.Key3: 0x3, // 3 3
		ebiten.Key4: 0xC, // 4 C
		ebiten.KeyQ: 0x4, // Q 4
		ebiten.KeyW: 0x5, // W 5
		ebiten.KeyE: 0x6, // E 6
		ebiten.KeyR: 0xD, // R D
		ebiten.KeyA: 0x7, // A 7
		ebiten.KeyS: 0x8, // S 8
		ebiten.KeyD: 0x9, // D 9
		ebiten.KeyF: 0xE, // F E
		ebiten.KeyZ: 0xA, // Z A
		ebiten.KeyX: 0x0, // X 0
		ebiten.KeyC: 0xB, // C B
		ebiten.KeyV: 0xF, // V F
	}
)

func KeyListener() {
	for key, value := range KeyBoardMaps {
		if inpututil.IsKeyJustPressed(key) { //listener here
			fmt.Println(key, value)
		}
	}
}
