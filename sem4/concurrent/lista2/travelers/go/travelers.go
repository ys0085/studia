package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	Nr_Of_Travelers      = 25
	Nr_Of_Wild_Travelers = 6
	Nr_Of_Trap_Tiles     = 2

	Min_Steps = 5
	Max_Steps = 40

	Min_Delay = 40 * time.Millisecond
	Max_Delay = 200 * time.Millisecond

	Wild_Traveler_Duration  = 1000 * time.Millisecond
	Wild_Traveler_Min_Delay = 100 * time.Millisecond
	Wild_Traveler_Max_Delay = 2 * time.Second

	Board_Width  = 8
	Board_Height = 8

	Trapped_Timeout = 100 * time.Millisecond // Timeout before a traveler is considered trapped
	Trap_Duration   = 200 * time.Millisecond // Duration of the trap effect
)

var Start_Time time.Time

var wg sync.WaitGroup

type Spot struct {
	occupied bool
	traveler *Traveler
	request  chan chan bool
	occupy   chan *Traveler
	free     chan bool
}

func NewSpot() *Spot {
	s := &Spot{
		occupied: false,
		traveler: nil,
		request:  make(chan chan bool),
		occupy:   make(chan *Traveler),
		free:     make(chan bool),
	}
	go s.run()
	return s
}

func (s *Spot) run() {
	for {
		select {
		case reply := <-s.request:
			reply <- s.occupied
		case t := <-s.occupy:
			s.occupied = true
			s.traveler = t
		case <-s.free:
			s.occupied = false
			s.traveler = nil
		}
	}
}

type Board struct {
	spots [][]*Spot
}

func NewBoard() *Board {
	board := &Board{spots: make([][]*Spot, Board_Width)}
	for i := 0; i < Board_Width; i++ {
		board.spots[i] = make([]*Spot, Board_Height)
		for j := 0; j < Board_Height; j++ {
			board.spots[i][j] = NewSpot()
		}
	}
	return board
}

func (b *Board) IsOccupied(pos Position) bool {
	reply := make(chan bool)
	b.spots[pos.X][pos.Y].request <- reply
	return <-reply
}

func (b *Board) Occupy(pos Position, traveler *Traveler) {
	b.spots[pos.X][pos.Y].occupy <- traveler
}

func (b *Board) Free(pos Position) {
	b.spots[pos.X][pos.Y].free <- true
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
	IsWild       bool
	IsTrap       bool
	Traces       Trace_Sequence
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
	directions := []int{0, 1, 2, 3}
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
			board.Occupy(t.Position, t)
			board.Free(oldPos)
			t.LastMoveTime = time.Now()
			t.Traces = append(t.Traces, Trace{
				timestamp: time.Since(Start_Time),
				id:        t.Id,
				position:  t.Position,
				symbol:    t.Symbol,
			})
			return true
		} else {
			opponent := board.spots[t.Position.X][t.Position.Y].traveler
			if opponent.IsWild {
				opponentSuccess := opponent.tryMove(board)
				if opponentSuccess {
					board.Occupy(t.Position, t)
					board.Free(oldPos)
					t.LastMoveTime = time.Now()
					t.Traces = append(t.Traces, Trace{
						timestamp: time.Since(Start_Time),
						id:        t.Id,
						position:  t.Position,
						symbol:    t.Symbol,
					})
					return true
				}
			} else if opponent.IsTrap {
				if t.IsWild {
					t.Symbol = '*'
				} else {
					t.Symbol = toLower(t.Symbol)
				}
				t.IsTrapped = true
				t.Traces = append(t.Traces, Trace{
					timestamp: time.Since(Start_Time),
					id:        t.Id,
					position:  t.Position,
					symbol:    t.Symbol,
				})
				board.Free(oldPos)
				board.Occupy(t.Position, t)
				go func() {
					time.Sleep(Trap_Duration)
					board.Occupy(t.Position, opponent)
					opponent.Traces = append(opponent.Traces, Trace{
						timestamp: time.Since(Start_Time),
						id:        opponent.Id,
						position:  opponent.Position,
						symbol:    opponent.Symbol,
					})
				}()
				return false
			}
		}
		t.Position = oldPos
	}

	if time.Since(t.LastMoveTime) > Trapped_Timeout && !t.IsTrapped {
		t.IsTrapped = true
		t.Symbol = toLower(t.Symbol)
	}

	t.Traces = append(t.Traces, Trace{
		timestamp: time.Since(Start_Time),
		id:        t.Id,
		position:  t.Position,
		symbol:    t.Symbol,
	})
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
	for {
		initialPosition = Position{rand.Intn(Board_Width), rand.Intn(Board_Height)}
		if !board.IsOccupied(initialPosition) {
			break
		}
	}

	traveler := Traveler{
		Id:           id,
		Symbol:       symbol,
		Position:     initialPosition,
		LastMoveTime: time.Now(),
		IsTrapped:    false,
		IsWild:       false,
		Traces:       make(Trace_Sequence, 0),
	}

	defer func() { report <- traveler.Traces }()

	board.Occupy(traveler.Position, &traveler)

	nrOfSteps := Min_Steps + rand.Intn(Max_Steps-Min_Steps)

	for range nrOfSteps {
		if traveler.IsTrapped {
			return
		}
		delay := Min_Delay + time.Duration(rand.Intn(int(Max_Delay-Min_Delay)))
		time.Sleep(delay)

		traveler.tryMove(board)

	}
}

