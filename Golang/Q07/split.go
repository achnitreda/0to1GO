package piscine

func Split(s, sep string) []string {
	var res []string
	start := 0
	for end := 0; end < len(s)-len(sep); end++ {
		if s[end:end+len(sep)] == sep {
			res = append(res, s[start:end])
			start = end + len(sep)
			end = start
		}
	}
	res = append(res, s[start:])
	return res
}
