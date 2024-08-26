package methods

import (
	"fmt"
	"os"
)

func WriteOutput(outputFile string, resultContent string) {
	err := os.WriteFile(outputFile, []byte(resultContent), 0o666)
	if err != nil {
		fmt.Printf("Error writing to %s: %v\n", outputFile, err)
		os.Exit(1)
	}
}
