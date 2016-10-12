package main

import "fmt"

func isPrime(n int) bool {
	// brute force method to check if n is prime
	prime := true
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}

	return prime
}

func main() {

	maximumChainLenOfPrimes := 0
	solA := 0
	solB := 0
	
	for a := -999; a < 1000; a++ {
		for b := -1000; b < 1000; b++ {
			currentChainLenOfPrimes := 0
			n := -1
			for {
				n++
				if isPrime(n*n + a*n + b) {
					currentChainLenOfPrimes++
				} else {
					break
				}
			}
			if currentChainLenOfPrimes > maximumChainLenOfPrimes {
				maximumChainLenOfPrimes = currentChainLenOfPrimes
				solA = a
				solB = b
			}
		}
	}

	fmt.Println("a: ", solA, "b: ", solB, " with length: ", maximumChainLenOfPrimes)
	fmt.Println("Solution: ", solA * solB)
}