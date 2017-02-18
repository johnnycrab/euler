/*
	Problem 95

	the sum of divisors sigma is a multiplicative function for coprime numbers.
	for p^n, we have sigma(p^n) = (p^(n+1)-1)/(p-1)
*/

package main

import (
	"fmt"
	"./crabMath"
)

const N int = 1000*1000
const Million int = 1000*1000
var divisorSums [N + 1]int
var primeSieve []int

func getSumOfDivisors(n int) int {
	
	_, primeFactors := crabMath.SimplePrimeFactorizationWithSieve(n, primeSieve)

	factorPairs := [][2]int{}

	last := 0
	num := 0
	for _, v := range primeFactors {

		if v != last {
			if last != 0 {
				factorPairs = append(factorPairs, [2]int{last, num})
			}
			last = v
			num = 1
		} else {
			num++
		}
	}
	factorPairs = append(factorPairs, [2]int{last, num})

	sum := 1
	
	for _, factorPair := range factorPairs {
		p := factorPair[0]
		a := factorPair[1]
		pow := crabMath.Power(p, a)
		sum *= (pow * p - 1)/(p-1)
		if (sum - n) > Million {
			return -1
		}
	}


	return sum - n
}

func main() {
	primeSieve = crabMath.PrimeSieve(N)
	divisorSums = [N+1]int{}

	// First step: For every number get the divisor sum and set it on the divisor-sum array
	for i := 1; i<=N; i++ {
		divisorSums[i] = getSumOfDivisors(i)
	}

	longestChainLength := 0
	smallestMember := 0

	for i := 2; i<=N; i++ {
		chain := []int{i}
		smallestInChain := i

		n := i
		for divisorSums[n] != 1 && divisorSums[n] != -1 {
			
			n = divisorSums[n]
		
			if n < smallestInChain {
				smallestInChain = n
			}

			// we have a real chain
			if n == chain[0] {			
				if longestChainLength < len(chain) || longestChainLength == 0 {
					longestChainLength = len(chain)
					smallestMember = smallestInChain
				}
				break
			} else if crabMath.IntSliceContains(chain, n) {
				break
			} else {
				chain = append(chain, n)
			}

		}
	}

	fmt.Println(smallestMember, longestChainLength)
}