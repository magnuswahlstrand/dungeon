package game

import "github.com/kyeett/gomponents/components"

func (g *Game) postStep() {

	for _, e := range g.filteredEntities(components.FollowingType, components.PosType) {
		following := g.entities.GetUnsafe(e, components.FollowingType).(*components.Following)

		pos := g.Pos(e)

		// Set position to whatever is being followed
		pos.Vec = g.Pos(following.ID).Vec.Add(following.Offset)
	}
}
