package piscine

// import "fmt"

func f(a, b int) int {
	return a - b
}

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) == 1 {
		return true
	}
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			for j := i + 1; j < len(a)-1; j++ {
				if f(a[j], a[j+1]) < 0 {
					return false
				}
			}
			return true
		}
		if f(a[i], a[i+1]) < 0 {
			for j := i + 1; j < len(a)-1; j++ {
				if f(a[j], a[j+1]) > 0 {
					return false
				}
			}
			return true
		}
		if f(a[i], a[i+1]) == 0 {
			return true
		}
	}
	return false
}

// func main(){
// 	a := []int{-452270, -851482, -146241, -307938, 570992, -734673, 697071, -674243}
// 	fmt.Println(IsSorted(Compare,a))
// }
