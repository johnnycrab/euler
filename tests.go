package main

import (
	"fmt"
	"./crabMath"
)

func main() {
	
	r1 := crabMath.Rational{A: 8, B: 2}
	r2 := crabMath.Rational{A: 4, B: 1}
	fmt.Println(crabMath.RationalsEqual(r1, r2), r1, r2)
}