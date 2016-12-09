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

func (sudoku *Sudoku) getNumberPossibilities(i, j int) []int {
	poss := sudoku.possibilities[rowCol(i,j)]
	possibleNums := []int{}

	for i, v := range poss {
		if v == true {
			possibleNums = append(possibleNums, i+1)
		}
	}

	return possibleNums
}

func (sudoku *Sudoku) removePossibility(i, j, num int) {
	if num > 0 {
		sudoku.possibilities[rowCol(i,j)][num - 1] = false
	}
}

// returns whether the i,j-field has num as possibility
func (sudoku *Sudoku) fieldHasPossibility(i, j, num int) bool {
	return sudoku.possibilities[rowCol(i,j)][num - 1]
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

func (sudoku *Sudoku) returnHiddenSingle(i, j int) int {
	hiddenSingle := 0

	for p := 1; p<=9; p++ {
		isHiddenSingle := true

		// Row: Check that no field on the same row has this value and no field with a zero can take this value
		for k := 1; k<=9; k++ {
			if k != j {
				if sudoku.grid[rowCol(i,k)] == p || sudoku.fieldHasPossibility(i, k, p) {
					isHiddenSingle = false
				}
			}
		}

		if isHiddenSingle {
			hiddenSingle = p
			break
		}

		// Column: Check that no field on the same column has this value and no field with a zero can take this value
		for k := 1; k<=9; k++ {
			if k != i {
				if sudoku.grid[rowCol(k,j)] == p || sudoku.fieldHasPossibility(k, j, p) {
					isHiddenSingle = false
				}
			}
		}

		if isHiddenSingle {
			hiddenSingle = p
			break
		}

		// Sub-square
		smallGrid_i, smallGrid_j := smallGridIndices(i, j)
		for n := 1; n<=3; n++ {
			for m := 1; m<=3; m++ {
				sg_i := n + 3*(smallGrid_i - 1)
				sg_j := m + 3*(smallGrid_j - 1)
				
				if (sg_i == i && sg_j == j) || sudoku.fieldHasPossibility(sg_i, sg_j, p) {
					isHiddenSingle = false
				}
			}
		}

		if isHiddenSingle {
			hiddenSingle = p

			break
		}
	}

	return hiddenSingle
}

// Naked single: Look through possibilites of zeros and see if there is only one left. This is basic.
func (sudoku *Sudoku) deduce_nakedSingle() bool {
	nakedSingleCouldProgress := false

	for i := 1; i<=9; i++ {
		for j := 1; j<=9; j++ {
			if sudoku.grid[rowCol(i,j)] == 0 {
				p := sudoku.returnOnlyPossibility(i,j)
				if p != 0 {
					nakedSingleCouldProgress = true
					sudoku.setNumber(i, j, p)
				}
			}
		}
	}

	return nakedSingleCouldProgress
}

// Hidden single: A field must take a value because all other fields may not take the value
func (sudoku *Sudoku) deduce_hiddenSingle() bool {
	hiddenSingleCouldProgress := false

	for i := 1; i<=9; i++ {
		for j := 1; j<=9; j++ {
			for k := 0; k<9; k++ {
				if sudoku.grid[rowCol(i,j)] == 0 {
					p := sudoku.returnHiddenSingle(i,j)
					if p != 0 {
						hiddenSingleCouldProgress = true
						sudoku.setNumber(i, j, p)
					}
				}
			}
		}
	}

	return hiddenSingleCouldProgress
}

func (sudoku *Sudoku) hasZeros() bool {
	for i := 1; i<=9; i++ {
		for j := 1; j<=9; j++ {
			if sudoku.grid[rowCol(i,j)] == 0 {
				return true	
			}
		}
	}

	return false
}

// given a set of numbers, checks if all numbers from 1 to 9 present
// exactly once. Otherwise it is considered conflicted
func subsetIsConflicted(subset [9]int) bool {
	count := 0
	for i := 1; i<=9; i++ {
		for _, v := range subset {
			if v == 1 {
				count++
				break
			}
		}
	}

	return count != 9
}

func (sudoku *Sudoku) hasConflict() bool {
	subset := [9]int{}

	// check rows
	for i := 1; i<=9; i++ {
		for j:=1; j<=9; j++ {
			subset[j-1] = sudoku.grid[rowCol(i,j)]
		}
		if subsetIsConflicted(subset) {
			return true
		}
	}

	// check cols
	for j := 1; j<=9; j++ {
		for i:=1; i<=9; i++ {
			subset[i-1] = sudoku.grid[rowCol(i,j)]
		}
		if subsetIsConflicted(subset) {
			return true
		}
	}

	// check subsquares
	// Sub-square
	for smallGrid_i := 1; smallGrid_i<= 3; smallGrid_i++ {
		for smallGrid_j := 1; smallGrid_j<= 3; smallGrid_j++ {
			index := 0

			for n := 1; n<=3; n++ {
				for m := 1; m<=3; m++ {
					sg_i := n + 3*(smallGrid_i - 1)
					sg_j := m + 3*(smallGrid_j - 1)
					index++
					
					subset[index - 1] = sudoku.grid[rowCol(sg_i,sg_j)]
				}
			}

			if subsetIsConflicted(subset) {
				return true
			}			
		}		
	}

	return false
}

func (sudoku *Sudoku) setup(initial [81]int) {
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
}

// get the first number with the least possibilities, set a number and then recursively
// solve the sudokus
func (sudoku *Sudoku) tryAndError() {
	if sudoku.hasZeros() == false {
		return
	}

	maxNumOfPossibilities := 2

	// this is 
	var iTry, jTry int
	var possibleNumsTry []int

	doBreak := false

	for {

		for i := 1; i<=9; i++ {
			for j := 1; j<=9; j++ {
				if sudoku.grid[rowCol(i,j)] == 0 {

					possibleNums := sudoku.getNumberPossibilities(i,j)

					if len(possibleNums) == maxNumOfPossibilities {
						possibleNumsTry = possibleNums
						iTry = i
						jTry = j
						doBreak = true
						break
					} else if len(possibleNums) == 0 {
						// this is conflicted then
						return
					}
				}
			}
			if doBreak {
				break
			}
		}

		if doBreak {
			break
		} else {
			maxNumOfPossibilities++
		}
	}
	
	// we have a field we will bruteforce
	for _, numToTry := range possibleNumsTry {
		// Step 1. Setup a test sudoku
		var testSudoku Sudoku
		testSudoku.grid = sudoku.grid
		testSudoku.possibilities = sudoku.possibilities

		// Step 2. Set the number to try and let it solve
		testSudoku.setNumber(iTry, jTry, numToTry)

		// Step 3. Let it solve itself
		conflicted := testSudoku.solve()

		if conflicted == false && testSudoku.solved {
			sudoku.grid = testSudoku.grid
			break
		}
	}
}

// solves a sudoku and returns if it is in conflict
func (sudoku *Sudoku) solve() bool {

	// Step 1: Try deduction. Use mulitple mechanisms and repeat them until none of them
	// yield any result.
	for {
		nakedSingleCouldProgress := sudoku.deduce_nakedSingle()
		hiddenSingleCouldProgress := sudoku.deduce_hiddenSingle()

		if !nakedSingleCouldProgress && !hiddenSingleCouldProgress {
			break
		}
	}

	// otherwise we tryAndError
	sudoku.tryAndError()

	if sudoku.hasZeros() == false {

		if sudoku.hasConflict() == false {
			sudoku.solved = true
			return false
		}

		return true
	}


	return true
}

func main() {
	sudokus := [50][81]int{[81]int{0,0,3,0,2,0,6,0,0,9,0,0,3,0,5,0,0,1,0,0,1,8,0,6,4,0,0,0,0,8,1,0,2,9,0,0,7,0,0,0,0,0,0,0,8,0,0,6,7,0,8,2,0,0,0,0,2,6,0,9,5,0,0,8,0,0,2,0,3,0,0,9,0,0,5,0,1,0,3,0,0,},[81]int{2,0,0,0,8,0,3,0,0,0,6,0,0,7,0,0,8,4,0,3,0,5,0,0,2,0,9,0,0,0,1,0,5,4,0,8,0,0,0,0,0,0,0,0,0,4,0,2,7,0,6,0,0,0,3,0,1,0,0,7,0,4,0,7,2,0,0,4,0,0,6,0,0,0,4,0,1,0,0,0,3,},[81]int{0,0,0,0,0,0,9,0,7,0,0,0,4,2,0,1,8,0,0,0,0,7,0,5,0,2,6,1,0,0,9,0,4,0,0,0,0,5,0,0,0,0,0,4,0,0,0,0,5,0,7,0,0,9,9,2,0,1,0,8,0,0,0,0,3,4,0,5,9,0,0,0,5,0,7,0,0,0,0,0,0,},[81]int{0,3,0,0,5,0,0,4,0,0,0,8,0,1,0,5,0,0,4,6,0,0,0,0,0,1,2,0,7,0,5,0,2,0,8,0,0,0,0,6,0,3,0,0,0,0,4,0,1,0,9,0,3,0,2,5,0,0,0,0,0,9,8,0,0,1,0,2,0,6,0,0,0,8,0,0,6,0,0,2,0,},[81]int{0,2,0,8,1,0,7,4,0,7,0,0,0,0,3,1,0,0,0,9,0,0,0,2,8,0,5,0,0,9,0,4,0,0,8,7,4,0,0,2,0,8,0,0,3,1,6,0,0,3,0,2,0,0,3,0,2,7,0,0,0,6,0,0,0,5,6,0,0,0,0,8,0,7,6,0,5,1,0,9,0,},[81]int{1,0,0,9,2,0,0,0,0,5,2,4,0,1,0,0,0,0,0,0,0,0,0,0,0,7,0,0,5,0,0,0,8,1,0,2,0,0,0,0,0,0,0,0,0,4,0,2,7,0,0,0,9,0,0,6,0,0,0,0,0,0,0,0,0,0,0,3,0,9,4,5,0,0,0,0,7,1,0,0,6,},[81]int{0,4,3,0,8,0,2,5,0,6,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,9,4,9,0,0,0,0,4,0,7,0,0,0,0,6,0,8,0,0,0,0,1,0,2,0,0,0,0,3,8,2,0,5,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,3,4,0,9,0,7,1,0,},[81]int{4,8,0,0,0,6,9,0,2,0,0,2,0,0,8,0,0,1,9,0,0,3,7,0,0,6,0,8,4,0,0,1,0,2,0,0,0,0,3,7,0,4,1,0,0,0,0,1,0,6,0,0,4,9,0,2,0,0,8,5,0,0,7,7,0,0,9,0,0,6,0,0,6,0,9,2,0,0,0,1,8,},[81]int{0,0,0,9,0,0,0,0,2,0,5,0,1,2,3,4,0,0,0,3,0,0,0,0,1,6,0,9,0,8,0,0,0,0,0,0,0,7,0,0,0,0,0,9,0,0,0,0,0,0,0,2,0,5,0,9,1,0,0,0,0,5,0,0,0,7,4,3,9,0,2,0,4,0,0,0,0,7,0,0,0,},[81]int{0,0,1,9,0,0,0,0,3,9,0,0,7,0,0,1,6,0,0,3,0,0,0,5,0,0,7,0,5,0,0,0,0,0,0,9,0,0,4,3,0,2,6,0,0,2,0,0,0,0,0,0,7,0,6,0,0,1,0,0,0,3,0,0,4,2,0,0,7,0,0,6,5,0,0,0,0,6,8,0,0,},[81]int{0,0,0,1,2,5,4,0,0,0,0,8,4,0,0,0,0,0,4,2,0,8,0,0,0,0,0,0,3,0,0,0,0,0,9,5,0,6,0,9,0,2,0,1,0,5,1,0,0,0,0,0,6,0,0,0,0,0,0,3,0,4,9,0,0,0,0,0,7,2,0,0,0,0,1,2,9,8,0,0,0,},[81]int{0,6,2,3,4,0,7,5,0,1,0,0,0,0,5,6,0,0,5,7,0,0,0,0,0,4,0,0,0,0,0,9,4,8,0,0,4,0,0,0,0,0,0,0,6,0,0,5,8,3,0,0,0,0,0,3,0,0,0,0,0,9,1,0,0,6,4,0,0,0,0,7,0,5,9,0,8,3,2,6,0,},[81]int{3,0,0,0,0,0,0,0,0,0,0,5,0,0,9,0,0,0,2,0,0,5,0,4,0,0,0,0,2,0,0,0,0,7,0,0,1,6,0,0,0,0,0,5,8,7,0,4,3,1,0,6,0,0,0,0,0,8,9,0,1,0,0,0,0,0,0,6,7,0,8,0,0,0,0,0,0,5,4,3,7,},[81]int{6,3,0,0,0,0,0,0,0,0,0,0,5,0,0,0,0,8,0,0,5,6,7,4,0,0,0,0,0,0,0,2,0,0,0,0,0,0,3,4,0,1,0,2,0,0,0,0,0,0,0,3,4,5,0,0,0,0,0,7,0,0,4,0,8,0,3,0,0,9,0,2,9,4,7,1,0,0,0,8,0,},[81]int{0,0,0,0,2,0,0,4,0,0,0,8,0,3,5,0,0,0,0,0,0,0,7,0,6,0,2,0,3,1,0,4,6,9,7,0,2,0,0,0,0,0,0,0,0,0,0,0,5,0,1,2,0,3,0,4,9,0,0,0,7,3,0,0,0,0,0,0,0,0,1,0,8,0,0,0,0,4,0,0,0,},[81]int{3,6,1,0,2,5,9,0,0,0,8,0,9,6,0,0,1,0,4,0,0,0,0,0,0,5,7,0,0,8,0,0,0,4,7,1,0,0,0,6,0,3,0,0,0,2,5,9,0,0,0,8,0,0,7,4,0,0,0,0,0,0,5,0,2,0,0,1,8,0,6,0,0,0,5,4,7,0,3,2,9,},[81]int{0,5,0,8,0,7,0,2,0,6,0,0,0,1,0,0,9,0,7,0,2,5,4,0,0,0,6,0,7,0,0,2,0,3,0,1,5,0,4,0,0,0,9,0,8,1,0,3,0,8,0,0,7,0,9,0,0,0,7,6,2,0,5,0,6,0,0,9,0,0,0,3,0,8,0,1,0,3,0,4,0,},[81]int{0,8,0,0,0,5,0,0,0,0,0,0,0,0,3,4,5,7,0,0,0,0,7,0,8,0,9,0,6,0,4,0,0,9,0,3,0,0,7,0,1,0,5,0,0,4,0,8,0,0,7,0,2,0,9,0,1,0,2,0,0,0,0,8,4,2,3,0,0,0,0,0,0,0,0,1,0,0,0,8,0,},[81]int{0,0,3,5,0,2,9,0,0,0,0,0,0,4,0,0,0,0,1,0,6,0,0,0,3,0,5,9,0,0,2,5,1,0,0,8,0,7,0,4,0,8,0,3,0,8,0,0,7,6,3,0,0,1,3,0,8,0,0,0,1,0,4,0,0,0,0,2,0,0,0,0,0,0,5,1,0,4,8,0,0,},[81]int{0,0,0,0,0,0,0,0,0,0,0,9,8,0,5,1,0,0,0,5,1,9,0,7,4,2,0,2,9,0,4,0,1,0,6,5,0,0,0,0,0,0,0,0,0,1,4,0,5,0,8,0,9,3,0,2,6,7,0,9,5,8,0,0,0,5,1,0,3,6,0,0,0,0,0,0,0,0,0,0,0,},[81]int{0,2,0,0,3,0,0,9,0,0,0,0,9,0,7,0,0,0,9,0,0,2,0,8,0,0,5,0,0,4,8,0,6,5,0,0,6,0,7,0,0,0,2,0,8,0,0,3,1,0,2,9,0,0,8,0,0,6,0,5,0,0,7,0,0,0,3,0,9,0,0,0,0,3,0,0,2,0,0,5,0,},[81]int{0,0,5,0,0,0,0,0,6,0,7,0,0,0,9,0,2,0,0,0,0,5,0,0,1,0,7,8,0,4,1,5,0,0,0,0,0,0,0,8,0,3,0,0,0,0,0,0,0,9,2,8,0,5,9,0,7,0,0,6,0,0,0,0,3,0,4,0,0,0,1,0,2,0,0,0,0,0,6,0,0,},[81]int{0,4,0,0,0,0,0,5,0,0,0,1,9,4,3,6,0,0,0,0,9,0,0,0,3,0,0,6,0,0,0,5,0,0,0,2,1,0,3,0,0,0,5,0,6,8,0,0,0,2,0,0,0,7,0,0,5,0,0,0,2,0,0,0,0,2,4,3,6,7,0,0,0,3,0,0,0,0,0,4,0,},[81]int{0,0,4,0,0,0,0,0,0,0,0,0,0,3,0,0,0,2,3,9,0,7,0,0,0,8,0,4,0,0,0,0,9,0,0,1,2,0,9,8,0,1,3,0,7,6,0,0,2,0,0,0,0,8,0,1,0,0,0,8,0,5,3,9,0,0,0,4,0,0,0,0,0,0,0,0,0,0,8,0,0,},[81]int{3,6,0,0,2,0,0,8,9,0,0,0,3,6,1,0,0,0,0,0,0,0,0,0,0,0,0,8,0,3,0,0,0,6,0,2,4,0,0,6,0,3,0,0,7,6,0,7,0,0,0,1,0,8,0,0,0,0,0,0,0,0,0,0,0,0,4,1,8,0,0,0,9,7,0,0,3,0,0,1,4,},[81]int{5,0,0,4,0,0,0,6,0,0,0,9,0,0,0,8,0,0,6,4,0,0,2,0,0,0,0,0,0,0,0,0,1,0,0,8,2,0,8,0,0,0,5,0,1,7,0,0,5,0,0,0,0,0,0,0,0,0,9,0,0,8,4,0,0,3,0,0,0,6,0,0,0,6,0,0,0,3,0,0,2,},[81]int{0,0,7,2,5,6,4,0,0,4,0,0,0,0,0,0,0,5,0,1,0,0,3,0,0,6,0,0,0,0,5,0,8,0,0,0,0,0,8,0,6,0,2,0,0,0,0,0,1,0,7,0,0,0,0,3,0,0,7,0,0,9,0,2,0,0,0,0,0,0,0,4,0,0,6,3,1,2,7,0,0,},[81]int{0,0,0,0,0,0,0,0,0,0,7,9,0,5,0,1,8,0,8,0,0,0,0,0,0,0,7,0,0,7,3,0,6,8,0,0,4,5,0,7,0,8,0,9,6,0,0,3,5,0,2,7,0,0,7,0,0,0,0,0,0,0,5,0,1,6,0,3,0,4,2,0,0,0,0,0,0,0,0,0,0,},[81]int{0,3,0,0,0,0,0,8,0,0,0,9,0,0,0,5,0,0,0,0,7,5,0,9,2,0,0,7,0,0,1,0,5,0,0,8,0,2,0,0,9,0,0,3,0,9,0,0,4,0,2,0,0,1,0,0,4,2,0,7,1,0,0,0,0,2,0,0,0,8,0,0,0,7,0,0,0,0,0,9,0,},[81]int{2,0,0,1,7,0,6,0,3,0,5,0,0,0,0,1,0,0,0,0,0,0,0,6,0,7,9,0,0,0,0,4,0,7,0,0,0,0,0,8,0,1,0,0,0,0,0,9,0,5,0,0,0,0,3,1,0,4,0,0,0,0,0,0,0,5,0,0,0,0,6,0,9,0,6,0,3,7,0,0,2,},[81]int{0,0,0,0,0,0,0,8,0,8,0,0,7,0,1,0,4,0,0,4,0,0,2,0,0,3,0,3,7,4,0,0,0,9,0,0,0,0,0,0,3,0,0,0,0,0,0,5,0,0,0,3,2,1,0,1,0,0,6,0,0,5,0,0,5,0,8,0,2,0,0,6,0,8,0,0,0,0,0,0,0,},[81]int{0,0,0,0,0,0,0,8,5,0,0,0,2,1,0,0,0,9,9,6,0,0,8,0,1,0,0,5,0,0,8,0,0,0,1,6,0,0,0,0,0,0,0,0,0,8,9,0,0,0,6,0,0,7,0,0,9,0,7,0,0,5,2,3,0,0,0,5,4,0,0,0,4,8,0,0,0,0,0,0,0,},[81]int{6,0,8,0,7,0,5,0,2,0,5,0,6,0,8,0,7,0,0,0,2,0,0,0,3,0,0,5,0,0,0,9,0,0,0,6,0,4,0,3,0,2,0,5,0,8,0,0,0,5,0,0,0,3,0,0,5,0,0,0,2,0,0,0,1,0,7,0,4,0,9,0,4,0,9,0,6,0,7,0,1,},[81]int{0,5,0,0,1,0,0,4,0,1,0,7,0,0,0,6,0,2,0,0,0,9,0,5,0,0,0,2,0,8,0,3,0,5,0,1,0,4,0,0,7,0,0,2,0,9,0,1,0,8,0,4,0,6,0,0,0,4,0,1,0,0,0,3,0,4,0,0,0,7,0,9,0,2,0,0,6,0,0,1,0,},[81]int{0,5,3,0,0,0,7,9,0,0,0,9,7,5,3,4,0,0,1,0,0,0,0,0,0,0,2,0,9,0,0,8,0,0,1,0,0,0,0,9,0,7,0,0,0,0,8,0,0,3,0,0,7,0,5,0,0,0,0,0,0,0,3,0,0,7,6,4,1,2,0,0,0,6,1,0,0,0,9,4,0,},[81]int{0,0,6,0,8,0,3,0,0,0,4,9,0,7,0,2,5,0,0,0,0,4,0,5,0,0,0,6,0,0,3,1,7,0,0,4,0,0,7,0,0,0,8,0,0,1,0,0,8,2,6,0,0,9,0,0,0,7,0,2,0,0,0,0,7,5,0,4,0,1,9,0,0,0,3,0,9,0,6,0,0,},[81]int{0,0,5,0,8,0,7,0,0,7,0,0,2,0,4,0,0,5,3,2,0,0,0,0,0,8,4,0,6,0,1,0,5,0,4,0,0,0,8,0,0,0,5,0,0,0,7,0,8,0,3,0,1,0,4,5,0,0,0,0,0,9,1,6,0,0,5,0,8,0,0,7,0,0,3,0,1,0,6,0,0,},[81]int{0,0,0,9,0,0,8,0,0,1,2,8,0,0,6,4,0,0,0,7,0,8,0,0,0,6,0,8,0,0,4,3,0,0,0,7,5,0,0,0,0,0,0,0,9,6,0,0,0,7,9,0,0,8,0,9,0,0,0,4,0,1,0,0,0,3,6,0,0,2,8,4,0,0,1,0,0,7,0,0,0,},[81]int{0,0,0,0,8,0,0,0,0,2,7,0,0,0,0,0,5,4,0,9,5,0,0,0,8,1,0,0,0,9,8,0,6,4,0,0,0,2,0,4,0,3,0,6,0,0,0,6,9,0,5,1,0,0,0,1,7,0,0,0,6,2,0,4,6,0,0,0,0,0,3,8,0,0,0,0,9,0,0,0,0,},[81]int{0,0,0,6,0,2,0,0,0,4,0,0,0,5,0,0,0,1,0,8,5,0,1,0,6,2,0,0,3,8,2,0,6,7,1,0,0,0,0,0,0,0,0,0,0,0,1,9,4,0,7,3,5,0,0,2,6,0,4,0,5,3,0,9,0,0,0,2,0,0,0,7,0,0,0,8,0,9,0,0,0,},[81]int{0,0,0,9,0,0,0,0,2,0,5,0,1,2,3,4,0,0,0,3,0,0,0,0,1,6,0,9,0,8,0,0,0,0,0,0,0,7,0,0,0,0,0,9,0,0,0,0,0,0,0,2,0,5,0,9,1,0,0,0,0,5,0,0,0,7,4,3,9,0,2,0,4,0,0,0,0,7,0,0,0,},[81]int{3,8,0,0,0,0,0,0,0,0,0,0,4,0,0,7,8,5,0,0,9,0,2,0,3,0,0,0,6,0,0,9,0,0,0,0,8,0,0,3,0,2,0,0,9,0,0,0,0,4,0,0,7,0,0,0,1,0,7,0,5,0,0,4,9,5,0,0,6,0,0,0,0,0,0,0,0,0,0,9,2,},[81]int{0,0,0,1,5,8,0,0,0,0,0,2,0,6,0,8,0,0,0,3,0,0,0,0,0,4,0,0,2,7,0,3,0,5,1,0,0,0,0,0,0,0,0,0,0,0,4,6,0,8,0,7,9,0,0,5,0,0,0,0,0,8,0,0,0,4,0,7,0,1,0,0,0,0,0,3,2,5,0,0,0,},[81]int{0,1,0,5,0,0,2,0,0,9,0,0,0,0,1,0,0,0,0,0,2,0,0,8,0,3,0,5,0,0,0,3,0,0,0,7,0,0,8,0,0,0,5,0,0,6,0,0,0,8,0,0,0,4,0,4,0,1,0,0,7,0,0,0,0,0,7,0,0,0,0,6,0,0,3,0,0,4,0,5,0,},[81]int{0,8,0,0,0,0,0,4,0,0,0,0,4,6,9,0,0,0,4,0,0,0,0,0,0,0,7,0,0,5,9,0,4,6,0,0,0,7,0,6,0,8,0,3,0,0,0,8,5,0,2,1,0,0,9,0,0,0,0,0,0,0,5,0,0,0,7,8,1,0,0,0,0,6,0,0,0,0,0,1,0,},[81]int{9,0,4,2,0,0,0,0,7,0,1,0,0,0,0,0,0,0,0,0,0,7,0,6,5,0,0,0,0,0,8,0,0,0,9,0,0,2,0,9,0,4,0,6,0,0,4,0,0,0,2,0,0,0,0,0,1,6,0,7,0,0,0,0,0,0,0,0,0,0,3,0,3,0,0,0,0,5,7,0,2,},[81]int{0,0,0,7,0,0,8,0,0,0,0,6,0,0,0,0,3,1,0,4,0,0,0,2,0,0,0,0,2,4,0,7,0,0,0,0,0,1,0,0,3,0,0,8,0,0,0,0,0,6,0,2,9,0,0,0,0,8,0,0,0,7,0,8,6,0,0,0,0,5,0,0,0,0,2,0,0,6,0,0,0,},[81]int{0,0,1,0,0,7,0,9,0,5,9,0,0,8,0,0,0,1,0,3,0,0,0,0,0,8,0,0,0,0,0,0,5,8,0,0,0,5,0,0,6,0,0,2,0,0,0,4,1,0,0,0,0,0,0,8,0,0,0,0,0,3,0,1,0,0,0,2,0,0,7,9,0,2,0,7,0,0,4,0,0,},[81]int{0,0,0,0,0,3,0,1,7,0,1,5,0,0,9,0,0,8,0,6,0,0,0,0,0,0,0,1,0,0,0,0,7,0,0,0,0,0,9,0,0,0,2,0,0,0,0,0,5,0,0,0,0,4,0,0,0,0,0,0,0,2,0,5,0,0,6,0,0,3,4,0,3,4,0,2,0,0,0,0,0,},[81]int{3,0,0,2,0,0,0,0,0,0,0,0,1,0,7,0,0,0,7,0,6,0,3,0,5,0,0,0,7,0,0,0,9,0,8,0,9,0,0,0,2,0,0,0,4,0,1,0,8,0,0,0,5,0,0,0,9,0,4,0,3,0,1,0,0,0,7,0,2,0,0,0,0,0,0,0,0,8,0,0,6,}}

	sum := 0

	for i, _ := range sudokus {
		var sudoku Sudoku
		sudoku.setup(sudokus[i])
		sudoku.solve()
		
		if sudoku.solved {
			sudoku.print()
			fmt.Println("solved")
			fmt.Println("------")
			sum += sudoku.grid[0] * 100 + sudoku.grid[1] * 10 + sudoku.grid[2]
		} else {
			fmt.Println("PROBLEM!")
			break
		}
	}
	
	fmt.Println(sum)
	

/*	var sudoku Sudoku
	sudoku.setup(sudokus[1])
	sudoku.solve()
	sudoku.print()
	
	if sudoku.solved {
		fmt.Println("SOLVED!")
	}*/

}