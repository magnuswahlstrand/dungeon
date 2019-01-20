package game

import (
	"github.com/kyeett/dungeon/resolvutil"

	"github.com/kyeett/gomponents/components"
)

func (g *Game) movement() {

	pos := g.entities.GetUnsafe(playerID, components.PosType).(*components.Pos)
	v := g.entities.GetUnsafe(playerID, components.VelocityType).(*components.Velocity)
	hb := g.entities.GetUnsafe(playerID, components.HitboxType).(*components.Hitbox)

	r := resolvutil.ScaledRect(hb.Rect.Moved(pos.Vec), collisionScaling)

	if res := g.staticSpace.Resolve(r, 0, int32(collisionScaling*v.Y)); res.Colliding() && !res.Teleporting {
		v.Y = 0
	} else {
		if abs(v.Y) > 0.1 {
			pos.Y += v.Y
		}
	}

	if res := g.staticSpace.Resolve(r, int32(collisionScaling*v.X), 0); res.Colliding() && !res.Teleporting {
		v.X = -0.1 * v.X
	} else {
		if abs(v.X) > 0.1 {
			pos.X += v.X
		}
	}

	// Movement
	// pos.X += v.X
	// pos.Y += v.Y
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
