package piscine

func SplitWhiteSpaces(s string) []string {
	var res []string
	start := 0
	end := 0
	for ; end < len(s); end++ {
		if s[end] == ' ' {
			if end > start {
				res = append(res, s[start:end])
			}
			start = end + 1
		}
	}
	if end > start {
		res = append(res, s[start:])
	}
	return res
}
