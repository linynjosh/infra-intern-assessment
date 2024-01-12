package main

func SolveSudoku(board [][]int) [][]int {
	// Check if board has valid dimensions
	if len(board) != 9 || len(board[0]) != 9 {
		return nil
	}

	// Find solution
	if Backtrack(board) {
		return board
	}

	return nil
}

// Constraint 1: Each row contains all the number from 1 to 9 without repetition
// Constraint 2: Each col contains all the number from 1 to 9 without repetition
// Constraint 3: Each sub-grid contains all the number from 1 to 9 without repetition
func Backtrack(board [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// Check if cell is empty
			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					// Check if filling the empty cell with the number breaks any of constraints 1, 2, 3
					if IsValid(board, row, col, num) {
						// If all constraints are met we put down the num and see if this choice leads to a solution
						board[row][col] = num
						// Return true if this choice leads to a solution
						if Backtrack(board) {
							return true
						} else {
							// If this number doesn't bring us to a solution, we undo this choice
							board[row][col] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

// Check if constaints 1, 2, 3 are held
func IsValid(board [][]int, row, col, num int) bool {
	return IsValidRow(board, row, col, num) && IsValidCol(board, row, col, num) && IsValidSubGrid(board, row, col, num)
}

// Checks constraint 1
func IsValidRow(board [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}
	return true
}

// Checks constraint 2
func IsValidCol(board [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	return true
}

// Checks constraint 3
func IsValidSubGrid(board [][]int, row, col, num int) bool {
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}