package game

import (
	"strings"

	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/colornames"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/gomponents/components"
	tiled "github.com/lafriks/go-tiled"
	"github.com/magnuswahlstrand/dungeon/draw"
	"github.com/peterhellberg/gfx"
)

func Version() string {
	return "0.6"
}

type Game struct {
	ID            string
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

var highscoreText string

func victoryScreen(g *Game, screen *ebiten.Image) error {
	screen.Fill(colornames.Black)
	draw.CenterText(screen, "Victory!", draw.FontFace11, colornames.White, -60)

	for i, v := range strings.Split(highscoreText, "\n") {
		text.Draw(screen, v, draw.FontFace5, 55, 70+i*10, colornames.White)
	}

	draw.CenterText(screen, "Press R/touch screen to restart", draw.FontFace5, colornames.White, 70)

	if mousePressed() {
		g.Reset()
		return nil
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return nil
}
