/*
	Problem 33
*/

package main

import "fmt"

var primes []int

func cancelFraction(nom int, denom int) (int, int) {
	for {
		couldCancel := false

		for _, prime := range primes {
			if nom%prime == 0 && denom%prime == 0 {
				couldCancel = true
				nom /= prime
				denom /= prime
			}
		}

		if !couldCancel {
			break
		}
	}

	return nom, denom
}

func main() {
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	resultNom := 1
	resultDenom := 1

	for nom := 10; nom <= 99; nom++ {
		for denom := nom+1; denom <= 99; denom++ {

			if nom%10 != 0 && denom%10 != 0 {

				// find if there is a common digit
				digitsNom := [2]int{nom%10, (nom-(nom%10))/10}
				digitsDenom := [2]int{denom%10, (denom-(denom%10))/10}

				possibilities := [][2]int{}
				var possibilityNom, possibilityDenom int

				if digitsNom[0] == digitsDenom[0] {
					possibilityNom, possibilityDenom = cancelFraction(digitsNom[1], digitsDenom[1])
					possibilities = append(possibilities, [2]int{possibilityNom, possibilityDenom})
				}

				if digitsNom[0] == digitsDenom[1] {
					possibilityNom, possibilityDenom = cancelFraction(digitsNom[1], digitsDenom[0])
					possibilities = append(possibilities, [2]int{possibilityNom, possibilityDenom})
				}

				if digitsNom[1] == digitsDenom[0] {
					possibilityNom, possibilityDenom = cancelFraction(digitsNom[0], digitsDenom[1])
					possibilities = append(possibilities, [2]int{possibilityNom, possibilityDenom})
				}

				if digitsNom[1] == digitsDenom[1] {
					possibilityNom, possibilityDenom = cancelFraction(digitsNom[0], digitsDenom[0])
					possibilities = append(possibilities, [2]int{possibilityNom, possibilityDenom})
				}

				cancelledNom, cancelledDenom := cancelFraction(nom, denom)

				for _, possibility := range possibilities {
					if possibility[0] == cancelledNom && possibility[1] == cancelledDenom {
						resultNom *= nom
						resultDenom *= denom
						break
					}
				}
			}
		}
	}

	_, cancelledResultDenom := cancelFraction(resultNom, resultDenom)

	fmt.Println(cancelledResultDenom)
}