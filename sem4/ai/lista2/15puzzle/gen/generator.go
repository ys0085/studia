package gen

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type BoardType [][]int

func RandomBoard(size int) BoardType {
	var board BoardType = make([][]int, size)
	pool := make([]int, size*size-1)
	for i := range pool {
		pool[i] = i + 1
	}
	rand.Shuffle(len(pool), func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})
	pool = append(pool, 0) // Add the empty tile (0) at the end

	for i := range board {
		board[i] = make([]int, size)
	}
	for i := range size {
		for j := 0; j < size; j++ {
			board[i][j] = pool[i*size+j]
		}
	}
	if !IsSolvable(board) {
		return RandomBoard(size) // Regenerate if not solvable
	}
	return board
}

func IsSolvable(board BoardType) bool {
	size := len(board)
	flatBoard := make([]int, 0, size*size)
	for i := range board {
		flatBoard = append(flatBoard, board[i]...)
	}
	inversions := 0
	for i := range flatBoard {
		for j := i + 1; j < len(flatBoard); j++ {
			if flatBoard[i] != 0 && flatBoard[j] != 0 && flatBoard[i] > flatBoard[j] {
				inversions++
			}
		}
	}
	return inversions%2 == 0
}

func ParseBoard(numbers string, size int) (BoardType, error) {
	var board BoardType = make([][]int, size)
	nums := strings.Split(numbers, " ")
	if len(nums) != size*size {
		return nil, fmt.Errorf("invalid input: expected %d numbers, got %d", size*size, len(nums))
	}
	for i := range board {
		board[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			num, err := strconv.Atoi(nums[i*size+j])
			if err != nil {
				return nil, fmt.Errorf("invalid number at position %d: %v", i*size+j, err)
			}
			board[i][j] = num
		}
	}
	if !VerifyBoard(board) {
		return nil, fmt.Errorf("invalid board: numbers out of range or duplicates")
	}
	return board, nil
}

func VerifyBoard(board BoardType) bool {
	size := len(board)
	pool := make([]int, size*size-1)
	for i := range pool {
		pool[i] = i + 1
	}
	if board[size-1][size-1] != 0 {
		return false
	}
	seen := make(map[int]bool)
	for i := range size {
		for j := range size {
			val := board[i][j]
			if val < 0 || val >= size*size || seen[val] {
				return false
			}
			seen[val] = true
		}
	}
	return true
}

func GenerateGoalState(size int) BoardType {
	var goalState BoardType = make([][]int, size)
	for i := range goalState {
		goalState[i] = make([]int, size)
	}
	for i := range size {
		for j := range size {
			if i == size-1 && j == size-1 {
				goalState[i][j] = 0
			} else {
				goalState[i][j] = i*size + j + 1
			}
		}
	}
	return goalState
}
