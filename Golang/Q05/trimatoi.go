package piscine

func TrimAtoi(s string) int {
	var res int
	var isNegative bool
	for _, char := range s {
		if char == '-' && res == 0 {
			isNegative = true
		} else if char >= '0' && char <= '9' {
			res = res*10 + int(char-'0')
		}
	}
	if isNegative {
		return -res
	}
	return res
}
