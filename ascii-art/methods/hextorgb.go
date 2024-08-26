package methods

import (
	"fmt"
	"os"
	"strconv"
)

func IsValidHex(color string) bool {
	if len(color) != 4 && len(color) != 7 {
		fmt.Println("Incorect hex ormat EX: (#000000) or (#000)")
		os.Exit(1)
	}
	hex := color[1:]
	for _, char := range hex {
		if (char < '0' || char > '9') && (char < 'a' || char > 'f') && (char < 'A' || char > 'F') {
			return false
		}
	}
	return true
}

func hex6char(color string) string {
	return ("#" + string(color[1]) + string(color[1]) + string(color[2]) + string(color[2]) + string(color[3]) + string(color[3]))
}

func HexToRgb(color string) (int, int, int) {
	var r, g, b int
	if IsValidHex(color) {
		if len(color) == 4 {
			color = hex6char(color) 
		}
		r64, _ := strconv.ParseInt(color[1:3], 16, 8)
		r = int(r64)
		g64, _ := strconv.ParseUint(color[3:5], 16, 8)
		g = int(g64)
		b64, _ := strconv.ParseUint(color[5:], 16, 8)
		b = int(b64)
	}

	return r, g, b
}
