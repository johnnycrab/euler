/*
	Problem 73
*/

package main

import (
	"fmt"
	"./crabMath"
)

const N int = 12000
var primeSieve []int
var factorization [N+1][][2]int

func isCoprime(factors1, factors2 [][2]int) bool {
	for _, factor1 := range factors1 {
		for _, factor2 := range factors2 {
			if factor1[0] == factor2[0] {
				return false
			}
		}
	}

	return true
}

func main() {
	primeSieve = crabMath.PrimeSieve(N)

	for i := 2; i<=N; i++ {
		_, factors := crabMath.SimplePrimeFactorizationWithSieve(i, primeSieve)
		factorization[i] = crabMath.PrimeFactorizationToPowers(factors)
	}

	count := 0

	for d := 4; d <= N; d++ {
		for n := (d/3)+1; 2*n < d; n++ {
			if isCoprime(factorization[d], factorization[n]) {
				count++
			}
		}
	}

	fmt.Println(count)
}

