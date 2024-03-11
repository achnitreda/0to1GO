package piscine

func StrRev(s string) string {
	runes := []rune(s)
	i := 0
	j := len(runes) - 1
	for i < j {
		temp := runes[j]
		runes[j] = runes[i]
		runes[i] = temp
		i++
		j--
	}
	s = string(runes)
	return s
}
