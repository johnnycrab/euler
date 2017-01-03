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

/*
	Given a slice of digits, returns the appropriate integer in reverse order
	Example: [1,2,3] produces 123
*/
func DigitsToInt(digits []int) int {
	power := 1
	result := 0
	for i := 0; i < len(digits); i++ {
		result += digits[i] * power
		power *= 10
	}

	return result
}