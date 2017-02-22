/*
	Problem 381
*/

package main

import (
	"fmt"
	"./crabMath"
)


// find multiplicative inverse of a in Z/pZ
func multiplicativeInverse(a, p int) int {
	t := 0
	newt := 1
	r := p
	newr := a
	for newr != 0 {
		quotient := r/newr
		_t := t
		t = newt
		newt = _t - quotient * newt

		_r := r
		r = newr
		newr = _r - quotient * newr
	}

	if t < 0 {
		t += p
	}

	return t%p
}

func main() {


	primes := crabMath.PrimesUpTo(1000*1000*100)

	sum := 0

	for _, p := range primes {
		if p >= 5 {
			inv_2 := multiplicativeInverse(2, p)
			inv_3 := multiplicativeInverse(3, p)
			inv_2_3 := (((inv_2*inv_2)%p)*inv_2)%p

			s := (-1 * inv_2 + inv_2*inv_3 - inv_2_3*inv_3)%p
			if s < 0 {
				s+=p
			}

			sum += s
		}
	}

	fmt.Println(sum)
}