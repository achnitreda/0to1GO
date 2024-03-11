package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			runes[i] = runes[i] - 'A' + 'a'
		}
	}
	ans := string(runes)
	return ans
}
