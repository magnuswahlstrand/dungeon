package game

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/resolvutil"
	"github.com/peterhellberg/gfx"
)

func (g *Game) preStep() {

	pos := g.Pos(playerID)
	v := g.V(playerID)

	// Gravity
	v.Y += gravity

	// Friction
	v.X = 0.98 * v.X
	v.Y = 0.96 * v.Y

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		pos.Vec.X++
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		pos.Vec.X--
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		v.Y += 0.1
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		v.Y = 0.1
	}

	// Apply rubber effect
	if rubberband {

		band := hook.Sub(g.Pos(hookID).Vec)
		fmt.Println(hook)
		fmt.Println(g.Pos(hookID).Vec)
		dist := band.Len()
		pw := (dist - 20) / 20
		if pw < 0 {
			pw = 0
		}

		if pw > 2*gravity {
			pw = 2 * gravity
		}

		band2 := band.Unit().Scaled(pw)

		v.X += band2.X
		v.Y += band2.Y
	}

	if mousePressed() {
		switch rubberband {
		case true:
			rubberband = false
		case false:
			c := mousePosition()
			rubberband = g.updateHook(c)
		}
	}
}

func mousePressed() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || len(inpututil.JustPressedTouchIDs()) > 0
}

func mousePosition() gfx.Vec {
	if len(inpututil.JustPressedTouchIDs()) > 0 {
		ID := inpututil.JustPressedTouchIDs()[0]
		x, y := ebiten.TouchPosition(ID)
		return gfx.V(float64(x), float64(y))
	}

	x, y := ebiten.CursorPosition()
	return gfx.V(float64(x), float64(y))
}

func (g *Game) updateHook(c gfx.Vec) bool {
	pos := g.Pos(hookID)
	target := c.Sub(pos.Vec).Unit().Scaled(1000).Add(pos.Vec)
	l := resolvutil.Line(pos.Vec, target)
	pts := l.IntersectionPoints(&g.staticSpace)
	if len(pts) == 0 {
		// no points found
		log.Fatal("yo")
		return false
	}

	hook = gfx.V(float64(pts[0].X), float64(pts[0].Y))
	return true
}
