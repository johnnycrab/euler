/*
	Problem 145
*/

package main

import (
	"fmt"
	"./crabMath"
)

func isReversible(n int) bool {
	digits := crabMath.GetDigits(n)
	
	pow := 1
	reversed := 0
	for i := 0; i<len(digits);i++ {
		reversed += digits[len(digits) - 1 - i]*pow
		pow *= 10
	}

	sum := n + reversed

	if sum%2 == 0 {
		return false
	} else {

		digits := crabMath.GetDigits(sum)
		for _, d := range digits {
			if d%2 == 0 {
				return false
			}
		}
	}

	return true
}

func main() {

	count := 0

	mc := 0 
	for n := 1; n<1000*1000*1000;n++ {
		if n%1000000 == 0 {
			mc++
			fmt.Println(mc)
		}
		if n%10 != 0 {
			if isReversible(n) {
				count++
			}
		}
	}

	fmt.Println(count)
}