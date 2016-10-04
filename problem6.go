/*
	Problem:
	Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import "fmt"

func main() {
	sumSquare := 0
	for i := 1; i<= 100; i++ {
		sumSquare += i*i
	}

	// small Gauss
	const squareSum int = 50 * 101 * 50 * 101

	fmt.Println(squareSum - sumSquare)
}