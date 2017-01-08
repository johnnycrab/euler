/*
	Problem 92
*/

package main

import (
	"fmt"
)

func numberChainArrivesAt89(startingPoint int) bool {
	number := startingPoint

	for !(number == 1 || number == 89) {
		digitSquareSum := 0
		for number > 0 {
			mod := number%10
			digitSquareSum += mod*mod
			number = (number - mod)/10
		}
		
		number = digitSquareSum
	}

	return number == 89
}

func main() {
	result := 0
	for i := 2; i < 10000000; i++ {
		if numberChainArrivesAt89(i) {
			result++
		}
	}

	fmt.Println(result)
}