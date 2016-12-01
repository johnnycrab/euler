/*
	Problem 39
*/

package main

import "fmt"

func getNumberOfSolutions(p int) int {

	numOfSols := 0

	for a := 1; a<p; a++ {
		for b := a; a+b<p; b++ {
			c := p - a - b
			if (c*c == a*a + b*b) && c > b {
				numOfSols++
			}
		}
	}

	return numOfSols
}

func main() {
	

	maxNumOfSols := 0
	theP := 3

	for p := 3; p<= 1000; p++ {

		n := getNumberOfSolutions(p)
		if n > maxNumOfSols {
			maxNumOfSols = n
			theP = p
		}
	}

	fmt.Println(theP)
}