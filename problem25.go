package main

import (
	"fmt"
	"math/big"
)

func main() {
	
	b := big.NewInt(1)
	c := big.NewInt(1)
	temp := big.NewInt(0)
	index := 2

	for len(c.String()) < 1000 {

		index++
		temp.Set(c)
		c.Add(c,b)
		b.Set(temp)

	}

	fmt.Println(index)
}