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

var GameColors = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xFF},
	color.RGBA{0x16, 0x10, 0x1e, 0xFF},
	color.RGBA{0x2e, 0x24, 0x40, 0xFF},
	color.RGBA{0x70, 0x57, 0x9c, 0xFF},
	color.RGBA{0xe0, 0x96, 0xa8, 0xFF},
	color.RGBA{0xff, 0xf1, 0xeb, 0xFF},
}

var GameColorsTransparent = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xAF},
	color.RGBA{0x16, 0x10, 0x1e, 0xAF},
	color.RGBA{0x2e, 0x24, 0x40, 0xAF},
	color.RGBA{0x70, 0x57, 0x9c, 0xAF},
	color.RGBA{0xe0, 0x96, 0xa8, 0xAF},
	color.RGBA{0xff, 0xf1, 0xeb, 0xAF},
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
