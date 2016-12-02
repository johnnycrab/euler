/*
	Problem 43
*/

package main

import "fmt"

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

func hasSubstringProperty(number []int, divisors *[7]int) bool {
	
	for i := 1; i<=7; i++ {
		num := (number[i] - 1) * 100 + (number[i+1] - 1) * 10 + (number[i+2] - 1)
		if num%(*divisors)[i-1] != 0 {
			return false
		}
	}

	return true
}

func makeNumFromPermutation(permutation []int) int {
	mul := 1
	res := 0

	for i := 9; i>=0; i-- {
		res += (permutation[i] -1) * mul
		mul *= 10
	}

	return res
}

func main() {
	S_10 := S_n(10)

	divisors := [7]int{2,3,5,7,11,13,17}

	result := 0

	for _, permutation := range S_10 {
		if permutation[0] != 1 {
			// it is a 0 to 9 pandigital number, let's check if it fulfils the property
			if hasSubstringProperty(permutation, &divisors) {
				num := makeNumFromPermutation(permutation)
				result += num
			}
		}
	}

	fmt.Println(result)
}