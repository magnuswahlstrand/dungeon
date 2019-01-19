package draw

import (
	"image/color"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

var DrawRect bool = true
var Debug bool = true

func ResolvPt(screen *ebiten.Image, pt resolv.IntersectionPoint, clr color.Color) {
	Pt(screen, gfx.V(float64(pt.X), float64(pt.Y)), clr)
}

func ResolvLine(screen *ebiten.Image, l *resolv.Line, clr color.Color) {
	if Debug {

		ebitenutil.DrawLine(screen, float64(l.X), float64(l.Y), float64(l.X2), float64(l.Y2), clr)
	}
}

func Rect(screen *ebiten.Image, r gfx.Rect, c color.Color) {
	if DrawRect {
		ebitenutil.DrawLine(screen, r.Min.X+1, r.Min.Y+1, r.Min.X+1, r.Max.Y, c)
		ebitenutil.DrawLine(screen, r.Min.X+1, r.Max.Y-1, r.Max.X, r.Max.Y-1, c)
		ebitenutil.DrawLine(screen, r.Max.X-1, r.Max.Y-1, r.Max.X-1, r.Min.Y, c)
		ebitenutil.DrawLine(screen, r.Max.X-1, r.Min.Y+1, r.Min.X, r.Min.Y+1, c)
	}
}

func Pt(screen *ebiten.Image, pt gfx.Vec, clr color.Color) {

	ebitenutil.DrawRect(screen, float64(pt.X-3), float64(pt.Y-3), float64(5), float64(5), clr)
	ebitenutil.DrawRect(screen, float64(pt.X-2), float64(pt.Y-2), float64(3), float64(3), colornames.Ghostwhite)
}

/*
	a := &gfx.Animation{
		frames,
		[]color.Palette{
			gamePalette,
			gamePalette,
			gamePalette,
		},
		50,
		0,
	}

	a.SaveGIF("animation.gif")
*/
