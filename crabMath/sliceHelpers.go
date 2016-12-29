package crabMath

// returns true iff a is contained in s
func IntSliceContains(s []int, a int) bool {
	for _, v := range s {
		if v == a {
			return true
		}
	}

	return false
}

func IntSlicesContainSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	bChecks := make([]int, len(b))

	ret := true

	for _, aVal := range a {
		found := false
		for i, bVal := range b {
			if bVal == aVal && bChecks[i] == 0 {
				found = true
				bChecks[i] = 1

				break
			}
		}
		if !found {
			ret = false
			break
		}
	}

	return ret
}

// returns a new slice where the order of elements has been reversed
func IntSliceReverse(s []int) []int {
	s_tilde := make([]int, len(s))

	for i, v := range s {
		s_tilde[len(s) - 1 - i] = v
	}

	return s_tilde
}