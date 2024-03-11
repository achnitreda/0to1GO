package piscine

func BasicAtoi2(s string) int {
	runes := []rune(s)
	allDigits := true
	var nb int
	for _, r := range runes {
		if 48 > r || r > 57 {
			allDigits = false
			return 0
		}
	}
	if allDigits {
		for _, r := range runes {
			digit := int(r - '0')
			nb = nb*10 + digit
		}
	}
	return nb
}
