package named

import (
	prism "../"
	"image/color"
)

const (
	AliceBlue            = "aliceblue"
	AntiqueWhite         = "antiquewhite"
	Aqua                 = "aqua"
	Aquamarine           = "aquamarine"
	Azure                = "azure"
	Beige                = "beige"
	Bisque               = "bisque"
	Black                = "black"
	BlanchedAlmond       = "blanchedalmond"
	Blue                 = "blue"
	BlueViolet           = "blueviolet"
	Brown                = "brown"
	BurlyWood            = "burlywood"
	CadetBlue            = "cadetblue"
	Chartreuse           = "chartreuse"
	Chocolate            = "chocolate"
	Coral                = "coral"
	CornflowerBlue       = "cornflowerblue"
	Cornsilk             = "cornsilk"
	Crimson              = "crimson"
	Cyan                 = "cyan"
	DarkBlue             = "darkblue"
	DarkCyan             = "darkcyan"
	DarkGoldenRod        = "darkgoldenrod"
	DarkGray             = "darkgray"
	DarkGreen            = "darkgreen"
	DarkKhaki            = "darkkhaki"
	DarkMagenta          = "darkmagenta"
	DarkOliveGreen       = "darkolivegreen"
	DarkOrange           = "darkorange"
	DarkOrchid           = "darkorchid"
	DarkRed              = "darkred"
	DarkSalmon           = "darksalmon"
	DarkSeaGreen         = "darkseagreen"
	DarkSlateBlue        = "darkslateblue"
	DarkSlateGray        = "darkslategray"
	DarkTurquoise        = "darkturquoise"
	DarkViolet           = "darkviolet"
	DeepPink             = "deeppink"
	DeepSkyBlue          = "deepskyblue"
	DimGray              = "dimgray"
	DodgerBlue           = "dodgerblue"
	FireBrick            = "firebrick"
	FloralWhite          = "floralwhite"
	ForestGreen          = "forestgreen"
	Fuchsia              = "fuchsia"
	Gainsboro            = "gainsboro"
	GhostWhite           = "ghostwhite"
	Gold                 = "gold"
	GoldenRod            = "goldenrod"
	Gray                 = "gray"
	Green                = "green"
	GreenYellow          = "greenyellow"
	HoneyDew             = "honeydew"
	HotPink              = "hotpink"
	IndianRed            = "indianred"
	Indigo               = "indigo"
	Ivory                = "ivory"
	Khaki                = "khaki"
	Lavender             = "lavender"
	LavenderBlush        = "lavenderblush"
	LawnGreen            = "lawngreen"
	LemonChiffon         = "lemonchiffon"
	LightBlue            = "lightblue"
	LightCoral           = "lightcoral"
	LightCyan            = "lightcyan"
	LightGoldenRodYellow = "lightgoldenrodyellow"
	LightGray            = "lightgray"
	LightGreen           = "lightgreen"
	LightPink            = "lightpink"
	LightSalmon          = "lightsalmon"
	LightSeaGreen        = "lightseagreen"
	LightSkyBlue         = "lightskyblue"
	LightSlateGray       = "lightslategray"
	LightSteelBlue       = "lightsteelblue"
	LightYellow          = "lightyellow"
	Lime                 = "lime"
	LimeGreen            = "limegreen"
	Linen                = "linen"
	Magenta              = "magenta"
	Maroon               = "maroon"
	MediumAquaMarine     = "mediumaquamarine"
	MediumBlue           = "mediumblue"
	MediumOrchid         = "mediumorchid"
	MediumPurple         = "mediumpurple"
	MediumSeaGreen       = "mediumseagreen"
	MediumSlateBlue      = "mediumslateblue"
	MediumSpringGreen    = "mediumspringgreen"
	MediumTurquoise      = "mediumturquoise"
	MediumVioletRed      = "mediumvioletred"
	MidnightBlue         = "midnightblue"
	MintCream            = "mintcream"
	MistyRose            = "mistyrose"
	Moccasin             = "moccasin"
	NavajoWhite          = "navajowhite"
	Navy                 = "navy"
	OldLace              = "oldlace"
	Olive                = "olive"
	OliveDrab            = "olivedrab"
	Orange               = "orange"
	OrangeRed            = "orangered"
	Orchid               = "orchid"
	PaleGoldenRod        = "palegoldenrod"
	PaleGreen            = "palegreen"
	PaleTurquoise        = "paleturquoise"
	PaleVioletRed        = "palevioletred"
	PapayaWhip           = "papayawhip"
	PeachPuff            = "peachpuff"
	Peru                 = "peru"
	Pink                 = "pink"
	Plum                 = "plum"
	PowderBlue           = "powderblue"
	Purple               = "purple"
	RebeccaPurple        = "rebeccapurple"
	Red                  = "red"
	RosyBrown            = "rosybrown"
	RoyalBlue            = "royalblue"
	SaddleBrown          = "saddlebrown"
	Salmon               = "salmon"
	SandyBrown           = "sandybrown"
	SeaGreen             = "seagreen"
	SeaShell             = "seashell"
	Sienna               = "sienna"
	Silver               = "silver"
	SkyBlue              = "skyblue"
	SlateBlue            = "slateblue"
	SlateGray            = "slategray"
	Snow                 = "snow"
	SpringGreen          = "springgreen"
	SteelBlue            = "steelblue"
	Tan                  = "tan"
	Teal                 = "teal"
	Thistle              = "thistle"
	Tomato               = "tomato"
	Turquoise            = "turquoise"
	Violet               = "violet"
	Wheat                = "wheat"
	White                = "white"
	WhiteSmoke           = "whitesmoke"
	YellowGreen          = "yellowgreen"
	Yellow               = "yellow"
)

