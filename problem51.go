/*
	Problem 51
*/

package main

import "fmt"

func factorial(n int) int {
	result := 1

	if n == 0 || n == 1 {
		return 1
	}

	for i := 2; i<=n; i++ {
		result *= i
	}

	return result
}

func nChooseK(n, k int) int {
	return factorial(n)/(factorial(k) * factorial(n-k))
}

// given a set of length of `lengthOfSuperset`, returns
// all subsets of length `l`
// subsets are represented as arrays of indices
func getSubsetsOfLength(lengthOfSuperset, l int) [][]int {
	sets := make([][]int, nChooseK(lengthOfSuperset, l))

	return sets
}

func main() {
	fmt.Println(nChooseK(10,7))
}