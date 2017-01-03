/*
	Problem 53
*/

package main

import (
	"fmt"
	"./crabMath"
)


func main() {
	count := 0
	var n, k int64
	for n = 1; n<= 100; n++ {
		for k = 1; k <= n; k++ {
			if len(crabMath.BigNChooseK(n, k).String()) >= 7 {
				count++
			}
		}
	}

	fmt.Println(count)
}