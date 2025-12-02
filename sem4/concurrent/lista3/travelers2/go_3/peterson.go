package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Configuration constants
const (
	NrOfProcesses = 2
	MinSteps      = 50
	MaxSteps      = 100
	MinDelay      = 10 * time.Millisecond
	MaxDelay      = 50 * time.Millisecond
)

// Process states
type ProcessState int

const (
	LocalSection ProcessState = iota
	EntryProtocol
	CriticalSection
	ExitProtocol
)

func (ps ProcessState) String() string {
	switch ps {
	case LocalSection:
		return "LOCAL_SECTION"
	case EntryProtocol:
		return "ENTRY_PROTOCOL"
	case CriticalSection:
		return "CRITICAL_SECTION"
	case ExitProtocol:
		return "EXIT_PROTOCOL"
	default:
		return "UNKNOWN"
	}
}

// Board dimensions
const (
	BoardWidth  = NrOfProcesses
	BoardHeight = int(ExitProtocol) + 1
)

// Position on the board
type Position struct {
	X int
	Y int
}

// Trace entry
type Trace struct {
	TimeStamp time.Duration
	ID        int
	Position  Position
	Symbol    rune
}

// Trace sequence
type TraceSequence struct {
	Traces []Trace
}

func (ts *TraceSequence) addTrace(trace Trace) {
	ts.Traces = append(ts.Traces, trace)
}

func (ts *TraceSequence) printTraces() {
	for _, trace := range ts.Traces {
		fmt.Printf("%.6f %d %d %d %c\n",
			trace.TimeStamp.Seconds(),
			trace.ID,
			trace.Position.X,
			trace.Position.Y,
			trace.Symbol)
	}
}

// Global variables

var startTime = time.Now()

// Printer goroutine
type Printer struct {
	traces chan TraceSequence
	done   chan struct{}
}

func NewPrinter() *Printer {
	return &Printer{
		traces: make(chan TraceSequence, NrOfProcesses),
		done:   make(chan struct{}),
	}
}

func (p *Printer) start() {
	go func() {
		defer close(p.done)

		for i := 0; i < NrOfProcesses; i++ {
			traces := <-p.traces
			traces.printTraces()
		}

		fmt.Printf("-1 %d %d %d ", NrOfProcesses, BoardWidth, BoardHeight)
		fmt.Printf("LOCAL_SECTION;ENTRY_PROTOCOL;CRITICAL_SECTION;EXIT_PROTOCOL;")
		fmt.Printf("EXTRA_LABEL;")
	}()
}

func (p *Printer) report(traces TraceSequence) {
	p.traces <- traces
}

func (p *Printer) wait() {
	<-p.done
}

// Process structure
type Process struct {
	ID        int
	Symbol    rune
	Position  Position
	rng       *rand.Rand
	traces    TraceSequence
	maxTicket int
}

func NewProcess(id int, symbol rune) *Process {
	return &Process{
		ID:     id,
		Symbol: symbol,
		Position: Position{
			X: id,
			Y: int(LocalSection),
		},
		rng:       rand.New(rand.NewSource(time.Now().UnixNano() + int64(id))),
		maxTicket: 0,
	}
}

func (p *Process) storeTrace() {
	timeStamp := time.Since(startTime)
	trace := Trace{
		TimeStamp: timeStamp,
		ID:        p.ID,
		Position:  p.Position,
		Symbol:    p.Symbol,
	}
	p.traces.addTrace(trace)
}

func (p *Process) changeState(state ProcessState) {
	p.Position.Y = int(state)
	p.storeTrace()
}

var (
	flag = make([]bool, 2)
	turn = 0
	once sync.Once
)

func (p *Process) run(printer *Printer, wg *sync.WaitGroup) {
	defer wg.Done()

	nrOfSteps := MinSteps + int(float64(MaxSteps-MinSteps)*p.rng.Float64())
	p.storeTrace()

	for range nrOfSteps / 4 {
		delay := MinDelay + time.Duration(float64(MaxDelay-MinDelay)*p.rng.Float64())
		time.Sleep(delay)

		p.changeState(EntryProtocol)

		other := 1 - p.ID
		flag[p.ID] = true
		turn = other
		for flag[other] && turn == other {
			time.Sleep(time.Millisecond)
		}

		p.changeState(CriticalSection)

		delay = MinDelay + time.Duration(float64(MaxDelay-MinDelay)*p.rng.Float64())
		time.Sleep(delay)

		p.changeState(ExitProtocol)

		flag[p.ID] = false

		time.Sleep(time.Millisecond)

		p.changeState(LocalSection)
	}

	printer.report(p.traces)
}

func main() {
	printer := NewPrinter()
	printer.start()

	processes := make([]*Process, NrOfProcesses)
	symbol := 'A'

	flag[0] = false
	flag[1] = false
	turn = 0

	for i := 0; i < NrOfProcesses; i++ {
		processes[i] = NewProcess(i, symbol)
		symbol++
	}

	var wg sync.WaitGroup
	wg.Add(NrOfProcesses)

	for i := 0; i < NrOfProcesses; i++ {
		go processes[i].run(printer, &wg)
	}

	wg.Wait()

	printer.wait()
}
