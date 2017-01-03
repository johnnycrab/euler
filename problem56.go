/*
	Problem 56
*/

package main

import (
	"fmt"
	"./crabMath"
	"math/big"
)

func digitSum(N *big.Int) int {
	s := N.String()

	digitSum := 0
	for _, v := range s {
		digitSum += int(v) - 48
	}

	return digitSum
}

func main() {
	var a, b int64

	biggest := 0

	for a = 1; a < 100; a++ {
		for b = 1; b < 100; b++ {
			sum := digitSum(crabMath.IntegerPowerBig(a, b))		
			if sum > biggest {
				biggest = sum
			}
		}
	}
	
	fmt.Println(biggest)
}