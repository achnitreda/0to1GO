package piscine

func StringToIntSlice(str string) []int {
	ans := []rune(str)
	var res []int
	for _, r := range ans {
		res = append(res, int(r))
	}
	return res
}
