/*
	Problem 62
*/

package main

import (
	"fmt"
	"./crabMath"
)

const digitLen int = 12
var cubes []int
var symmetricGroups [][][]int

func countPermutations(n int) int {
	count := 1
	digits := crabMath.GetDigits(n)
	symmGroup := symmetricGroups[len(digits) - 1]


	used := []int{n}

	for _, permutation := range symmGroup {
		permutedDigits := make([]int, len(digits))

		if digits[permutation[0] - 1] == 0 {
			continue
		}
		for i, v := range permutation {
			permutedDigits[i] = digits[v-1]
		}

		permutedNum := crabMath.DigitsToInt(permutedDigits)

		if cubes[permutedNum] == 1 && !crabMath.IntSliceContains(used, permutedNum) {
			count++
			used = append(used, permutedNum)
		}
	}

	return count
}

func main() {
	// this is just an estimation that the needed number has exactly ten digits (it doesn't have nine, checked that already)

	cubeDigits := [][]int{}

	n := 1
	for {
		n++
		cube := n*n*n
		cd := crabMath.GetDigits(cube)
		if len(cd) == digitLen {
			cubeDigits = append(cubeDigits, cd)	
		} else if len(cd) > digitLen {
			break
		}
	}
	
	for _, cubeDigits1 := range cubeDigits {
		count := 0
		partners := []int{}
		for _, cubeDigits2 := range cubeDigits {
			if crabMath.IntSlicesContainSame(cubeDigits1, cubeDigits2) {
				partners = append(partners, crabMath.DigitsToInt(cubeDigits2))
				count++
			}
		}
		
		if count == 5 {
			fmt.Println(crabMath.DigitsToInt(cubeDigits1))
			break
		}
	}
}