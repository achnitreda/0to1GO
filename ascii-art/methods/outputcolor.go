package methods

import (
	"log"
	"strings"
)

func ProccessOutputColor(color, subStr, str string, graphics map[byte][]string) string {
	if !strings.Contains(str, subStr) {
		log.Fatal("the substring provided do not match charachers of the original string\n Provide a valid substring")
	}
	input := ProccessTheInput(str)

	const graphHeight int = 8
	output := ""

	for _, element := range input {
		if element == "" {
			output += "\n"
		} else {
			var graphicLine string
			for i := 0; i < graphHeight; i++ {
				j := 0
				for j < len(element) {
					if subStr == "" {
						graphicLine += Colorize(graphics[element[j]][i], color)
						j++
					} else {
						// Check if we have a substring match starting at this position
						if j+len(subStr) <= len(element) && element[j:j+len(subStr)] == subStr {
							// Color the entire matching substring
							for k := 0; k < len(subStr); k++ {
								graphicLine += Colorize(graphics[element[j+k]][i], color)
							}
							// Skip the characters we've just colored
							j += len(subStr)
						} else {
							graphicLine += graphics[element[j]][i]
							j++
						}
					}
				}
				graphicLine += "\n"
			}
			output += graphicLine
			graphicLine = ""
		}
	}
	return output
}
