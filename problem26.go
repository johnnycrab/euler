package main

import (
	"fmt"
)

// returns reciprocal cycle for 1/n
func getLengthOfReciprocalCycleForFraction(n int) int {
	// we start by dividing 10 / n
	// for each iteration we store the result as well as the congruence class
	// if it repeats somewhere, we know the whole cycle will repeat

	store := [][2]int{}

	toDivide := 1

	result := -1

	for {
		var divisionResult int
		var congruence int

		toDivide *= 10

		// check if we can divide
		if toDivide < n {
			divisionResult = 0
			congruence = toDivide
		} else {
			congruence = toDivide % n
			divisionResult = (toDivide - congruence)/n
		}

		// check if we have a repetition
		for i, val := range store {
			
			if val[0] == divisionResult && val[1] == congruence {
				// we have a repetition
				result = len(store) - i
			}
		}

		if result > -1 {
			break
		}

		store = append(store, [2]int{divisionResult, congruence})
		toDivide = congruence
	}

	return result
}

func main() {
	res := 1
	cycleLen := 1

	for n := 2; n < 1000; n++ {
		length := getLengthOfReciprocalCycleForFraction(n)
		if length > cycleLen {
			res = n
			cycleLen = length
		}
	}

	fmt.Println(res)
}