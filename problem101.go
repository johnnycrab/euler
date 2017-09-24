/*
	Problem 101
*/

package main

import (
	"fmt"
	"./crabMath"
)

var generatingPolynomial []crabMath.Rational = []crabMath.Rational{crabMath.Rational{A: 1, B: 1}, crabMath.Rational{A: -1, B: 1}, crabMath.Rational{A: 1, B: 1}, crabMath.Rational{A: -1, B: 1}, crabMath.Rational{A: 1, B: 1}, crabMath.Rational{A: -1, B: 1}, crabMath.Rational{A: 1, B: 1}, crabMath.Rational{A: -1, B: 1}, crabMath.Rational{A: 1, B: 1}, crabMath.Rational{A: -1, B: 1}, crabMath.Rational{A: 1, B: 1}}

func main() {

	xs := make([]crabMath.Rational, 11)
	fs := make([]crabMath.Rational, 11)

	for i := 0; i<11; i++ {
		xs[i] = crabMath.Rational{A: i+1, B: 1}
		fs[i] = crabMath.InsertIntoRationalPolynomial(generatingPolynomial, xs[i])
	}

	result := crabMath.Rational{A: 0, B: 1}

	for i := 1; i<11; i++ {
		interpolation := crabMath.InterpolationRationalPolynomial(xs[:i], fs[:i])
		for j := 0; j<len(xs); j++ {
			inserted := crabMath.InsertIntoRationalPolynomial(interpolation, xs[j])
			if !crabMath.RationalsEqual(inserted, fs[j]) {
				result = crabMath.AddRationals(result, inserted)
				break
			}
		}
	}

	fmt.Println(result)
}