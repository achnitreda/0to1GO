package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("ERROR: Invalid Arguments\nUsage: go run your-program.go data.txt")
		os.Exit(1)
	}

	file := os.Args[1]

	// _________________________________________

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("ERROR file: ", err)
		os.Exit(1)
	}

	if len(data) == 0 {
		fmt.Println("Error: Empty DataSet")
		os.Exit(1)
	}

	strArr := strings.Split(string(data[:len(data)-1]), "\n")
	dataSet := make([]int, len(strArr))
	for i, v := range strArr {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error Atoi: ", err)
			os.Exit(1)
		}
		dataSet[i] = num
	}
	// fmt.Println(dataSet)

	// __________________________________________

	mean := (Average(dataSet))
	fmt.Println("Average:", int(math.Round(mean)))

	fmt.Println("Median:", Median(dataSet))

	V := Variance(dataSet, mean)
	fmt.Println("Variance:", V)

	// Standard deviation -> (square root) of the variance
	sd := math.Round(math.Sqrt(float64(V)))
	fmt.Println("Standard Deviation:", int(sd))
}

// Formula: (x₁ + x₂ + ... + xₙ) / n
func Average(dataSet []int) float64 {
	var sum float64
	N := float64(len(dataSet))
	for _, v := range dataSet {
		sum += float64(v)
	}
	avg := sum / N
	return avg
}

/*
sorted data
odd number [5]int -> middle num int[2]
even number [6]int -> avg of the two middle num ([2]+[3])/2
*/
func Median(dataSet []int) int {
	// sorted it
	sort.Ints(dataSet)

	med := len(dataSet) / 2

	if len(dataSet)%2 != 0 {
		return dataSet[med]
	}
	median := float64(dataSet[med-1]+dataSet[med]) / 2
	ans := math.Round(median)
	return int(ans)
}

/*
σ² = Σ(x - μ)² / N
x -> each value in dataset
μ ->  the mean
N -> entire population
N-1 -> num of values for sample data
  - why (x - μ)² ??
    -> we (square) the differences to avoid negative values,
    but we could we don't use the (absolute) because it less sensitive to outliers.
*/
func Variance(dataSet []int, mean float64) int {
	N := len(dataSet)
	var sum float64
	for _, v := range dataSet {
		diff := float64(v) - mean
		sum += diff * diff
	}
	v := math.Round(sum / float64(N))
	return int(v)
}
