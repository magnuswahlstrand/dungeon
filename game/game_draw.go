package game

import (
	"fmt"
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
	// draw.Shadow(tmpImg, player, g.pts, g.staticSpace)

	// Draw hitboxes
	if debug {

		g.drawHitboxes(screen)
	}

	// Draw hook
	if rubberband {
		pos := g.Pos(hookID)
		g.drawArm(screen, pos.Vec, hook)

	} else {
		cursor := mousePosition()
		pos := g.Pos(hookID)
		aim := cursor.Sub(pos.Vec).Unit().Scaled(10).Add(pos.Vec)

		g.drawArm(screen, pos.Vec, aim)
		// draw.ResolvLine(screen, resolvutil.Line(pos.Vec, aim), color.RGBA{255, 255, 255, 100})
	}
	g.drawEntities(screen)

	// screen.DrawImage(tmpImg, &ebiten.DrawImageOptions{})
}

func (g *Game) drawArm(screen *ebiten.Image, start, end gfx.Vec) {

	for _, offset := range []struct {
		v gfx.Vec
		c color.Color
	}{
		{v: gfx.V(-1, 0), c: draw.GameColorsTransparent[3]},
		{v: gfx.V(0, 0), c: draw.GameColors[4]},
		{v: gfx.V(1, 0), c: draw.GameColorsTransparent[4]},
	} {
		draw.ResolvLine(screen, resolvutil.Line(start.Add(offset.v), end.Add(offset.v)), offset.c)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(end.X-2, end.Y-3)
	screen.DrawImage(g.spriteImg.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), op)
}

func (g *Game) drawHitboxes(screen *ebiten.Image) {

	for _, e := range g.filteredEntities(components.HitboxType) {
		hb := g.entities.GetUnsafe(e, components.HitboxType).(*components.Hitbox)
		offset := gfx.V(0, 0)

		if g.entities.HasComponents(e, components.PosType) {
			offset = g.Pos(e).Vec
		}

		draw.Rect(screen, hb.Rect.Moved(offset), colornames.Red)
	}
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
			fmt.Println(w, h, x, y)
			img = img.SubImage(image.Rect(int(x), int(y), int(x+w), int(y+h))).(*ebiten.Image)
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pos.X, pos.Y)
		screen.DrawImage(img, op)
	}

	if debug {

		for _, e := range g.filteredEntities(components.FollowingType, components.PosType) {
			// following := g.entities.GetUnsafe(e, components.FollowingType).(*components.Following)
			draw.Pt(screen, g.Pos(e).Vec, colornames.Burlywood)
		}
	}
}
