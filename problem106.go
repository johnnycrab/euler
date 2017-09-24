/*
	Problem 106

	Also use same algorithms as in Problem 103, 105
*/

package main

import (
	"fmt"
	"./crabMath"
)

const N int = 12

// from n = 7 to 12, all n choose k sets for k = 2 to n-1 
var subsets [][][][]int

// from n = 7 to 12, all disjoint pairs of n choose k sets for k = 2 to n/2 
var distinctSubsetPairsOfSameLength [][][][2][]int

func main() {
	
	makeNeededSubsets()

	
	distinctNPairsOfSameLength := distinctSubsetPairsOfSameLength[N - 7]

	pairsThatNeedToBeChecked := 0

	for k := 2; k <= N/2; k++ {
		distinctPairsOfLengthK := distinctNPairsOfSameLength[k-2]
		
		for _, distinctPair := range distinctPairsOfLengthK {
			A := make([]int, k)
			B := make([]int, k)
			// check which is A and which is B
			if distinctPair[0][0] < distinctPair[1][0] {
				copy(A, distinctPair[0])
				copy(B, distinctPair[1])
			} else {
				copy(A, distinctPair[1])
				copy(B, distinctPair[0])
			}

			for _, v1 := range A {
				for i, v2 := range B {
					if v2 > v1 {
						B = append(B[:i], B[i+1:]...)
						break
					}
				}
			}

			if len(B) > 0 {
				pairsThatNeedToBeChecked++
			}
		}
	}

	fmt.Println(pairsThatNeedToBeChecked)
}


// from n = 7 to 12, gets all n choose k sets for k = 2 to n-1 
func makeNeededSubsets() {
	// Step 1: from n = 7 to 12, gets all n choose k sets for k = 2 to n-1 
	for n := 7; n <= 12; n++ {
		subsetsForN := [][][]int{}
		for k := 2; k<n; k++ {
			subsetsForN = append(subsetsForN, crabMath.SubsetsOfLength(n, k))
		}
		subsets = append(subsets, subsetsForN)
	}

	// Step 2: from n = 7 to 12, make all disjoint pairs of n choose k sets for k = 2 to n/2 
	for n := 7; n <= 12; n++ {
		subsetsForN := subsets[n-7]

		disjointPairsForN := [][][2][]int{}

		for k := 2; k <= n/2; k++ {
			nChooseKSets := subsetsForN[k - 2]
			disjointKSetPairs := [][2][]int{}

			for i, set1 := range nChooseKSets {
				for j := i+1; j<len(nChooseKSets); j++ {
					set2 := nChooseKSets[j]

					if setsAreDisjoint(set1, set2) {
						disjointKSetPairs = append(disjointKSetPairs, [2][]int{set1, set2})
					}
				}
			}

			disjointPairsForN = append(disjointPairsForN, disjointKSetPairs) 
		}

		distinctSubsetPairsOfSameLength = append(distinctSubsetPairsOfSameLength, disjointPairsForN)
	}

}


func setsAreDisjoint(a, b []int) bool {
	// sets have to be nonempty
	
	for _, k := range a {
		if crabMath.IntSliceContains(b, k) {
			return false
		}
	}

	return true
}
