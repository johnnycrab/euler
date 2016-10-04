/*
	Problem:
	By considering the terms in the Fibonacci sequence whose values do not exceed four million,
	find the sum of the even-valued terms.
*/

package main

import "fmt"

func main() {

	// final sum
	sum := 0

	// fibonacci number 1 and 2
	var (
		fib1 = 1
		fib2 = 1
		temp = 0
	)

	for fib2 <= 4000000 {
		temp = fib2
		fib2 += fib1
		fib1 = temp

		if fib2 % 2 == 0 {
			sum += fib2
		}
	}

	fmt.Println(sum)
}