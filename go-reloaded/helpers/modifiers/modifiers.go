package modifiers

import (
	// "fmt"
	"strings"
	"regexp"
)

func ApplyModifiers(text string) string {
	words := strings.Fields(text)
	// fmt.Printf("%q\n", words)

	for i := 0; i < len(words); i++ {
		switch {
		case words[i] == "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i-- 
			}
		case words[i] == "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i-- 
			}
		case words[i] == "(cap)":
			if i > 0 {
				if words[i-1] != "" {
					words[i-1] = strings.ToUpper(string(words[i-1][0])) + strings.ToLower(words[i-1][1:])
				}
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(words, " ")
}

func ConvertCase(text string) string {
	regex := regexp.MustCompile(`([!-~]+)\s+((\(up\)|\(low\)|\(cap\))[.,!?:;\s])`)
	
	return regex.ReplaceAllStringFunc(text, func(match string) string {
		matches := regex.FindStringSubmatch(match)
		// fmt.Printf("red -> %q\n", matches)
		str := matches[1]
		// fmt.Printf("-> %q\n", str)
		// fmt.Printf(" m2 -> %q\n", matches[2])

		switch matches[3] {
		case "(up)":
			return strings.ToUpper(str) + string(matches[2][len(matches[2])-1])
		case "(low)":
			return strings.ToLower(str) + string(matches[2][len(matches[2])-1])
		case "(cap)":
			return strings.ToUpper(string(str[0])) + strings.ToLower(str[1:]) + string(matches[2][len(matches[2])-1])
		default:
			return str 
		}
	})
}
