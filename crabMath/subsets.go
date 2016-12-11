/*
	Given a subset of length `n`
	this utility functions returns all subsets of length `k`

	Subsets are represented by arrays of indices, that is, function 
	`SubsetsOfLength(n=3, k=2)` returns following slice:
	[[1 2] [1 3] [2 3]]
*/

package crabMath

func fillInSubsetBatch(batch [][]int, withNumber int, finalSubsetLength int, originalSetLength int) {
	currentLen := len(batch[0])

	if currentLen < finalSubsetLength && withNumber <= originalSetLength {
		// calculate into how many of the batch we need to fill in the number
		remainingLength := finalSubsetLength - currentLen - 1
		countOfNumsToChooseFrom := originalSetLength - withNumber
		insertIntoCount := NChooseK(countOfNumsToChooseFrom, remainingLength)
			
		for i := 0; i<insertIntoCount; i++ {
			batch[i] = append(batch[i], withNumber)
		}

		// recursively call the function two times: 
		// 1x for incremented `withNumber`, called on the rest of the batch below
		// 1x for incremented `withNumber`, cllaed on the batch that has just been filled

		// call on batch that has just been filled
		fillInSubsetBatch(batch[0:insertIntoCount], withNumber + 1, finalSubsetLength, originalSetLength)

		if len(batch) > insertIntoCount {
			fillInSubsetBatch(batch[insertIntoCount:], withNumber + 1, finalSubsetLength, originalSetLength)
		}
	}
}

// given a set of length of `n`, returns
// all subsets of length `l`
// subsets are represented as arrays of indices
func SubsetsOfLength(n, l int) [][]int {
	sets := make([][]int, NChooseK(n, l))

	fillInSubsetBatch(sets, 1, l, n)

	return sets
}