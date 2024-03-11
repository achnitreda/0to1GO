package piscine

func Compact(ptr *[]string) int {
	count := 0
	var res []string
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			res = append(res, (*ptr)[i])
			count++
		}
	}
	*ptr = res
	return count
}
