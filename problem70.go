/*
	Problem 70

	We have phi(p) = p-1 for primes and phi(p*q) = (p-1)*(q-1), for semiprimes.
	They produce the minimal ratio n/phi(n). Thus if we have permutations in them, we are done.
	There won't be any permutations for primes. Thus only check semiprimes.
*/

package main

import (
	"fmt"
)

const n int = 10000000
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
/*func primeFactorization(N int) []int {
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

func power(p int, k int) int {
	if k == 0 {
		return 1
	}

	result := p
	for i := 2; i<=k; i++ {
		result *= p
	}

	return result
}

// returns the exact value of the euler totient function
func eulerTotient(N int) int {
	// 1. Split up N into prime factors and their powers.
	primeFactors := primeFactorization(N)

	// 2. the totient is the product over alle primes with power k, where each multiplicand is given by p^k-p^(k-1)
	totient := 1
	for i := 0; i < len(primeFactors) - 1; i += 2 {
		prime := primeFactors[i]
		totient *= (power(prime, primeFactors[i+1] - 1) * (prime - 1))
	}

	return totient
}*/

func digits(a int) []int {
	if a == 0 {
		return []int{0}
	}

	res := []int{}

	for a > 0 {
		res = append(res, a%10)
		a = (a - (a%10))/10
	}

	return res
}

func isPermutation(a int, b int) bool {
	digitsA := digits(a)
	digitsB := digits(b)

	if len(digitsA) != len(digitsB) {
		return false
	}

	isPerm := true

	for _, aDigit := range digitsA {
		found := false

		for i, bDigit := range digitsB {
			if bDigit == aDigit {
				found = true
				digitsB[i] = -1
				break
			}
		}

		if !found {
			isPerm = false
			break
		}
	}

	return isPerm
}

func main() {
	

	// 1. Get all prime numbers < 10^7
	makePrimes()

	var minimumRatio float64 = -1
	resultN := 0

	// 2. Calculate the exact euler totient function for semiprimes
	for i, prime1 := range primes {

		if prime1 * prime1 > n {
			break
		}
	
		for j := i+1; j < len(primes); j++ {
			prime2 := primes[j]
			if prime1 * prime2 > n {
				break
			}

			semiprime := prime1 * prime2
			totient := (prime1 - 1)*(prime2 - 1)
			if isPermutation(semiprime, totient) {
				var ratio float64 = float64(semiprime)/float64(totient)

				if (ratio < minimumRatio) || (minimumRatio == -1) {
					minimumRatio = ratio
					resultN = semiprime
				}
			}
		}

	}

	fmt.Println(resultN)
}