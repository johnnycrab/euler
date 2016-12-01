/*
	Problem 41
*/

package main

import "fmt"

const n int = 10000000

var sieve [n]int

func power(n, k int) int {
	if k == 0 {
		return 1
	}

	result := n
	for j := 2; j<=k; j++ {
		result *= n
	}

	return result
}

// gets primes up to n
func doSieve() {

	sieve[0] = 1

	p := 2

	for p*p <= n {
		for i := p; i*p <= n; i++ {
			sieve[i*p - 1] = 1
		}

		for i := 1; i < n - p; i++ {
			if sieve[p - 1 + i] == 0 {
				p = p + i
				
				break
			}
		}
	}
}

func digitsOfNumber(num int) []int {
	digits := []int{}

	for num > 0 {
		digits = append(digits, num%10)
		num = (num - num%10)/10
	}

	return digits
}

func isPandigital(k int) bool {
	digits := digitsOfNumber(k)

	digitsFound := 0

	for i := 1; i<=len(digits); i++ {
		for _, d := range digits {
			if d == i {
				digitsFound++
				break
			}
		}
	}

	return digitsFound == len(digits)
}

func main() {

	doSieve()

	k := n-1
	for {
		if sieve[k] == 0 && isPandigital(k+1) {
			fmt.Println(k+1)
			break
		}
		k--
	}
}