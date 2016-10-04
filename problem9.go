/*
	Problem:
	There exists exactly one Pythagorean triplet for which a + b + c = 1000.
	Find the product abc.
*/

package main

import (
	"fmt"
	"time"
)


func main() {

	// measure execution time
	start := time.Now()

	doBreak := false
	for a := 1; a <= 997; a++ {
		var i int = 1

		// remark: there is one thing that can never be possible: a odd, and b odd
		if a%2 != 0 {
			i = 2
		}

		// we have a < b < c
		for b := a+1; b < (1000-a)/2; b+= i {
			c := 1000 - a - b

			if a*a + b*b == c*c {
				fmt.Println("a: ", a, "b: ", b, "c: ", c)
				fmt.Println("a*b*c = ", a * b * c)
				doBreak = true
				break
			}
		}

		if doBreak {
			break
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Took ", elapsed)
}