package piscine

func convertNumToBase(n int) []byte {
	base := "01"
	baselen := len(base)
	var res []byte
	if n < 0 {
		n = -n
	}
	for n >= baselen {
		i := n % baselen
		if i < 0 {
			i += baselen
		}
		res = append(res, base[i])
		n /= baselen
	}
	return res
}

func ActiveBits(n int) int {
	ans := convertNumToBase(n)
	count := 1
	for i := 0; i < len(ans); i++ {
		if ans[i] == '1' {
			count++
		}
	}
	return count
}
