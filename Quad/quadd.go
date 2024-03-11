package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args
	if len(arg) == 3 {
		x, _ := strconv.Atoi(os.Args[1])
		y, _ := strconv.Atoi(os.Args[2])
		if x > 0 && y > 0 {
			QuadD(x, y)
		} 
	}
}

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for a := 0; a < y; a++ {
		for b := 0; b < x; b++ {
			if (a == 0 || a == y-1) && b == 0 {
				fmt.Print("A")
			} else if (a == 0 && b == x-1) || (a == y-1 && b == x-1) {
				fmt.Print("C")
			} else if (a == 0 || a == y-1) || (b == 0 || b == x-1) {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
