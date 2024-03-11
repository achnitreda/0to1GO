package piscine

func Join(strs []string, sep string) string {
	var res string
	for i := 0; i < len(strs); i++ {
		if i < len(strs)-1 {
			res += strs[i] + sep
		} else if i == len(strs)-1 {
			res += strs[i]
		}
	}
	return res
}
