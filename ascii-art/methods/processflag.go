package methods

import (
	"strings"
)

func OutputFile(str string) string {
	outputFile := ""
	if strings.HasPrefix(str, "--output=") && strings.HasSuffix(str, ".txt") {
		if !strings.Contains(str[9:], "/") {
			outputFile = str[9:]
		}
	}
	return outputFile
}

func AlignColorOutput(str string) string {
	alignColorOutput := ""
	if strings.HasPrefix(str, "--align=") || strings.HasPrefix(str, "--color=") {
		alignColorOutput = str[8:]
	}
	return alignColorOutput
}

func CheckFlag(s string) bool {
	if len(s) < 7 {
		return false
	}
	return strings.HasPrefix(s, "--output") || strings.HasPrefix(s, "--align") || strings.HasPrefix(s, "--color")
}
