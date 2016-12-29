/*
	Problem 55
*/

package main

import (
	"fmt"
	"./crabMath"
)

func isPalindrome(num []int) bool {
	for i := 0; i < (len(num) / 2); i++ {
		if num[i] != num[len(num) - 1 - i] {
			return false
		}
	}

	return true
}

func isLychrel(n int) bool {
	n_slice := crabMath.GetDigits(n)

	for i := 1; i <= 50; i++ {
		n_slice = crabMath.SliceAdd(n_slice, crabMath.IntSliceReverse(n_slice))

		if isPalindrome(n_slice) {
			return false
		}
	}

	return true
}

func main() {
	count := 0

	for n := 1; n <= 10000; n++ {
		if isLychrel(n) {
			count++
		}
	}

	fmt.Println(count)
}