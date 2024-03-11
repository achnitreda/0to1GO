package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	for i := 0; i < len(args); i++ {
		runes := []rune(args[i])
		for j := 0; j < len(runes); j++ {
			z01.PrintRune(runes[j])
		}
		z01.PrintRune('\n')
	}
}
