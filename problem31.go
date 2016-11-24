/*
	Problem 31
*/

package main

import (
	"fmt"
)

func getTotalSum(ofSlice [8]int, coinValues *[8]int) int {
	sum := 0
	for i, v := range ofSlice {
		sum += (*coinValues)[i] * v
	}

	return sum
}

func fillUp(coinValues *[8]int, indexToUse int, results *[][8]int, sliceToFill [8]int) {
	// 1. Get how many pences are left to fill up, and get the pence value we use to fill up
	penceValue := (*coinValues)[indexToUse]
	pencesLeft := 200 - getTotalSum(sliceToFill, coinValues)
	
	// 2. Get the maximum number of coins we may use to reach a maximum of 200 pences (this is floored)
	var maximumNum int = pencesLeft / penceValue

	// 3. For each value from 0 up to `maximumNum`, make new arrays and fill them up
	for i := 0; i<= maximumNum; i++ {
		newSlice := [8]int{}
		// fill it up with the current array
		for j, v := range sliceToFill {
			newSlice[j] = v
		}

		newSlice[indexToUse] = i

		// continue with new type of coin, or – if we are at 200 – finish
		if indexToUse < 7 {
			fillUp(coinValues, indexToUse + 1, results, newSlice)
		} else if getTotalSum(newSlice, coinValues) == 200 {
			*results = append(*results, newSlice)
		}
	}
}

func main() {
	
	// these are the possible pence values
	coinValues := [8]int{1,2,5,10,20,50,100,200}
	results := [][8]int{}

	// Push an empty slice into `fillUp`, from which it propagates

	fillUp(&coinValues, 0, &results, [8]int{})

	fmt.Println(len(results))
}