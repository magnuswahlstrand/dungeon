package game

import (
	"fmt"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/gomponents/components"
	"github.com/peterhellberg/gfx"
)

func (g *Game) filteredEntities(types ...components.Type) []string {
	var IDs []string
	for _, ID := range g.entityList {
		if g.entities.HasComponents(ID, types...) {
			IDs = append(IDs, ID)
		}
	}
	return IDs
}

func (g *Game) Width() int {
	return g.m.Width
}

func (g *Game) Height() int {
	return g.m.Height
}

var tmpImg *ebiten.Image
var playerImg *ebiten.Image
var camera *ebiten.Image

func (g *Game) Reset() {
	g.currentScene = "game"
	rubberband = false
	g.removeEntity(playerID)
	g.newPlayer()
	g.initMap()
}

func (g *Game) checkTriggers() {
	pos := g.entities.GetUnsafe(playerID, components.PosType).(*components.Pos)
	hb := g.entities.GetUnsafe(playerID, components.HitboxType).(*components.Hitbox)

	playerShape := rectToShape(hb.Moved(pos.Vec))

	for _, e := range g.filteredEntities(components.TriggerType) {
		t := g.entities.GetUnsafe(e, components.TriggerType).(*components.Trigger)

		tRect := rectToShape(t.Rect)
		if playerShape.WouldBeColliding(tRect, 0, 0) {
			fmt.Println("triggered!", t.Scenario)

			switch {
			case t.Scenario == "victory":
				g.currentScene = "victory"
			}
		}
	}
}

func rectToShape(hb gfx.Rect) *resolv.Rectangle {
	return resolv.NewRectangle(int32(hb.Min.X), int32(hb.Min.Y), int32(hb.W()), int32(hb.H()))
}
