/*
	Problem 75


	UNSOLVED YET!!
*/

package main

import (
	"fmt"
	"math"
)

const n = 1500 * 1000

// we keep track of in how many ways a length can be formed into a right-angled triangle
var lengths [n]int

func main() {

	fmt.Println("foo")
	// we iterate over all integer values and count up the lengths of the right sided triangle they form
	// a triangle is of the form (a,b,c) with a < b < c
	// a = 439339 is the highest value that `a` can reach

	for a := 1; a <= 439339; a++ {
		b := a
		
		aDivByThree := a%3 == 0

		for {

			// calculate c
			pyth := a*a + b*b

			c := int(math.Sqrt(float64(pyth)))

			length := a+b+c
			if length <= n {
				// length is still in range. check if we have a right triangle with all integer sides
				if pyth == c*c {
					lengths[length-1]++
				}
			} else {
				break
			}
			
			b++
		}
	}

	count := 0
	for _, v := range lengths {
		if v == 1 {
			count++
		}
	}

	fmt.Println(count)

}