/*
	Problem 74
*/

package main

import (
	"fmt"
	"./crabMath"
)

func getFactorialSum(n int) int {
	digits := crabMath.GetDigits(n)
	sum := 0
	for _, d := range digits {
		sum += crabMath.Factorial(d)
	}

	return sum
}

func getChainLength(n int) int {
	chain := []int{}
	chain = append(chain, n)

	for {
		f := getFactorialSum(n)

		if crabMath.IntSliceContains(chain, f) {
			break
		}

		chain = append(chain, f)
		n = f
	}

	return len(chain)
}

func main() {
	res := 0

	for i := 0; i<1000000; i++ {
		if getChainLength(i) == 60 {
			res++
		}
	}

	fmt.Println(res)
}