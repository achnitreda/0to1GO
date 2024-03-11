package piscine

func JumpOver(str string) string {
	var res []byte
	for i := 0; i < len(str); i++ {
		if (i+1)%3 == 0 {
			res = append(res, str[i])
		}
	}
	return string(res) + "\n"
}
