package solve

import (
	"fmt"
	"time"
)

type BoardType [][]int

type PositionType [2]int

type StateType struct {
	Board       BoardType
	Cost        int
	Heuristic   int
	Previous    *StateType
	emptyRow    int
	emptyColumn int
}

func SolveBoard(board BoardType) (StateType, int, time.Duration, error) {
	var size int = len(board)
	startTime := time.Now()
	counter := 0
	var currentState StateType

	pq := PriorityQueue{}
	initialState := StateType{Board: board, Cost: 0, Heuristic: calculateHeuristic(StateType{Board: board}),
		emptyRow: size - 1, emptyColumn: size - 1}
	pq.enqueue(initialState)

	visited := make(map[uint64]bool, 0)

	for len(pq) > 0 {
		counter++
		currentState := pq.dequeue()

		boardHash := hashBoard(currentState.Board)
		if visited[boardHash] {
			continue
		}
		visited[boardHash] = true

		if currentState.Heuristic == 0 {
			elapsed := time.Since(startTime)
			return currentState, counter, elapsed, nil
		}

		deltaRow := []int{-1, 1, 0, 0}
		deltaColumn := []int{0, 0, -1, 1}

		//DEBUG
		// fmt.Printf("%07d, %03d\n", counter, currentState.Heuristic+currentState.Cost)

		for i := range 4 {
			newRow := currentState.emptyRow + deltaRow[i]
			newColumn := currentState.emptyColumn + deltaColumn[i]
			if newRow >= 0 && newRow < size && newColumn >= 0 && newColumn < size {
				var next StateType
				next.Board = make(BoardType, len(currentState.Board))
				for i := range currentState.Board {
					next.Board[i] = make([]int, size)
					copy(next.Board[i], currentState.Board[i])
				}
				next.Previous = &currentState

				next.Board[currentState.emptyRow][currentState.emptyColumn] = currentState.Board[newRow][newColumn]
				next.Board[newRow][newColumn] = 0

				next.emptyRow = newRow
				next.emptyColumn = newColumn
				next.Cost = currentState.Cost + 1
				next.Heuristic = calculateHeuristic(next)
				pq.enqueue(next)
			}
		}
	}
	elapsed := time.Since(startTime)
	return currentState, counter, elapsed, fmt.Errorf("no solution found")
}

func hashBoard(board BoardType) uint64 {
	var hash uint64
	const prime uint64 = 31
	for _, row := range board {
		for _, val := range row {
			hash = hash*prime + uint64(val)
		}
	}
	return hash
}
