package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/magnuswahlstrand/dungeon/game"
)

func main() {
	// create game
	rand.Seed(time.Now().UnixNano())

	g, err := game.New()
	if err != nil {
		log.Fatal("Could not create game", err)
	}

	// start game
	ebiten.SetFullscreen(true)
	if err := ebiten.Run(g.Update, 16*16, 12*16, 3, "Dungeon"); err != nil {
		log.Fatal(err)
	}
}
