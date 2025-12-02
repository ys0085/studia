package main

import (
	"15puzzle/gen"
	"15puzzle/solve"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var BOARD_SIZE = 4

const ANIMATION_SPEED = 200 * time.Millisecond

// Wydaje mi się, że działa nawet dla 5x5, dla 6x6 trwa bardzo długo

func Presentation() {
	input := bufio.NewScanner(os.Stdin)

	var numbers string

	fmt.Print("Enter size (3-5): ")
	input.Scan()
	size := input.Text()
	BOARD_SIZE = int(size[0] - '0')
	if BOARD_SIZE < 3 || BOARD_SIZE > 6 {
		fmt.Println("Invalid size. Exiting...")
		return
	}
	fmt.Println("Enter numbers: (0 for empty tile)")
	input.Scan()
	numbers = input.Text()
	board, e := gen.ParseBoard(numbers, BOARD_SIZE)
	if e != nil {
		fmt.Println("Error parsing board: ", e)
		fmt.Println("Generating random board...")
		board = gen.RandomBoard(BOARD_SIZE)
	}
	fmt.Println("Initial board:")
	printBoard(solve.BoardType(board))
	fmt.Println("Is solvable: ", gen.IsSolvable(board))
	if !gen.IsSolvable(board) {
		fmt.Println("Board is not solvable. Exiting...")
		return
	}
	solved, counter, elapsed, _ := solve.SolveBoard(solve.BoardType(board))
	states := []solve.StateType{}
	for solved.Previous != nil {
		states = append(states, solved)
		solved = *solved.Previous
	}
	totalSteps := len(states)

	fmt.Println("Number of iterations: ", counter)
	fmt.Println("Time elapsed: ", elapsed)
	fmt.Println("Number of steps: ", len(states))

	fmt.Println("Press Enter to start the animation...")
	input.Scan()

	for i := len(states) - 1; i >= 0; i-- {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Printf("Step: %d/%d Heuristic: %d\n", states[i].Cost, totalSteps, states[i].Heuristic)
		printBoard(states[i].Board)
		time.Sleep(ANIMATION_SPEED)
	}
}

func printBoard(board solve.BoardType) {
	for _, row := range board {
		for _, tile := range row {
			if tile == 0 {
				fmt.Print("   ")
			} else {
				fmt.Printf("%2d ", tile)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
