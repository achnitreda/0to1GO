package main

import (
	"fmt"
	"os"
)

func main() {
	params := os.Args[1:]
	for _, param := range params {
		if param == "01" || param == "galaxy" || param == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
	return
}
