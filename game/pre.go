package game

import (
	"fmt"
	"time"

	"github.com/magnuswahlstrand/dungeon/audio"

	"github.com/kyeett/gomponents/direction"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/gomponents/components"
	"github.com/magnuswahlstrand/dungeon/resolvutil"
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
		slashID := UUID()
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

	// if ebiten.IsKeyPressed(ebiten.KeyW) {
	// 	if v.Y == 0 {
	// 		v.Y = -jumpSpeed
	// 	}
	// }

	d := g.entities.GetUnsafe(playerID, components.DirectedType).(*components.Directed)
	if ebiten.IsKeyPressed(ebiten.KeyD) || rightPadPressed() {
		v.X += accX
		d.D = direction.Right
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || leftPadPressed() {
		v.X -= accX
		d.D = direction.Left
	}

	if mouseJustPressed() {

		if gamePadPressed() {
			return
		}

		switch rubberband {
		case true:
			rubberband = false
		default:
			c := g.mousePositionCameraAdjusted()
			// Todo, clean this up
			rubberband = g.updateHook(c)
		}
	}

}

func gamePadPressed() bool {
	return leftPadPressed() || rightPadPressed()
}

func leftPadPressed() bool {
	p := mousePosition()
	return (p.X > 0 && p.X < 50) && (p.Y > 150 && p.Y < 190) && mousePressed()
}
func rightPadPressed() bool {
	p := mousePosition()
	return (p.X > 50 && p.X < 100) && (p.Y > 150 && p.Y < 190) && mousePressed()
}

// if len(inpututil.JustPressedTouchIDs()) > 0 {
// 	ID := inpututil.JustPressedTouchIDs()[0]
// 	x, y := ebiten.TouchPosition(ID)
// 	return gfx.V(float64(x), float64(y))
// }

func (g *Game) mousePositionCameraAdjusted() gfx.Vec {
	cr := g.getCameraPosition()
	c := mousePosition()
	return c.Add(gfx.V(float64(cr.Min.X), float64(cr.Min.Y)))
}

func (g *Game) preStep() {
	v := g.V(playerID)
	g.handleControls()

	// Gravity
	v.Y += gravity

	// Friction
	v.X = frictionX * v.X
	v.Y = frictionY * v.Y

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

	max := 4.0
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

func mouseJustPressed() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || len(inpututil.JustPressedTouchIDs()) > 0
}

func mousePressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || len(ebiten.TouchIDs()) > 0
}

func mousePosition() gfx.Vec {
	if len(ebiten.TouchIDs()) > 0 {
		ID := ebiten.TouchIDs()[0]
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
	pts := l.GetIntersectionPoints((&g.staticSpace).FilterByTags("hookable"))
	if len(pts) == 0 {
		// no points found
		fmt.Println("missed")
		audio.Play(audio.MissedSound)
		return false
	}

	hook = gfx.V(float64(pts[0].X)/collisionScaling, float64(pts[0].Y)/collisionScaling)

	audio.Play(audio.HookSound, 0.2)
	return true
}
