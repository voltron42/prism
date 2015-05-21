package prism

import "image/color"

type ColorFactory interface {
	GetColor(id string) color.Color
}

type RGB struct {
	R, G, B uint32
}

func (c *RGB) RGBA() (r, g, b, a uint32) {
	r = c.R
	g = c.G
	b = c.B
	a = 255
	return
}
