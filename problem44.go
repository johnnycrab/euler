/*
	Problem 44
*/

package main

import (
	"fmt"
	"math"
)

// n and m are the generators
// see the equation http://www.wolframalpha.com/input/?i=solve:+3x%5E2+-+x+%3D+3*(n%5E2%2Bm%5E2)+-+(n%2Bm)
// and http://www.wolframalpha.com/input/?i=solve:+3x%5E2+-+x+%3D+3*(n%5E2-m%5E2)+-+(n-m)

// the solutions must be integers
func sumAndDifferenceArePentagonal(n float64, m float64) bool {
	var mSq float64 = m*m
	var nSq float64 = n*n
	x1 := (math.Sqrt(36*mSq - 12*m - 36*nSq + 12*n + 1) + 1)/6
	x2 := (math.Sqrt(36*mSq - 12*m + 36*nSq - 12*n + 1) + 1)/6

	almostZero := 0.00001
	// check if they are approximately integers
	if (x1 - math.Floor(x1) < almostZero) && (x2 - math.Floor(x2) < almostZero) {
		return true
	}

	return false
}


func main() {

	var n float64 = 1
	var m float64 = 1


	for {
		m++

		if sumAndDifferenceArePentagonal(n,m) {
			fmt.Println(int(m*(3*m-1)-n*(3*n-1))/2)
			break
		}

		if (m-n) > 100000 {
			n++
			m = n
		}


	}
}