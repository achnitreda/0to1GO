package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var res []rune

	if n == -9223372036854775808 {
		ans := "-9223372036854775808"
		for i := 0; i < len(ans); i++ {
			z01.PrintRune(rune(ans[i]))
		}
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		digit := n % 10
		res = append(res, rune(digit+'0'))
		n /= 10
	}

	for i := len(res) - 1; i >= 0; i-- {
		z01.PrintRune(res[i])
	}
}
