package game

import (
	"github.com/kyeett/gomponents/components"
)

func (g *Game) Pos(e string) *components.Pos {
	return g.entities.GetUnsafe(e, components.PosType).(*components.Pos)
}

func (g *Game) V(e string) *components.Velocity {
	return g.entities.GetUnsafe(e, components.VelocityType).(*components.Velocity)
}

func (g *Game) Directed(e string) *components.Directed {
	return g.entities.GetUnsafe(e, components.DirectedType).(*components.Directed)
}

func (g *Game) Drawable(e string) *components.Drawable {
	return g.entities.GetUnsafe(e, components.DrawableType).(*components.Drawable)
}

func (g *Game) Animated(e string) *components.Animated {
	return g.entities.GetUnsafe(e, components.AnimatedType).(*components.Animated)
}
