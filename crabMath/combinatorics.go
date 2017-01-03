/*
	Basic combinatoric functions
*/

package crabMath

import (
	"math/big"
)

func Factorial(n int) int {
	result := 1

	if n == 0 || n == 1 {
		return 1
	}

	for i := 2; i<=n; i++ {
		result *= i
	}

	return result
}

func NChooseK(n, k int) int {

	return int(BigNChooseK(int64(n), int64(k)).Uint64())
}

func BigNChooseK(n, k int64) *big.Int {
	if k == 0 {
		return big.NewInt(1)	
	}

	if k == 1 {
		return big.NewInt(n)
	}
	
	NMinKF := big.NewInt(n)
	KFactorial := big.NewInt(k)

	var i int64 = 1
	for ; i <= k - 1; i++ {
		NMinKF.Mul(NMinKF, big.NewInt(n-i))
		KFactorial.Mul(KFactorial, big.NewInt(k-i))
	}

	return NMinKF.Quo(NMinKF, KFactorial)
}

// Symmetric groups

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
	permutations := make([][]int, Factorial(n))
	fillSlice(permutations, n)

	return permutations
}