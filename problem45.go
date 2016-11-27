/*
	Problem 45
*/

package main

import (
	"fmt"
	"math"
)


func main() {

	var k float64 = 143

	for {
		k++
		var n float64 = 2*k + 1
		// given an n that generates a triangular number, see if there is an integer solution that also generates the same 
		// Pentagonal: http://www.wolframalpha.com/input/?i=solve:+x*(3x-1)+%3D+n*(n%2B1)
		// Hexagonal: http://www.wolframalpha.com/input/?i=solve:+2*x*(2x-1)+%3D+n*(n%2B1): this is always (n+1)/2
		// Thus it suffices to check only the odd ones

		pentaX := (math.Sqrt(12*n*n + 12*n + 1) + 1)/6
		
		if pentaX - math.Floor(pentaX) < 0.000001 {
			fmt.Println((int(n)*(int(n) + 1))/2)
			break
		}
	}	

}