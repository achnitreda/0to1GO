package methods

import (
	"strings"
	"log"
	"errors"
)

 var errNonAsciiChar = errors.New("string contain a non ascii character");

func ProccessTheInput(str string) []string {
	if !ValidInput(str) {
		log.Fatal(errNonAsciiChar)
	}
	splitedInput := strings.Split(str, "\\n")

	if CheckNewLine(splitedInput) {
		return splitedInput[1:]
	}
	
	return splitedInput
}

func ValidInput(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] < 32 || str[i] > 126 {
			return false
		}
	}
	return true
}

func CheckNewLine(slice []string) bool {
	for i := 0; i < len(slice); i++ {
		if len(slice[i]) != 0 {
			return false
		}
	}
	return true
}
