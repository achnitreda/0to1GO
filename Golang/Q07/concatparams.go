package piscine

func ConcatParams(args []string) string {
	var res []byte
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			res = append(res, args[i][j])
		}
		if i != len(args)-1 {
			res = append(res, '\n')
		}
	}
	ans := string(res)
	return ans
}
