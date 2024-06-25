package converter

import (
	"fmt"
	"go-reloaded/helpers/atoan"
	"go-reloaded/helpers/base"
	"go-reloaded/helpers/modifierWithNum"
	"go-reloaded/helpers/modifiers"
	"go-reloaded/helpers/punctuation"
	"go-reloaded/helpers/apostrophe"
	"regexp"
	"strings"
)

func ConvertBase(text string) string {
	res := strings.Fields(text)
	fmt.Printf("%q\n", res)

	text = punctuation.FormatPunctuation(text)
	text = punctuation.FormatWordPunctuation(text)
	text = punctuation.NoPunctuationNumber(text)
	text = punctuation.NoPunctuationBNumber(text)

	type functionCondition struct {
		condition func(string) bool
		function  func(string) string
	}
	
	conditions := []functionCondition{
		{func(s string) bool {
			state, _ := regexp.MatchString(`\(hex\)`, s)
			return state
		}, func(s string) string {
			return base.ConvertHexToDecimal(base.ConvertHexBToDecimal(s))
		}},
		{func(s string) bool {
			state, _ := regexp.MatchString(`\(bin\)`, s)
			return state
		}, func(s string) string {
			return base.ConvertBinToDecimal(base.ConvertBinBToDecimal(s))
		}},
		{func(s string) bool {
			matchedCaseN, _ := regexp.MatchString(`\((low|up|cap),`, s)
			return matchedCaseN
		}, func(s string) string {
			return modifierWithNum.ModifyWordsWithNumber(modifierWithNum.ApplyModifNum(s)) 
		}},
		{func(s string) bool {
			matchedCase, _ := regexp.MatchString(`\((low|up|cap)\)`, s)
			return matchedCase
		}, func(s string) string {
			return modifiers.ConvertCase(modifiers.ApplyModifiers(s))
		}},
		{func(s string) bool {
			matchedCaseA, _ := regexp.MatchString(`[aA]$`, s)
			return matchedCaseA
		}, func(s string) string {
			return atoan.AAA(atoan.FixIndefiniteArticles(s))
		}},
		{func(s string) bool {
			state, _ := regexp.MatchString(`[']`, s)
			return state
		}, func(s string) string {
			return apostrophe.FixApostrophes(s)
		}},
	}
	
	for _, r := range res {
		for _, cond := range conditions {
			if cond.condition(r) {
				text = cond.function(text)
				break
			}
		}
	}	

	return text
}
