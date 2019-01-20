package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/game"
)

func main() {
	// create game
	g, err := game.New(game.OptionFromDisk)
	if err != nil {
		log.Fatal("Could not create game", err)
	}

	// start game
	if err := ebiten.Run(g.Update, g.Width()*16, g.Height()*16, 3, "Dungeon"); err != nil {
		log.Fatal(err)
	}
}
