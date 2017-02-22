/*
	Problem 100
*/

package main

import (
	"fmt"
	//"math"
)

func main() {

	trillion := 1000*1000*1000*1000

	x_0 := 1
	y_0 := 1
	for {
		x := 3*x_0 + 4*y_0
		y := 3*y_0 + 2*x_0

		x_0 = x
		y_0 = y

		n := (x+1)/2
		if n > trillion {
			fmt.Println((y+1)/2)
			break
		}
	}
}