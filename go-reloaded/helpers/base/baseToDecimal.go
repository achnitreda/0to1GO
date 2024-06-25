package base

import (
	"regexp"
	"strconv"
)

func ConvertBinToDecimal(text string) string {
	binRegex := regexp.MustCompile(`\s([01]+)\s+\(bin\)`)

	return binRegex.ReplaceAllStringFunc(text, func(match string) string {
		binStr := binRegex.FindStringSubmatch(match)[1]
		decimalValue, _ := strconv.ParseInt(binStr, 2, 64)
		return " " + strconv.FormatInt(decimalValue, 10)
	})
}

func ConvertBinBToDecimal(text string) string {
	binBRegex := regexp.MustCompile(`^([01]+)\s+\(bin\)`)

	return binBRegex.ReplaceAllStringFunc(text, func(match string) string {
		binStr := binBRegex.FindStringSubmatch(match)[1]
		decimalValue, _ := strconv.ParseInt(binStr, 2, 64)
		return strconv.FormatInt(decimalValue, 10)
	})
}

func ConvertHexToDecimal(text string) string {
	hexRegex := regexp.MustCompile(`\s([0-9A-Fa-f]+)\s+\(hex\)`)
	return hexRegex.ReplaceAllStringFunc(text, func(match string) string {
		hexStr := hexRegex.FindStringSubmatch(match)[1]
		decimalValue, _ := strconv.ParseInt(hexStr, 16, 64)
		return " " + strconv.FormatInt(decimalValue, 10)
	})
}

func ConvertHexBToDecimal(text string) string {
	hexRegex := regexp.MustCompile(`^([0-9A-Fa-f]+)\s+\(hex\)`)

	return hexRegex.ReplaceAllStringFunc(text, func(match string) string {
		hexStr := hexRegex.FindStringSubmatch(match)[1]
		decimalValue, _ := strconv.ParseInt(hexStr, 16, 64)
		return strconv.FormatInt(decimalValue, 10)
	})
}
