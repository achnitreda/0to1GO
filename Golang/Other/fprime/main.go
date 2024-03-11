package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	isPrime := true
	var res []int
	if n == 0 || n == 1 {
		return
	}
	if n%2 == 0 {
		isPrime = false
		res = append(res, 2)
		n /= 2
	}
	j := 3
	for n > 1 {
		if n%j == 0 {
			isPrime = false
			res = append(res, j)
			n /= j
		} else {
			j += 2
		}
	}
	if isPrime {
		fmt.Println(n)
	}

	for i := 0; i < len(res)-1; i++ {
		char := strconv.Itoa(res[i])
		fmt.Print(char)
		fmt.Print("*")
	}
	fmt.Println(strconv.Itoa(res[len(res)-1]))
}
