package main

import "fmt"

// returns how many numbers have been used when completing the nth spiral
func S_n(n int) int {
	if n == 1 {
		return 1
	} else {
		return 8*(n-1)
	}
}

// returns the number on the lower right corner on the nth spiral
/*func E_n(n int) int {
	if n == 1 {
		return 1
	} else {
		tempSum := 0
		for i := 1; i<= n-1; i++ {
			tempSum += S_n(i)
		}

		return tempSum + (n-1)*2
	}
}*/

func main() {
	// we have a 1001 & 1001 spiral, that is, we have 501 spirals
	sum := 0

	numbersUsed := 1

	for i:= 1; i<=501; i++ {
		lowerRight := 1

		if i > 1 {
			lowerRight = numbersUsed + (i-1)*2
		}

		sum += lowerRight

		if i > 1 {
			for j:=1; j<=3; j++ {
				sum += lowerRight + j*2*(i-1)
			}

			numbersUsed = lowerRight + 6*(i-1)
		}
	}

	fmt.Println(sum)
}