package crabMath

import "math/big"

/*
	Given two slices of integers where each entry is expected to be in {0,1,...,9} (i.e. the slice
	represents a big integers where each entry is a digit),
	adds both of them and returns a new slice. 

	Note: Lowest digit is expected to be a the end of the slice.
*/

func SliceAdd(a, b []int) []int {
	result := []int{}

	minLength := Min(len(a), len(b))
	subtractIndex := 1

	remain := 0
	for subtractIndex <= minLength {
		digitSum := a[len(a) - subtractIndex] + b[len(b) - subtractIndex] + remain

		if digitSum >= 10 {
			remain = 1
			digitSum %= 10
		} else {
			remain = 0
		}

		result = append([]int{digitSum}, result...)
		subtractIndex++
	}

	if len(a) == len(b) {
		if remain == 1 {
			result = append([]int{1}, result...)
		}

		return result
	}

	bigger := b
	if len(a) > len(b) {
		bigger = a
	}

	for subtractIndex <= len(bigger) {
		digitSum := bigger[len(bigger) - subtractIndex] + remain

		if digitSum >= 10 {
			remain = 1
			digitSum %= 10
		} else {
			remain = 0
		}

		result = append([]int{digitSum}, result...)
		subtractIndex++
	}

	return result
}

/*
	Power function a^b of positive integers, returns big.Int
*/
func IntegerPowerBig(a, b int64) *big.Int {
	if b == 0 {
		return big.NewInt(1)
	}

	if b == 1 {
		return big.NewInt(a)
	}

	res := big.NewInt(a)
	A := big.NewInt(a)
	var i int64 = 1
	for ; i < b; i++ {
		res.Mul(res, A)
	}

	return res
}