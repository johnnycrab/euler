/*
	Problem 72
*/

package main

import (
	"fmt"
	"./crabMath"
)

const N int = 1000*1000
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
	primeSieve = crabMath.PrimeSieve(N)

	/*sum := 0
	for d := 2; d <= N; d++ {
		sum += eulerPhi(d)
	} 

	fmt.Println(sum)*/

	for n := 2; n<=1000000; n++ {
		if eulerPhi(n) == 16 {
			fmt.Println(n)
		}
	}
}