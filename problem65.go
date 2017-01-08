/*
	Problem 65
*/

package main

import (
	"fmt"
	"math/big"
)

var cs [101][2](*big.Int)
var bs [101](*big.Int)

const N int = 100

func main() {
	for n := 2; n<101; n++ {
		i := 1
		if n%3 == 0 {
			i = (n/3)*2
		}
		bs[n] = big.NewInt(int64(i))
	}

	cs[N-1] = [2](*big.Int){ new(big.Int).Set(bs[N]), new(big.Int).Add(big.NewInt(1), new(big.Int).Mul(bs[N-1], bs[N])) }

	for n := N-2; n >= 2; n-- {
		cs[n] = [2](*big.Int){ new(big.Int).Set(cs[n+1][1]), new(big.Int).Add(cs[n+1][0], new(big.Int).Mul(bs[n], cs[n+1][1])) }
	}

	nominator := new(big.Int).Add(new(big.Int).Add(cs[2][1],cs[2][1]),cs[2][0]).String()


	sum := 0
	for _, v := range nominator {
		sum += int(v) - 48
	}

	fmt.Println(sum)
}