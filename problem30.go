package main

import (
	"fmt"
	"math"
)

func getDigitPowerSum(i int) int {
	digitPowerSum := 0
	temp := i
	
	for temp > 0 {
		digitPowerSum += int(math.Pow(float64(temp % 10), 5))

		temp = (temp - temp%10)/10
	}

	return digitPowerSum
}

func main() {
	// if a number fulfils the requirement, it must be leq 999999 (because for all subsequent numbers, the sum of power of digits will
	// always be smaller than the actual number)
	result := 0

	for i := 2; i<= 999999; i++ {

		if getDigitPowerSum(i) == i {
			result += i
		}
	}

	fmt.Println(result)
}