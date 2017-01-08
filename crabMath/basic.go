package crabMath

import (
	"strconv"
	"math/big"
)

func Abs(a int) int {
	if a < 0 {
		return a*(-1)
	}

	return a
}

func Max(a int, b int) int {
	if a >= b {
		return a
	} 
	return b
}

func Min(a int, b int) int {
	if a <= b {
		return a
	} 
	return b
}

// concatenates two integers a, b
func Concatenate(a, b int) int {
	res, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))

	return res
}

// for small integer powers
func Power(n, pow int) int {
	if pow == 0 {
		return 1
	}

	if pow == 1 {
		return n
	}

	result := n
	for i := 2; i<= pow; i++ {
		result *= n
	}

	return result
}

func BigPower(n *big.Int, power int) *big.Int {
	if power == 0 {
		return big.NewInt(0)
	}

	if power == 1 {
		return n
	}

	result := new(big.Int).Set(n)
	for i := 2; i <= power; i++ {
		result.Mul(result, n)
	}

	return result
}