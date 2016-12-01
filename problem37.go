/*
	Problem 37
*/

package main

import "fmt"

const n int = 1000000
var sieve [n]int

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

// p must be prime
func truncatablePrime(p int) bool {
	// truncate from right to left
	r := p
	pow := 1

	for r > 10 {
		pow *= 10
		r = (r - r%10)/10

		if sieve[r-1] == 1 {
			return false
		}
	}

	l := p
	for l > 10 {
		l = l%pow

		if sieve[l-1] == 1 {
			return false
		}

		pow /= 10
	}

	return true
}

func main() {
	sum := 0

	doSieve()

	for i := 11; i<n; i++ {
		if sieve[i-1] == 0 && truncatablePrime(i) == true {
			sum += i
		}
	}

	fmt.Println(sum)
}