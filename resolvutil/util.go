package resolvutil

import (
	"github.com/SolarLune/resolv/resolv"
	"github.com/peterhellberg/gfx"
)

func Line(a, b gfx.Vec) *resolv.Line {
	return resolv.NewLine(int32(a.X), int32(a.Y), int32(b.X), int32(b.Y))
}

func ScaledLine(a, b gfx.Vec, factor float64) *resolv.Line {
	return resolv.NewLine(int32(a.X*factor), int32(a.Y*factor), int32(b.X*factor), int32(b.Y*factor))
}

func Rect(r gfx.Rect) *resolv.Rectangle {
	return resolv.NewRectangle(int32(r.Min.X), int32(r.Min.Y), int32(r.W()), int32(r.H()))
}

func ScaledRect(r gfx.Rect, factor float64) *resolv.Rectangle {
	return resolv.NewRectangle(int32(r.Min.X*factor), int32(r.Min.Y*factor), int32(r.W()*factor), int32(r.H()*factor))
}
