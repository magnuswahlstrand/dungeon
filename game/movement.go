package game

import (
	"github.com/kyeett/gomponents/components"
)

func (g *Game) movement() {

	pos := g.entities.GetUnsafe(playerID, components.PosType).(*components.Pos)
	v := g.entities.GetUnsafe(playerID, components.VelocityType).(*components.Velocity)
	_ = g.entities.GetUnsafe(playerID, components.HitboxType).(*components.Hitbox)

	// Movement
	pos.X += v.X
	pos.Y += v.Y
}
