/*
	Problem 52
*/

package main

import (
	"fmt"
	"./crabMath"
)	

func main() {
	i := 0
	for {
		i++

		num := i
		is := true
		for j := 1; j<=5; j++ {
			if !crabMath.IntSlicesContainSame(crabMath.GetDigits(num), crabMath.GetDigits(num + i)) {
				is = false
				break
			}
			num += i
		}

		if is {
			fmt.Println(i)
			break
		}
	}
}