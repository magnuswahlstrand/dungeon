package game

import (
	"github.com/hajimehoshi/ebiten"
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

func (g *Game) Width() int {
	return g.m.Width
}

func (g *Game) Height() int {
	return g.m.Height
}

var tmpImg *ebiten.Image
var playerImg *ebiten.Image

// func init() {
// 	tmpImg, _ = ebiten.NewImage(12*16, 12*16, ebiten.FilterDefault)

// 	path := "assets/animation/hero_animated.png"
// 	img, err := gfx.DecodePNG(assets.FileReaderFatal(path))
// 	// img, err := gfx.OpenPNG(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	playerImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
// }
