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