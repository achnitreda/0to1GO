package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var res []int
	for min < max {
		res = append(res, min)
		min++
	}
	return res
}
