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
)

var Start_Time time.Time

var wg sync.WaitGroup

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
	Id       int
	Symbol   rune
	Position Position
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

func (t *Traveler) move() {
	switch rand.Intn(4) {
	case 0:
		Move_Up(&t.Position)
	case 1:
		Move_Down(&t.Position)
	case 2:
		Move_Left(&t.Position)
	case 3:
		Move_Right(&t.Position)
	}
}

func traveler(id int, symbol rune, report chan<- Trace_Sequence) {
	defer wg.Done()

	traveler := Traveler{
		Id:       id,
		Symbol:   symbol,
		Position: Position{rand.Intn(Board_Width), rand.Intn(Board_Height)},
	}

	nrOfSteps := Min_Steps + rand.Intn(Max_Steps-Min_Steps)
	var traces Trace_Sequence

	for range nrOfSteps {
		delay := Min_Delay + time.Duration(rand.Intn(int(Max_Delay-Min_Delay)))
		time.Sleep(delay)

		traveler.move()
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

	Start_Time = time.Now()
	report := make(chan Trace_Sequence, Nr_Of_Travelers)
	done := make(chan bool)

	go printer(report, done)

	for i := range Nr_Of_Travelers {
		wg.Add(1)
		go traveler(i, rune('A'+i), report)
	}

	wg.Wait()
	close(report)
	<-done
}
