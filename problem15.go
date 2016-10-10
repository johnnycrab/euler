/*
	Problem 15:
	https://projecteuler.net/problem=15
*/

package main

import "fmt"

// maps a point in the 20x20 grid (that is, we have (i,j) with i,j \in {0,...,20}) to an array index
func m(i,j int) int {
	return i*21 + j
}

func main() {

	// each point in the 20x20 grid can only be reached from above and / or from the left
	// so we only need to iterate over all points and add the number of ways from above and left

	ways := [21*21]int{}

	// first row all have 1 way to come from
	for i:=0; i<= 20; i++ {
		ways[m(0,i)] = 1
	}

	// now iterate over all rows
	for i:=1; i<= 20; i++ {
		for j:= 0; j<= 20; j++ {
			// point in the middle
			if j > 0 {
				ways[m(i,j)] = ways[m(i,j-1)] + ways[m(i-1,j)]
			} else {
				// point at left edge
				ways[m(i,j)] = ways[m(i-1,j)]
			}
			
		}
	}



	fmt.Println(ways[21*21 - 1])
}