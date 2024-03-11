package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for i := 0; i < len(a); i += 2 {
		if i == len(a)-1 {
			return a[i]
		}
		if i+1 < len(a) && a[i] != a[i+1] {
			return a[i]
		}
	}
	return -1
}
