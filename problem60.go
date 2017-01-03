/*
	Problem 60
*/

package main

import (
	"fmt"
	"./crabMath"
)

const n int = 100000000
const checkPrimesUpTo int = 10000 // this is an estimate
var sieve []int
var primes []int

func concBothWaysIsPrime(a, b int) bool {
	return sieve[crabMath.Concatenate(a,b) - 1] == 0 && sieve[crabMath.Concatenate(b,a) - 1] == 0
}

// given a set of primes, checks if this set fulfills the property
func setOfPrimesIsSpecial(set []int) bool {
	for i, prime1 := range set {
		for j, prime2 := range set {
			if i != j && !concBothWaysIsPrime(prime1, prime2) {
				return false
			}
		}
	}

	return true
}

// given a slice that contains subsets of primes, it sieves out within the subset all primes which
// have less than 4 different prime pairs which fulfil the needed property.
// If a subset then still has more than 20 elements, the subset is split up again into its appropriate pairs
// and the process repeated
func splitSieve(primeSubsets [][]int) [][]int {
	updated := [][]int{}

	for _, primeSubset := range primeSubsets {
		
		primeSubset = usefulSieve(primeSubset)

		if len(primeSubset) > 20 {

			for _, prime := range primeSubset {
				split := usefulSieve(getPartnerPrimes(prime, primeSubset))

				if len(split) >= 5 {
					updated = append(updated, split)
				}
			}
		} else if len(primeSubset) >= 5 {
			updated = append(updated, primeSubset)
		}
	}

	return updated
}


// given a prime and a set of primes, returns all partners of the prime within the set
func getPartnerPrimes(prime int, set []int) []int {
	partners := []int{prime}

	for _, prime1 := range set {
		if prime != prime1 && concBothWaysIsPrime(prime, prime1) {
			partners = append(partners, prime1)
		}
	}

	return partners
}

// given primes, sieves out all primes that have less than 4 different prime pairs that fulfil the property
func usefulSieve(ps []int) []int {

	usefulPrimes := []int{}
	for i, prime1 := range ps {
		count := 0
		for j, prime2 := range ps {
			if i != j && concBothWaysIsPrime(prime1, prime2) {
				count++
			}
		}

		if count >= 4 {
			usefulPrimes = append(usefulPrimes, prime1)
		}
	}

	if len(usefulPrimes) != len(ps) {
		return usefulSieve(usefulPrimes)
	}

	return usefulPrimes
}

func main() {
	sieve = crabMath.PrimeSieve(n)
	primes = crabMath.PrimesUpTo(checkPrimesUpTo)
	partners := splitSieve([][]int{primes})

	smallest := 0

	for _, matching := range partners {
		if len(matching) >= 5 && len(matching) <= 20 {
			// get subsets
			generalSubsets := crabMath.SubsetsOfLength(len(matching), 5)
			for _, generalSubset := range generalSubsets {
				subset := []int{}
				for _, v := range generalSubset {
					subset = append(subset, matching[v-1])
				}
			
				if setOfPrimesIsSpecial(subset) {

					sum := 0
					for _, p := range subset {
						sum += p
					}

					if smallest == 0 || sum < smallest {
						smallest = sum
					}
				}
			}
			
		}	
	}

	fmt.Println(smallest)
}