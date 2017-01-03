/*
	Problem 63
*/

package main

import (
	"fmt"
	"./crabMath"
	"math/big"
)

func main() {

	result := 9

	// for n >= 2, only nth-powers of 4-9 are possible


	for a := int64(4); a <= 9; a++ {
		A := big.NewInt(a)

		n := 2

		for {
			pow := crabMath.BigPower(A, n)

			length := len(pow.String())

			if n > length {
				break
			}

			if length == n {
				result++
			}

			n++
		}
	}

	fmt.Println(result)
}