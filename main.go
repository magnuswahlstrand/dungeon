package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	// create game
	g, err := New()
	if err != nil {
		log.Fatal("Could not create game", err)
	}

	// start game
	if err := ebiten.Run(g.Update, 12*16, 12*16, 3, "Dungeon"); err != nil {
		log.Fatal(err)
	}
}
