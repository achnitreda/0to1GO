package methods

import (
	"strings"
)

func ProccessTheInput(str string) []string {
	splitedInput := strings.Split(str, "\r\n")

	if CheckNewLine(splitedInput) {
		return splitedInput[1:]
	}

	return splitedInput
}

func ValidInput(str string) string {
	var res string
	for i := 0; i < len(str); i++ {
		if str[i] < 32 || str[i] > 126 {
			if str[i] != '\r' && str[i] != '\n' {
				continue
			}
		}
		res += string(str[i])
	}
	return res
}

func CheckNewLine(slice []string) bool {
	for i := 0; i < len(slice); i++ {
		if len(slice[i]) != 0 {
			return false
		}
	}
	return true
}
