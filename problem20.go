/*
	Problem:
	Find the sum of the digits in the number 100!
*/

package main

import (
	"fmt"
	"strconv"
)

func increaseBitInSlice(bitPos int, bits []int) {
		for bits[bitPos] == 1 {
			bits[bitPos] = 0
			bitPos += 1
		}
		bits[bitPos] = 1
}

func divide10WithDecimalRemainder(bitstring string) (string, int) {
	remainder := 0
	hasOne := false
	retString := ""
	for i, r := range bitstring {
		bit := int(r - '0')
		
		remainder += remainder + bit

		if remainder >= 10 {
			remainder -= 10
			retString += "1"
			hasOne = true
		} else if hasOne || (i==len(bitstring)-1) {
			retString += "0"
		}
	}

	return retString, remainder
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

func binaryStringToDecimalString(bitstring string) string {
	decimalString := ""

	remainder := 0

	for bitstring != "0" {
		bitstring, remainder = divide10WithDecimalRemainder(bitstring)
		decimalString = strconv.Itoa(remainder) + decimalString
	}

	return decimalString
}

func multiplyBinaryStrings(bitstring1, bitstring2 string) string {

	// convert bitstrings to little endian slices
	bin1 := make([]bool, len(bitstring1))
	bin2 := make([]bool, len(bitstring2))

	for i, r := range bitstring1 {
		if int(r - '0') == 1 {
			bin1[len(bin1) - i - 1] = true
		} else {
			bin1[len(bin1) - i - 1] = false
		}	
	}

	for i, r := range bitstring2 {
		if int(r - '0') == 1 {
			bin2[len(bin2) - i - 1] = true
		} else {
			bin2[len(bin2) - i - 1] = false
		}	
	}

	// this will store the big endian bits of our result
	resultBits := make([]int, len(bitstring1)*len(bitstring2))

	for i := 0; i<len(bin1); i++ {
		for j:=0; j<len(bin2); j++ {
			
			if bin1[i] && bin2[j] {
				increaseBitInSlice(i+j, resultBits)
			}
		}
	}

	retVal := ""
	hasOne := false
	for i := len(resultBits) - 1; i>=0; i-- {
		if resultBits[i] == 1 {
			hasOne = true
		}

		if hasOne {
			retVal += strconv.Itoa(resultBits[i])
		}
	}

	return retVal
}

func main() {
	factorial := "1"
	for i:=2; i<=100; i++ {
		factorial = multiplyBinaryStrings(factorial, smallDecimalToBinaryString(i))
	}

	decimal := binaryStringToDecimalString(factorial)

	sum := 0

	for _, r := range decimal {
		sum += int(r - '0')
	}

	fmt.Println(sum)
}