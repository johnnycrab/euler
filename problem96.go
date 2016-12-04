/*
	Problem 96

	Sudoku Solver

	Sodokus are always given as a 81 long integer array, where the numbers in the sudoku are read from left to right, top to bottom
	Missing numbers are represented by a zero
*/

package main

import "fmt"


type Sudoku struct {
	grid [81]int 						// stores the numbers in the grid
	possibilities [81][9]bool		 	// for each number in the grid, stores the possibilities for numbers (true for possible, false for not possible)
	solved bool
}

// transforms row-column notation to linear array index (using 81 as base)
func rowCol(i, j int) int {
	return (j-1)*9 + i - 1
}

// prints the sudoku
func (sudoku *Sudoku) print() {
	for i := 0; i<9; i++ {
		for j := 0; j<9; j++ {
			fmt.Print(sudoku.grid[i*9 + j], " ")

			if j == 8 {
				fmt.Println("")
			}
		}
		
	}
}

// we distribute our grid into nine different small grids labeled just as in a matrix (i,j), with 1<=i,j<= 3
// returns the indices of the small grid the given number is in
func smallGridIndices(i, j int) (int, int) {
	var smallGrid_i, smallGrid_j int
	if i <= 3 {
		smallGrid_i = 1
	} else if i <= 6 {
		smallGrid_i = 2
	} else {
		smallGrid_i = 3
	}

	if j <= 3 {
		smallGrid_j = 1
	} else if j <= 6 {
		smallGrid_j = 2
	} else {
		smallGrid_j = 3
	}

	return smallGrid_i, smallGrid_j
}

// sets a number within the sudoku, and already removes the possibility in the row and column, as well as within the same
// small grid
func (sudoku *Sudoku) setNumber(i, j, num int) {
	sudoku.grid[rowCol(i,j)] = num

	if num != 0 {
		// remove the number possibility within the same row and column
		for k := 1; k<=9; k++ {
			sudoku.removePossibility(i, k, num)
			sudoku.removePossibility(k, j, num)
		}

		// now remove the possibility in the same small grid
		smallGrid_i, smallGrid_j := smallGridIndices(i, j)
		for n := 1; n<=3; n++ {
			for m := 1; m<=3; m++ {
				sudoku.removePossibility(n + 3*(smallGrid_i - 1), m + 3*(smallGrid_j - 1), num)
			}
		}
	}
}

func (sudoku *Sudoku) removePossibility(i, j, num int) {
	if num > 0 {
		sudoku.possibilities[rowCol(i,j)][num - 1] = false
	}
}

// checks within the array of possibilities, if there is an only possible number left
// if still ambiguous, returns 0
func (sudoku *Sudoku) returnOnlyPossibility(i, j int) int {
	possibleArray := sudoku.possibilities[rowCol(i,j)]

	possibleNum := 0
	for i, v := range possibleArray {
		if v == true {
			if possibleNum == 0 {
				possibleNum = i + 1
			} else {
				// ambiguours
				possibleNum = 0
				break
			}
		}
	}

	return possibleNum
}

func (sudoku *Sudoku) solve(initial [81]int) {
	// Step 1. Setup all possibilities
	for i := 1; i<=9; i++ {
		for j := 1; j<=9; j++ {
			for k := 0; k<9; k++ {
				sudoku.possibilities[rowCol(i,j)][k] = true
			}
		}
	}

	// Step 2. Add all initial numbers.
	for i := 1; i<=9; i++ {
		for j := 1; j<=9; j++ {
			sudoku.setNumber(i,j, initial[rowCol(i,j)])
		}
	}

	// Step 3. Iterate over all numbers, check if they are zero, and if yes,
	// see if we can already deduce the number from the possibilities. If yes, set it.
	// Repeat until no longer possible
	for {
		foundAZero := false // keeps track if we need to solve anything anyway
		couldProgress := false // keeps track if we could do anything new 

		for i := 1; i<=9; i++ {
			for j := 1; j<=9; j++ {
				for k := 0; k<9; k++ {
					if sudoku.grid[rowCol(i,j)] == 0 {
						foundAZero = true
						p := sudoku.returnOnlyPossibility(i,j)
						if p != 0 {
							couldProgress = true
							sudoku.setNumber(i, j, p)
						}
					}
				}
			}
		}

		if !foundAZero {
			// we are done!
			fmt.Println("solved!")
			sudoku.solved = true
		}

		if !foundAZero || !couldProgress {
			// nothing we could do
			break
		}
	}
}

func main() {
	testSudoku := [81]int{0,0,3,0,2,0,6,0,0,9,0,0,3,0,5,0,0,1,0,0,1,8,0,6,4,0,0,0,0,8,1,0,2,9,0,0,7,0,0,0,0,0,0,0,8,0,0,6,7,0,8,2,0,0,0,0,2,6,0,9,5,0,0,8,0,0,2,0,3,0,0,9,0,0,5,0,1,0,3,0,0,}

	var sudoku Sudoku
	sudoku.solve(testSudoku)
	sudoku.print()

}