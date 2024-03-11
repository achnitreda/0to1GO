package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(0)
		return
	}
	if n <= 0 {
		fmt.Println(0)
		return
	}
	isPrime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
		}
	}

	if isPrime {
		res := 0
		for i := 3; i <= n; i += 2 {
			res += i
		}
		res += 2
		fmt.Println(res)
	} else {
		fmt.Println(0)
	}
}
