/*
	Problem 80
*/

package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

const PRECISION uint = 500

func main() {
	result := 0

	for n := 1; n<=100; n++ {
		if math.Sqrt(float64(n)) != float64(int(math.Sqrt(float64(n)))) {
			// is irrational
			result += getDigitSumOfSqrt(n)
		}
	}

	fmt.Println(result)
}

func getDigitSumOfSqrt(a int) int {
	sum := 0

	b := bigSquareRootNewton(int64(a), PRECISION)
	sqrtString := fmt.Sprint(b)


	digitString := strings.Split(sqrtString, ".")[0] + strings.Split(sqrtString, ".")[1]

	for i, r := range digitString {
		if i < 100 {
			sum += int(r) - 48
		} else {
			break
		}
		
	}

	return sum
}

// calculate square root using newton's method (x_{n+1} = 0.5 * (x_n + a/x_n))
func bigSquareRootNewton(a int64, precisionBits uint) (*big.Float) {
	// calculate number of steps needed
	steps := int(math.Log2(float64(precisionBits)))

	a_big := new(big.Float).SetPrec(precisionBits).SetInt64(a)
	half := new(big.Float).SetPrec(precisionBits).SetFloat64(0.5)

	// we take x_1 to be the square root
	x_n := new(big.Float).SetPrec(precisionBits).SetFloat64(math.Sqrt(float64(a)))

	x_n1 := new(big.Float).SetPrec(precisionBits)

	for i := 0; i<steps; i++ {
		x_n1.Quo(a_big, x_n)
		x_n1.Add(x_n1, x_n)
		x_n1.Mul(half, x_n1)

		x_n.Set(x_n1)
	}

	return x_n
}