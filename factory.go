package prism

import (
	"encoding/xml"
	"image/color"
	"strconv"
)

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

func (rgb *RGB) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		num, err := strconv.Atoi(attr.Value)
		if err != nil {
			return err
		}
		switch attr.Name.Local {
		case "red":
			rgb.R = uint32(num)
		case "green":
			rgb.G = uint32(num)
		case "blue":
			rgb.B = uint32(num)
		}
	}
	d.Token()
	return nil
}
