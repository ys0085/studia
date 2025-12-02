package eval

import (
	"sort"
)

func (board *Board) countV(player int) int {
	total := 0
	for i := range 5 {
		counter := 0
		for j := range 5 {
			if board[i][j] == player {
				counter++
			} else if board[i][j] == 3-player && !(j == 0 || j == 4) {
				counter = 0
				break
			}
		}
		total += counter
	}
	return total
}

func (board *Board) countH(player int) int {
	total := 0
	for i := range 5 {
		counter := 0
		for j := range 5 {
			if board[j][i] == player {
				counter++
			} else if board[j][i] == 3-player && !(j == 0 || j == 4) {
				counter = 0
				break
			}
		}
		total += counter
	}
	return total
}

func (board *Board) countD(player int) int {
	total := 0
	for k := range 2 {
		counter := 0
		for i := range 5 - k {
			j := i + k
			if board[i][j] == player {
				counter++
			} else if board[i][j] == 3-player && !(i == 0 || i == 4 || j == 0 || j == 4) {
				counter = 0
				break
			}
		}
		total += counter
	}

	counter := 0
	for i := range 4 {
		j := i
		if board[i+1][j] == player {
			counter++
		} else if board[i+1][j] == 3-player && !(i+1 == 0 || i+1 == 4 || j == 0 || j == 4) {
			counter = 0
			break
		}
	}
	total += counter

	for k := range 2 {
		counter := 0
		for i := range 5 - k {
			j := 4 - (i + k)
			if board[i][j] == player {
				counter++
			} else if board[i][j] == 3-player && !(i == 0 || i == 4 || j == 0 || j == 4) {
				counter = 0
				break
			}
		}
		total += counter
	}

	counter = 0
	for i := range 4 {
		j := 4 - i
		if board[i+1][j] == player {
			counter++
		} else if board[i+1][j] == 3-player && !(i+1 == 0 || i+1 == 4 || j == 0 || j == 4) {
			counter = 0
			break
		}
	}
	total += counter

	return total
}

func (board *Board) countOneOffWinConditionSpots(player int) int {
	count := 0
	for i := range 28 {
		winCon1 := board[win[i][0][0]][win[i][0][1]] == player &&
			board[win[i][1][0]][win[i][1][1]] == player &&
			board[win[i][3][0]][win[i][3][1]] == player
		winCon2 := board[win[i][0][0]][win[i][0][1]] == player &&
			board[win[i][2][0]][win[i][2][1]] == player &&
			board[win[i][3][0]][win[i][3][1]] == player

		if winCon1 || winCon2 {
			count++
		}
	}
	return count
}

func (board *Board) countOneOffLossConditionSpots(player int) int {
	count := 0
	for i := range 48 {
		loseCon1 := board[lose[i][0][0]][lose[i][0][1]] == 0 &&
			board[lose[i][1][0]][lose[i][1][1]] == player &&
			board[lose[i][2][0]][lose[i][2][1]] == player
		loseCon2 := board[lose[i][0][0]][lose[i][0][1]] == player &&
			board[lose[i][1][0]][lose[i][1][1]] == 0 &&
			board[lose[i][2][0]][lose[i][2][1]] == player
		loseCon3 := board[lose[i][0][0]][lose[i][0][1]] == player &&
			board[lose[i][1][0]][lose[i][1][1]] == player &&
			board[lose[i][2][0]][lose[i][2][1]] == 0

		if loseCon1 || loseCon2 || loseCon3 {
			count++
		}
	}
	return count
}

func (board *Board) calculateHeuristic(player int) float64 {
	switch EVALUATION_MODE {
	case 1:
		h := board.countH(player)
		v := board.countV(player)
		d := board.countD(player)
		return float64(h + v + d)

	case 2:
		return float64(board.countOneOffWinConditionSpots(player))
	case 3:
		return 0.7*float64(board.countOneOffLossConditionSpots(3-player)) + 1.1*float64(board.countOneOffWinConditionSpots(player))
	default:
		return 0.0
	}
}

func sortByHeuristic(moves []int, player int) []int {
	sort.Slice(moves, func(i, j int) bool {
		board := Board{}
		board.SetEmptyBoard()
		board.SetMove(moves[i], player)
		h1 := board.calculateHeuristic(player)

		board.SetEmptyBoard()
		board.SetMove(moves[j], player)
		h2 := board.calculateHeuristic(player)

		return h1 > h2
	})
	return moves
}
