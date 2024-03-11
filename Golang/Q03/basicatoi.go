package piscine

func BasicAtoi(s string) int {
	var nb int
	runes := []rune(s)
	for _, r := range runes {
		digit := int(r - '0')
		nb = nb*10 + digit
	}
	return nb
}
