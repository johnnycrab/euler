/*
	Problem 46
*/

package main

import (
	"fmt"
	"math"
)

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

func main() {
	makePrimes()

	k := 1

	for {
		k++
		n := 2*k + 1

		isGoldbachsch := false
		isPrime := false

		for _, prime := range primes {
			if prime == n {
				isPrime = true
				break
			}

			a := n - prime
			
			if a < 2 {
				break
			}

			a /= 2

			sqrt := int(math.Sqrt(float64(a)))
			if sqrt * sqrt == a {
				isGoldbachsch = true
				break
			}
		}

		if !isGoldbachsch && !isPrime {
			fmt.Println(n)
			break
		}
	}
}