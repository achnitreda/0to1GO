package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	params := os.Args
	if len(params) != 2 {
		fmt.Printf("ERROR: cannot convert to roman digit\n")
		return
	}
	n,err := strconv.Atoi(params[1])
	if err != nil || 0 >= n || n >= 4000 {
		fmt.Printf("ERROR: cannot convert to roman digit\n")
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	res := ""
	for _,conversion := range conversions{
		for n >= conversion.value{
			res += conversion.digit
			n -= conversion.value
		}
	}
	fmt.Printf("%s\n",res)
}