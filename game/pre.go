package game

import (
	"fmt"
	"time"

	"github.com/kyeett/gomponents/direction"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/resolvutil"
	"github.com/kyeett/gomponents/components"
	"github.com/peterhellberg/gfx"
)

func (g *Game) playerDead() bool {
	a := g.Animated(playerID)
	return a.Ase.CurrentAnimation.Name == "Dead"
}

func (g *Game) playerDying() bool {
	a := g.Animated(playerID)
	return a.Ase.CurrentAnimation.Name == "Death"
}

func (g *Game) handleControls() {
	if g.playerDead() {
		return
	}

	v := g.V(playerID)
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		// Slash
		slashID := g.UUID()
		g.entityList = append(g.entityList, slashID)
		g.entities.Add(slashID, components.Pos{Vec: g.Pos(playerID).Vec})
		g.entities.Add(slashID, components.Drawable{Image: slashImg})
		g.entities.Add(slashID, components.Animated{Ase: slashFile})
		g.entities.Add(slashID, components.Timed{Time: time.Now().Add(400 * time.Millisecond)})
		g.entities.Add(slashID, components.Following{ID: playerID, Offset: gfx.V(0, 0)})

		// Update animation
		a := g.entities.GetUnsafe(playerID, components.AnimatedType).(*components.Animated)
		a.Ase.Play("Slash")

	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if v.Y == 0 {
			v.Y = -jumpSpeed
		}
	}

	d := g.entities.GetUnsafe(playerID, components.DirectedType).(*components.Directed)
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		v.X += accX
		if v.X > 2 {
			v.X = 2
		}
		d.D = direction.Right
		fmt.Println("Turn!")
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		v.X -= accX
		if v.X < -2 {
			v.X = -2
		}
		d.D = direction.Left
		fmt.Println("Turn!")
	}

	if mousePressed() {
		switch rubberband {
		case true:
			rubberband = false
		case false:
			c := g.mousePositionCameraAdjusted()
			// Todo, clean this up
			rubberband = g.updateHook(c)
		}
	}
}

func (g *Game) mousePositionCameraAdjusted() gfx.Vec {
	cr := g.getCameraPosition()
	c := mousePosition()
	return c.Add(gfx.V(float64(cr.Min.X), float64(cr.Min.Y)))
}

func (g *Game) preStep() {
	v := g.V(playerID)

	// Gravity
	v.Y += gravity

	// Friction
	v.X = 0.93 * v.X
	v.Y = 0.96 * v.Y

	g.handleControls()

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
	pts := l.IntersectionPoints((&g.staticSpace).FilterByTags("hookable"))
	if len(pts) == 0 {
		// no points found
		fmt.Println("missed")
		return false
	}

	hook = gfx.V(float64(pts[0].X)/collisionScaling, float64(pts[0].Y)/collisionScaling)
	return true
}
