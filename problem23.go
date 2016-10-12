package main

import "fmt"

func isAbundant(n int) bool {
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

	return sum > n
}

func main() {
	// all integers greater than 28123 can be written as sum of two abundant numbers
	// thus we will just check all numbers up to 28123 if they are abundant or not
	isAbundantList := [28123]bool{}

	for i := 0; i<28123; i++ {
		if isAbundant(i+1) {
			isAbundantList[i] = true
		}
	}

	// iterate over all numbers and check if they can be written as sum of two abundant sums
	solution := 0
	for n := 1; n<=28123; n++ {
		canBeWrittenAsSum := false
		j := n/2

		for i := 1; i <= j; i++ {
			if isAbundantList[i-1] && isAbundantList[n-i-1] {
				canBeWrittenAsSum = true
				break
			}
		}

		if !canBeWrittenAsSum {
			solution += n
		}
	}

	fmt.Println(solution)
}