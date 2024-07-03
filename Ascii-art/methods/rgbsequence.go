package methods

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func RGB(input string) (int, int, int) {
	re := regexp.MustCompile(`^rgb\( *(\d+) *, *(\d+) *, *(\d+) *\)$`)

	matches := re.FindStringSubmatch(input)
	if matches == nil {
		fmt.Println("ERROR: Incorrect RGB() Format")
		os.Exit(1)
	}
	r, _ := strconv.Atoi(matches[1])
	g, _ := strconv.Atoi(matches[2])
	b, _ := strconv.Atoi(matches[3])
	if r > 255 || g > 255 || b > 255 {
		fmt.Println("ERROR: rbg values start from 0 to 255")
		os.Exit(1)
	}
	return r, g, b
}
