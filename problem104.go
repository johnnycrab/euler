/*
	Problem 104
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {

	x_1 := big.NewInt(1)
	x_2 := big.NewInt(1)
	x_n := big.NewInt(0)

	n := 2 

	for {
		n++
		x_n.Add(x_1, x_2)

		x_1.Set(x_2)
		x_2.Set(x_n)

		s := x_n.String()

		if len(s) >= 18 {
			lastDigits := s[len(s) - 9:]
			firstDigits := s[:9]

			if isPanDigital(lastDigits) && isPanDigital(firstDigits) {
				fmt.Println(n)
				break
			}	
		}
		
	}
}

func isPanDigital(s string) bool {
	check := make([]int, 9)

	for _, v := range s {
		if v > 48 {
			check[int(v) - 49] = 1	
		}
		
	}

	for _, v := range check {
		if v == 0 {
			return false
		}
	}

	return true
}