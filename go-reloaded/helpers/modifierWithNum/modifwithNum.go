package modifierWithNum

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ApplyModifNum(text string) string {
    words := strings.Fields(text)
    fmt.Printf("%q\n", words)

    for i:=0; i<len(words); i++{
        if words[i] == "(cap," {
            regex := regexp.MustCompile(`(^[0-9]+\))`)
            if i+1 <= len(words)-1 {
                if regex.MatchString(words[i+1]) && i > 0 && words[i-1] != "" {
                    // fmt.Printf("%d\n", index)
                    numWords, _ := strconv.Atoi(string(words[i+1][:len(words[i+1])-1])) // 4)i => 0
                    // fmt.Printf("%d\n", numWords)
                    if numWords == 0 {
                        break
                    }
                    if numWords > len(words[:i]) {
                        numWords = len(words[:i])
                    }
                    for j := len(words[:i]) - numWords; j < len(words[:i]); j++ {
                        words[j] = strings.ToUpper(string(words[j][0])) + strings.ToLower(words[j][1:])
                    }
                    words = append(words[:i], words[i+2:]...)
                    i--
                }
            }
        } else if words[i] == "(up," {
            regex := regexp.MustCompile(`(^[0-9]+\))`)
            if i+1 <= len(words)-1 {
                if regex.MatchString(words[i+1]) && i > 0 && words[i-1] != ""{
                    // fmt.Printf("%d\n", index)
                    numWords, _ := strconv.Atoi(string(words[i+1][:len(words[i+1])-1]))
                    // fmt.Printf("%d\n", numWords)
                    if numWords == 0 {
                        break
                    }
                    if numWords > len(words[:i]) {
                        numWords = len(words[:i])
                    }
                    for j := len(words[:i]) - numWords; j < len(words[:i]); j++ {
                        words[j] = strings.ToUpper(string(words[j]))
                    }
                    words = append(words[:i], words[i+2:]...)
                    i--
                }
            }
        } else if words[i] == "(low," {
            regex := regexp.MustCompile(`(^[0-9]+\))`)
            if i+1 <= len(words)-1 && i > 0 && words[i-1] != "" {
                if regex.MatchString(words[i+1]){
                    numWords, _ := strconv.Atoi(string(words[i+1][:len(words[i+1])-1]))
                    if numWords == 0 {
                        break
                    }
                    if numWords > len(words[:i]) {
                        numWords = len(words[:i])
                    }
                    for j := len(words[:i]) - numWords; j < len(words[:i]); j++ {
                        words[j] = strings.ToLower(string(words[j]))
                    }
                    words = append(words[:i], words[i+2:]...)
                    i--
                }
            }
        }
    }
    return strings.Join(words, " ")
}

func ModifyWordsWithNumber(text string) string {
	numModifierRegex := regexp.MustCompile(`([\w\s'.,!?;:]+)\s+\((low|up|cap),\s(\d+)\)`)

	return numModifierRegex.ReplaceAllStringFunc(text, func(match string) string {
		fmt.Printf("match -> %q\n", match)
		parts := numModifierRegex.FindStringSubmatch(match)
		fmt.Printf("#Parts -> %q\n", parts)
		// words := strings.Split(parts[1][:len(parts[1])-1], " ")
		re := regexp.MustCompile(`[ \n]`)
		words := re.Split(parts[1][:len(parts[1])], -1)	
		// fmt.Printf("#words -> %q\n", words)
		modifier := parts[2]
		numWords, _ := strconv.Atoi(parts[3])

		if numWords > len(words) {
			numWords = len(words)
		}

		for i := len(words) - numWords; i < len(words); i++ {
			switch modifier {
			case "low":
				words[i] = strings.ToLower(words[i])
			case "up":
				words[i] = strings.ToUpper(words[i])
			case "cap":
				if words[i] == "" {
					continue
				} else {
					words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
				}
			}
		}
		return strings.Join(words, " ")
	})
}
