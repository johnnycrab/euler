/*
	Problem:
	Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {

	// measure execution time
	start := time.Now()

	greatestPalindromeNum := 0

	// we don't have too many possibilities, so let's just check them all
	for i := 100; i<1000; i++ {
		for j := 100; j<1000; j++ {
			num := i*j
			// check if it is a palindrome
			asString := strconv.Itoa(num)
			isPalindrome := true
			count := len(asString)
			for k := 0; k < count / 2.0; k++ {
				if asString[k] != asString[count - k - 1] {
					isPalindrome = false
					break
				}
			}

			if isPalindrome && num > greatestPalindromeNum {
				greatestPalindromeNum = num
			}
		}
	}

	fmt.Println(greatestPalindromeNum)

	elapsed := time.Since(start)

	fmt.Println("Took ", elapsed)
}