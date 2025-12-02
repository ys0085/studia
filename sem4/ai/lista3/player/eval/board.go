package eval

import "fmt"

type Board [5][5]int

// Win positions - 28 patterns of 4 cells for a win
var win = [28][4][2]int{
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	{{1, 0}, {1, 1}, {1, 2}, {1, 3}},
	{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
	{{3, 0}, {3, 1}, {3, 2}, {3, 3}},
	{{4, 0}, {4, 1}, {4, 2}, {4, 3}},
	{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
	{{1, 1}, {1, 2}, {1, 3}, {1, 4}},
	{{2, 1}, {2, 2}, {2, 3}, {2, 4}},
	{{3, 1}, {3, 2}, {3, 3}, {3, 4}},
	{{4, 1}, {4, 2}, {4, 3}, {4, 4}},
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
	{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
	{{0, 3}, {1, 3}, {2, 3}, {3, 3}},
	{{0, 4}, {1, 4}, {2, 4}, {3, 4}},
	{{1, 0}, {2, 0}, {3, 0}, {4, 0}},
	{{1, 1}, {2, 1}, {3, 1}, {4, 1}},
	{{1, 2}, {2, 2}, {3, 2}, {4, 2}},
	{{1, 3}, {2, 3}, {3, 3}, {4, 3}},
	{{1, 4}, {2, 4}, {3, 4}, {4, 4}},
	{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
	{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
	{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
	{{1, 0}, {2, 1}, {3, 2}, {4, 3}},
	{{0, 3}, {1, 2}, {2, 1}, {3, 0}},
	{{0, 4}, {1, 3}, {2, 2}, {3, 1}},
	{{1, 3}, {2, 2}, {3, 1}, {4, 0}},
	{{1, 4}, {2, 3}, {3, 2}, {4, 1}},
}

// Lose positions - 48 patterns of 3 cells for a loss
var lose = [48][3][2]int{
	{{0, 0}, {0, 1}, {0, 2}}, {{0, 1}, {0, 2}, {0, 3}}, {{0, 2}, {0, 3}, {0, 4}},
	{{1, 0}, {1, 1}, {1, 2}}, {{1, 1}, {1, 2}, {1, 3}}, {{1, 2}, {1, 3}, {1, 4}},
	{{2, 0}, {2, 1}, {2, 2}}, {{2, 1}, {2, 2}, {2, 3}}, {{2, 2}, {2, 3}, {2, 4}},
	{{3, 0}, {3, 1}, {3, 2}}, {{3, 1}, {3, 2}, {3, 3}}, {{3, 2}, {3, 3}, {3, 4}},
	{{4, 0}, {4, 1}, {4, 2}}, {{4, 1}, {4, 2}, {4, 3}}, {{4, 2}, {4, 3}, {4, 4}},
	{{0, 0}, {1, 0}, {2, 0}}, {{1, 0}, {2, 0}, {3, 0}}, {{2, 0}, {3, 0}, {4, 0}},
	{{0, 1}, {1, 1}, {2, 1}}, {{1, 1}, {2, 1}, {3, 1}}, {{2, 1}, {3, 1}, {4, 1}},
	{{0, 2}, {1, 2}, {2, 2}}, {{1, 2}, {2, 2}, {3, 2}}, {{2, 2}, {3, 2}, {4, 2}},
	{{0, 3}, {1, 3}, {2, 3}}, {{1, 3}, {2, 3}, {3, 3}}, {{2, 3}, {3, 3}, {4, 3}},
	{{0, 4}, {1, 4}, {2, 4}}, {{1, 4}, {2, 4}, {3, 4}}, {{2, 4}, {3, 4}, {4, 4}},
	{{0, 2}, {1, 3}, {2, 4}}, {{0, 1}, {1, 2}, {2, 3}}, {{1, 2}, {2, 3}, {3, 4}},
	{{0, 0}, {1, 1}, {2, 2}}, {{1, 1}, {2, 2}, {3, 3}}, {{2, 2}, {3, 3}, {4, 4}},
	{{1, 0}, {2, 1}, {3, 2}}, {{2, 1}, {3, 2}, {4, 3}}, {{2, 0}, {3, 1}, {4, 2}},
	{{0, 2}, {1, 1}, {2, 0}}, {{0, 3}, {1, 2}, {2, 1}}, {{1, 2}, {2, 1}, {3, 0}},
	{{0, 4}, {1, 3}, {2, 2}}, {{1, 3}, {2, 2}, {3, 1}}, {{2, 2}, {3, 1}, {4, 0}},
	{{1, 4}, {2, 3}, {3, 2}}, {{2, 3}, {3, 2}, {4, 1}}, {{2, 4}, {3, 3}, {4, 2}},
}

// SetBoard initializes the board to empty state
func (board *Board) SetEmptyBoard() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			board[i][j] = 0
		}
	}
}

// PrintBoard displays the current state of the board
func PrintBoard(board Board) {
	fmt.Println("  1 2 3 4 5")
	for i := 0; i < 5; i++ {
		fmt.Print(i + 1)
		for j := 0; j < 5; j++ {
			switch board[i][j] {
			case 0:
				fmt.Print(" -")
			case 1:
				fmt.Print(" X")
			case 2:
				fmt.Print(" O")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// SetMove places a player's piece on the board
func (board *Board) SetMove(move, player int) bool {
	i := (move / 10) - 1
	j := (move % 10) - 1
	if i < 0 || i > 4 || j < 0 || j > 4 {
		return false
	}
	if board[i][j] != 0 {
		return false
	}
	board[i][j] = player
	return true
}

// WinCheck checks if the player has won
func (board *Board) WinCheck(player int) bool {
	w := false
	for i := 0; i < 28; i++ {
		if board[win[i][0][0]][win[i][0][1]] == player &&
			board[win[i][1][0]][win[i][1][1]] == player &&
			board[win[i][2][0]][win[i][2][1]] == player &&
			board[win[i][3][0]][win[i][3][1]] == player {
			w = true
		}
	}
	return w
}

// LoseCheck checks if the player has lost
func (board *Board) LoseCheck(player int) bool {
	l := false
	for i := 0; i < 48; i++ {
		if board[lose[i][0][0]][lose[i][0][1]] == player &&
			board[lose[i][1][0]][lose[i][1][1]] == player &&
			board[lose[i][2][0]][lose[i][2][1]] == player {
			l = true
		}
	}
	return l
}

func (board *Board) FreeTiles() []int {
	var free []int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == 0 {
				free = append(free, (i+1)*10+(j+1))
			}
		}
	}
	return free
}

func (board *Board) FreeCenterTiles() []int {
	var free []int
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if board[i][j] == 0 {
				free = append(free, (i+1)*10+(j+1))
			}
		}
	}
	return free
}

func (board *Board) Copy() Board {
	var newBoard Board
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			newBoard[i][j] = board[i][j]
		}
	}
	return newBoard
}

func (board *Board) Hash() uint64 {
	var hash uint64
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			hash = hash*3 + uint64(board[i][j])
		}
	}
	return hash
}
