package piscine

func ReverseMenuIndex(menu []string) []string {
	reversed := make([]string, len(menu))
	for i, j := len(menu)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversed[j] = menu[i]
	}
	return reversed
}
