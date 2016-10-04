/*
	Problem:
	What is the 10 001st prime number?
*/

package main

import (
	"fmt"
	"time"
)

// construct a bijective map between the index set of our array storing the flags and 
// odd natural numbers >= 3:
// f: IndexSet -> OddNumbers >= 3
// i |-> 2*i + 3

// This will go up to 2*(i-1) +3 


func m(a int) int {
	return 2*a + 3
}

func m_inv(a int) int {
	return (a-3)/2
}

func main() {
	// measure execution time
	start := time.Now()


	// use sieve of Eratosthenes, but only keep track of odd numbers >= 3

	const store = 100000
	var upTo = m(store - 1) // our numbers go up until there

	// we see 2 already as prime and only look at odd numbers
	var numOfPrimesFound int = 1

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

	for i, val := range flags {
		if val == false {
			numOfPrimesFound++
		}

		if numOfPrimesFound == 10001 {
			fmt.Println("10001st prime: ", m(i))
			break
		}
	}

	//fmt.Println(numOfPrimesFound, " primes found")
	elapsed := time.Since(start)
	
	fmt.Println("Took ", elapsed)
}