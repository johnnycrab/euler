/*
	Basic combinatoric functions
*/

package crabMath

func Factorial(n int) int {
	result := 1

	if n == 0 || n == 1 {
		return 1
	}

	for i := 2; i<=n; i++ {
		result *= i
	}

	return result
}

func NChooseK(n, k int) int {
	return Factorial(n)/(Factorial(k) * Factorial(n-k))
}