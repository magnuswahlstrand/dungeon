package game

import (
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/gomponents/components"
	tiled "github.com/lafriks/go-tiled"
	"github.com/peterhellberg/gfx"
)

type Game struct {
	entityList    []string
	entities      *components.Map
	Width, Height int
	baseDir       string
	m             *tiled.Map
	spriteImg     *ebiten.Image
	backgroundImg *ebiten.Image
	staticSpace   resolv.Space
	pts           []gfx.Vec
}

func (g *Game) Update(screen *ebiten.Image) error {

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		return gfx.ErrDone
	}

	// Pre-step
	g.preStep()

	// Movement
	g.movement()

	// Post-step
	g.postStep()

	// Draw
	g.draw(screen)

	// rend.SaveAsPng(img
	return nil
}
