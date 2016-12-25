/*
	Problem 75
*/

package main

import (
	"fmt"
	"./crabMath"
)

const n int = 5000
var primeSieve []int

func main() {

	primitivePythagoreanTriples := crabMath.GetPrimitivePythagoreanTriples(n)

	generatedLengths := [1500000]int{}
	
	for _, triple := range primitivePythagoreanTriples {

		a := triple[0]
		b := triple[1]
		c := triple[2]
		
		k := 0

		for {
			k++
			newA := a*k
			newB := b*k
			newC := c*k

			sum := newA + newB + newC
			if sum > 1500000 {
				break
			}

			generatedLengths[sum-1]++
		}
	}

	counter := 0
	for _, v := range generatedLengths {
		if v == 1 {
			counter++
		}
	}

	fmt.Println(counter)
}