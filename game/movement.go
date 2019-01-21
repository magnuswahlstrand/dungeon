package game

import (
	"fmt"
	"time"

	"github.com/kyeett/dungeon/audio"
	"github.com/kyeett/dungeon/resolvutil"

	"github.com/kyeett/gomponents/components"
)

func (g *Game) movement() {

	g.handleMovement(playerID)
}

func (g *Game) handleMovement(ID string) {
	pos := g.entities.GetUnsafe(ID, components.PosType).(*components.Pos)
	v := g.entities.GetUnsafe(ID, components.VelocityType).(*components.Velocity)
	hb := g.entities.GetUnsafe(ID, components.HitboxType).(*components.Hitbox)

	r := resolvutil.ScaledRect(hb.Rect.Moved(pos.Vec), collisionScaling)

	if res := g.staticSpace.Resolve(r, 0, int32(collisionScaling*v.Y)); res.Colliding() && !res.Teleporting {
		v.Y = 0
		if res.ShapeB.HasTags("hazard") {
			g.handleDeath(ID)
			return
		}

	} else {

		if abs(v.Y) > 0.1 {
			pos.Y += v.Y
		}
	}

	if res := g.staticSpace.Resolve(r, int32(collisionScaling*v.X), 0); res.Colliding() && !res.Teleporting {
		v.X = -0.1 * v.X

		if res.ShapeB.HasTags("hazard") {
			g.handleDeath(ID)
			return
		}
	} else {
		if abs(v.X) > 0.1 {
			pos.X += v.X
		}
	}
}

func (g *Game) handleDeath(ID string) {
	if g.playerDead() || g.playerDying() {
		return
	}

	a := g.Animated(playerID)
	a.Ase.Play("Death")
	rubberband = false
	audio.Play(audio.DeathSound)

	// Add timer until reset
	fmt.Println("add timer")
	end := time.Now().Add(1 * time.Second)
	scenarioID := g.UUID()
	g.entities.Add(scenarioID, components.Scenario{
		F: func() bool {
			if time.Now().After(end) {
				g.Reset()
				return true
			}

			return false
		},
	})
	g.entityList = append(g.entityList, scenarioID)
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
