/*
	Problem 47
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
	primes = append(primes, p)

	for p*p <= n {
		for i := p; i*p <= n; i++ {
			nums[i*p - 1] = 1
		}

		for {
			p++
			if nums[p - 1] == 0 {
				primes = append(primes, p)
				break
			}
		}
	}

	// now add the rest of the primes
	for ; p<=n; p++ {
		if nums[p - 1] == 0 {
			primes = append(primes, p)
		}
	}

}

// Split up N into prime factors and their powers. Store this in an array where at pos i is the prime, at position i+1 is the power
func primeFactorization(N int) []int {
	factors := []int{}

	temp := N
	for _, prime := range primes {

		if temp == 1 {
			break
		}

		if temp%prime == 0 {

			// we have a prime, let's get the power
			power := 1
			temp /= prime

			for temp%prime == 0 {
				power += 1
				temp /= prime
			}

			factors = append(factors, prime, power)
		}
	}

	return factors
}

func hasFourDistinctFactors(N int) bool {
	factors := primeFactorization(N)

	return len(factors) == 8
}

func main() {
	makePrimes()

	N := 1

	for {
		if hasFourDistinctFactors(N) && hasFourDistinctFactors(N+1) && hasFourDistinctFactors(N+2) && hasFourDistinctFactors(N+3) {
			fmt.Println(N)
			break
		}
		N++

		if N > n {
			fmt.Println(N)
			break
		}
	}
}