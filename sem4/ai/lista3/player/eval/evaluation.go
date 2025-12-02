package eval

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
)

const WIN_VALUE float64 = 100
const LOSS_VALUE float64 = -100

var EVALUATION_MODE int = 1

func SetEvaluationMode(mode int) {
	EVALUATION_MODE = mode
}

func Move(board Board, maxDepth int, player int) int {
	if maxDepth > 9 {
		maxDepth = 9
	}
	type moveEval struct {
		move int
		eval float64
	}

	moves := board.FreeTiles()
	if len(moves) > 20 {
		moves = board.FreeCenterTiles()
	}

	moves = sortByHeuristic(moves, player)

	evals := make([]moveEval, 0, len(moves))
	var wg sync.WaitGroup
	for _, move := range moves {
		newBoard := board.Copy()
		newBoard.SetMove(move, player)
		e := newBoard.evaluate(player)
		if e == LOSS_VALUE {
			evals = append(evals, moveEval{move, LOSS_VALUE * 2})
			continue
		}
		if e == WIN_VALUE {
			evals = append(evals, moveEval{move, WIN_VALUE * 2})
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			eval := minmax(newBoard, maxDepth, 1, LOSS_VALUE, WIN_VALUE, false, 3-player, player)
			evals = append(evals, moveEval{move, eval})
		}()
	}
	wg.Wait()

	sort.Slice(evals, func(i, j int) bool {
		if evals[i].eval == evals[j].eval {
			return rand.Intn(2) == 0
		}
		return evals[i].eval > evals[j].eval
	})

	fmt.Println("Top moves and evaluations:")
	for i := 0; /*i < 5 && */ i < len(evals); i++ {
		fmt.Printf("Move: %v, Eval: %v\n", evals[i].move, evals[i].eval)
	}

	return evals[0].move
}

func minmax(board Board, maxDepth, depth int, alpha, beta float64, maximizingPlayer bool, currentPlayer int, originalPlayer int) float64 {
	if depth == maxDepth || board.isTerminal() {
		return board.evaluate(originalPlayer)
	}

	moves := board.FreeTiles()
	if len(moves) > 20 {
		moves = board.FreeCenterTiles()
	}
	//moves = sortByHeuristic(moves, currentPlayer)
	result := 0.0
	if maximizingPlayer {
		maxEval := LOSS_VALUE
		for _, move := range moves {
			newBoard := board.Copy()
			newBoard.SetMove(move, currentPlayer)
			eval := minmax(newBoard, maxDepth, depth+1, alpha, beta, false, 3-currentPlayer, originalPlayer)
			maxEval = math.Max(maxEval, eval)
			alpha = math.Max(alpha, eval)
			if beta <= alpha {
				break
			}
		}
		result = maxEval
	} else {
		minEval := WIN_VALUE
		for _, move := range moves {
			newBoard := board.Copy()
			newBoard.SetMove(move, currentPlayer)
			eval := minmax(newBoard, maxDepth, depth+1, alpha, beta, true, 3-currentPlayer, originalPlayer)
			minEval = math.Min(minEval, eval)
			beta = math.Min(beta, eval)
			if beta <= alpha {
				break
			}
		}
		result = minEval
	}
	return result
}

func (board *Board) evaluate(player int) float64 {
	if board.WinCheck(player) || (board.LoseCheck(3-player) && !board.WinCheck(3-player)) {
		return WIN_VALUE
	}
	if board.WinCheck(3-player) || (board.LoseCheck(player) && !board.WinCheck(player)) {
		return LOSS_VALUE
	}
	if len(board.FreeTiles()) == 0 {
		return 0
	}

	return board.calculateHeuristic(player) - board.calculateHeuristic(3-player)
}

func (board *Board) isTerminal() bool {
	return board.LoseCheck(1) || board.LoseCheck(2) || len(board.FreeTiles()) == 0
}
