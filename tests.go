package main

import (
	"fmt"
	"./crabMath"
	"math/big"
)


func main() {
	n := big.NewInt(10)

	fmt.Println(len(crabMath.BigPower(n, 3).String()))
	
}