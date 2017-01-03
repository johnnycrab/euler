/*
	Problem 61
*/

package main

import "fmt"

// this stores triangle, square, pentagonal, hexagonal, heptagonal, octagonal numbers in slices.
// each slice represents numbers from 0-9999, where 0 marks that a number is NOT polygonal, and 1 if it is
var polygonals [6][10000]int

type Cycle struct {
	numbers [6]int
	usedPolygonals [6]int // marks if a polygonal has been used
}

func makePolygonals() {
	for n := 1; (n*(n+1))/2 <= 9999; n++ {
		polygonals[0][(n*(n+1))/2] = 1

		square := n*n
		penta := (n*(3*n-1))/2
		hexa := n*(2*n-1)
		hepta := (n*(5*n-3))/2
		octa := n*(3*n-2)

		if square <= 9999 {
			polygonals[1][square] = 1
		}

		if penta <= 9999 {
			polygonals[2][penta] = 1
		}

		if hexa <= 9999 {
			polygonals[3][hexa] = 1
		}

		if hepta <= 9999 {
			polygonals[4][hepta] = 1
		}

		if octa <= 9999 {
			polygonals[5][octa] = 1
		}
	}
}

// given numbers a and b, checks if the last two digits of a are the first two digits of b
func bFollowsA(a,b int) bool {
	return a%100 == b/100
}

// given a number, returns a list of possible next numbers (where the last two digits of num are the same as the beginning two digits)
// together with an according list which contains the polygonal index of the number used
func returnNextPossible(num int, polygonalCheck [6]int) ([]int, []int) {
	possibleNums := []int{}
	accordingPolygonals := []int{}

	lastTwoDigs := num%100

	if lastTwoDigs >= 10 {
		for n := lastTwoDigs*100; n <= lastTwoDigs*100 + 99; n++ {
			for i := 0; i<5; i++ {
				if polygonalCheck[i] == 0 && polygonals[i][n] == 1 {
					possibleNums = append(possibleNums, n)
					accordingPolygonals = append(accordingPolygonals, i)
				}
			}
		}	
	}

	return possibleNums, accordingPolygonals
}

func makeNewCycle(oldCycle Cycle, num int, polygonal int) Cycle {
	newCycle := Cycle{}
	set := true
	for i := 0; i<6; i++ {
		if oldCycle.numbers[i] == 0 {
			if set {
				newCycle.numbers[i] = num
				set = false	
			}
		} else {
			newCycle.numbers[i] = oldCycle.numbers[i]
		}

		newCycle.usedPolygonals[i] = oldCycle.usedPolygonals[i]
	}

	newCycle.usedPolygonals[polygonal] = 1
	
	return newCycle
}

// given an octagonal number, it tries to build up a fitting cycle. If not possible, returns false.
func fulfillCycle(cycle Cycle) (bool, Cycle) {
	// check if we are done
	if cycle.numbers[5] != 0 {
		// we have five numbers. Now we need to check if they build up a cycle
		return !bFollowsA(cycle.numbers[5], cycle.numbers[0]), cycle
	}

	for i := 0; i < 6; i++ {
		if cycle.numbers[i] == 0 {
			// we need to get the next possible number
			possibleNums, accordingPolygonals := returnNextPossible(cycle.numbers[i-1], cycle.usedPolygonals)

			if len(possibleNums) == 0 {
				return true, Cycle{}
			}

			for j, possibleNum := range possibleNums {
				newCycle := makeNewCycle(cycle, possibleNum, accordingPolygonals[j])

				err, res := fulfillCycle(newCycle)

				if !err {
					return false, res
				}
			}

			break
		}
	}

	return true, Cycle{}
}

func main() {
	makePolygonals()
	
	for n := 1000; n<=9999;n++ {
		if polygonals[5][n] == 1 {
			
			err, cycle := fulfillCycle(Cycle{numbers: [6]int{n}, usedPolygonals: [6]int{0,0,0,0,0,1}})

			if !err {
				sum := 0
				for _, v := range cycle.numbers {
					sum += v
				}
				fmt.Println(sum)
				break
			}
		}		
	}

	
}