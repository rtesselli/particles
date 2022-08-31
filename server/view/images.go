package view

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

func CircleImage(radius int, r, g, b, a float64) *ebiten.Image {
	dc := gg.NewContext(2*radius, 2*radius)
	dc.DrawCircle(float64(radius), float64(radius), float64(radius))
	dc.SetRGBA(r, g, b, a)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}
