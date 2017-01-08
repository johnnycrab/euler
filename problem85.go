/*
	Problem 85

	for formula given a rectangle of size n*x, check 
	http://www.wolframalpha.com/input/?i=sum+(sum+(n-a%2B1)*(m-b%2B1),+b+%3D+1+to+m),+a+%3D+1+to+n
*/

package main

import (
	"fmt"
	"./crabMath"
)

func numOfRects(n, m int) int {
	return (m*(m+1)*n*(n+1))/4
}

func main() {
	N := 2*1000*1000

	minDist := -1
	closest := 0

	for n := 1; n<=10000; n++ {
		for m := 1; m<=10000;m++ {
			num := numOfRects(n,m)
			abs := crabMath.Abs(N - num)
			if abs < minDist || minDist == -1 {
				minDist = abs
				closest = n*m
			}
		}
	}

	fmt.Println(closest)
}
