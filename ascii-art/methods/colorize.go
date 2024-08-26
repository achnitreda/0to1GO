package methods

import (
	"fmt"
	"log"
	"strings"
)

func Colorize(text string, color string) string {
	myMap := map[string]string{
		"red":      "rgb(255,0,0)",
		"green":    "rgb(0,255,0)",
		"blue":     "rgb(0,0,255)",
		"yellow":   "rgb(255,255,0)",
		"cyan":     "rgb(0,255,255)",
		"magenta":  "rgb(255,0,255)",
		"black":    "rgb(0,0,0)",
		"white":    "rgb(255,255,255)",
		"gray":     "rgb(128,128,128)",
		"orange":   "rgb(255,165,0)",
		"purple":   "rgb(128,0,128)",
		"brown":    "rgb(165,42,42)",
		"pink":     "rgb(255,192,203)",
		"navy":     "rgb(0,0,128)",
		"teal":     "rgb(0,128,128)",
		"olive":    "rgb(128,128,0)",
		"maroon":   "rgb(128,0,0)",
		"silver":   "rgb(192,192,192)",
		"gold":     "rgb(255,215,0)",
		"beige":    "rgb(245,245,220)",
		"coral":    "rgb(255,127,80)",
		"lavender": "rgb(230,230,250)",
		"salmon":   "rgb(250,128,114)",
	}

	// name | #ff0000 | hsl(0, 100%, 50%)
	if color[0] == '#' {
		r, g, b := HexToRgb(color)
		colorCode := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		reset := "\033[0m"
		return colorCode + text + reset
	} else if strings.HasPrefix(color, "hsl") {
		r, g, b := HslToRgb(color)
		colorCode := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		reset := "\033[0m"
		return colorCode + text + reset
	} else if strings.HasPrefix(color, "rgb") || strings.HasPrefix(color, "RGB") {
		r, g, b := RGB(color)
		colorCode := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		reset := "\033[0m"
		return colorCode + text + reset
	} else {
		if colorName, exists := myMap[color]; exists {
			r, g, b := RGB(colorName)
			colorCode := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
			reset := "\033[0m"
			return colorCode + text + reset
		} else {
			log.Fatal("This color do not exists, Try an other Color")
			return ""
		}
	}
}
