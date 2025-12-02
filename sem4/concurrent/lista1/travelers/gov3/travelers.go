package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	Nr_Of_Travelers = 15
	Min_Steps       = 10
	Max_Steps       = 100
	Min_Delay       = 10 * time.Millisecond
	Max_Delay       = 50 * time.Millisecond
	Board_Width     = 15
	Board_Height    = 15
	Trapped_Timeout = 100 * time.Millisecond // Timeout before a traveler is considered trapped
)

var Start_Time time.Time

var wg sync.WaitGroup

// Board structure to track occupied positions
type Board struct {
	occupied [][]bool
	mutex    sync.Mutex
}

func NewBoard() *Board {
	board := &Board{
		occupied: make([][]bool, Board_Height),
	}
	for i := range board.occupied {
		board.occupied[i] = make([]bool, Board_Width)
	}
	return board
}

func (b *Board) IsOccupied(pos Position) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.occupied[pos.Y][pos.X]
}

func (b *Board) Occupy(pos Position) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.occupied[pos.Y][pos.X] = true
}

func (b *Board) Free(pos Position) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.occupied[pos.Y][pos.X] = false
}

type Position struct {
	X int
	Y int
}

func Move_Down(pos *Position) {
	pos.Y = (pos.Y + 1) % Board_Height
}

func Move_Up(pos *Position) {
	pos.Y = (pos.Y + Board_Height - 1) % Board_Height
}

func Move_Right(pos *Position) {
	pos.X = (pos.X + 1) % Board_Width
}

func Move_Left(pos *Position) {
	pos.X = (pos.X + Board_Width - 1) % Board_Width
}

type Trace struct {
	timestamp time.Duration
	id        int
	position  Position
	symbol    rune
}

type Trace_Sequence []Trace

type Traveler struct {
	Id           int
	Symbol       rune
	Position     Position
	LastMoveTime time.Time
	IsTrapped    bool
}

func Print_Trace(trace Trace) {
	fmt.Printf("%f %d %d %d %c\n", float64(trace.timestamp.Microseconds())/1e6, trace.id, trace.position.X, trace.position.Y, trace.symbol)
}

func Print_Trace_Sequence(sequence Trace_Sequence) {
	for _, trace := range sequence {
		Print_Trace(trace)
	}
}

func printer(report <-chan Trace_Sequence, done chan<- bool) {
	for traces := range report {
		Print_Trace_Sequence(traces)
	}
	done <- true
}

func (t *Traveler) tryMove(board *Board) bool {
	var directions []int
	if t.Id%2 == 0 {
		directions = []int{0, 1}
	} else {
		directions = []int{2, 3}
	}
	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	for _, dir := range directions {
		oldPos := t.Position

		switch dir {
		case 0:
			Move_Up(&t.Position)
		case 1:
			Move_Down(&t.Position)
		case 2:
			Move_Left(&t.Position)
		case 3:
			Move_Right(&t.Position)
		}

		if !board.IsOccupied(t.Position) {
			board.Occupy(t.Position)
			board.Free(oldPos)
			t.LastMoveTime = time.Now()

			return true
		}

		t.Position = oldPos
	}

	if time.Since(t.LastMoveTime) > Trapped_Timeout && !t.IsTrapped {
		t.IsTrapped = true
		t.Symbol = toLower(t.Symbol)
	}

	return false
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func traveler(id int, symbol rune, report chan<- Trace_Sequence, board *Board) {
	defer wg.Done()

	var initialPosition Position
	initialPosition.X, initialPosition.Y = id, id

	traveler := Traveler{
		Id:           id,
		Symbol:       symbol,
		Position:     initialPosition,
		LastMoveTime: time.Now(),
		IsTrapped:    false,
	}

	board.Occupy(traveler.Position)

	nrOfSteps := Min_Steps + rand.Intn(Max_Steps-Min_Steps)
	var traces Trace_Sequence
	traces = append(traces, Trace{
		timestamp: time.Since(Start_Time),
		id:        traveler.Id,
		position:  traveler.Position,
		symbol:    traveler.Symbol,
	})
	for range nrOfSteps {
		if traveler.IsTrapped {
			break
		}
		delay := Min_Delay + time.Duration(rand.Intn(int(Max_Delay-Min_Delay)))
		time.Sleep(delay)

		traveler.tryMove(board)
		traces = append(traces, Trace{
			timestamp: time.Since(Start_Time),
			id:        traveler.Id,
			position:  traveler.Position,
			symbol:    traveler.Symbol,
		})

	}

	report <- traces
}

func main() {
	fmt.Println("-1", Nr_Of_Travelers, " ", Board_Width, " ", Board_Height)

	board := NewBoard()

	Start_Time = time.Now()
	report := make(chan Trace_Sequence, Nr_Of_Travelers)
	done := make(chan bool)

	go printer(report, done)

	for i := range Nr_Of_Travelers {
		wg.Add(1)
		go traveler(i, rune('A'+i), report, board)
	}

	wg.Wait()
	close(report)
	<-done
}
