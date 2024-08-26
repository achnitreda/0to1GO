package methods

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func HslToRgb(input string) (int, int, int) {
	re := regexp.MustCompile(`^hsl\( *(\d+) *, *(\d+)% *, *(\d+)% *\)$`)

	matches := re.FindStringSubmatch(input)
	if matches == nil {
		fmt.Println("ERROR: Incorrect Hsl() Format")
		os.Exit(1)
	}
	h, _ := strconv.Atoi(matches[1])
	s, _ := strconv.Atoi(matches[2])
	l, _ := strconv.Atoi(matches[3])
	if h > 360 || s > 100 || l > 100 {
		fmt.Println("ERROR: HSL values: Hue (0-360°), Saturation (0-100%), Lightness (0-100%)")
		os.Exit(1)
	}
	// Normalize HSL values:
	H := float64(h)
	S := float64(s) / 100
	L := float64(l) / 100
	// Calculate Chroma: (C) = (1 - |2L - 1|) * S
	C := (1 - math.Abs(2*L-1)) * S
	// Calculate X: X = C * (1 - |(H / 60) mod 2 - 1|)
	X := C * (1 - math.Abs(math.Mod(H/60, float64(2))-1))
	// Calculate m: m = L - C / 2
	m := L - C/2

	var r, g, b float64
	switch {
	case 0 <= h && h < 60:
		r, g, b = C, X, 0
	case 60 <= h && h < 120:
		r, g, b = X, C, 0
	case 120 <= h && h < 180:
		r, g, b = 0, C, X
	case 180 <= h && h < 240:
		r, g, b = 0, X, C
	case 240 <= h && h < 300:
		r, g, b = X, 0, C
	case 300 <= h && h <= 360:
		r, g, b = C, 0, X
	default:
		r, g, b = 0, 0, 0
	}

	r = (r + m) * 255
	g = (g + m) * 255
	b = (b + m) * 255

	return int(r), int(g), int(b)
}
