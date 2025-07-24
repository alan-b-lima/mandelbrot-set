package color

import (
	"math"
)

func HSLToRGB(hue, sat, lum float64) (R, G, B uint8) {
	hue /= 60
	if hue < 0 {
		hue = -hue
	}

	C := (1 - math.Abs(2*lum-1)) * sat
	X := C * (1 - math.Abs(math.Mod(hue, 2)-1))

	var r, g, b float64

	switch int(hue) % 6 {
	case 0: r, g, b = C, X, 0
	case 1: r, g, b = X, C, 0
	case 2: r, g, b = 0, C, X
	case 3: r, g, b = 0, X, C
	case 4: r, g, b = X, 0, C
	case 5: r, g, b = C, 0, X
	}

	m := lum - C/2
	
	R = uint8((r + m) * 255)
	G = uint8((g + m) * 255)
	B = uint8((b + m) * 255)
	return
}

