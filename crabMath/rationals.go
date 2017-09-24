/*
	Rationals
*/

package crabMath


/*
	rational number, represented as a/b
*/
type Rational struct {

	A, B int
}

func (r *Rational) Cancel() {
	if r.A == 0 {
		r.B = 1
		return
	}

	n := 2

	for (n*n < Abs(r.A) && n*n < Abs(r.B)) {
		if r.A%n == 0 && r.B%n == 0 {
			r.A /= n
			r.B /= n
		} else {
			n++
		}
	}

	if r.A%r.B == 0 {
		r.A /= r.B
		r.B = 1
	}

	if r.B%r.A == 0 {
		r.B /= r.A
		r.A = 1
	}
}

func AddRationals(r1, r2 Rational) Rational {
	r := Rational{ A: r1.A*r2.B + r2.A*r1.B , B: r1.B*r2.B}
	r.Cancel()
	return r
}

func MultiplyRationals(r1, r2 Rational) Rational {
	r := Rational{ A: r1.A*r2.A, B: r1.B*r2.B }
	r.Cancel()
	return r
}

// r1-r2
func SubtractRationals(r1, r2 Rational) Rational {
	r := Rational{ A: r1.A*r2.B - r2.A*r1.B , B: r1.B*r2.B}
	r.Cancel()
	return r
}

// r1/r2
func DivideRationals(r1, r2 Rational) Rational {
	r := Rational{ A: r1.A*r2.B, B: r1.B*r2.A}
	r.Cancel()
	return r
}

func FillRationalSliceWithZeros(r []Rational) {
	for i := 0; i<len(r); i++ {
		r[i] = Rational{ A: 0, B: 1}
	}
}

func RationalsEqual(r1, r2 Rational) bool {
	r1.Cancel()
	r2.Cancel()

	return r1.A == r2.A && r1.B == r2.B
}