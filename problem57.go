/*
	Problem 57

	If p_n is the n-th iteration, each p_n is of the the form 1 + a_n
	Let a_(n-1) be the fraction k/m, then have a_n = m/(2m+k) and p_n = (3m+k)/(2m+k)
*/

package main

import (
	"fmt"
	"math/big"
)

var p_ns, a_ns [1000][2](*big.Int) // 2 element slice represents fraction


func main() {
	p_ns[0] = [2](*big.Int){big.NewInt(3),big.NewInt(2)}
	a_ns[0] = [2](*big.Int){big.NewInt(1),big.NewInt(2)}

	count := 0

	for n := 2; n<=1000; n++ {
		k := a_ns[n-2][0]
		m := a_ns[n-2][1]
		two_m := new(big.Int).Add(m,m)
		three_m := new(big.Int).Add(two_m,m)
		a_ns[n-1] = [2](*big.Int){new(big.Int).Set(m), new(big.Int).Add(two_m,k)}
		p_ns[n-1] = [2](*big.Int){new(big.Int).Add(three_m, k), new(big.Int).Add(two_m,k)}

		if len(p_ns[n-1][0].String()) > len(p_ns[n-1][1].String()) {
			count++
		}
	}

	fmt.Println(count)
}