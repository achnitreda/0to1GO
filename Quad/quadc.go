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
			QuadC(x, y)
		} 
	}
}

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 && j == 1) || (i == 1 && j == x) {
				fmt.Print("A")
			} else if (i == y && j == 1) || (i == y && j == x) {
				fmt.Print("C")
			} else if (i == 1 || i == y) || (j == 1 || j == x) {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

}