func wild_traveler(id int, symbol rune, report chan<- Trace_Sequence, board *Board) {
	defer wg.Done()

	time.Sleep(Wild_Traveler_Min_Delay + time.Duration(rand.Intn(int(Wild_Traveler_Max_Delay-Wild_Traveler_Min_Delay))))

	var initialPosition Position
	for {
		initialPosition = Position{rand.Intn(Board_Width), rand.Intn(Board_Height)}
		if !board.IsOccupied(initialPosition) {
			break
		}
	}

	traveler := Traveler{
		Id:           id,
		Symbol:       symbol,
		Position:     initialPosition,
		LastMoveTime: time.Now(),
		IsTrapped:    false,
		IsWild:       true,
		Traces:       make(Trace_Sequence, 0),
	}

	defer func() { report <- traveler.Traces }()

	board.Occupy(traveler.Position, &traveler)

	traveler.Traces = append(traveler.Traces, Trace{
		timestamp: time.Since(Start_Time),
		id:        traveler.Id,
		position:  traveler.Position,
		symbol:    traveler.Symbol,
	})

	time.Sleep(Wild_Traveler_Duration)

	if traveler.IsTrapped {
		return
	}

	traveler.Traces = append(traveler.Traces, Trace{
		timestamp: time.Since(Start_Time),
		id:        traveler.Id,
		position:  Position{X: Board_Width, Y: Board_Height},
		symbol:    traveler.Symbol,
	})

}

func trap_tile_traveler(id int, symbol rune, report chan<- Trace_Sequence, board *Board) {

	var initialPosition Position
	for {
		initialPosition = Position{rand.Intn(Board_Width), rand.Intn(Board_Height)}
		if !board.IsOccupied(initialPosition) {
			break
		}
	}

	traveler := Traveler{
		Id:           id,
		Symbol:       symbol,
		Position:     initialPosition,
		LastMoveTime: time.Now(),
		IsTrapped:    false,
		IsWild:       false,
		IsTrap:       true,
		Traces:       make(Trace_Sequence, 0),
	}

	defer func() { report <- traveler.Traces }()

	board.Occupy(traveler.Position, &traveler)

	traveler.Traces = append(traveler.Traces, Trace{
		timestamp: time.Since(Start_Time),
		id:        traveler.Id,
		position:  traveler.Position,
		symbol:    traveler.Symbol,
	})

	wg.Wait()
}

func main() {
	Total_Travelers := Nr_Of_Travelers + Nr_Of_Wild_Travelers + Nr_Of_Trap_Tiles
	fmt.Println("-1", Total_Travelers, " ", Board_Width, " ", Board_Height)

	board := NewBoard()

	Start_Time = time.Now()
	report := make(chan Trace_Sequence, Nr_Of_Travelers)
	done := make(chan bool)

	go printer(report, done)

	id := 0

	wg.Add(1)

	for range Nr_Of_Trap_Tiles {
		go trap_tile_traveler(id, rune('#'), report, board)
		id++
	}

	for i := range Nr_Of_Travelers {
		wg.Add(1)
		go traveler(id, rune('A'+i), report, board)
		id++
	}

	for i := range Nr_Of_Wild_Travelers {
		wg.Add(1)
		go wild_traveler(id, rune('0'+i), report, board)
		id++
	}

	wg.Done()

	wg.Wait()

	time.Sleep(10 * time.Millisecond)
	close(report)
	<-done
}
