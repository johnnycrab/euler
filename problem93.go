/*
	Problem 93

	We choose four distinct digits as a 9 choose 4 set. 
	Arithmetic operators are expressed as 1, 2, 3, 4. Distribution of those is then {1,...,4}^3
	The brackets are represented as S_4, i.e. a permutation represents the order in which the arithmetic operations
	should be executed.

	This makes it (9 choose 4)*4^3*4*3*2 = 193.536 possible arithmetic expressions to check, which is doable.
	Plus we need to check extra the order (a+b)+(c+d)

	For each set of four distinct digits we just mark off the results of each expression in a boolean array of size 3100, which is sufficient.

	1: +
	2: -
	3: *
	4: /

*/

package main

import (
	"fmt"
	"./crabMath"
	"math"
)

func main() {
	highestResultObtained := 0
	var digitsResult []int

	digitsDist := crabMath.NChooseKSets(9, 4)

	arithmeticOperatorsDist := crabMath.SetNPowerK(4,3)
	arithmeticOrderDist := crabMath.S_n(4)


	for _, digits := range digitsDist {
		results := make([]bool, 3500)

		digitsFloats := []float64{float64(digits[0]), float64(digits[1]), float64(digits[2]), float64(digits[3])}

		for _, arithmeticOperators := range arithmeticOperatorsDist {

			for _, arithmeticOrder := range arithmeticOrderDist {
		
				t1 := operate(arithmeticOperators[0], digitsFloats[arithmeticOrder[0] - 1], digitsFloats[arithmeticOrder[1] - 1])
				t2 := operate(arithmeticOperators[1], t1, digitsFloats[arithmeticOrder[2] - 1])
				resultFloat := operate(arithmeticOperators[2], t2, digitsFloats[arithmeticOrder[3] - 1])

				if resultFloat >= 0 && resultFloat == math.Trunc(resultFloat) {					
					// is integer

					results[int64(math.Trunc(resultFloat))] = true
				}

				// at last there is one possibility we haven't considered yet, i.e. (a+b)+(c+d) (taking the result of two pairs, then operating on them)
				resultFloat = operate(arithmeticOperators[1], operate(arithmeticOperators[0], digitsFloats[arithmeticOrder[0] - 1], digitsFloats[arithmeticOrder[1] - 1]), operate(arithmeticOperators[2], digitsFloats[arithmeticOrder[2] - 1], digitsFloats[arithmeticOrder[3] - 1]))

				if resultFloat >= 0 && resultFloat == math.Trunc(resultFloat) {					
					// is integer

					results[int64(math.Trunc(resultFloat))] = true
				}
			}
		}

		i := 1
		for results[i] == true {
			i++
		}

		if i-1 > highestResultObtained {
			highestResultObtained = i-1
			digitsResult = digits
		}
	}

	fmt.Println(digitsResult)
	fmt.Println(highestResultObtained)
}

func prettyPrint(digits []int, operators []int, order []int) {
	fmt.Println(digits[order[0] - 1], operatorAsString(operators[0]) , digits[order[1] - 1], operatorAsString(operators[1]), digits[order[2] - 1], operatorAsString(operators[2]), digits[order[3] - 1])
}

func operatorAsString(o int) string {
	if o == 1 {
		return "+"
	} else if o == 2 {
		return "-"
	} else if o == 3 {
		return "*"
	} else if o == 4 {
		return "/"
	}
	return ""
}

func operate(operator int, a float64, b float64) float64 {
	if operator == 1 {
		return a+b
	} else if operator == 2 {
		return a-b
	} else if operator == 3 {
		return a*b
	} else if operator == 4 {
		return a/b
	}

	return .0
}