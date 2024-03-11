package piscine

import (
	"github.com/01-edu/z01"
)

func IsValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if !IsValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	if nbr == -9223372036854775808 {
		s := "9223372036854775808"
		ans := []rune(s)
		for i := 0; i < len(ans); i++ {
			z01.PrintRune(ans[i])
		}
		return
	}

	baselen := len(base)
	PrintRecursively(nbr, baselen, base)
}

func PrintRecursively(nbr, baselen int, base string) {
	if nbr >= baselen {
		PrintRecursively(nbr/baselen, baselen, base)
	}
	i := nbr % baselen
	if i < 0 {
		i += baselen
	}
	z01.PrintRune(rune(base[i]))
}
