/*
	https://projecteuler.net/problem=18
*/

package main

import (
	"fmt"
)

// We store the triangle in an []int and think of it in the following way:
// 1
// | \
// 2  3
// | \| \
// 3  4  5

// gets an element in the tree. `i` specifies number of row, `j` specifies column

func get(i, j int, triangle []int) int {
	return triangle[((i-1)*i)/2 + j - 1]
}

func set(i, j int, triangle []int, val int) {
	triangle[((i-1)*i)/2 + j - 1] = val	
}

func main() {
	// we store the longest way from bottom to each number
	triangle := []int{75, 95, 64, 17, 47, 82, 18, 35, 87, 10, 20, 04, 82, 47, 65, 19, 1, 23, 75, 3, 34, 88, 2, 77, 73, 07, 63, 67, 99, 65, 4, 28, 06, 16, 70, 92, 41, 41, 26, 56, 83, 40, 80, 70, 33, 41, 48, 72, 33, 47, 32, 37, 16, 94, 29, 53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14, 70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57, 91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48, 63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31, 4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23}

	// now iterate over each row, starting from one before the last
	for i := 14; i>=1; i-- {
		// go through the columns
		for j := 1; j<=i; j++ {
			current := get(i, j, triangle)
			// we have two possibilites to come here. once from directly below, once from one below-one right. store the biggest one
			below := get(i+1, j, triangle)
			belowRight := get(i+1, j+1, triangle)

			if below > belowRight {
				current += below
			} else {
				current += belowRight
			}

			set(i, j, triangle, current)
		}
	}

	fmt.Println(triangle[0])
}