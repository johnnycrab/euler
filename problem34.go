/*
	Problem 34
*/

package main

import "fmt"

var factorialDigitMap [10]int

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	result := 2

	for i := 3; i<=n; i++ {
		result *= i
	} 

	return result
}

func factorialOfDigits(n int) int {
	sum := 0

	for n != 0 {
		sum += factorialDigitMap[n%10]
		n = (n - n%10)/10
	}

	return sum
}

func main() {
	for i:=0; i<10; i++ {
		factorialDigitMap[i] = factorial(i)
	}


	result := 0
	for i:=3; i<=1000000; i++ {
		if factorialOfDigits(i) == i {
			result += i
		}
	}

	fmt.Println(result)
}