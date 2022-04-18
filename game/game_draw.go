package game

import (
	"image"
	"image/color"
	"log"

	"github.com/kyeett/gomponents/direction"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/gomponents/components"
	"github.com/magnuswahlstrand/dungeon/assets"
	"github.com/magnuswahlstrand/dungeon/draw"
	"github.com/magnuswahlstrand/dungeon/resolvutil"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

func initMobileControls() {
	// controlLeftImg, _ := ebiten.NewImage(100, 100, ebiten.FilterDefault)

	path := "assets/animation/button_right.png"
	img, err := gfx.DecodePNG(assets.FileReaderFatal(path))
	if err != nil {
		log.Fatal(err)
	}

	controlRightImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	controlLeftImg, _ = ebiten.NewImage(img.Bounds().Dx(), img.Bounds().Dy(), ebiten.FilterDefault)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(img.Bounds().Dx())/2, 0)
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(+float64(img.Bounds().Dx())/2, 0)
	controlLeftImg.DrawImage(controlRightImg, op)

	// w := 13.0
	// pts := gfx.Polygon{
	// 	gfx.V(0, 1*w),
	// 	gfx.V(w, 0),
	// 	gfx.V(2*w, 0),
	// 	gfx.V(2*w, 3*w),
	// 	gfx.V(w, 3*w),
	// 	gfx.V(0, 2*w),
	// }

	// // offset := gfx.V(10, 140)

	// var poly gfx.Polygon
	// for _, p := range pts {
	// 	poly = append(poly, p)
	// }

	// b := poly.Bounds()
	// draw.Rect(controlLeftImg, gfx.R(float64(b.Min.X), float64(b.Min.Y), float64(b.Dx()), float64(b.Dy())), colornames.Red)
	// gfx.DrawPolygon(controlLeftImg, poly, 0, draw.GameColors[5])
	// gfx.DrawPolygon(controlLeftImg, poly, 1, color.Black)

	// poly = []gfx.Vec{}
	// for _, p := range pts {
	// 	poly = append(poly, gfx.V(50, 3*w).Sub(p)) //.Add(gfx.V(10, 140)))
	// }

	// b = poly.Bounds() //.Add(image.Pt(10, 0))
	// draw.Rect(controlRightImg, gfx.R(float64(b.Min.X), float64(b.Min.Y), float64(b.Dx()), float64(b.Dy())), colornames.Red)
	// gfx.DrawPolygon(controlRightImg, poly, 0, draw.GameColors[5])
	// gfx.DrawPolygon(controlRightImg, poly, 1, color.Black)
}

var controlLeftImg *ebiten.Image
var controlRightImg *ebiten.Image

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

	} else if !g.playerDead() && !g.playerDying() {
		// cursor := mousePosition()
		cursor := g.mousePositionCameraAdjusted()
		pos := g.Pos(hookID)
		aim := cursor.Sub(pos.Vec).Unit().Scaled(10).Add(pos.Vec)

		g.drawArm(screen, pos.Vec, aim)
		// draw.ResolvLine(screen, resolvutil.Line(pos.Vec, aim), color.RGBA{255, 255, 255, 100})
	}
	g.drawEntities(screen)

	// screen.DrawImage(tmpImg, &ebiten.DrawImageOptions{})
}

func (g *Game) drawControl(screen *ebiten.Image) {

	scale := 0.7
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(5, 150)
	if !leftPadPressed() {
		op.ColorM.Scale(1, 1, 1, 0.5)
	}
	screen.DrawImage(controlLeftImg, op)

	op.ColorM.Reset()
	op.GeoM.Translate(float64(controlLeftImg.Bounds().Dx())*scale+5, 0)
	if !rightPadPressed() {
		op.ColorM.Scale(1, 1, 1, 0.5)
	}
	screen.DrawImage(controlRightImg, op)

	// gfx.DrawFilledCircle(screen, gfx.V(100, 100), 20, draw.GameColors[5])
	// gfx.DrawCircle(screen, gfx.V(100, 100), 21, 2, colornames.Black)
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

		// Handle animated entites
		if g.entities.HasComponents(e, components.AnimatedType) {
			a := g.entities.GetUnsafe(e, components.AnimatedType).(*components.Animated)
			w, h := a.Ase.FrameWidth, a.Ase.FrameHeight
			x, y := a.Ase.GetFrameXY()
			img = img.SubImage(image.Rect(int(x), int(y), int(x+w), int(y+h))).(*ebiten.Image)
		}

		// Handle directed entities
		op := &ebiten.DrawImageOptions{}
		if g.entities.HasComponents(e, components.DirectedType) {
			d := g.Directed(e)

			switch d.D {
			case direction.Left:
				w, h := img.Size()
				op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
				op.GeoM.Scale(-1, 1)
				op.GeoM.Translate(float64(w)/2, float64(h)/2)
			}
		}

		// Handle following entities following directed
		if g.entities.HasComponents(e, components.FollowingType) {
			following := g.entities.GetUnsafe(e, components.FollowingType).(*components.Following)

			d := g.Directed(following.ID)
			switch d.D {
			case direction.Left:
				w, h := img.Size()
				op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
				op.GeoM.Scale(-1, 1)
				op.GeoM.Translate(float64(w)/2, float64(h)/2)
			}
		}

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

func (g *Game) getCameraPosition() image.Rectangle {
	var cameraWidth float64 = 16 * 16.0
	var cameraHeight float64 = 12 * 16.0
	pos := g.entities.GetUnsafe(playerID, components.PosType).(*components.Pos)
	cx := pos.X - cameraWidth/2
	cx = min(float64(g.Width()*16)-cameraWidth, cx)
	cx = max(0, cx)
	return image.Rect(int(cx), 0, int(cx+cameraWidth), int(cameraHeight))
}
