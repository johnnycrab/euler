/*
	Problem:
	https://projecteuler.net/problem=17
*/

package main

import (
	"fmt"
)

func main() {
	m := [100]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	m[30] = "thirty"
	m[40] = "forty"
	m[50] = "fifty"
	m[60] = "sixty"
	m[70] = "seventy"
	m[80] = "eighty"
	m[90] = "ninety"

	sum := 0

	for i := 1; i<=1000; i++ {
		num := ""

		if i == 1000 {
			num = "onethousand"
		} else {

			ones := i % 10
			tens := ((i%100) - (i%10))
			hundreds := (i-(i%100)) / 100

			if hundreds > 0 {
				num += m[hundreds] + "hundred"

				if i % 100 != 0 {
					num += "and"
				}
			}

			if tens >= 20 {
				num += m[tens]
				if ones > 0 {
					num += m[ones]
				}
			} else if tens == 10 {
				num += m[tens + ones]
			} else if tens == 0 && ones > 0 {
				num += m[ones]
			}
		}

		sum += len(num)
	}

	fmt.Println(sum)
}