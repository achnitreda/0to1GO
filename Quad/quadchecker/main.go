package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Counter(inputStr string) (x, y int) {
	flag := true
	for i := 0; i < len(inputStr); i++ {
		if inputStr[i] == '\n' {
			y++
			flag = false
		} else if flag == true {
			x++
		}
	}
	return x, y
}

func quadA(x, y int) string {
	rs := ""
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i != 0 && i != y-1 {
				if j == 0 || j == x-1 {
					rs += "|"
				} else {
					rs += " "
				}
			} else if j == 0 || j == x-1 {
				rs += "o"
			} else {
				rs += "-"
			}
		}
		rs += "\n"
	}
	return rs
}

func quadB(x, y int) string {
	rs := ""
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i != 0 && i != y-1 {
				if j == 0 || j == x-1 {
					rs += "*"
				} else {
					rs += " "
				}
			}
			if i == 0 && j == 0 {
				rs += "/"
			} else if (j == 0 && i == y-1) || (i == 0 && j == x-1) {
				rs += "\\"
			} else if i == y-1 && j == x-1 {
				rs += "/"
			} else if i == 0 || i == y-1 {
				rs += "*"
			}
		}
		rs += "\n"
	}
	return rs
}

func quadC(x, y int) string {
	rs := ""
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i != 0 && i != y-1 {
				if j == 0 || j == x-1 {
					rs += "B"
				} else {
					rs += " "
				}
			} else if i == 0 && (j == 0 || j == x-1) {
				rs += "A"
			} else if i == y-1 && (j == 0 || j == x-1) {
				rs += "C"
			} else {
				rs += "B"
			}
		}
		rs += "\n"
	}
	return rs
}

func quadD(x, y int) string {
	rs := ""
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i != 0 && i != y-1 {
				if j == 0 || j == x-1 {
					rs += "B"
				} else {
					rs += " "
				}
			}
			if (i == 0 && j == 0) || (i == y-1 && j == 0) {
				rs += "A"
			} else if i == 0 && j == x-1 {
				rs += "C"
			} else if i == y-1 && j == x-1 {
				rs += "C"
			} else if i == 0 || i == y-1 {
				rs += "B"
			}
		}
		rs += "\n"
	}
	return rs
}

func quadE(x, y int) string {
	rs := ""
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i != 0 && i != y-1 {
				if j == 0 || j == x-1 {
					rs += "B"
				} else {
					rs += " "
				}
			}
			if i == 0 && j == 0 {
				rs += "A"
			} else if i == 0 && j == x-1 {
				rs += "C"
			} else if i == y-1 && j == 0 {
				rs += "C"
			} else if i == y-1 && j == x-1 {
				rs += "A"
			} else if i == 0 || i == y-1 {
				rs += "B"
			}
		}
		rs += "\n"
	}
	return rs
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	inputStr := ""
	for reader.Scan() {
		inputStr += reader.Text() + "\n"
	}
	if inputStr == ""{
		fmt.Println("Not a quad function")
		return
	}
	x, y := Counter(inputStr)
	strX := strconv.Itoa(x)
	strY := strconv.Itoa(y)

	str := []string{}

	if quadA(x, y) == inputStr {
		str = append(str, "[quadA]"+" "+"["+strX+"]"+" "+"["+strY+"]")
	}
	if quadB(x, y) == inputStr {
		str = append(str, "[quadB]"+" "+"["+strX+"]"+" "+"["+strY+"]")
	}
	if quadC(x, y) == inputStr {
		str = append(str, "[quadC]"+" "+"["+strX+"]"+" "+"["+strY+"]")
	}
	if quadD(x, y) == inputStr {
		str = append(str, "[quadD]"+" "+"["+strX+"]"+" "+"["+strY+"]")
	}
	if quadE(x, y) == inputStr {
		str = append(str, "[quadE]"+" "+"["+strX+"]"+" "+"["+strY+"]")
	}
	if len(str) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	for i := 0; i < len(str)-1; i++ {
		fmt.Print(str[i] + " || ")
	}
	fmt.Println(str[len(str)-1])
}
