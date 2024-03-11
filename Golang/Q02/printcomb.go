package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var v [5]rune
	v[0] = '0'
	v[3] = ','
	v[4] = ' '
	for v[0] <= '7' {
		v[1] = v[0] + 1
		for v[1] <= '8' {
			v[2] = v[1] + 1
			for v[2] <= '9' {
				z01.PrintRune(v[0])
				z01.PrintRune(v[1])
				z01.PrintRune(v[2])
				if v[0] == '7' && v[1] == '8' && v[2] == '9' {
					z01.PrintRune('\n')
					break
				}
				z01.PrintRune(v[3])
				z01.PrintRune(v[4])
				v[2]++
			}
			v[1]++
		}
		v[0]++
	}
}
