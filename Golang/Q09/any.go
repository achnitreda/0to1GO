package piscine

// func IsNumeric(s string) bool {
// 	for i := 0; i < len(s); i++ {
// 		if s[i] < '0' || s[i] > '9' {
// 			return false
// 		}
// 	}
// 	return true
// }

func Any(f func(string) bool, a []string) bool {
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			return true
		}
	}
	return false
}
