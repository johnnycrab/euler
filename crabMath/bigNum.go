package crabMath

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