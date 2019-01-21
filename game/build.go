// +build !js

package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var javascriptBuild = false

func quitButtonPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyTab)
}
