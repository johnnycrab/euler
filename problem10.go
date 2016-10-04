/*
	Problem:
	Find the sum of all the primes below two million.
*/

package main

import (
	"fmt"
	"time"
)

func m(a int) int {
	return 2*a + 3
}

func m_inv(a int) int {
	return (a-3)/2
}

func main() {

	// measure execution time
	start := time.Now()

	sum := 2

	// we use the sieve from problem 7 !

	// use sieve of Eratosthenes, but only keep track of odd numbers >= 3

	const store = 10000000
	var upTo = m(store - 1) // our numbers go up until there

	flags := [store]bool{}

	// only go up to square root
	for i := 0; m(i)*m(i) <= upTo; i +=1 {
		// go to the nearest one, this is a prime
		if flags[i] == true {
			continue
		}

		for j := i; m(j)*m(i) <= upTo; j++ {
			// this is basically m_inv
			flags[2*i*j + 3*i + 3*j + 3] = true
		}
	}


	greatestPrime := 0

	for i, val := range flags {
		
		if val == false {
			x := m(i)
			greatestPrime = x
			if x < 2000000 {
				sum += x	
			} else {
				break
			}
		}
	}

	if greatestPrime < 2000000 {
		fmt.Println("Not enough primes found. Greatest prime: ", greatestPrime)
	} else {
		fmt.Println("Sum: ", sum)
	}

	elapsed := time.Since(start)

	fmt.Println("Took ", elapsed)
}