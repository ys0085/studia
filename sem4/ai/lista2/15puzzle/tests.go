package main

import (
	"15puzzle/gen"
	"15puzzle/solve"
	"fmt"
	"sync"
	"time"
)

func singleTest(size int) (int, int, time.Duration) {
	board := gen.RandomBoard(size)
	state, iter, elapsed, _ := solve.SolveBoard(solve.BoardType(board))
	return iter, state.Cost, elapsed
}

func RunTests(boardSize int, n int) {
	var totalTime time.Duration
	var totalIterations int64
	var totalMoves int
	var times []time.Duration = make([]time.Duration, n)
	var iterationCounters []int = make([]int, n)
	var moveCounters []int = make([]int, n)
	var wg sync.WaitGroup

	counter := 0
	for i := range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			iterationCounters[i], moveCounters[i], times[i] = singleTest(boardSize)
			totalIterations += int64(iterationCounters[i])
			totalTime += times[i]
			totalMoves += moveCounters[i]
			counter++
			fmt.Println(counter, "out of", n, "in", times[i].Round(time.Millisecond), "and", iterationCounters[i], "iterations,", moveCounters[i], "moves.")
		}()
	}

	wg.Wait()

	averageTime := totalTime / time.Duration(n)
	averageIterations := totalIterations / int64(n)
	averageMoves := totalMoves / n

	fmt.Println("Avg time: ", averageTime.Round(time.Millisecond))
	fmt.Println("Avg iterations: ", averageIterations)
	fmt.Println("Avg moves: ", averageMoves)

}
