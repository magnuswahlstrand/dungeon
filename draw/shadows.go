package draw

import (
	"image/color"
	"sort"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/dungeon/resolvutil"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

var shadowImg *ebiten.Image
var trianglesImg *ebiten.Image
var whiteImg *ebiten.Image

func init() {
	trianglesImg, _ = ebiten.NewImage(12*16, 12*16, ebiten.FilterDefault)
	shadowImg, _ = ebiten.NewImage(12*16, 12*16, ebiten.FilterDefault)
	whiteImg, _ = ebiten.NewImage(12, 12, ebiten.FilterDefault)

	whiteImg.Fill(colornames.White)
}

func Shadow(dstImg *ebiten.Image, start gfx.Vec, pts []gfx.Vec, space resolv.Space) {
	trianglesImg.Clear()
	shadowImg.Fill(color.Black)

	foundPts := []gfx.Vec{}
	for _, p := range pts {
		Pt(dstImg, p, colornames.Red)

		for _, a := range []float64{-0.01, 0.0, 0.01} {
			target := p.Sub(start).Rotated(a).Unit().Scaled(1000).Add(start)
			l := resolvutil.Line(start, target)
			pts := l.IntersectionPoints(&space)
			if len(pts) > 0 {
				p := gfx.V(float64(pts[0].X), float64(pts[0].Y))
				foundPts = append(foundPts, p.Sub(start))
			}

			ResolvLine(dstImg, l, color.RGBA{255, 255, 255, 100})
		}
	}

	opt := &ebiten.DrawTrianglesOptions{}
	opt.Address = ebiten.AddressRepeat

	sort.Slice(foundPts, func(i int, j int) bool {
		return foundPts[i].Angle() < foundPts[j].Angle()
	})

	pPrev := foundPts[len(foundPts)-1]
	for _, p := range foundPts {
		// drawPt(dstImg, p.Add(player), colornames.Turquoise)

		vertices := []ebiten.Vertex{
			newVertex(start.X, start.Y),
			newVertex(start.Add(p).X, start.Add(p).Y),
			newVertex(start.Add(pPrev).X, start.Add(pPrev).Y),
		}
		trianglesImg.DrawTriangles(vertices, []uint16{0, 1, 2}, whiteImg, opt)
		pPrev = p
	}

	// tmpImg.DrawImage(shadowImg, &ebiten.DrawImageOptions{})
	op := &ebiten.DrawImageOptions{}
	op.CompositeMode = ebiten.CompositeModeDestinationOut
	shadowImg.DrawImage(trianglesImg, op)
	op = &ebiten.DrawImageOptions{}
	// op.CompositeMode = ebiten.CompositeModeDestinationOut
	op.ColorM.Scale(1, 1, 1, 0.7)
	dstImg.DrawImage(shadowImg, op)
}

// Used by drawShadow
func newVertex(x, y float64) ebiten.Vertex {
	return ebiten.Vertex{
		DstX:   float32(x),
		DstY:   float32(y),
		SrcX:   0,
		SrcY:   0,
		ColorR: 1,
		ColorG: 1,
		ColorB: 1,
		ColorA: 1,
	}
}
