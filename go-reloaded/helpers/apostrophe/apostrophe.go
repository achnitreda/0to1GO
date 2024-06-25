package apostrophe

import (
	// "fmt"
	"strings"
	"unicode"
)

func FixApostrophes(text string) string {
	var x []int
	var res string
	// fmt.Printf("text:%v\n", text)
	if len(text) == 0{
		return ""
	}
	if text == "'"{
		return text
	}
	for i := 0; i < len(text); i++ {
        if i == 0 && text[i] == '\'' {
            x = append(x, i)
        }
        if i == len(text)-1 && text[i] == '\'' {
            x = append(x, i)
        }
        if i-1 >= 0 && i+1 <= len(text)-1 {
            if text[i] == '\'' && !(isValidApostrophe(text[i-1]) && isValidApostrophe(text[i+1])) {
                x = append(x, i)
            }
        }
    }
	// fmt.Printf("x -> %v\n", x)
	if len(x) > 0 {
		if x[0] != 0{
			res += text[:x[0]]
		} 
		for i := 0; i < len(x)-1; i++ {
			if i%2 == 0{
				res += "'"+strings.TrimSpace(text[x[i]+1:x[i+1]])+"'"
			}else{
				res += text[x[i]+1:x[i+1]]
			}
		}
		if len(x)%2 == 0 {
			res += text[x[len(x)-1]+1:]
		} else {
			res += text[x[len(x)-1]:]
		}
	}else {
		return text
	}
	return res
}

func isValidApostrophe(r byte) bool {
    return unicode.IsLetter(rune(r)) || unicode.IsDigit(rune(r))
}
