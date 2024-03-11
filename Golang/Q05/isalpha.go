package piscine

func IsAlpha(s string) bool {
	for _, char := range s {
		if !isAlphanumerical(char) && char != ' ' {
			return false
		}
	}
	return true
}

func isAlphanumerical(char rune) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9')
}
