/*
	Problem 38
*/

package main

import (
	"fmt"
	"strconv"
)

// only call if you ar sure the number has nine digits
func getNineDigits(n int) [9]int {
	digits := [9]int{}

	i := 0
	for n > 0 {
		digits[i] = n%10
		n = (n - n%10)/10
		i++
	}

	return digits
}

// check if the given number is 1 to 9 pandigital
func isPandigital(num int) bool {
	digits := getNineDigits(num)

	count := 0
	for i := 1; i<=9; i++ {
		for _,d := range digits {
			if d == i {
				count++
				break
			}
		}
	}

	return count == 9
}

func main() {

	largestPandigital := 0
	
	// this is our integer we multiply everything with. as we have at least i o i*2, i will be smaller than 10000
	for i := 1; i < 10000; i++ {
		concat := i
		n := 1

		// multiply until we get an at least nine digit num
		for concat < 100000000 {
			n++
			concat, _ = strconv.Atoi(strconv.Itoa(concat) + strconv.Itoa(i*n))
		}

		// check if we have exactly nine digits and its pandigital
		if concat < 1000000000 && isPandigital(concat) && concat > largestPandigital {
			largestPandigital = concat
		}
	}

	fmt.Println(largestPandigital)
}