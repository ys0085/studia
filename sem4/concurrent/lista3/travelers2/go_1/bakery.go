package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Configuration constants
const (
	NrOfProcesses = 15
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

// Protected maximum ticket counter
type MaxTicketProtected struct {
	mu        sync.RWMutex
	maxTicket int
}

func (mtp *MaxTicketProtected) updateMax(value int) {
	mtp.mu.Lock()
	defer mtp.mu.Unlock()
	if value > mtp.maxTicket {
		mtp.maxTicket = value
	}
}

func (mtp *MaxTicketProtected) getMax() int {
	mtp.mu.RLock()
	defer mtp.mu.RUnlock()
	return mtp.maxTicket
}

// Global variables
var (
	startTime          = time.Now()
	maxTicketProtected = &MaxTicketProtected{}
	choosing           = make([]bool, NrOfProcesses)
	number             = make([]int, NrOfProcesses)
)

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
		fmt.Printf("MAX_TICKET=%d;\n", maxTicketProtected.getMax())
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

func (p *Process) run(printer *Printer, wg *sync.WaitGroup) {
	defer wg.Done()

	// Calculate number of steps
	nrOfSteps := MinSteps + int(float64(MaxSteps-MinSteps)*p.rng.Float64())

	// Store initial trace
	p.storeTrace()

	for range nrOfSteps / 4 {
		delay := MinDelay + time.Duration(float64(MaxDelay-MinDelay)*p.rng.Float64())
		time.Sleep(delay)

		p.changeState(EntryProtocol)

		choosing[p.ID] = true
		max := 0
		for i := range NrOfProcesses {
			if number[i] > max {
				max = number[i]
			}
		}

		number[p.ID] = max + 1

		if number[p.ID] > p.maxTicket {
			p.maxTicket = number[p.ID]
		}

		choosing[p.ID] = false

		for j := range NrOfProcesses {
			if j == p.ID {
				continue
			}
			// Wait if process j is choosing its number
			for choosing[j] {
				time.Sleep(time.Millisecond)
			}
			// Wait if process j has a ticket and (its ticket is lower, or same ticket but lower ID)
			for number[j] != 0 && (number[j] < number[p.ID] || (number[j] == number[p.ID] && j < p.ID)) {
				time.Sleep(time.Millisecond)
			}
		}

		p.changeState(CriticalSection)

		// CRITICAL_SECTION - start
		delay = MinDelay + time.Duration(float64(MaxDelay-MinDelay)*p.rng.Float64())
		time.Sleep(delay)
		// CRITICAL_SECTION - end

		p.changeState(ExitProtocol)

		number[p.ID] = 0
		time.Sleep(time.Millisecond) // ensure delay in traces

		p.changeState(LocalSection)
	}

	maxTicketProtected.updateMax(p.maxTicket)
	printer.report(p.traces)
}

func main() {
	printer := NewPrinter()
	printer.start()

	processes := make([]*Process, NrOfProcesses)
	symbol := 'A'

	for i := range NrOfProcesses {
		processes[i] = NewProcess(i, symbol)
		symbol++
	}

	var wg sync.WaitGroup
	wg.Add(NrOfProcesses)

	for i := range NrOfProcesses {
		go processes[i].run(printer, &wg)
	}

	wg.Wait()

	printer.wait()
}
