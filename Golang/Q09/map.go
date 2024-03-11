package piscine

func IsPrime(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func Map(f func(int) bool, a []int) []bool {
	var res []bool
	for i := 0; i < len(a); i++ {
		res = append(res, f(a[i]))
	}
	return res
}
