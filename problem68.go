package main

import (
	"fmt"
	"strconv"
)

/*
	In the 5-gon, no matter if the distribution is valid or not, we have 10! = 3.628.800 possibilites – this is easily brute-forcable.
	We will check each permutation in S_10 for validity in the magic 5-gon. If it is valid, check the 16-digit string
*/

func faculty(n int) int {
	if n == 0 {
		return 1
	}

	res := 1
	for i:=1; i<=n; i++ {
		res *= i
	}

	return res
}

func getPossibleDistNums(permutation []int, n int) []int {
	retVal := []int{}

	for i := 1; i<=n; i++ {
		present := false
		for _, el := range permutation {
			if el == i {
				present = true
				break
			}
		}

		if !present {
			retVal = append(retVal, i)
		}
	}

	return retVal
}

func fillSlice(slice [][]int, n int) {
	// get all possible numbers we can distribute on the next position
	possibleNums := getPossibleDistNums(slice[0], n)

	if len(possibleNums) > 0 {
		newSliceSize := len(slice) / len(possibleNums)

		for i, possibleNum := range possibleNums {
			for j := 0; j<newSliceSize; j++ {
				slice[i*newSliceSize + j] = append(slice[i*newSliceSize + j], possibleNum)
			}
			fillSlice(slice[i*newSliceSize:(i+1)*newSliceSize], n)
		}
	}
}

func S_n(n int) [][]int {
	permutations := make([][]int, faculty(n))
	fillSlice(permutations, n)

	return permutations
}

func makePossibleSolutionFromPermutation(permutation []int) (string, bool) {
	smallestIndex := 0
	smallest := permutation[0]
	for i := 0; i<=4; i++ {
		if permutation[i] < smallest {
			smallest = permutation[i]
			smallestIndex = i
		}
	}

	indexOrder := [15]int{0,6,7,1,7,8,2,8,9,3,9,5,4,5,6}

	startIndex := 0
	for i := 0; i<15; i++ {
		if indexOrder[i] == smallestIndex {
			startIndex = i
		}
	}

	result := ""
	for i := 0; i<15; i++ {
		result += strconv.Itoa(permutation[indexOrder[(startIndex + i)%15]])
	}

	if len(result) == 16 {
		return result, false
	}

	return "", true
}

func main() {
	S_10 := S_n(10)

	possibleSolutions := []string{}

	for _, permutation := range S_10 {
		// check if it's a valid arrangement. See problem68.jpg for distribution
		sum := permutation[0] + permutation[6] + permutation[7]

		if sum == (permutation[1] + permutation[7] + permutation[8]) && sum == (permutation[2] + permutation[8] + permutation[9]) && sum == (permutation[3] + permutation[9] + permutation[5]) && sum == (permutation[4] + permutation[5] + permutation[6]) {
			// it is valid
			sol, error := makePossibleSolutionFromPermutation(permutation)
			if !error {
				possibleSolutions = append(possibleSolutions, sol)	
			}
		
		}
	}

	maximumString := possibleSolutions[0]

	for _, solution := range possibleSolutions {
		bigger := true
		for i, r := range solution {
			if int(r - '0') < int(maximumString[i] - '0') {
				bigger = false
				break
			} else if int(r - '0') > int(maximumString[i] - '0') {
				break
			}
		}

		if bigger {
			maximumString = solution
		}
	}

	fmt.Println(maximumString)
}