/*
	Polynomial functions

	A polynomial \sum_{i=0}^{n}{a_i*X^i} is represented as a slice of size n+1 where the i-th
	entry corresponds to a_i

*/

package crabMath

func InsertIntoPolynomial(p []float64, x float64) float64 {
	result := 0.0
	for i := 0; i<len(p); i++ {
		pow := 1.0
		for j := 1; j<=i; j++ {
			pow *= x
		}
		result += pow*p[i]
	}

	return result
}

func MultiplyPolynomials(p1, p2 []float64) []float64 {
	result := make([]float64, len(p1) + len(p2) - 1)

	for i, a_i := range p1 {
		for j, a_j := range p2 {
			result[i+j] += a_i * a_j
		}
	}

	return result
}

func AddPolynomials(p1, p2 []float64) []float64 {
	m := Max(len(p1), len(p2))
	result := make([]float64, m)

	highestNonZero := 0

	for i := 0; i<m; i++ {
		var a1, a2 float64
		if i >= len(p1) {
			a1 = 0
		} else {
			a1 = p1[i]
		}

		if i >= len(p2) {
			a2 = 0
		} else {
			a2 = p2[i]
		}

		s := a1 + a2
		result[i] = s

		if s != 0 {
			highestNonZero = i
		}
	}

	return result[:highestNonZero + 1]
}

/*
	Calculates the interpolation polynomial p(X) where p(xs[i]) = fs[i]
	using Newton's divided differences method
*/
func InterpolationPolynomial(xs, fs []float64) []float64 {

	dividedDiffs := make([][]float64, len(xs))

	// prepopulate the divided diffs
	for i := 0; i<len(dividedDiffs); i++ {
		dividedDiffs[i] = []float64{fs[i]}
	}

	// calculate the remaining diffs from row to row
	for i := 1; i<len(dividedDiffs); i++ {
		for j := 1; j <= i; j++ {
			diff := (dividedDiffs[i][j-1] - dividedDiffs[i-1][j-1]) / (xs[i] - xs[i-j])

			dividedDiffs[i] = append(dividedDiffs[i], diff)
		}
	}

	// calculate the final interpolation polynomial
	interpolationPolynomial := []float64{fs[0]}
 
	for i := 1; i<len(xs); i++ {
		p := []float64{-1 * xs[0],1}

		// first calculate (X-x_0)*...*(X-x_0)
		for j := 1; j<i; j++ {
			p = MultiplyPolynomials(p, []float64{-1 * xs[j] ,1})
		}

		// multiply with divided diff
		p = MultiplyPolynomials(p, []float64{dividedDiffs[i][i]})

		interpolationPolynomial = AddPolynomials(interpolationPolynomial, p)
	}

	return interpolationPolynomial
}