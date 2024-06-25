package punctuation

import (
	// "fmt"
	"regexp"
	"strings"
)

func FormatPunctuation(text string) string {
	punctuationRegex := regexp.MustCompile(`\s*([.,!?:;])`)

	return punctuationRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := punctuationRegex.FindStringSubmatch(match)
		return parts[1]
	})
}

func FormatWordPunctuation(text string) string {
	wPunctuationRegex := regexp.MustCompile(`([-"-+\/-9<->@-~]+)[\s]*([.,!?;:]+)[\s]*`)

	return wPunctuationRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := wPunctuationRegex.FindStringSubmatch(match)
		if strings.HasSuffix(text, match) {
			return parts[1] + parts[2]+ "\n"
		}
		return parts[1] + parts[2] + " "
	})
}

func NoPunctuationNumber(text string) string { // OPTIONAL 	??
	punctuationRegex := regexp.MustCompile(`\s(\d+[.,])(\s\d)`)
	// fmt.Printf("%q\n", text)
	return punctuationRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := punctuationRegex.FindStringSubmatch(match)
		return " "+parts[1]+parts[2][1:]
	})
}

func NoPunctuationBNumber(text string) string { // OPTIONAL 	??
	punctuationRegex := regexp.MustCompile(`^(\d+[.,])(\s\d)`)
	// fmt.Printf("%q\n", text)
	return punctuationRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := punctuationRegex.FindStringSubmatch(match)
		return parts[1]+parts[2][1:]
	})
}