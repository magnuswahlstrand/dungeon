package resolvutil

import (
	"github.com/SolarLune/resolv/resolv"
	"github.com/peterhellberg/gfx"
)

func Line(a, b gfx.Vec) *resolv.Line {
	return resolv.NewLine(int32(a.X), int32(a.Y), int32(b.X), int32(b.Y))
}

func Rect(r gfx.Rect) *resolv.Rectangle {
	return resolv.NewRectangle(int32(r.Min.X), int32(r.Min.Y), int32(r.W()), int32(r.H()))
}
