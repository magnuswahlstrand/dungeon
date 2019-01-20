package game

import (
	"fmt"
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

	// Animated
	for _, e := range g.filteredEntities(components.AnimatedType) {
		fmt.Println("Updating", e)
		a := g.entities.GetUnsafe(e, components.AnimatedType).(*components.Animated)
		if e == playerID && a.Ase.IsPlaying("Slash") {
			if a.Ase.FinishedAnimation() {
				fmt.Println("Finished")
				a.Ase.Play("Stand")
			}
		}
		// Update animation time
		a.Ase.Update(float32(diffTime.Nanoseconds()) / 1000000000)
	}

	// Check timers
	for _, e := range g.filteredEntities(components.TimedType) {
		t := g.entities.GetUnsafe(e, components.TimedType).(*components.Timed)

		if t.Time.Sub(time.Now()) < 0 {
			// Remove entity
			g.removeEntity(e)
			fmt.Println("Remove")
		} else {
			fmt.Println("Not remove")
		}
	}
}

func (g *Game) removeEntity(e string) {
	g.entities.RemoveAll(e)
	var entities []string

	// Remove entity from list
	for _, s := range g.entityList {
		if s == e {
			continue
		}
		entities = append(entities, s)
	}
	g.entityList = entities
}

var currentTime time.Time
