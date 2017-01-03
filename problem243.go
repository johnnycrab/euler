/*
	Problem 243
*/

package main

import (
	"fmt"
	"./crabMath"
)

const N int = 1000*1000*1000
const limit float64 = 15499.0 / 94744.0
var primeSieve []int

func eulerPhiForPrimePower(prime, power int) int {
	phi := prime - 1
	for k := 1; k <= power - 1; k++ {
		phi *= prime
	}

	return phi
}

func eulerPhi(n int) int {
	_, factorization := crabMath.SimplePrimeFactorizationWithSieve(n, primeSieve)

	phi := 1

	lastPrime := factorization[0]
	
	power := 1
	for i := 1; i < len(factorization); i++ {
		prime := factorization[i]
		if lastPrime != prime {
			phi *= eulerPhiForPrimePower(lastPrime, power)

			lastPrime = prime
			power = 1
		} else {
			power++
		}
	}

	phi *= eulerPhiForPrimePower(lastPrime, power)

	return phi
}

func main() {
	primeSieve := crabMath.PrimeSieve(N)
	fmt.Println(len(primeSieve))
	/*for n := 2; n < N; n++ {
		
		if float64(eulerPhi(n))/float64(n) <= 0.17 {
			fmt.Println(n)
			break
		}
	}*/
}