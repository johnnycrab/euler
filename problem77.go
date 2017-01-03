/*
	Problem 77

	See: http://math.stackexchange.com/questions/89240/prime-partition
*/

package main

import (
	"fmt"
	"./crabMath"
)

const N int = 1000 * 1000
var primeSieve []int
var sopfs [N]int
var primePartitions [N]int

func makeSopfs() {
	for i := 2; i < N; i++ {
		sum := 0
		_, factors := crabMath.SimplePrimeFactorizationWithSieve(i, primeSieve)
		
		lastFactor := 0
		for _, factor := range factors {
			if factor == lastFactor {
				continue
			} else {
				sum += factor
				lastFactor = factor
			}
		}

		sopfs[i] = sum
	}
}

func getPrimePartitionAmount(n int) int {
	sum := sopfs[n]
	for j := 1; j<= n-1; j++ {
		sum += sopfs[j]*primePartitions[n-j]
	}

	return sum/n
}

func main() {
	primeSieve = crabMath.PrimeSieve(N)
	makeSopfs()

	for n := 2; n < N; n++ {
		primePartition := getPrimePartitionAmount(n)
		if primePartition > 5000 {
			fmt.Println(n)
			break
		}
		primePartitions[n] = primePartition
	}
}