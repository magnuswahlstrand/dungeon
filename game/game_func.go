package game

import (
	"bytes"
	"fmt"
	"log"
	"text/tabwriter"
	"time"

	"github.com/kyeett/dungeon/highscore"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/audio"
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
	for _, e := range g.entityList {
		g.entities.RemoveAll(e)
	}
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
			fmt.Println("triggered!", t.Scenario, t.Direction, e)

			switch {
			case t.Scenario == "victory":
				g.currentScene = "victory"
				audio.Play(audio.VictorySound)
				g.handleHighscore()
			}
		}
	}
}

func rectToShape(hb gfx.Rect) *resolv.Rectangle {
	return resolv.NewRectangle(int32(hb.Min.X), int32(hb.Min.Y), int32(hb.W()), int32(hb.H()))
}

func (g *Game) handleHighscore() {
	t := time.Since(startTime)

	highscore.SaveScore(UUID(), g.ID, t, Version())

	score, err := highscore.GetScore()
	if err != nil {
		log.Fatal("Could not retrieve highscore")
	}
	buf := bytes.NewBufferString(fmt.Sprintf("    Your time was %0.2fs\n\n", t.Seconds()))
	w := tabwriter.NewWriter(buf, 3, 0, 4, ' ', tabwriter.AlignRight)
	for k, v := range score.ByVersion(Version()).Limit(5) {
		fmt.Fprintf(w, "%d\t%s\t%0.2fs\t\n", k+1, v.UserID, float64(v.Time)/1000000000)
	}
	w.Flush()

	highscoreText = buf.String()
}
