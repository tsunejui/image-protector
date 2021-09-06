package image

import (
	"fmt"
	"image/color"
)

type Pixel struct {
	X     int
	Y     int
	Color *Color
}

const (
	max      = 256
	interval = 32 // 256 / 8
)

func NewPixel(x, y int) *Pixel {
	return &Pixel{
		X: x,
		Y: y,
	}
}

func (p *Pixel) WithColor(c color.Color) *Pixel {
	r, g, b, a := c.RGBA()
	p.Color = &Color{
		R: newColor(r),
		G: newColor(g),
		B: newColor(b),
		A: newColor(a),
	}
	return p
}

func newColor(v uint32) *ColorValue {
	m := v >> 8
	if m > max {
		panic(fmt.Errorf("invalid color: %d", v))
	}
	return NewColorValue(uint8(m))
}
