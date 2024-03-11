package main

import (
	"fmt"
	"os"
)

func main() {
	params := os.Args
	if len(params) == 1 {
		fmt.Println("File name missing")
		return
	}
	if len(params) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	file, err := os.ReadFile(params[1])
	if err != nil {
		fmt.Println("File name missing")
		return
	}
	res := string(file[:len(file)-1])
	fmt.Println(res)
}
