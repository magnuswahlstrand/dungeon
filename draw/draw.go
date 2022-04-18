package draw

import (
	"image/color"
	"log"

	"github.com/SolarLune/resolv/resolv"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
)

var DrawRect bool = true
var Debug bool = true

//func ResolvPt(screen *ebiten.Image, pt resolv.IntersectionPoint, clr color.Color) {
//	Pt(screen, gfx.V(float64(pt.X), float64(pt.Y)), clr)
//}

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

func init() {
	fnt, err := truetype.Parse(gomono.TTF)
	if err != nil {
		log.Fatal("loading font:", err)
	}
	const dpi = 144
	FontFace5 = truetype.NewFace(fnt, &truetype.Options{
		Size:    5,
		DPI:     dpi,
		Hinting: font.HintingNone,
	})
	FontFace7 = truetype.NewFace(fnt, &truetype.Options{
		Size:    7,
		DPI:     dpi,
		Hinting: font.HintingNone,
	})
	FontFace9 = truetype.NewFace(fnt, &truetype.Options{
		Size:    9,
		DPI:     dpi,
		Hinting: font.HintingNone,
	})
	FontFace11 = truetype.NewFace(fnt, &truetype.Options{
		Size:    11,
		DPI:     dpi,
		Hinting: font.HintingNone,
	})
}

var FontFace5 font.Face
var FontFace7 font.Face
var FontFace9 font.Face
var FontFace11 font.Face

func CenterText(screen *ebiten.Image, txt string, face font.Face, c color.Color, offsetY ...int) {
	y := 0
	for _, o := range offsetY {
		y += o
	}
	size := face.Metrics().Height.Ceil() / 2
	width := int(1.135 * float64(len(txt)*size))
	w, h := screen.Size()
	text.Draw(screen, txt, face, (w-width)/2, (h+size)/2+y, c)
}
