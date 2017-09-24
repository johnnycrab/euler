/*
	Problem 88
*/

package main

import (
	"fmt"
	"./crabMath"
)

const maxK int = 12000
const N int = 10000000
var primeSieve []int
var prodSums [][]int
var missing []int
var result int

func main() {
	prodSums = make([][]int, maxK +1)
	prodSums[0] = make([]int, 1)
	prodSums[1] = make([]int, 1)

	primeSieve = crabMath.PrimeSieve(N)

	checkedForMissing := false
	missing = []int{}

	for n := 4; n<=N; n++ {
		_, factors := crabMath.SimplePrimeFactorizationWithSieve(n, primeSieve)
		k := getNeededK(n, factors)

		if n%5000 == 0 {
			fmt.Println("k: ", k, "factors: ", factors)
		}

		if k > maxK {
			if !checkedForMissing {
				checkedForMissing = true
				// check which ones are missing. then let it run until all have been removed
				fmt.Println("k is greater than ", maxK, ". Checking for missing ks")
				for i, v := range prodSums {
					if len(v) == 0 {
						missing = append(missing, i)
					}
				}
				fmt.Println("Missing: ", missing)
			}

			checkMinProdSums(n, k, factors)

			if len(missing) == 0 {
				break
			}
		} else {
			if len(prodSums[k]) == 0 {
				// we just store the factors. Everything else is then considered to be filled up with 1's
				prodSums[k] = factors
			}

			// now see what other k values we could satisfy with this "n"
			checkMinProdSums(n, k, factors)
		}
			
	}


	fmt.Println("got it, calculating result")
	result = 0
	used := []int{}
	for k := 2; k<=maxK; k++ {
		prodSum := prodSums[k]

		prod := 1
		for _, v := range prodSum {
			prod *= v
		}

		if !crabMath.IntSliceContains(used, prod) {
			result += prod
			used = append(used, prod)
		}
	}

	fmt.Println(result)
}

// n is the productSum, k is the current length for a min prod sum, factors are factors of n
func checkMinProdSums(n int, k int, factors []int) {
	// now filter out all factor pairs
	if len(factors) >= 3 {
		// keep track of the factors we have already checked, so that we don't always repeat ourselves (for example for n = 2^m)
		usedFactorsSlice := [][2]int{}

		for i := 0; i<len(factors); i++ {
			for j := i+1; j<len(factors); j++ {
				a := factors[i]
				b := factors[j]

				used := false
				for _, usedFactors := range usedFactorsSlice {
					if (usedFactors[0] == a && usedFactors[1] == b) || (usedFactors[0] == b && usedFactors[1] == a) {
						used = true
					}
				}

				if !used {
					usedFactorsSlice = append(usedFactorsSlice, [2]int{a,b})

					// calculate the number of 1's we need to fill a*b up
					q := (a*b) - (a+b)
					// q must be <= k - len(factors), which is the number of 1's we can use
					if q <= k - len(factors) {
						// this is the new k we can have with the same min prod sum
						k_tilde := k - 1 - q
						
						// we build up the new factors, where a,b are replaced by a*b
						newFactors := make([]int, len(factors) - 1)
						index := -1
						for c, val := range factors {
							if c != i && c != j {
								index++
								newFactors[index] = val
							}
						}
						newFactors[len(factors) - 2] = a*b
						
						// set it into the prodSumSlice
						if k_tilde <= maxK && len(prodSums[k_tilde]) == 0 {
							prodSums[k_tilde] = newFactors

							// check if the new k_tilde can erase a missing k
							if len(missing) > 0 {
								
								for i, v := range missing {
									if k_tilde == v {
										
										fmt.Println("Found missing k: ", k_tilde, "With n: ", n)
										
										missing = append(missing[:i], missing[i+1:]...)
										fmt.Println("Remaining missing: ", missing)
										break
									}
								}
							}
						}

						// do the same procedure again
						checkMinProdSums(n, k_tilde, newFactors)
					} 
				}

				

			}
		}
	}
}

func getNeededK(n int, factors []int) int {
	s := 0
	for _, v := range factors {
		s += v
	}

	return len(factors) + n - s
}