package methods

func ProccessOutput(str string, graphics map[byte][]string) string {
	const graphHeight int = 8
	input := ProccessTheInput(str)
	output := ""

	for _, element := range input {
		if element == "" {
			output += "\n"
		} else {
			for i := 0; i < graphHeight; i++ {
				for j := 0; j < len(element); j++ {
					if j == len(element)-1 {
						output += graphics[element[j]][i] + "\n"
					} else {
						output += graphics[element[j]][i]
					}
				}
			}
		}
	}
	return output
}
