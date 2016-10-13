package main

import (
	"fmt"
	"math"
)

func checkSolutions(M int, c chan [2]int) {

	numOfSols := 0

	for a := 1; a <= M; a++ {
		for b := a; b <= M; b++ {
			for c := b; c <= M; c++ {

				shortestSquare := a*a + b*b + c*c + 2*a*b // this is the shortest, as a < b < c

				p := int(math.Sqrt(float64(shortestSquare)))

				if p*p == shortestSquare {
					numOfSols++
				}
			}
		}
	}

	c <- [2]int{M, numOfSols}
} 

func main() {
	channel := make(chan [2]int)
	toCheck := 0
	// through trial and error we have found that it's something between 1810 and 1820

	M := 1820

	for i := 1810; i <= M; i++ {
		toCheck++
		go checkSolutions(i, channel)		
	}

	for toCheck > 0 {
		sols := <- channel
		toCheck--

		if sols[1] > 1000000 && sols[0] < M {
			M = sols[0]
		}
	}

	close(channel)

	fmt.Println(M)
}