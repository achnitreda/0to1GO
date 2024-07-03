package methods

import (
	"errors"
	"log"
	"strings"
)

var (
	errWidthOfTerminal     = errors.New("graph width of string than  graph width of terminal")
	widthOfTerminal    int = TerminalWidth()
)

func ProccessOutput(align, str string, graphics map[byte][]string) string {
	input := ProccessTheInput(str)
	var graphicWidth int
	var spaces string
	const graphHeight int = 8
	output := ""
	for _, element := range input {
		if align != "" {
			element = strings.TrimSpace(element)
			graphicWidth = GraphicWidth(element, align, graphics)
			spaces = SpacesToAdd(graphicWidth, align, element)
		}
		if element == "" {
			output += "\n"
		} else {
			var graphicLine string
			for i := 0; i < graphHeight; i++ {
				for j := 0; j < len(element); j++ {
					if j == len(element)-1 {
						graphicLine += graphics[element[j]][i] + "\n"
					} else {
						if element[j] == ' ' && align == "justify" {
							if element[j-1] != ' ' {
								graphicLine += spaces
							}
						} else {
							graphicLine += graphics[element[j]][i]
						}
					}
				}
				if align == "center" || align == "right" {
					output += spaces + graphicLine
					graphicLine = ""
				} else {
					output += graphicLine
					graphicLine = ""
				}
			}
		}
	}
	return output
}

func GraphicWidth(str, align string, grafphics map[byte][]string) int {
	grafphicline := ""
	for i := 0; i < len(str); i++ {
		if align == "justify" && str[i] == ' ' {
			continue
		}
		grafphicline += grafphics[str[i]][0]
	}
	return len(grafphicline)
}

func SpacesToAdd(graphicWidth int, align, str string) string {
	spaces := ""
	switch align {
	case "right":
		numberOfSpaces := widthOfTerminal - graphicWidth
		spaces = strings.Repeat(" ", numberOfSpaces)
	case "center":
		numberOfSpaces := (widthOfTerminal - graphicWidth) / 2
		spaces = strings.Repeat(" ", numberOfSpaces)
	case "justify":
		if graphicWidth > widthOfTerminal {
			log.Fatal(errWidthOfTerminal)
		}
		numberOfSpaces := (widthOfTerminal - graphicWidth) / (len(strings.Fields(str)) - 1)
		spaces = strings.Repeat(" ", numberOfSpaces)
	}
	return spaces
}
