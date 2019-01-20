package game

import (
	"time"

	"github.com/kyeett/gomponents/components"
)

func (g *Game) postStep() {

	for _, e := range g.filteredEntities(components.FollowingType, components.PosType) {
		following := g.entities.GetUnsafe(e, components.FollowingType).(*components.Following)

		pos := g.Pos(e)

		// Set position to whatever is being followed
		pos.Vec = g.Pos(following.ID).Vec.Add(following.Offset)
	}

	diffTime := time.Since(currentTime)
	currentTime = time.Now()
	for _, e := range g.filteredEntities(components.AnimatedType) {
		a := g.entities.GetUnsafe(e, components.AnimatedType).(*components.Animated)

		// Update animation time
		a.Ase.Update(float32(diffTime.Nanoseconds()) / 1000000000)
	}
}

var currentTime time.Time
