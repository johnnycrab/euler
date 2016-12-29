/*
	Problem 79
*/

package main

import "fmt"

// we have at least 8 digits

var codes [50][3]int

func increase(code []int) []int {
	index := len(code)
	doStop := false

	for !doStop {
		index -= 1
		code[index] = (code[index] + 1)%10

		if code[index] == 0 {
			if index == 0 {
				code = append([]int{0}, code...)
			}
		} else {
			doStop = true
		}
	}

	return code
}

func working(code []int) bool {
	for _, try := range codes {
		index := 0
		for _, digit := range code {
			if digit == try[index] {
				index++
				if index == 3 {
					break
				}
			}
		}

		if index != 3 {
			return false
		}	
	}

	return true
}

func main() {
	codes = [50][3]int{{3,1,9},{6,8,0},{1,8,0},{6,9,0},{1,2,9},{6,2,0},{7,6,2},{6,8,9},{7,6,2},{3,1,8},{3,6,8},{7,1,0},{7,2,0},{7,1,0},{6,2,9},{1,6,8},{1,6,0},{6,8,9},{7,1,6},{7,3,1},{7,3,6},{7,2,9},{3,1,6},{7,2,9},{7,2,9},{7,1,0},{7,6,9},{2,9,0},{7,1,9},{6,8,0},{3,1,8},{3,8,9},{1,6,2},{2,8,9},{1,6,2},{7,1,8},{7,2,9},{3,1,9},{7,9,0},{6,8,0},{8,9,0},{3,6,2},{3,1,9},{7,6,0},{3,1,6},{7,2,9},{3,8,0},{3,1,9},{7,2,8},{7,1,6}}

	code := []int{1,0,0,0,0,0,0,0}

	for !working(code) {
		code = increase(code)
	}

	for _, d := range code {
		fmt.Print(d)
	}
	fmt.Println("")
}