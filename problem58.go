/*
	Problem 58
*/

package main

import (
	"fmt"
	"./crabMath"
)

var primeSieve []int

// `ofLayer` is the number of layer around 1, i.e. 1 does not count as layer itself
func returnCornerNums(ofLayer int) []int {
	upperRight := 0

	// calculate the number on the upper right corner
	if ofLayer == 1 {
		upperRight = 3
	} else {
		upperRight = 1 + 2*ofLayer + 4*(ofLayer-1)*ofLayer
	}

	return []int{upperRight, upperRight + ofLayer*2, upperRight + ofLayer*4, upperRight + ofLayer*6}
}

func primeRatio(set []int) float64 {
	numOfPrimes := 0
	for _, v := range set {
		
		if primeSieve[v-1] == 0 {
			numOfPrimes++
		}
	}

	
	//fmt.Println(float64(numOfPrimes) / float64(len(set)))
	return float64(numOfPrimes) / float64(len(set))
}

func main() {
	primeSieve = crabMath.PrimeSieve(700000000)

	diagonals := []int{1}

	numOfLayers := 0
	last := 1
	for {
		numOfLayers++
		upperRight := last + 2*numOfLayers
		diagonals = append(diagonals, upperRight, upperRight + numOfLayers*2, upperRight + numOfLayers*4, upperRight + numOfLayers*6)
		last = upperRight + numOfLayers*6


		if primeRatio(diagonals) < 0.1 {
			fmt.Println(2*numOfLayers + 1)
			break
		}
	}

}