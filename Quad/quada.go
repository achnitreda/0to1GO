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
			QuadA(x, y)
		} 
	}
}

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 || i == y) && (j == 1 || j == x) {
				fmt.Print("o")
			} else if i == 1 || i == y {
				fmt.Print("-")
			} else if j == 1 || j == x {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

}
