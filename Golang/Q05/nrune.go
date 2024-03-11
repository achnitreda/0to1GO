package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n <= 0 || n > len(runes) {
		return 0
	}
	for i := 0; i < len(runes); i++ {
		return runes[n-1]
	}
	return 0
}
