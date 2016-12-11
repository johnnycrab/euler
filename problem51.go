/*
	Problem 51
*/

package main

import (
	"fmt"
	"./crabMath"
	"strconv"
)

const n int = 1000 * 1000
var subsetMap [][][][]int
var primeSieve []int

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}

// replaces all digits in the given subset of numString with `with`
func getReplacedNum(numString string, subset []int, with int) int {
	var r rune = rune(with + 48)
	for _, v := range subset {
		numString = replaceAtIndex(numString, r, v - 1)
	}

	i, _ := strconv.Atoi(numString)

	return i
}

func numOfPrimeFamilyBySubset(primeString string, subset []int) int {
	// we need to check that a digits in the given subset are the same. otherwise return 0
	r := primeString[subset[0] - 1]
	for _, v := range subset {
		if primeString[v - 1] != r {
			return 0
		}
	}

	// otherwise it is the same
	current := int(r) - 48
	count := 1
	
	for i := 0; i<=9; i++ {
		if i != current && !(i == 0 && subset[0] == 1) {
			replacedNum := getReplacedNum(primeString, subset, i)

			if replacedNum > 0 && primeSieve[replacedNum - 1] == 0 {
				count++
			}
		}
	}

	return count
}

func getMaxNumOfPrimeFamily(p int) int {
	primeString := strconv.Itoa(p)

	subsets := subsetMap[len(primeString) - 1]

	maxNum := 0

	// iterate over the number of digits and check replacement of digits
	for _, subsetsOfLengthK := range subsets {
		for _, subset := range subsetsOfLengthK {
			primeSubsetNum := numOfPrimeFamilyBySubset(primeString, subset)
			if primeSubsetNum > maxNum {
				maxNum = primeSubsetNum
			}
		}
	}

	return maxNum
}

// this is of the form
// numberOfDigits -> subsets of length k (from k=1,...,numOfDigits) -> the subsets
func getSubsetMap(maxNum int) [][][][]int {
	result := [][][][]int{}

	for numOfDigits := 1; numOfDigits<=maxNum; numOfDigits++ {
		byDigit := [][][]int{}

		for k:=1; k<=numOfDigits; k++ {
			subsets := crabMath.SubsetsOfLength(numOfDigits, k)
			byDigit = append(byDigit, subsets)
		}

		result = append(result, byDigit)
	}

	return result
}

func main() {
	subsetMap = getSubsetMap(7)
	primeSieve = crabMath.PrimeSieve(n)


	for i, v := range primeSieve {
		if v == 0 {
			num := getMaxNumOfPrimeFamily(i+1)
			if num == 8 {
				fmt.Println(i+1)
				break
			}
		}
	}

}