/*
	Prime functions
*/

package crabMath

/*
	Returns a prime sieve of length `n`, represented as slice
	Sieve goes from 1 - n
	1 represents not-prime
	0 represents prime
*/
func PrimeSieve(n int) []int {

	sieve := make([]int, n)

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

	return sieve
}

/*
	Returns a slice containing all primes up to the given integer.
	This is not a sieve!
*/
func PrimesUpTo(n int) []int {
	primes := []int{}
	sieve := PrimeSieve(n)

	for i, v := range sieve {
		if v == 0 {
			primes = append(primes, i+1)
		}
	}

	return primes
}