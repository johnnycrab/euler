/*
	Problem:
	https://projecteuler.net/problem=21
*/

package main

import (
	"fmt"
)

func findSum(n int, channel chan [2]int) {

	limit := n
	i := 1
	sum := 0

	for i < limit {
		if n%i == 0 {
			sum += i

			if i > 1 && limit == n {
				// found a proper divisor smaller than n. we can set our limit
				limit = (n / i) + 1
			}
		}
		i++
	}

	channel <- [2]int{n,sum}
}

func main() {
	
	channel := make(chan [2]int)
	divisorSum := [9999]int{}

	solution := 0

	for i := 2; i <= 10000; i++ {
		go findSum(i, channel)
	}

	for i := 2; i <= 10000; i++ {
		res := <- channel

		divisorSum[res[0] - 2] = res[1]
	}

	close(channel)

	for i, sum := range divisorSum {
		a := i+2
		b := sum

		if a != b && b > 2 && b < 10001 && divisorSum[b-2] == a {
			solution += a
		}
	}

	fmt.Println(solution)
}