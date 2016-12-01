/*
	Problem 40
*/

package main

import "fmt"

const n int = 1000000

func digitsOfNumber(num int) []int {
	digits := []int{}

	for num > 0 {
		digits = append(digits, num%10)
		num = (num - num%10)/10
	}

	return digits
}

// returns number of added digits
func addDigits(startIndex int, digits *[n]int, ofNumber int) int {
	digitsOfNum := digitsOfNumber(ofNumber)

	for i := len(digitsOfNum) - 1; i >= 0 && startIndex < n; i-- {
		(*digits)[startIndex] = digitsOfNum[i]
		startIndex++
	}

	return len(digitsOfNum)
}

func main() {
	
	digits := [n]int{}

	index := 0
	numberToAddDigitsOf := 1

	for index < n {
		numOfDigitsAdded := addDigits(index, &digits, numberToAddDigitsOf)
		index += numOfDigitsAdded
		numberToAddDigitsOf++
	}

	fmt.Println(digits[0] * digits[9] * digits[99] * digits[999] * digits[9999] * digits[99999] * digits[999999])

}