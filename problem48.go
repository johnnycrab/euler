package main

import (
	"fmt"
	"math/big"
)

func power(num *big.Int, n int) {
	toPow := new(big.Int)

	toPow.Set(num)

	for i := 2; i<=n; i++ {
		num.Mul(num, toPow)
	}
}

func main() {
	result := big.NewInt(0)

	for i := 1; i <= 1000; i++ {

		num := big.NewInt(int64(i))

		power(num, i)

		result.Add(num, result)

	}

	resString := result.String()

	for i := len(resString) - 10; i < len(resString); i++ {
		fmt.Print(string(resString[i]))
	}

	fmt.Println("")
}