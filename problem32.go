/*
	Problem 32:
	Find sum of all pandigital products
*/
package main

import (
	"fmt"
	"strconv"
)


// computes a*b=c, checks if a o b c is pandigital
func isPandigital(a,b int) bool {
	c := a * b
	
	fullString := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c)
	
	if len(fullString) != 9 {
		return false
	}

	check := [9]int{}

	for _, r := range fullString {
		if r-49 >= 0 {
			check[r-49] = 1
		}
	}
	
	for _, i := range check {
		if i == 0 {
			return false
		} 
	}

	return true
}

func main() {
	pandigitalProducts := []int{}

	for i := 2; i < 1000; i++ {
		for j := 1; j < 10000; j++ {
			if (isPandigital(i,j)) {
				product := i*j
				present := false
				for _, val := range pandigitalProducts {
					if val == product {
						present = true
						break
					}
				}

				if !present {
					pandigitalProducts = append(pandigitalProducts, product)	
				}
			}
		}
	}

	sum := 0
	for _, val := range pandigitalProducts {
		sum += val
	}

	fmt.Println(sum)
}