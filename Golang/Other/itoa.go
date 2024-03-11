package piscine

import "fmt"

func Itoa(n int) string {
	res := ""
	if n < 0 {
		res += "-"
		n = -n
	}
	if n == 0 {
		return "0"
	}
	for n != 0 {
		digit := rune(n%10) + '0'
		res += string(digit)
		n /= 10
	}
	fmt.Printf("res -> %s\n", res)
	i := 0
	j := len(res) - 1
	runes := []rune(res)
	if runes[0] == '-' {
		i = 1
	}
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}
