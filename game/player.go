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
