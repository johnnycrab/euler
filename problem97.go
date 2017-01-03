/*
	Problem 97
*/

package main

import (
	"fmt"
)

var tenDigits [10]int

func timesTwo() {

	remain := 0
	for i := 9; i>=0; i-- {
		digitTimesTwo := tenDigits[i]*2 + remain
		if digitTimesTwo >= 10 {
			remain = 1
		} else {
			remain = 0
		}

		tenDigits[i] = digitTimesTwo%10
	}
}

func main() {
	tenDigits = [10]int{0,0,0,0,0,5,6,8,6,6} // this is 28433 * 2

	for i := 1; i <= 7830456; i++ {
		timesTwo()
	}

	tenDigits[9] += 1
	for _, v := range tenDigits {
		fmt.Print(v)
	}

	fmt.Println("")
}