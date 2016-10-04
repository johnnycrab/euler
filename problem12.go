/*
	Problem:
	What is the value of the first triangle number to have over five hundred divisors?
*/

package main

import (
	"fmt"
)

func getNumOfDivisors(a int) int {
	if a == 1 {
		return 1
	}

	p_factors := primeFactors(a)

	num := 1

	// we need to count the exponents for each prime factor.
	currentFactor := p_factors[0]
	count := 0
	for _, fac := range p_factors {
		if currentFactor != fac {
			currentFactor = fac
			num *= (count + 1)
			count = 0
		}
		count++
	}

	num *= (count + 1)

	return num
}

func primeFactors(toCheck int) []int {
	primeFactors := []int{}

	done := false
	for done != true {
		// check the 2 manually, so later we can skip all even numbers
		if toCheck % 2 == 0 {
			primeFactors = append(primeFactors, 2)
			toCheck /= 2
			continue
		}

		// we only need to go up to max the square root
		var i int = 3
		var j int = i*i
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

func main() {

	found := false
	i := 0

	var triangle int = 0

	for found == false {
		i++

		
		// get ith triangle number
		triangle += i

		divisors := getNumOfDivisors(triangle)

		if divisors > 500 {
			fmt.Println(triangle)
			found = true
		}
	}
}

// 22014930