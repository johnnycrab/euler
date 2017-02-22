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

/*
	Simple prime factorization of an integer, using the primes passed in as sieve
	Returns an error (first return value `true`) if the sieve is not big enough
*/
func SimplePrimeFactorizationWithSieve(n int, primeSieve []int) (bool, []int) {

	factors := []int{}

	if len(primeSieve) < n || (n == 1 || n == 0) {
		return true, factors
	}

	for primeSieve[n - 1] == 1 {
		for i, v := range primeSieve {
			if v == 0 && n%(i+1) == 0 {
				factors = append(factors, i+1)
				n = n/(i+1)
				break
			}
		}
	}

	factors = append(factors, n)

	return false, factors
}

/*
	Given a simple prime factorization, it groups prime factors together 
	by representing each prime and its power as a pair.
	E.g. [2,2,2,3,3] becomes [[2,3],[3,2]]
*/
func PrimeFactorizationToPowers(factorization []int) [][2]int {
	grouped := [][2]int{}

	currentPrime := factorization[0]
	count := 1

	for i := 1; i<len(factorization); i++ {
		p := factorization[i]
		if p != currentPrime {
			grouped = append(grouped, [2]int{currentPrime, count})
			count = 1
			currentPrime = p
		} else {
			count++
		}
	}

	grouped = append(grouped, [2]int{currentPrime, count})

	return grouped
}
