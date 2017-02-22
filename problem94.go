/*
	Problem 94

	We are looking for almost-equilateral Heronian triangles where `a` is a given side.

	- `a` must be odd and a unit in Z/16Z
*/

package main

import (
	"fmt"
	//"math/big"
	//"./crabMath"
)

const N int = 1000*1000*1000


func main() {

	x_k := 2
	y_k := 1

	perimeterSum := 0

	for x_k < N {
		x_k1 := 2*x_k + 3*y_k
		y_k1 := 2*y_k + x_k

		x_k = x_k1
		y_k = y_k1

		a3_1 := 2*x_k + 1
		a3_2 := 2*x_k - 1

		if a3_1%3 == 0 {
			a := a3_1/3
			
			if (a+1)%2 == 0 || y_k%2 == 0 {
				// a is a solution!
				perimeter := 3*a + 1

				if perimeter <= N {
					perimeterSum += 3*a + 1	
				}
				
			}
		}

		if a3_2%3 == 0 {
			a := a3_2/3
			
			if (a+1)%2 == 0 || y_k%2 == 0 {
				// a is a solution!

				perimeter := 3*a - 1

				if perimeter <= N {
					perimeterSum += 3*a - 1	
				}
			}
		}
	}

	fmt.Println(perimeterSum)

	
}