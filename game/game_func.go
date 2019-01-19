package game

import (
	"log"

	"github.com/peterhellberg/gfx"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/assets"
	"github.com/kyeett/gomponents/components"
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

var tmpImg *ebiten.Image
var playerImg *ebiten.Image

func init() {
	tmpImg, _ = ebiten.NewImage(12*16, 12*16, ebiten.FilterDefault)

	path := "assets/animation/hero.png"
	img, err := gfx.DecodePNG(assets.FileReaderFatal(path))
	// img, err := gfx.OpenPNG(path)
	if err != nil {
		log.Fatal(err)
	}

	playerImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}
