/*
	Problem 50
*/

package main

import "fmt"

const n int = 1000000
var primes []int


// gets primes up to n
func makePrimes() {

	nums := [n]int{}

	nums[0] = 1
	
	p := 2

	for p*p <= n {
		for i := p; i*p <= n; i++ {
			nums[i*p - 1] = 1
		}

		p++
		for nums[p-1] == 1 {
			p++
		}
	}

	// add the raw primes
	for i := 1; i<=n; i++ {
		if nums[i - 1] == 0 {
			primes = append(primes, i)
		}
	}

}

func getHighestChainLength(prime int, index int) int {
	highestChainLength := 0

	for j := index-1; j>=0; j-- {
		// check if this prime can be the end of a possible chain
		k := j
		chainSum := 0
		for chainSum < prime && k >= 0 {
			
			chainSum += primes[k]

			if chainSum == prime {
				if highestChainLength < (j - k + 1) {
					highestChainLength = j - k + 1
					break
				}
			}

			k--
		}
	}

	return highestChainLength
}

func main() {
	makePrimes()

	highestChainLength := 0
	bestPrime := 0

	for i := len(primes) - 1; i >= 0; i-- {
		// iterate over all primes
		primeToCheck := primes[i]

		cl := getHighestChainLength(primeToCheck, i)
		if cl > highestChainLength {
			fmt.Println("chain length: ", cl, "prime: ", primeToCheck)
			highestChainLength = cl
			bestPrime = primeToCheck
		}
	}

	fmt.Println(bestPrime)
}