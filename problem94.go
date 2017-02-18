/*
	Problem 94

	We are looking for almost-equilateral Heronian triangles where `a` is a given side.

	- `a` must be odd
*/

package main

import (
	"fmt"
	"./crabMath"
)

const N int = 1000*1000*1000


func main() {
	sum := 0

	pythTriples := crabMath.GetPrimitivePythagoreanTriples(20000)

	for _, triple := range pythTriples {
		// a is even. Don't need that
		a := triple[2]
		if a%2 == 0 {
			continue
		}

		b_1 := (a+1)/2
		b_2 := (a-1)/2

		if triple[0] == b_1 || triple[0] == b_2 {
			//fmt.Println(a, triple[0]*2)
			perimeter := a+a+triple[0]
			if perimeter <= N {
				sum += a+a+triple[0]	
			}
			
		} 
		if triple[1] == b_1 || triple[1] == b_2 {
			//fmt.Println(a, triple[1]*2)

			perimeter := a+a+triple[1]
			if perimeter <= N {
				sum += a+a+triple[1]	
			}
		}
	}

	fmt.Println(sum)
}