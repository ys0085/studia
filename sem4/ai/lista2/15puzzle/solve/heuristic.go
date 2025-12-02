package solve

func manhattanDistance(board BoardType) int {
	distance := 0
	size := len(board)
	for i := range size {
		for j := range size {
			if board[i][j] != 0 {
				targetRow := (board[i][j] - 1) / size
				targetCol := (board[i][j] - 1) % size
				if i > targetRow {
					distance += i - targetRow
				} else {
					distance += targetRow - i
				}
				if j > targetCol {
					distance += j - targetCol
				} else {
					distance += targetCol - j
				}
			}
		}
	}
	return distance
}

func mismatchedTiles(board BoardType) int {
	mismatches := 0
	size := len(board)
	for i := range size {
		for j := range size {
			if board[i][j] != 0 && board[i][j] != i*size+j+1 {
				mismatches++
			}
		}
	}
	return mismatches
}

func linearConflict(board BoardType) int {
	size := len(board)
	linearConflict := 0

	// Row conflicts
	for row := range size {
		for column1 := range size - 1 {
			val1 := board[row][column1]
			if val1 == 0 {
				continue
			}
			targetRow1 := (val1 - 1) / size
			if targetRow1 != row {
				continue
			}
			for column2 := column1 + 1; column2 < size; column2++ {
				val2 := board[row][column2]
				if val2 == 0 {
					continue
				}
				targetRow2 := (val2 - 1) / size
				if targetRow2 != row {
					continue
				}

				targetColumn1 := (val1 - 1) % size
				targetColumn2 := (val2 - 1) % size

				if targetColumn1 > targetColumn2 {
					linearConflict++
				}
			}
		}
	}

	// Column conflicts
	for column := range size {
		for row1 := range size - 1 {
			val1 := board[row1][column]
			if val1 == 0 {
				continue
			}
			targetColumn1 := (val1 - 1) % size
			if targetColumn1 != column {
				continue
			}

			for row2 := row1 + 1; row2 < size; row2++ {
				val2 := board[row2][column]
				if val2 == 0 {
					continue
				}
				targetColumn2 := (val2 - 1) % size
				if targetColumn2 != column {
					continue
				}

				targetRow1 := (val1 - 1) / size
				targetRow2 := (val2 - 1) / size

				if targetRow1 > targetRow2 {
					linearConflict++
				}
			}
		}
	}

	return linearConflict
}

func calculateHeuristic(state StateType) int {
	return manhattanDistance(state.Board) + 2*linearConflict(state.Board) + mismatchedTiles(state.Board)
}
