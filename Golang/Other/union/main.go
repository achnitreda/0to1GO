package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	p1 := []byte(os.Args[1])
	p2 := []byte(os.Args[2])
	var res []byte
	for i := 0; i < len(p1); i++ {
		for j := i + 1; j < len(p1); j++ {
			if p1[i] == p1[j] {
				p1[j] = '+'
			}
		}
		res = append(res, p1[i])
	}

	for i := 0; i < len(p2); i++ {
		for j := i + 1; j < len(p2); j++ {
			if p2[i] == p2[j] {
				p2[j] = '+'
			}
		}
		res = append(res, p2[i])
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

	for _, char := range ans {
		if char != '+' {
			fmt.Printf("%c", char)
		}
	}
	fmt.Println()
}
