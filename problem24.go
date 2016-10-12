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

func main() {
	S_10 := S_n(10)

	for _, i := range S_10[999999] {
		fmt.Print(i - 1)
	}
	fmt.Println("")

}