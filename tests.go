package main

import (
	"fmt"
	"./crabMath"
)

func main() {
	
	a := crabMath.NChooseKSets(10,6)
	fmt.Println(a)
}