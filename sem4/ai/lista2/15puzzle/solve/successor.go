package solve

func findEmptyTile(board BoardType) PositionType {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				return PositionType{i, j}
			}
		}
	}
	return PositionType{-1, -1} // Should never happen if the board is valid
}

func isValidPosition(pos PositionType, board BoardType) bool {
	return pos[0] >= 0 && pos[0] < len(board) && pos[1] >= 0 && pos[1] < len(board[0])
}

func swapTiles(board BoardType, pos1, pos2 PositionType) {
	board[pos1[0]][pos1[1]], board[pos2[0]][pos2[1]] = board[pos2[0]][pos2[1]], board[pos1[0]][pos1[1]]
}
