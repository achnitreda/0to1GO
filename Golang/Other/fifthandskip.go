package piscine

func FifthAndSkip(str string) string {
	if str == ""{
		return "\n"
	}
	if len(str) < 5{
		return "Invalid Input\n"
	}
	j := 0
	var res []rune
	runes := []rune(str)
	for i:=0; i<len(runes);i++{
		if runes[i] == ' '{
			continue
		}else {
			if (j+1)%6 == 0{
				runes[i] = ' '
			}
			res = append(res,runes[i])
			j++
		}
	}
	return string(res) + "\n"
}