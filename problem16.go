/*
	Problem 16:
	What is the sum of the digits of the number 2^1000?
*/

package main

import(
	"fmt"
	"strconv"
)

func divideWithDecimalRemainder(bitstring string, divideBy int) (string, int) {
	remainder := 0
	hasOne := false
	retString := ""
	for i, r := range bitstring {
		bit := int(r - '0')
		
		remainder += remainder + bit

		if remainder >= divideBy {
			remainder -= divideBy
			retString += "1"
			hasOne = true
		} else if hasOne || (i==len(bitstring)-1) {
			retString += "0"
		}
	}

	return retString, remainder
}

func binaryStringToDecimalString(bitstring string) string {
	decimalString := ""

	remainder := 0

	for bitstring != "0" {
		bitstring, remainder = divideWithDecimalRemainder(bitstring, 10)
		decimalString = strconv.Itoa(remainder) + decimalString
	}

	return decimalString
}

func main() {
	pow := "1"

	for i:=1; i<=1000; i++ {
		pow += "0"
	}

	decimal := binaryStringToDecimalString(pow)
	sum := 0

	for _, r := range decimal {
		sum += int(r - '0')
	}

	fmt.Println(sum)
}