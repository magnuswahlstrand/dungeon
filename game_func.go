package main

import (
	"image"
	"image/color"
	"log"

	"github.com/peterhellberg/gfx"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/draw"
	"github.com/kyeett/dungeon/resolvutil"
	"github.com/kyeett/gomponents/components"
	"golang.org/x/image/colornames"
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

	img, err := gfx.OpenPNG("assets/animation/hero.png")
	if err != nil {
		log.Fatal(err)
	}

	playerImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func (g *Game) draw(screen *ebiten.Image) {
	screen.DrawImage(g.backgroundImg, &ebiten.DrawImageOptions{})
	tmpImg.Clear()
	// draw.Shadow(tmpImg, player, g.pts, g.staticSpace)

	// Draw hitboxes
	for _, e := range g.filteredEntities(components.HitboxType) {
		hb := g.entities.GetUnsafe(e, components.HitboxType).(*components.Hitbox)

		draw.Rect(screen, hb.Rect, colornames.Red)
	}

	// Draw player
	draw.Pt(screen, player, colornames.Dodgerblue)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(player.X-16, player.Y-16)
	screen.DrawImage(playerImg.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image), op)

	// Draw target
	target := cursor.Sub(player).Unit().Scaled(30).Add(player)
	draw.ResolvLine(screen, resolvutil.Line(player, target), color.RGBA{255, 255, 255, 100})

	// Draw hook
	if rubberband {
		draw.Pt(screen, hook, colornames.Dodgerblue)
		draw.ResolvLine(screen, resolvutil.Line(player, hook), color.RGBA{255, 255, 255, 100})
	}

	// screen.DrawImage(tmpImg, &ebiten.DrawImageOptions{})
}
