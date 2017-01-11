/*
	Problem 71
*/

package main

import (
	"fmt"
	//"./crabMath"
)

const N int = 1000*1000

func main() {

	//primes := crabMath.PrimesUpTo(N)

	the_n := 2
	the_d := 5

	for d := 2; d <= N; d++ {

		for n := 1; n < d; n++ {

			diff_n := 3*d - 7*n

			if diff_n <= 0 {
				break
			} else if n*the_d - the_n*d > 0 {

				the_n = n
				the_d = d
			}
		}
	}

	fmt.Println(the_n, the_d)
}