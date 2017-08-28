/*
	Problem 90
*/

package main

import (
	"fmt"
	"./crabMath"
)

var squares = [][]int{{0,1}, {0,4}, {0,9}, {1,6}, {2,5}, {3,6}, {4,9}, {6,4}, {8,1}}

func main() {
	// keeping track of good pairs.
	count := 0

	// we first generate all 10 choose 6 sets
	choiceSets := crabMath.NChooseKSets(10, 6)

	for i := 0; i<len(choiceSets); i++ {
		set1 := updateSet(choiceSets[i])

		for j := i + 1; j<len(choiceSets); j++ {
			set2 := updateSet(choiceSets[j])
			if checkSets(set1, set2) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func checkSets(set1, set2 []int) bool {
	isOkay := true

	for _, square := range squares {
		if !(crabMath.IntSliceContains(set1, square[0]) && crabMath.IntSliceContains(set2, square[1]) || crabMath.IntSliceContains(set2, square[0]) && crabMath.IntSliceContains(set1, square[1])) {
			isOkay = false
			break
		}
	}

	return isOkay
}

// updates the numbers in the set (subtracting each entry by one)
// and adds a 6 or a 9 to it, if needed
func updateSet(set []int) []int {
	cpySet := make([]int, 6)
	copy(cpySet, set)

	hasSix := false
	hasNine := false
	for i, v := range cpySet {
		n := v-1
		cpySet[i] = n
		if n == 6 {
			hasSix = true
		} else if n == 9 {
			hasNine = true
		}
	}

	if hasSix && !hasNine {
		cpySet = append(cpySet, 9)
	} else if hasNine && !hasSix {
		cpySet = append(cpySet, 6)
	}

	return cpySet
}