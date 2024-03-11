package piscine

func Capitalize(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] = runes[i] - 'A' + 'a'
		}
	}

	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] = runes[0] - 'a' + 'A'
	}

	for i := 0; i < len(runes)-1; i++ {
		if (runes[i] < 'A' || runes[i] > 'Z') && (runes[i] < 'a' || runes[i] > 'z') && (runes[i] < '0' || runes[i] > '9') {
			if runes[i+1] >= 'a' && runes[i+1] <= 'z' {
				runes[i+1] = runes[i+1] - 'a' + 'A'
			}
		}
	}

	res := string(runes)
	return res
}
