/*
	Rational Polynomials.
	Same as polynomials.go, coefficients are rationals however
*/

package crabMath

// import "fmt"

func InsertIntoRationalPolynomial(p []Rational, x Rational) Rational {
	result := Rational{A: 0, B: 1}
	for i := 0; i<len(p); i++ {
		pow := Rational{A: 1, B: 1}
		for j := 1; j<=i; j++ {
			pow = MultiplyRationals(pow, x)
		}
		result = AddRationals(result, MultiplyRationals(pow, p[i]))
	}

	return result
}

func MultiplyRationalPolynomials(p1, p2 []Rational) []Rational {
	result := make([]Rational, len(p1) + len(p2) - 1)
	
	FillRationalSliceWithZeros(result)

	for i, a_i := range p1 {
		for j, a_j := range p2 {
			result[i+j] = AddRationals(result[i+j], MultiplyRationals(a_i, a_j))
		}
	}

	return result
}

func AddRationalPolynomials(p1, p2 []Rational) []Rational {
	m := Max(len(p1), len(p2))
	result := make([]Rational, m)

	highestNonZero := 0

	for i := 0; i<m; i++ {
		var a1, a2 Rational
		if i >= len(p1) {
			a1 = Rational{A: 0, B: 1}
		} else {
			a1 = p1[i]
		}

		if i >= len(p2) {
			a2 = Rational{A: 0, B: 1}
		} else {
			a2 = p2[i]
		}

		s := AddRationals(a1, a2)
		result[i] = s

		if s.A != 0 {
			highestNonZero = i
		}
	}

	return result[:highestNonZero + 1]
}

func InterpolationRationalPolynomial(xs, fs []Rational) []Rational {

	dividedDiffs := make([][]Rational, len(xs))

	// prepopulate the divided diffs
	for i := 0; i<len(dividedDiffs); i++ {
		dividedDiffs[i] = []Rational{fs[i]}
	}

	// calculate the remaining diffs from row to row
	for i := 1; i<len(dividedDiffs); i++ {
		for j := 1; j <= i; j++ {
			diff := DivideRationals(SubtractRationals(dividedDiffs[i][j-1], dividedDiffs[i-1][j-1]), SubtractRationals(xs[i], xs[i-j]))

			dividedDiffs[i] = append(dividedDiffs[i], diff)
		}
	}

	// calculate the final interpolation polynomial
	interpolationPolynomial := []Rational{fs[0]}
 
	for i := 1; i<len(xs); i++ {
		p := []Rational{MultiplyRationals(Rational{A: -1, B: 1}, xs[0]), Rational{A: 1, B:1 }}

		// first calculate (X-x_0)*...*(X-x_0)
		for j := 1; j<i; j++ {
			p = MultiplyRationalPolynomials(p, []Rational{MultiplyRationals(Rational{A: -1, B: 1}, xs[j]), Rational{A: 1, B:1 }})
		}

		// multiply with divided diff
		p = MultiplyRationalPolynomials(p, []Rational{dividedDiffs[i][i]})

		interpolationPolynomial = AddRationalPolynomials(interpolationPolynomial, p)
	}

	return interpolationPolynomial
}