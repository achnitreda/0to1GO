package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	var res []byte
	for i := 0; i < len(os.Args[1]); i++ {
		for j := 0; j < len(os.Args[2]); j++ {
			if os.Args[1][i] == os.Args[2][j] {
				res = append(res, os.Args[1][i])
				break
			}
		}
	}

	var ans []byte
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] == res[j] {
				res[j] = '+'
			}
		}
		ans = append(ans, res[i])
	}

	for _, c := range ans {
		if c != '+' {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
