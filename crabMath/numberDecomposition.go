package crabMath

/*
	Returns the digits of the given integer in reverse order
*/
func GetDigits(ofNum int) []int {
	ret := []int{}

	for ofNum > 0 {
		ret = append(ret, ofNum%10)

		ofNum = (ofNum - ofNum%10) / 10
	}

	return ret
}