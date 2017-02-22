/*
	Problem 205
*/

package main

import (
	"fmt"
	"./crabMath"
)

// num of possibilities: 4^9*6^6
const numOfPoss float64 = 12230590464

func makeSums(throws [][]int) [37]int {
	sums := [37]int{}

	for _, throw := range throws {
		sum := 0
		for _, v := range throw {
			sum += v
		}

		sums[sum]++
	}

	return sums
}

func main() {

	pyramidThrows := crabMath.SetNPowerK(4, 9)
	cubeThrows := crabMath.SetNPowerK(6,6)
	
	pyramidSums := makeSums(pyramidThrows)
	cubeSums := makeSums(cubeThrows)

	won := 0
	for i, v := range cubeSums {
		for j := i+1; j < 37; j++ {
			won += pyramidSums[j]*v
		}
	}
	
	fmt.Println(float64(won)/numOfPoss)
}

