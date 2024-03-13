package sudoku

import (
	"fmt"
	"os"
	"strings"
)

var grid [][]rune

func printMatrix() {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if j != len(grid[i])-1 {
				fmt.Printf("%c ", grid[i][j])
			} else {
				fmt.Printf("%c", grid[i][j])
			}
		}
		fmt.Println()
	}
}

func checkArg(arg string) bool {
	if len(arg) != 9 {
		return false
	}
	for i := 0; i < len(arg); i++ {
		if arg[i] != '.' && (arg[i] < '1' || arg[i] > '9') {
			return false
		}
		for j := i + 1; j < len(arg); j++ {
			if arg[i] != '.' && arg[i] == arg[j] {
				return false
			}
		}
	}
	return true
}

func checkArgs(args []string) bool {
	if len(args) != 9 {
		return false
	}
	for i := 0; i < len(args[1:]); i++ {
		if !checkArg(args[1:][i]) {
			return false
		}
	}
	return true
}

func possibleValue(y, x int, n rune) bool {
	for i := 0; i < 9; i++ {
		if grid[y][i] == n {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if grid[i][x] == n {
			return false
		}
	}
	x0 := (x / 3) * 3
	y0 := (y / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[y0+i][x0+j] == n {
				return false
			}
		}
	}
	return true
}

func Solve() bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x] == '.' {
				for n := '1'; n <= '9'; n++ {
					if possibleValue(y, x, n) {
						grid[y][x] = n
						if Solve() {
							return true
						}
						grid[y][x] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func Sudoku() {
	var params []string
	for _, arg := range os.Args[1:] {
		param := strings.Split(arg, " ")
		params = append(params, param...)
	}
	if !checkArgs(params) {
		fmt.Println("Error")
		return
	}
	grid = make([][]rune, 9)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]rune, 9)
		for j := 0; j < len(params[i]); j++ {
			grid[i][j] = rune(params[i][j])
		}
	}

	if Solve() {
		printMatrix()
	} else {
		fmt.Println("Error")
	}
}
