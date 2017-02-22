/*
	Problem 112
*/

package main

import (
	"fmt"
	"./crabMath"
)

func isBouncy(n int) bool {
	digits := crabMath.GetDigits(n)
	//fmt.Println(digits)

	decided := false
	ascending := false

	for i, v := range digits {
		if i-1 >= 0 {
			if decided {
				if ascending && v < digits[i-1] {
					return true
				}

				if !ascending && v > digits[i-1] {
					return true
				}
			} else if digits[i-1] != v {
				decided = true
				ascending = digits[i-1] < v
			}
		}
	}

	return false
}

func main() {

	n := 10
	count := 0
	ratio := 0

	for {
		n++

		if isBouncy(n) {
			count++
		}

		ratio = (count*100)/n

		if ratio == 99 {
			fmt.Println(count)
			fmt.Println(n)
			break
		}
	}
}