/*
	Problem 78
*/

package main

import "fmt"

// save congruences of p(n) mod 10^6
const mill int = 1000*1000
var p_congs [mill]int
var pentagonalNums [mill]int

func makePentagonalNums() {
	for k := 1; k < mill/2; k++ {
		// -k is at index k + 500000
		pentagonalNums[k] = (k*(3*k - 1))/2
		pentagonalNums[mill/2 + k] = (k*(3*k + 1))/2
	}
}

func p_cong(n int) int {
	k := 0
	signum := -1
	congruence := 0
	for {
		k++
		signum *= -1

		penta_pos := pentagonalNums[k]
		penta_neg := pentagonalNums[mill/2+k]

		if n - penta_pos < 0 {
			break
		} else {
			congruence = (congruence + signum * p_congs[n-penta_pos])%mill
			
			if n - penta_neg >= 0 {
				congruence = (congruence + signum * p_congs[n-penta_neg])%mill				
			}
		}
	}

	return congruence
}

func main() {
	makePentagonalNums()

	p_congs = [1000000]int{}
	p_congs[0] = 1

	n := 0
	for {
		n++
		p_n_cong := p_cong(n)
		if p_n_cong == 0 {
			fmt.Println(n)
			break
		}
		p_congs[n] = p_n_cong
	}
}