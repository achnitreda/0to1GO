package main

import (
	"os"
)

func parseArg(arg string) int {
	result := 0
	negative := false

	for _, char := range arg {
		if char == '-' {
			negative = true
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		} else {
			return 9223372036854775807
		}
	}

	if negative {
		result = -result
	}
	return result
}

func intToString(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	negative := false

	if num < 0 {
		negative = true
		num = -num
	}

	for num > 0 {
		digit := num % 10
		result = string(digit+'0') + result
		num /= 10
	}

	if negative {
		result = "-" + result
	}

	return result
}

func main() {
	if len(os.Args) != 4 {
		os.Exit(0)
	}

	if os.Args[1] >= "9223372036854775807" || os.Args[3] >= "9223372036854775807" || os.Args[1] == "-9223372036854775809" || os.Args[3] == "-9223372036854775809" {
		os.Exit(0)
	}

	value1 := parseArg(os.Args[1])
	value2 := parseArg(os.Args[3])

	result := 0
	operator := os.Args[2]

	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			os.Exit(0)
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			os.Exit(0)
		}
		result = value1 % value2
	default:
		os.Exit(0)
	}
	os.Stdout.WriteString(intToString(result) + "\n")
}