var names = map[string]prism.RGB{
	"aliceblue":            prism.RGB{R: 240, G: 248, B: 255},
	"antiquewhite":         prism.RGB{R: 250, G: 235, B: 215},
	"aqua":                 prism.RGB{R: 0, G: 255, B: 255},
	"aquamarine":           prism.RGB{R: 127, G: 255, B: 212},
	"azure":                prism.RGB{R: 240, G: 255, B: 255},
	"beige":                prism.RGB{R: 245, G: 245, B: 220},
	"bisque":               prism.RGB{R: 255, G: 228, B: 196},
	"black":                prism.RGB{R: 0, G: 0, B: 0},
	"blanchedalmond":       prism.RGB{R: 255, G: 235, B: 205},
	"blue":                 prism.RGB{R: 0, G: 0, B: 255},
	"blueviolet":           prism.RGB{R: 138, G: 43, B: 226},
	"brown":                prism.RGB{R: 165, G: 42, B: 42},
	"burlywood":            prism.RGB{R: 222, G: 184, B: 135},
	"cadetblue":            prism.RGB{R: 95, G: 158, B: 160},
	"chartreuse":           prism.RGB{R: 127, G: 255, B: 0},
	"chocolate":            prism.RGB{R: 210, G: 105, B: 30},
	"coral":                prism.RGB{R: 255, G: 127, B: 80},
	"cornflowerblue":       prism.RGB{R: 100, G: 149, B: 237},
	"cornsilk":             prism.RGB{R: 255, G: 248, B: 220},
	"crimson":              prism.RGB{R: 220, G: 20, B: 60},
	"cyan":                 prism.RGB{R: 0, G: 255, B: 255},
	"darkblue":             prism.RGB{R: 0, G: 0, B: 139},
	"darkcyan":             prism.RGB{R: 0, G: 139, B: 139},
	"darkgoldenrod":        prism.RGB{R: 184, G: 134, B: 11},
	"darkgray":             prism.RGB{R: 169, G: 169, B: 169},
	"darkgreen":            prism.RGB{R: 0, G: 100, B: 0},
	"darkkhaki":            prism.RGB{R: 189, G: 183, B: 107},
	"darkmagenta":          prism.RGB{R: 139, G: 0, B: 139},
	"darkolivegreen":       prism.RGB{R: 85, G: 107, B: 47},
	"darkorange":           prism.RGB{R: 255, G: 140, B: 0},
	"darkorchid":           prism.RGB{R: 153, G: 50, B: 204},
	"darkred":              prism.RGB{R: 139, G: 0, B: 0},
	"darksalmon":           prism.RGB{R: 233, G: 150, B: 122},
	"darkseagreen":         prism.RGB{R: 143, G: 188, B: 143},
	"darkslateblue":        prism.RGB{R: 72, G: 61, B: 139},
	"darkslategray":        prism.RGB{R: 47, G: 79, B: 79},
	"darkturquoise":        prism.RGB{R: 0, G: 206, B: 209},
	"darkviolet":           prism.RGB{R: 148, G: 0, B: 211},
	"deeppink":             prism.RGB{R: 255, G: 20, B: 147},
	"deepskyblue":          prism.RGB{R: 0, G: 191, B: 255},
	"dimgray":              prism.RGB{R: 105, G: 105, B: 105},
	"dodgerblue":           prism.RGB{R: 30, G: 144, B: 255},
	"firebrick":            prism.RGB{R: 178, G: 34, B: 34},
	"floralwhite":          prism.RGB{R: 255, G: 250, B: 240},
	"forestgreen":          prism.RGB{R: 34, G: 139, B: 34},
	"fuchsia":              prism.RGB{R: 255, G: 0, B: 255},
	"gainsboro":            prism.RGB{R: 220, G: 220, B: 220},
	"ghostwhite":           prism.RGB{R: 248, G: 248, B: 255},
	"gold":                 prism.RGB{R: 255, G: 215, B: 0},
	"goldenrod":            prism.RGB{R: 218, G: 165, B: 32},
	"gray":                 prism.RGB{R: 128, G: 128, B: 128},
	"green":                prism.RGB{R: 0, G: 128, B: 0},
	"greenyellow":          prism.RGB{R: 173, G: 255, B: 47},
	"honeydew":             prism.RGB{R: 240, G: 255, B: 240},
	"hotpink":              prism.RGB{R: 255, G: 105, B: 180},
	"indianred":            prism.RGB{R: 205, G: 92, B: 92},
	"indigo":               prism.RGB{R: 75, G: 0, B: 130},
	"ivory":                prism.RGB{R: 255, G: 255, B: 240},
	"khaki":                prism.RGB{R: 240, G: 230, B: 140},
	"lavender":             prism.RGB{R: 230, G: 230, B: 250},
	"lavenderblush":        prism.RGB{R: 255, G: 240, B: 245},
	"lawngreen":            prism.RGB{R: 124, G: 252, B: 0},
	"lemonchiffon":         prism.RGB{R: 255, G: 250, B: 205},
	"lightblue":            prism.RGB{R: 173, G: 216, B: 230},
	"lightcoral":           prism.RGB{R: 240, G: 128, B: 128},
	"lightcyan":            prism.RGB{R: 224, G: 255, B: 255},
	"lightgoldenrodyellow": prism.RGB{R: 250, G: 250, B: 210},
	"lightgray":            prism.RGB{R: 211, G: 211, B: 211},
	"lightgreen":           prism.RGB{R: 144, G: 238, B: 144},
	"lightpink":            prism.RGB{R: 255, G: 182, B: 193},
	"lightsalmon":          prism.RGB{R: 255, G: 160, B: 122},
	"lightseagreen":        prism.RGB{R: 32, G: 178, B: 170},
	"lightskyblue":         prism.RGB{R: 135, G: 206, B: 250},
	"lightslategray":       prism.RGB{R: 119, G: 136, B: 153},
	"lightsteelblue":       prism.RGB{R: 176, G: 196, B: 222},
	"lightyellow":          prism.RGB{R: 255, G: 255, B: 224},
	"lime":                 prism.RGB{R: 0, G: 255, B: 0},
	"limegreen":            prism.RGB{R: 50, G: 205, B: 50},
	"linen":                prism.RGB{R: 250, G: 240, B: 230},
	"magenta":              prism.RGB{R: 255, G: 0, B: 255},
	"maroon":               prism.RGB{R: 128, G: 0, B: 0},
	"mediumaquamarine":     prism.RGB{R: 102, G: 205, B: 170},
	"mediumblue":           prism.RGB{R: 0, G: 0, B: 205},
	"mediumorchid":         prism.RGB{R: 186, G: 85, B: 211},
	"mediumpurple":         prism.RGB{R: 147, G: 112, B: 219},
	"mediumseagreen":       prism.RGB{R: 60, G: 179, B: 113},
	"mediumslateblue":      prism.RGB{R: 123, G: 104, B: 238},
	"mediumspringgreen":    prism.RGB{R: 0, G: 250, B: 154},
	"mediumturquoise":      prism.RGB{R: 72, G: 209, B: 204},
	"mediumvioletred":      prism.RGB{R: 199, G: 21, B: 133},
	"midnightblue":         prism.RGB{R: 25, G: 25, B: 112},
	"mintcream":            prism.RGB{R: 245, G: 255, B: 250},
	"mistyrose":            prism.RGB{R: 255, G: 228, B: 225},
	"moccasin":             prism.RGB{R: 255, G: 228, B: 181},
	"navajowhite":          prism.RGB{R: 255, G: 222, B: 173},
	"navy":                 prism.RGB{R: 0, G: 0, B: 128},
	"oldlace":              prism.RGB{R: 253, G: 245, B: 230},
	"olive":                prism.RGB{R: 128, G: 128, B: 0},
	"olivedrab":            prism.RGB{R: 107, G: 142, B: 35},
	"orange":               prism.RGB{R: 255, G: 165, B: 0},
	"orangered":            prism.RGB{R: 255, G: 69, B: 0},
	"orchid":               prism.RGB{R: 218, G: 112, B: 214},
	"palegoldenrod":        prism.RGB{R: 238, G: 232, B: 170},
	"palegreen":            prism.RGB{R: 152, G: 251, B: 152},
	"paleturquoise":        prism.RGB{R: 175, G: 238, B: 238},
	"palevioletred":        prism.RGB{R: 219, G: 112, B: 147},
	"papayawhip":           prism.RGB{R: 255, G: 239, B: 213},
	"peachpuff":            prism.RGB{R: 255, G: 218, B: 185},
	"peru":                 prism.RGB{R: 205, G: 133, B: 63},
	"pink":                 prism.RGB{R: 255, G: 192, B: 203},
	"plum":                 prism.RGB{R: 221, G: 160, B: 221},
	"powderblue":           prism.RGB{R: 176, G: 224, B: 230},
	"purple":               prism.RGB{R: 128, G: 0, B: 128},
	"rebeccapurple":        prism.RGB{R: 102, G: 51, B: 153},
	"red":                  prism.RGB{R: 255, G: 0, B: 0},
	"rosybrown":            prism.RGB{R: 188, G: 143, B: 143},
	"royalblue":            prism.RGB{R: 65, G: 105, B: 225},
	"saddlebrown":          prism.RGB{R: 139, G: 69, B: 19},
	"salmon":               prism.RGB{R: 250, G: 128, B: 114},
	"sandybrown":           prism.RGB{R: 244, G: 164, B: 96},
	"seagreen":             prism.RGB{R: 46, G: 139, B: 87},
	"seashell":             prism.RGB{R: 255, G: 245, B: 238},
	"sienna":               prism.RGB{R: 160, G: 82, B: 45},
	"silver":               prism.RGB{R: 192, G: 192, B: 192},
	"skyblue":              prism.RGB{R: 135, G: 206, B: 235},
	"slateblue":            prism.RGB{R: 106, G: 90, B: 205},
	"slategray":            prism.RGB{R: 112, G: 128, B: 144},
	"snow":                 prism.RGB{R: 255, G: 250, B: 250},
	"springgreen":          prism.RGB{R: 0, G: 255, B: 127},
	"steelblue":            prism.RGB{R: 70, G: 130, B: 180},
	"tan":                  prism.RGB{R: 210, G: 180, B: 140},
	"teal":                 prism.RGB{R: 0, G: 128, B: 128},
	"thistle":              prism.RGB{R: 216, G: 191, B: 216},
	"tomato":               prism.RGB{R: 255, G: 99, B: 71},
	"turquoise":            prism.RGB{R: 64, G: 224, B: 208},
	"violet":               prism.RGB{R: 238, G: 130, B: 238},
	"wheat":                prism.RGB{R: 245, G: 222, B: 179},
	"white":                prism.RGB{R: 255, G: 255, B: 255},
	"whitesmoke":           prism.RGB{R: 245, G: 245, B: 245},
	"yellow":               prism.RGB{R: 255, G: 255, B: 0},
	"yellowgreen":          prism.RGB{R: 154, G: 205, B: 50},
}

type NamedColorFactory struct {
}

func NewFactory() prism.ColorFactory {
	return &NamedColorFactory{}
}

func (n *NamedColorFactory) GetColor(id string) color.Color {
	c, ok := names[id]
	if !ok {
		return nil
	}
	return &c
}
