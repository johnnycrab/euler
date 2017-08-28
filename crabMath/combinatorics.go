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

func fillSlicePowerSet(slice [][]int, atPos int, n int, k int) {
	if atPos == k {
		return
	}

	dist := len(slice)/n


	for i := 1; i <= n; i++ {

		startAt := (i-1)*dist
		for m := startAt; m < startAt + dist; m++ {
			slice[m][atPos] = i
		}
		//fmt.Println(slice)
		fillSlicePowerSet(slice[startAt:startAt+dist], atPos + 1, n, k)
	}
}

func S_n(n int) [][]int {
	permutations := make([][]int, Factorial(n))
	fillSlice(permutations, n)

	return permutations
}

/*
	Returns all elements of {1,...,n}^k 
*/
func SetNPowerK(n, k int) [][]int {
	neededLength := Power(n, k)
	elements := make([][]int, neededLength)

	for i := 0; i<neededLength; i++ {
		elements[i] = make([]int, k)
	}

	fillSlicePowerSet(elements, 0, n, k)

	return elements
}

/*
	Returns NChooseK-sets, that is all distinct tuples of the form {p_1, ..., p_k} with 1 <= p_i <= n and p_i > p_{i+1}  
	Works only for small n
*/
func NChooseKSets(n, k int) [][]int {

	neededLength := NChooseK(n,k)
	sets := make([][]int, neededLength)

	for i := 0; i<neededLength; i++ {
		sets[i] = make([]int, k)
	}

	nChooseKSetsFillSlice(sets, 1, n, k)

	return sets
}

// given number m on position p of k slots, how many slices are needed?
func nChooseKSetsGetNumOfSlices(sum *int, m int, p int, k int) {
	if p == k {
		*sum++
	} else {
		for l := m-1; l >= k-p; l-- {
			nChooseKSetsGetNumOfSlices(sum, l, p+1, k)
		} 
	}
}

func nChooseKSetsFillSlice(slice [][]int, atPos int, n int, k int) {
	previousNum := n+1
	if atPos != 1 {
		previousNum = slice[0][atPos - 2]
	}


	downTo := k - atPos + 1
	nextCut := 0
	for m := previousNum - 1; m >= downTo; m-- {
		slicesNeeded := NChooseK(m-1, k - atPos)

		sliceShare := slice[nextCut: nextCut + slicesNeeded]
		nextCut += slicesNeeded

		for i := 0; i<len(sliceShare); i++ {
			sliceShare[i][atPos - 1] = m
		}

		if atPos != k {
			nChooseKSetsFillSlice(sliceShare, atPos + 1, n, k)
		}
	}
}