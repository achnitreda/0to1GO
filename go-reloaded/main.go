package main

import (
	"fmt"
	"go-reloaded/converter"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inputContent, err := os.ReadFile(inputFile)

	if err != nil {
		fmt.Printf("Error reading %s: %v\n", inputFile, err)
		return
	}

	resultContent := converter.ConvertBase(string(inputContent))

	err = os.WriteFile(outputFile, []byte(resultContent), 0666)
	if err != nil {
		fmt.Printf("Error writing to %s: %v\n", outputFile, err)
		return
	}

	fmt.Printf("Contents of %s successfully copied to %s\n", inputFile, outputFile)
}
