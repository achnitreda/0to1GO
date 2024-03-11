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
			QuadE(x, y)
		} 
	}
}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	if y == 1 {
		for a := 0; a < y; a++ {
			for b := 0; b < x; b++ {
				if b == 0 {
					fmt.Print("A")
				} else if b == x-1 {
					fmt.Print("C")
				} else {
					fmt.Print("B")
				}
			}
			fmt.Print("\n")
		}
	} else if x == 1 {
		for a := 0; a < y; a++ {
			for b := 0; b < x; b++ {
				if a == 0 {
					fmt.Print("A")
				} else if a == y-1 {
					fmt.Print("C")
				} else {
					fmt.Print("B")
				}
			}
			fmt.Print("\n")
		}
	} else {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if (i == 0 && j == 0) || (i == y-1 && j == x-1) {
					fmt.Print("A")
				} else if (i == y-1 && j == 0) || (i == 0 && j == x-1) {
					fmt.Print("C")
				} else if (i == 0 || i == y-1) || (j == 0 || j == x-1) {
					fmt.Print("B")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
}
