package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	res := 1
	for i := res; i <= nb; i++ {
		if res > (1<<63-1)/nb {
			return 0
		}
		res *= i
	}
	return res
}
