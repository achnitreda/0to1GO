package piscine

func LoafOfBread(str string) string {
	j := 0
	var res []byte
	for i := 0; i < len(str); i++ {
		if j < 5 && str[i] == ' ' {
			continue
		}
		if j == 5 {
			if i != len(str)-1 && str[i+1] == ' ' {
				continue
			}
			if i == len(str)-1 {
				break
			}
			res = append(res, ' ')
			j = 0
			continue
		}
		res = append(res, str[i])
		j++
	}
	return string(res) + "\n"
}
