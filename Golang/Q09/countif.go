package piscine

func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			count++
		}
	}
	return count
}
