/*
	Problem 64
*/

package main

import (
	"fmt"
	"math"
)

const LIMIT float64 = 0.0001

func main() {

	count := 0
	for n := 2.0; n <= 10000; n++ {
		period := getContinuedFractionPeriod(n)
		//fmt.Println(int(n), period)
		fmt.Println(int(n), len(period))
		if len(period)%2 == 1 {
			count++
		}
	}

	fmt.Println(count)
}

func getContinuedFractionPeriod(n float64) []int {
	contFrac := []int{}

	// start
	sqrt := math.Sqrt(n)

	a_0 := int(sqrt)
	b_0 := 1

	if sqrt == float64(a_0) {
		return contFrac
	}

	a_k := a_0
	b_k := b_0
	for {

		f_a_k := float64(a_k)
		f_b_k := float64(b_k)
		f_n := float64(n)
		f_b_k1 := (f_n - f_a_k * f_a_k) / f_b_k

		if float64(int(f_b_k1)) != f_b_k1 {
			fmt.Println("PROBLEM!")
		}

		c := int( (sqrt + f_a_k)/f_b_k1 )
		f_a_k1 := float64(c) * f_b_k1 - f_a_k

		contFrac = append(contFrac, c)

		a_k = int(f_a_k1)
		b_k = int(f_b_k1)

		if a_k == a_0 && b_k == b_0 {
			break
		}
	}

	return contFrac
}

func addDigitsToSlice(s *[]int, n int) {
	digits := []int{}
	
	for n > 0 {
		digits = append(digits, n%10)
		n = (n - (n%10)) / 10
	}

	for i := 0; i<len(digits); i++ {
		*s = append(*s, digits[len(digits) - 1 - i])
	}
}

// given a_1,...,a_n of continued fractions, calculates the remainder
func calculateRemainder(sqrt float64, contFrac []int) float64 {
	n := len(contFrac)
	fraction := [2]int{1, contFrac[n - 1]}

	for i := 1; i< n; i++ {
		fraction = add([2]int{contFrac[n - 1 - i], 1}, fraction)

		if i != n-1 {
			fraction = invert(fraction)
		}
	}

	fmt.Println((float64(fraction[0])/float64(fraction[1])))
	return sqrt - (float64(fraction[0])/float64(fraction[1]))
}

// Helper functions for working with fractions
func invert(frac [2]int) [2]int {
	return [2]int{frac[1], frac[0]}
}

func add(a, b [2]int) [2]int {
	return [2]int{a[0]*b[1]+a[1]*b[0], a[1]*b[1]}
}