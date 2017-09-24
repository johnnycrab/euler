/*
	Problem 108
*/

package main

import (
	"fmt"
	//"./crabMath"
)




func main() {

	nSolutions := make([]int, 1000000)

	x := 4

	largest := 0

	for {
		x++

		for n := 4; n < x; n++ {
			if (n*n)%(x-n) == 0 {
				nSolutions[n]++

				nS := nSolutions[n]

				if nS > largest {
					largest = nS
					fmt.Println("largest: ", largest)
				}

				if nSolutions[n] > 1000 {
					fmt.Println("n: ", n)
				}
			}
		}
	}
}