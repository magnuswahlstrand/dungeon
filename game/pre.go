package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/resolvutil"
	"github.com/peterhellberg/gfx"
)

func (g *Game) preStep() {

	v := g.V(playerID)

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if v.Y == 0 {
			v.Y = -jumpSpeed
		}
	}

	// Gravity
	v.Y += gravity

	// Friction
	v.X = 0.93 * v.X
	v.Y = 0.96 * v.Y

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		v.X += accX
		if v.X > 2 {
			v.X = 2
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		v.X -= accX
		if v.X < -2 {
			v.X = -2
		}
	}

	// Apply rubber effect
	if rubberband {

		band := hook.Sub(g.Pos(hookID).Vec)
		dist := band.Len()
		pw := (dist - 15) / 50

		if pw < 0 {
			pw = 0
		}

		if pw > 3*gravity {
			pw = 3 * gravity
		}

		band2 := band.Unit().Scaled(pw)

		v.X += band2.X
		v.Y += band2.Y
	}

	max := 5.5
	if v.Len() > max {
		v.Vec = v.Unit().Scaled(max)
		fmt.Println("Maex")
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

func limit(v float64, mx float64) float64 {
	if v < 0 {
		return -min(mx, -v)
	}
	return min(mx, v)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
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
	target := c.Sub(pos.Vec).Unit().Scaled(200).Add(pos.Vec)
	l := resolvutil.ScaledLine(pos.Vec, target, collisionScaling)
	pts := l.IntersectionPoints(&g.staticSpace)
	fmt.Println(l)
	if len(pts) == 0 {
		// no points found
		fmt.Println("missed")
		return false
	}

	hook = gfx.V(float64(pts[0].X)/collisionScaling, float64(pts[0].Y)/collisionScaling)
	return true
}
