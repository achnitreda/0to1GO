package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}
	var res []int
	for max > min {
		res = append(res, max)
		max--
	}
	return res
}
