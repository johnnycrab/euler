package main

import (
	"fmt"
	"math/big"
)

// a^b
func power(a *big.Int, b int) *big.Int {
	res := new(big.Int).Set(a)
	
	if b >= 2 {
		for i := 2; i<=b; i++ {
			res.Mul(a, res)
		}
	}

	return res
}

func main() {
	
	var powers [99 * 99](*big.Int)

	var a int64 = 2
	b := 2
	i := 0

	for a <= 100 {
		aBig := big.NewInt(a)

		for b <= 100 {
			
			powers[i] = new(big.Int).Set(power(aBig, b))

			b++
			i++
		}
		b = 2
		a++
	}

	// we have all powers. now iterate over all of them and remove duplicates
	var results []string
	for _, val := range powers {
		s := val.String()

		found := false
		for _, r := range results {
			if r == s {
				found = true
				break		
			}
		}

		if !found {
			results = append(results, s)
		}
	}

	fmt.Println(len(results))
}