/*
	Problem 49
*/

package main

import (
	"fmt"
	"strconv"
)

const n int = 9999
var sieve [n]int

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

func digitArray(a int) [4]int {
	arr := [4]int{}

	i := 0
	for a != 0 {
		arr[i] = a%10
		a = (a - a%10)/10
		i++
	}

	return arr
}

func arePermutationsOfOneAnother(a int, b int) bool {
	digitsA := digitArray(a)
	digitsB := digitArray(b)

	count := 0
	for _, v := range digitsA {
		for j, w := range digitsB {
			if v == w {
				digitsB[j] = -1
				count++
				break
			}
		}
	}

	return count == 4
}

func main() {
	doSieve()

	for n := 1000; n<9998; n++ {
		if sieve[n] == 0 {
			prime1 := n+1

			for m := n+1; m<9999; m++ {
				if sieve[m] == 0 {
					prime2 := m+1

					k := 2*m - n
					if k < 9999 && arePermutationsOfOneAnother(prime1, prime2) && sieve[k] == 0 && arePermutationsOfOneAnother(k + 1, prime1) {
						prime3 := k + 1
						
						if prime1 != 1487 {
							fmt.Println(strconv.Itoa(prime1) + strconv.Itoa(prime2) + strconv.Itoa(prime3))
						}
					}
				}
			}
		}
	}
}