package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/draw"
	"github.com/kyeett/dungeon/resolvutil"
	"github.com/kyeett/gomponents/components"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

func (g *Game) draw(screen *ebiten.Image) {
	screen.DrawImage(g.backgroundImg, &ebiten.DrawImageOptions{})
	tmpImg.Clear()
	// draw.Shadow(tmpImg, player, g.pts, g.staticSpace)

	// Draw hitboxes
	for _, e := range g.filteredEntities(components.HitboxType) {
		hb := g.entities.GetUnsafe(e, components.HitboxType).(*components.Hitbox)
		offset := gfx.V(0, 0)

		if g.entities.HasComponents(e, components.PosType) {
			offset = g.Pos(e).Vec
		}

		draw.Rect(screen, hb.Rect.Moved(offset), colornames.Red)
	}

	g.drawEntities(screen)

	// Draw target
	cursor := mousePosition()
	pos := g.Pos(hookID)
	aim := cursor.Sub(pos.Vec).Unit().Scaled(30).Add(pos.Vec)
	draw.ResolvLine(screen, resolvutil.Line(pos.Vec, aim), color.RGBA{255, 255, 255, 100})

	// Draw hook
	if rubberband {
		draw.Pt(screen, hook, colornames.Dodgerblue)
		target := hook
		draw.ResolvLine(screen, resolvutil.Line(pos.Vec, target), color.RGBA{255, 255, 255, 100})
	}

	// screen.DrawImage(tmpImg, &ebiten.DrawImageOptions{})
}

func (g *Game) drawEntities(screen *ebiten.Image) {
	// Draw entitie
	for _, e := range g.filteredEntities(components.DrawableType, components.PosType) {
		pos := g.entities.GetUnsafe(e, components.PosType).(*components.Pos)
		s := g.entities.GetUnsafe(e, components.DrawableType).(*components.Drawable)
		img := s.Image
		// If animated
		if g.entities.HasComponents(e, components.AnimatedType) {
			a := g.entities.GetUnsafe(e, components.AnimatedType).(*components.Animated)
			w, h := a.Ase.FrameWidth, a.Ase.FrameHeight
			x, y := a.Ase.GetFrameXY()
			img = img.SubImage(image.Rect(int(x), int(y), int(x+w), int(y+h))).(*ebiten.Image)
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pos.X, pos.Y)
		screen.DrawImage(img, op)
	}

	for _, e := range g.filteredEntities(components.FollowingType, components.PosType) {
		// following := g.entities.GetUnsafe(e, components.FollowingType).(*components.Following)
		draw.Pt(screen, g.Pos(e).Vec, colornames.Burlywood)
	}
}
