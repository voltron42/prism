package hex

import (
	"../"
	"image/color"
	"regexp"
	"strings"
)

type HexColorFactory struct {
}

func NewFactory() prism.ColorFactory {
	return &HexColorFactory{}
}

var pattern = regexp.MustCompile("[#]?[0-9a-fA-F]{6}")

const chars = "0123456789abcdef"

func (h *HexColorFactory) GetColor(id string) color.Color {
	if !pattern.Match([]byte(id)) {
		return nil
	}
	r := h.getHex(id, -6)
	g := h.getHex(id, -4)
	b := h.getHex(id, -2)
	return &prism.RGB{R: r, G: g, B: b}
}

func (h *HexColorFactory) getHex(id string, offset int) uint32 {
	size := len(id)
	hex := strings.Split(id, "")
	a := strings.Index(chars, hex[size+offset])
	b := strings.Index(chars, hex[size+offset+1])
	return uint32(a*256 + b)
}
