/*
	Problem 35
*/

package main

import (
	"fmt"
)

const n int = 1000000
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


func digitArray(a int) []int {
	arr := []int{}

	i := 0
	for a != 0 {
		arr = append(arr, a%10)
		a = (a - a%10)/10
		i++
	}
	return arr
}

func makeRotations(p int) []int {
	digits := digitArray(p)

	if len(digits) == 1 {
		return []int{}
	}

	rotations := make([]int, len(digits) - 1)

	for i := 1; i < len(digits); i++ {
		rotation := 0

		for j := 0; j < len(digits); j++ {
			rotation += power(10, j) * digits[(j + i)%len(digits)]
		}

		rotations[i - 1] = rotation
	}

	return rotations
}

func isRotational(p int) bool {
	if p < 10 {
		return true
	}

	// checks if prime p is rotational, but only returns true if all rotations are larger than p
	rotations := makeRotations(p)

	for _, rotation := range rotations {
		if sieve[rotation - 1] == 1 {
			return false
		}
	}

	return true
}

func main() {
	doSieve()

	count := 0

	for m := 0; m < n; m++ {
		if sieve[m] == 0 && isRotational(m+1) {
			count++
		}
	}

	fmt.Println(count)
}