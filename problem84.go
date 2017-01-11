/*
	Problem 84
	Monopoly odds

	We see the game as a Markov process.
	As there are 40 fields, we use a 40x40 probability matrix, where the i,j-th entry specifies the probability
	to land on the j-th field if you started on the i-th field and rolled the dices.

	Building up the matrix:
	Store all possible dice throws.

	Populate each row in the following way:
		Iterate over each dice roll. 
		Take the dice roll and its probability, and pass it to a "populate function", where the rules are carried out.
*/

package main

import (
	"fmt"
)

var diceThrows [][2]int // stores all possible throws of two dices
var oneRollProbability float64
var probabilityMatrix [40][40]float64


const jailField int = 10
const goField int = 0
const go2JailField int = 30
const sixteenth float64 = float64(1)/float64(16)

func main() {
	storeDiceThrows(4)
	populateProbabilityMatrix()

	vector := [40]float64{}

	// probability to be at GO in the beginning is 1
	vector[0] = 1

	for i := 1; i<= 100000; i++ {
		vector = multiplyVectorFromLeft(vector)	
	}
	
	// look for the greatest three
	first := 0
	second := 0
	third := 0

	for i, v := range vector {
		if v > vector[first] {
			first = i
		} else if v > vector[second] {
			second = i
		} else if v > vector[third] {
			third = i
		}
	}

	fmt.Println(first, second, third)
	fmt.Println(vector[first], vector[second], vector[third])
}

func multiplyVectorFromLeft(vector [40]float64) [40]float64 {
	result := [40]float64{}

	for j := 0; j<40; j++ {
		var sum float64 = 0
		for i := 0; i<40; i++ {
			sum += probabilityMatrix[i][j] * vector[i]
		}
		result[j] = sum
	}

	return result
}

func nextFieldOf(possibilities []int, field int) int {
	for i, v := range possibilities {
		next := possibilities[(i+1)%len(possibilities)]
		if (field >= v && i == len(possibilities) - 1) || (field < v && i == 0) {
			return possibilities[0]
		}

		if field >= v && field < next {
			return next
		}
	}

	return possibilities[0]
}

func nextRailway(field int) int {
	railways := []int{5, 15, 25, 35}
	return nextFieldOf(railways, field)
}

func nextUtility(field int) int {
	utilities := []int{12,28}
	return nextFieldOf(utilities, field)
}

func isCommunityChest(fieldNum int) bool {
	return fieldNum == 2 || fieldNum == 17 || fieldNum == 33
}

func isChance(fieldNum int) bool {
	return fieldNum == 7 || fieldNum == 22 || fieldNum == 36
}

func directlyMoveTo(landingField int, isPair bool, numOfPairs int, eventProbability float64, probabilityRow []float64) {
	if !isPair {
		probabilityRow[landingField] += eventProbability
	} else {
		// we have a pair, so we may roll again again
		for _, diceRoll := range diceThrows {
			carryOutRoll(diceRoll, landingField, numOfPairs, eventProbability * oneRollProbability, probabilityRow)
		}
	}
}

// `carries` out a roll, i.e. given a start field and a roll and a probability that this has happened (together the number of pairs that have been rolled so far),
// carries out the whole move until its end and populates the final probability in the given probabilityRow
func carryOutRoll(diceRoll [2]int, startField int, numOfPairs int, eventProbability float64, probabilityRow []float64) {
	isPair := diceRoll[0] == diceRoll[1] // if it is a pair, we do not directly add the probabilities
	if isPair {
		numOfPairs++

		// when you have three consecutive pairs, move directly to jail. Move is over
		if numOfPairs == 3 {
			probabilityRow[jailField] += eventProbability
			return
		}
	}

	landingField := (startField + diceRoll[0] + diceRoll[1])%40

	// Go2Jail Field
	if landingField == go2JailField {

		directlyMoveTo(jailField, isPair, numOfPairs, eventProbability, probabilityRow)
	} else if isCommunityChest(landingField) {

		newProb := eventProbability*sixteenth

		// in 14 cases, nothing happens
		for i := 1; i<=14; i++ {
			directlyMoveTo(landingField, isPair, numOfPairs, newProb, probabilityRow)
		}

		// in 1 case, we move to "Go" directly
		directlyMoveTo(goField, isPair, numOfPairs, newProb, probabilityRow)

		// in 1 case, we move to "Jail" directly
		directlyMoveTo(jailField, isPair, numOfPairs, newProb, probabilityRow)		

	} else if isChance(landingField) {
		newProb := eventProbability*sixteenth

		// in 6 cases, nothing happens
		for i := 1; i<=6; i++ {
			directlyMoveTo(landingField, isPair, numOfPairs, newProb, probabilityRow)
		}


		// Advance to go
		directlyMoveTo(goField, isPair, numOfPairs, newProb, probabilityRow)
		// go to jail
		directlyMoveTo(jailField, isPair, numOfPairs, newProb, probabilityRow)
		// go to c1
		directlyMoveTo(11, isPair, numOfPairs, newProb, probabilityRow)
		// go to e3
		directlyMoveTo(24, isPair, numOfPairs, newProb, probabilityRow)
		// go to h2
		directlyMoveTo(39, isPair, numOfPairs, newProb, probabilityRow)
		// go to r1
		directlyMoveTo(5, isPair, numOfPairs, newProb, probabilityRow)
		// go to next railway
		directlyMoveTo(nextRailway(landingField), isPair, numOfPairs, newProb, probabilityRow)
		// go to next railway
		directlyMoveTo(nextRailway(landingField), isPair, numOfPairs, newProb, probabilityRow)
		// go to next utility
		directlyMoveTo(nextUtility(landingField), isPair, numOfPairs, newProb, probabilityRow)
		// go back three squares
		if (landingField - 3) == go2JailField {
			directlyMoveTo(jailField, isPair, numOfPairs, newProb, probabilityRow)	
		} else {
			directlyMoveTo(landingField - 3, isPair, numOfPairs, newProb, probabilityRow)	
		}

	} else {

		directlyMoveTo(landingField, isPair, numOfPairs, eventProbability, probabilityRow)
	}

}

func populateProbabilityMatrix() {
	for i := 0; i<40; i++ {
		if i == go2JailField {
			continue
		}

		probabilityRow := make([]float64, 40)

		for _, diceRoll := range diceThrows {
			carryOutRoll(diceRoll, i, 0, oneRollProbability, probabilityRow)
		}

		for j, v := range probabilityRow {
			probabilityMatrix[i][j] = v
		}
	}
}


func storeDiceThrows(numOfSides int) {
	for i := 1; i <= numOfSides; i++ {
		for j := 1; j <= numOfSides; j++ {
			diceThrows = append(diceThrows, [2]int{i,j})
		}
	}

	oneRollProbability = float64(1)/float64(len(diceThrows))
}