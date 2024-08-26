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



// package piscine

// func FifthAndSkip(str string) string {
// 	if str == "" {
// 		return "\n"
// 	}
// 	if len(str) < 5 {
// 		return "Invalid Input\n"
// 	}
// 	runes := []rune(str)
// 	i := 0
// 	e := 0
// 	res := ""
// 	for e < len(runes) {
// 		if runes[e] != ' ' {
// 			i++
// 			res += string(runes[e])
// 		}
// 		if e == len(runes)-1{
// 			break
// 		}
// 		if e+1 < len(runes) && runes[e+1] == ' ' && i == 5 {
// 			res += " "
// 			i = 0
// 			e += 2
// 		} else {
// 			if i == 5 {
// 				res += " "
// 				i = 0
// 				e++
// 			}
// 		}
// 		e++
// 	}
// 	res += "\n"
// 	return res
// }
