/*
	Problem 69

	We work with floating point numbers here – an mere approximation is alright for us, we don't need the
	exact integer value of the totient function
*/

package main

import "fmt"

const n int = 1000*1000
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

		for i := 1; i < n - p; i++ {
			if nums[p - 1 + i] == 0 {
				p = p + i
				primes = append(primes, p)
				break
			}
		}
	}
}

func eulerTotient(N int) float64 {
	primeProd := 1.0

	for _, prime := range primes {
		if prime*prime > N {
			break
		}
		if N%prime == 0 {
			primeProd *= (1.0 - (1.0/float64(prime)))
		}
	}

	return float64(N) * primeProd
} 

func main() {
	makePrimes()

	var largestTotientDivision float64 = 0
	nValue := 0

	for i := 1; i<=n; i++ {
		val := float64(i)/eulerTotient(i)
		if val > largestTotientDivision {
			nValue = i
			largestTotientDivision = val
		}
	}

	fmt.Println(nValue)
}

