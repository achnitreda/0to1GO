package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for j := 1; j <= 4; j++ {
		fmt.Printf("Player %d:", j)
		for i := 0; i < 3; i++ {
			index := (j-1)*3 + i
			if i < 2 {
				fmt.Printf(" %d,", deck[index])
			} else {
				fmt.Printf(" %d\n", deck[index])
			}
		}
	}
}
