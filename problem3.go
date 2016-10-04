/*
	Problem:
	What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	// measure execution time
	start := time.Now()

	// stores prime factors
	primeFactors := []uint64{}

	var toCheck uint64 = 600851475143
	//var toCheck uint64 = 200

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
	
	fmt.Println(primeFactors[len(primeFactors) - 1])
	elapsed := time.Since(start)

	fmt.Println("Took ", elapsed)
}