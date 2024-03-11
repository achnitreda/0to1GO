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
			QuadB(x, y)
		} 
	}
}

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	if y == 1 {
		for a := 0; a < y; a++ {
			for b := 0; b < x; b++ {
				if b == 0 {
					fmt.Print("/")
				} else if b == x-1 {
					fmt.Print("\\")
				} else {
					fmt.Print("*")
				}
			}
			fmt.Print("\n")
		}
	} else if x == 1 {
		for a := 0; a < y; a++ {
			for b := 0; b < x; b++ {
				if a == 0 {
					fmt.Print("/")
				} else if a == y-1 {
					fmt.Print("\\")
				} else {
					fmt.Print("*")
				}
			}
			fmt.Print("\n")
		}
	} else {
		for a := 0; a < y; a++ {
			for b := 0; b < x; b++ {
				if (a == 0 && b == 0) || (a == y-1 && b == x-1) {
					fmt.Print("/")
				} else if (a == 0 && b == x-1) || (b == 0 && a == y-1) {
					fmt.Print("\\")
				} else if (a == 0 || a == y-1) || (b == 0 || b == x-1) {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
}
