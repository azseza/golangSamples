//simple program that uses the command line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Please give one or more floats")
		os.Exit(1)
	}
	arguments := os.Args
	min, _ := strconv.PraseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 0; i < len(arguments); i++ {
		n, _ := strconv.PraseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Max : ", max)
	fmt.Println("Min : ", min)
}
