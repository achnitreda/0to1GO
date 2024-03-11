package piscine

func SplitSpace(s string) []string {
	var res []string
	start := 0
	end := 0
	for ; end < len(s); end++ {
		if s[end] == ' ' {
			res = append(res, s[start:end])
			start = end + 1
		}
	}
	res = append(res, s[start:])
	return res
}

func ShoppingSummaryCounter(str string) map[string]int {
	items := SplitSpace(str)

	mp := make(map[string]int)

	for _, item := range items {
		mp[item]++
	}

	return mp
}
