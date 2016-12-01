/*
	Problem 36
*/

package main

import (
	"fmt"
	"strconv"
)

func isAPalindrome(s string) bool {
	count := len(s)
	for k := 0; k < count / 2.0; k++ {
		if s[k] != s[count - k - 1] {
			return false
			break
		}
	}

	return true
}

func smallDecimalToBinaryString(decimal int) string {
	binaryString := ""

	highbit := 0
	cur := 1
	for cur * 2 <= decimal {
		highbit++
		cur *= 2
	}

	for decimal > 0 {
		if cur <= decimal {
			binaryString += "1"
			decimal -= cur
		} else {
			binaryString += "0"	
		}
		cur /= 2

	}

	for len(binaryString) < highbit + 1 {
		binaryString += "0"
	}

	return binaryString
}


func main() {
	sum := 0

	for i := 1; i < 1000000; i++ {
		if isAPalindrome(strconv.Itoa(i)) && isAPalindrome(smallDecimalToBinaryString(i)) {
			sum += i
		}
	}

	fmt.Println(sum)
}