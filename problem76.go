/*
	Problem 76
*/

package main

import "fmt"

// save congruences of p(n) mod 10^6
var ps [101]int
var pentagonalNums [101]int

func makePentagonalNums() {
	for k := 1; k < 50; k++ {
		// -k is at index k + 50
		pentagonalNums[k] = (k*(3*k - 1))/2
		pentagonalNums[50 + k] = (k*(3*k + 1))/2
	}
}

func p(n int) int {
	k := 0
	signum := -1
	theP := 0
	for {
		k++
		signum *= -1

		penta_pos := pentagonalNums[k]
		penta_neg := pentagonalNums[50+k]

		if n - penta_pos < 0 {
			break
		} else {
			theP = (theP + signum * ps[n-penta_pos])
			
			if n - penta_neg >= 0 {
				theP = (theP + signum * ps[n-penta_neg])
			}
		}
	}

	return theP
}

func main() {
	makePentagonalNums()

	ps = [101]int{}
	ps[0] = 1

	for n := 1; n<=100; n++ {
		ps[n] = p(n)
	}

	fmt.Println(ps[100] - 1)
}