/*
	Problem 206
*/

package main

import (
	"fmt"
	"math/big"
	//"./crabMath"
)

const N int64 = 1000000000
const Two int64 = 2

var indices [10]int

func check(s string) bool {
	if len(s) != 19 {
		return false
	}

	// 1_2_3_4_5_6_7_8_9_0

	for i, v := range indices {
		if int(s[v]) - 48 != (i + 1)%10 {
			//fmt.Println(int(s[v]) - 48, i, v)
			return false
		}
	}

	return true
}

func main() {
	indices = [10]int{0,2,4,6,8,10,12,14,16,18}

	//fmt.Println(check("1122434455667788990"))

	// has 10 digits and is a mutliple of ten

	for j := N; j < N*10; j+= 10 {
		a := big.NewInt(j)
		a.Mul(a, a)

		str := a.String()

		if check(str) {
			fmt.Println(j)
			break
		}
	}
}