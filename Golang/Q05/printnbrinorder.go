package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var runes []rune
	for n != 0 {
		digit := n % 10
		runes = append(runes, rune(digit))
		n /= 10
	}
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				temp := runes[i]
				runes[i] = runes[j]
				runes[j] = temp
			}
		}
	}
	for i := 0; i < len(runes); i++ {
		z01.PrintRune(runes[i] + '0')
	}
}
