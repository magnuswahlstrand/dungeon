package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/resolvutil"
	"github.com/peterhellberg/gfx"
)

var Gravity = 0.1

func (g *Game) preStep() {

	// Gravity
	v.Y += Gravity

	// Friction
	v.X = 0.98 * v.X
	v.Y = 0.96 * v.Y

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.X++
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.X--
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		v.Y += 0.1
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		v.Y = 0.1
	}

	x, y := ebiten.CursorPosition()
	cursor.X = float64(x)
	cursor.Y = float64(y)

	// Apply rubber effect
	if rubberband {

		band := hook.Sub(player)
		dist := band.Len()
		pw := (dist - 20) / 20
		if pw < 0 {
			pw = 0
		}

		if pw > 2*Gravity {
			pw = 2 * Gravity
		}

		band2 := band.Unit().Scaled(pw)

		v.X += band2.X
		v.Y += band2.Y
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {

		switch rubberband {
		case true:

			rubberband = false
		case false:
			rubberband = g.updateHook(cursor)
		}
	}

	// Movement
	player.X += v.X
	player.Y += v.Y

}

func (g *Game) updateHook(c gfx.Vec) bool {
	target := cursor.Sub(player).Unit().Scaled(1000).Add(player)
	l := resolvutil.Line(player, target)
	pts := l.IntersectionPoints(&g.staticSpace)
	if len(pts) == 0 {
		// no points found
		log.Fatal("yo")
		return false
	}

	hook = gfx.V(float64(pts[0].X), float64(pts[0].Y))
	fmt.Println(hook)
	return true
}

var rubberband bool = true

var v = gfx.Vec{0, 0}
