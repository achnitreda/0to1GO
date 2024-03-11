package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	return runes[len(s)-1]
}
