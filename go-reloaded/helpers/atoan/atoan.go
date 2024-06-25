package atoan

import (
	// "fmt"
	"regexp"
	"strings"
)

func FixIndefiniteArticles(text string) string {
	combinedRegex := regexp.MustCompile(`^(\W?)+[aA]\s*[aeiouhAEIOUH]`) // ((a orange))

	return combinedRegex.ReplaceAllStringFunc(text, func(match string) string {
		// fmt.Printf("%q\n",match)
		if strings.HasPrefix(match, "A") {
			return strings.Replace(match, "A ", "An ", 1)
		} else {
			// fmt.Printf("match -> %q\n",match)
			return strings.Replace(match, "a ", "an ", 1)
		}
	})
}

func ContainsRune(str string) bool{
	s := []rune(str)
	if str == ""{
		return false
	}
	runes := []rune{'a', 'e', 'i', 'o', 'u', 'h','A','E','I','O','U','H'}
	for _, rune :=range runes {
		if s[0] == rune{
			return true
		}
	}
	return false
}

func AAA(text string) string {
    words := strings.Split(text, " ")
	// fmt.Printf("%q\n",words)
    for i:=0; i<len(words)-1; i++ {
        if words[i] == "a" && ContainsRune(words[i+1])  {
            words[i] = "an"
        }
		if words[i] == "A" && ContainsRune(words[i+1])  {
            words[i] = "An"
        }
    }
    return strings.Join(words, " ")
}

