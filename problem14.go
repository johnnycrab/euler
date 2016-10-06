/*
	Problem:
	Which starting number, under one million, produces the longest Collatz sequence?
*/

package main

import (
	"fmt"
)

func collatz(n int) int {
	if n%2 == 0 {
		return n/2
	} else {
		return 3*n + 1
	}
}

func main() {

	longestStartingNumber := 0
	longestChain := 0

	for i := 1; i < 1000000; i++ {
		chainlen := 0
		j := i
		for j != 1 {
			chainlen++
			j = collatz(j)
		}

		if chainlen > longestChain {
			longestChain = chainlen
			longestStartingNumber = i
		}
	}

	fmt.Println(longestStartingNumber)

}