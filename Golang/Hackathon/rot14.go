package piscine

func Rot14(s string) string {
	var res []rune
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			res = append(res, rune((runes[i]-'a'+14)%26+'a'))
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			res = append(res, rune((runes[i]-'A'+14)%26+'A'))
		} else {
			res = append(res, runes[i])
		}
	}
	return string(res)
}
