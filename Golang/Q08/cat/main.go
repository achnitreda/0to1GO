package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func Error(err string) {
	runes := []rune(err)
	for i := 0; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}

func printFileContent(content []byte) {
	s := string(content)
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}

func main() {
	params := os.Args
	if len(params) == 1 {
		io.Copy(os.Stdout, os.Stdin)
	}
	if len(params) >= 2 {
		for _, param := range params[1:] {
			file, err := os.ReadFile(param)
			if err != nil {
				Error("ERROR: " + err.Error())
				os.Exit(1)
			} else {
				printFileContent(file)
			}
		}
	}
}
