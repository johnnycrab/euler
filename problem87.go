/*
	Problem 87
*/

package main

import "fmt"

// we are interested in sums of at least squares below 50 * 10^6, so get prime numbers only
// up to square-root
const n int = 70702
const max int = 50 * 1000 * 1000

var sieve [n]int
var primeSquares []int
var primeCubes []int
var primeFourthPowers []int


func power(n, k int) int {
	if k == 0 {
		return 1
	}

	result := n
	for j := 2; j<=k; j++ {
		result *= n
	}

	return result
}

// gets primes up to n
func doSieve() {

	sieve[0] = 1

	p := 2

	for p*p <= n {
		for i := p; i*p <= n; i++ {
			sieve[i*p - 1] = 1
		}

		for i := 1; i < n - p; i++ {
			if sieve[p - 1 + i] == 0 {
				p = p + i
				
				break
			}
		}
	}
}

// prepares powers of prime numbers
func makePrimePowers() {
	doSieve()

	// keep track of third powers – if they get too large, we do not need
	// to store them anymore
	doThird := true
	doFourth := true


	for i, v := range sieve {
		if v == 0 {
			// this is a prime
			prime := i + 1

			primeSquare := prime*prime
			primeSquares = append(primeSquares, primeSquare)

			if doThird {
				primeCube := prime*prime*prime
				if primeCube < max {
					primeCubes = append(primeCubes, primeCube)
				} else {
					doThird = false
				}
			}

			if doFourth {
				primeFourth := prime*prime*prime*prime
				if primeFourth < max {
					primeFourthPowers = append(primeFourthPowers, primeFourth)
				} else {
					doFourth = false
				}
			}
		}
	}
}



func main() {
	makePrimePowers()

	count := 0

	// keep track of all numbers below 50 mill.
	nums := make([]int, max)

	for _, sq := range primeSquares {
		for _, cu := range primeCubes {
			for _, fp := range primeFourthPowers {
				sum := sq + cu + fp
				if sum < max {
					nums[sum - 1] = 1
				}
			}		
		}		
	}

	for _, v := range nums {
		if v == 1 {
			count++
		}
	}

	fmt.Println(count)
}