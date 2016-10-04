/*
	Problem:
	What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
	"time"
)

func primeFactors(toCheck uint64) []uint64 {
	primeFactors := []uint64{}

	done := false
	for done != true {
		// check the 2 manually, so later we can skip all even numbers
		if toCheck % 2 == 0 {
			primeFactors = append(primeFactors, 2)
			toCheck /= 2
			continue
		}

		// we only need to go up to max the square root
		var i uint64 = 3
		var j uint64 = i*i
		for j <= toCheck {
			if toCheck % i == 0 {
				primeFactors = append(primeFactors, i)
				toCheck /= i
				break
			}

			i += 2
			j = i*i
		}

		// we are done here
		if j > toCheck {
			if toCheck != 1 {
				primeFactors = append(primeFactors, toCheck)
			}
			done = true
			break
		}
	}

	return primeFactors
}

func findFirstPositionOf(element uint64, inSlice []uint64) int {
	for i, j := range inSlice {
		if element == j {
			return i
		}
	}

	return -1
}

func main() {

	// measure execution time
	start := time.Now()

	// get all the prime factors from 2 to 20
	factorStorage := [19][]uint64{}

	for i := 2; i <= 20; i++ {
		factorStorage[i - 2] = primeFactors(uint64(i))
	}

	// iterate over all factors and find the smallest common multiple, that means
	// for each prime factor take the maximum occurence in one 
	var scm uint64 = 1
	
	for i := 0; i<19; i++ {
		if len(factorStorage[i]) == 0 {
			continue
		}
		
		factor := factorStorage[i][0]
		
		scm *= factor

		for j := i; j<19; j++ {
			k := findFirstPositionOf(factor, factorStorage[j])
			if k > -1 {
				factorStorage[j] = append(factorStorage[j][:k], factorStorage[j][(k+1):]...)
			}
		}
	}


	elapsed := time.Since(start)
	fmt.Println(scm)
	fmt.Println("Took ", elapsed)
}