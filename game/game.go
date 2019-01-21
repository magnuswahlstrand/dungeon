package game

import (
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/colornames"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/draw"
	"github.com/kyeett/gomponents/components"
	tiled "github.com/lafriks/go-tiled"
	"github.com/peterhellberg/gfx"
)

type Game struct {
	entityList    []string
	entities      *components.Map
	baseDir       string
	currentScene  string
	scenes        map[string]func(*Game, *ebiten.Image) error
	m             *tiled.Map
	spriteImg     *ebiten.Image
	backgroundImg *ebiten.Image
	staticSpace   resolv.Space
	pts           []gfx.Vec
}

func (g *Game) Update(screen *ebiten.Image) error {
	if !javascriptBuild && inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		return gfx.ErrDone
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
		return nil
	}

	return g.scenes[g.currentScene](g, screen)
}

func gameLoop(g *Game, screen *ebiten.Image) error {
	// Pre-step
	g.preStep()

	// Movement
	g.movement()

	// Post-step
	g.postStep()

	// Check for collision with triggers
	g.checkTriggers()

	// Draw
	camera.Clear()
	g.draw(camera)

	// Draw camera to screen
	cr := g.getCameraPosition()
	screen.DrawImage(camera.SubImage(cr).(*ebiten.Image), &ebiten.DrawImageOptions{})

	g.drawControl(screen)

	// gfx.SavePNG("map.png", screen)
	// return gfx.ErrDone
	return nil
}

func victoryScreen(g *Game, screen *ebiten.Image) error {
	screen.Fill(colornames.Black)
	draw.CenterText(screen, "Victory!", draw.FontFace11, colornames.White, -20)
	draw.CenterText(screen, "Press R/touch screen", draw.FontFace5, colornames.White, 20)
	draw.CenterText(screen, "to restart game", draw.FontFace5, colornames.White, 40)

	if mousePressed() {
		g.Reset()
		return nil
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return nil
}
