package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	res := make([]int, max-min)
	for i := range res {
		res[i] = min + i
	}
	return res
}
