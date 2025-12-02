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
	EntryProtocol1
	EntryProtocol2
	EntryProtocol3
	EntryProtocol4
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

var (
	flags     = make([]int, NrOfProcesses) // 0..4
	flagsLock sync.Mutex
)

// Global variables
var (
	startTime = time.Now()
	choosing  = make([]bool, NrOfProcesses)
	number    = make([]int, NrOfProcesses)
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
		fmt.Printf("LOCAL_SECTION;ENTRY_PROTOCOL_1;ENTRY_PROTOCOL_2;ENTRY_PROTOCOL_3;ENTRY_PROTOCOL_4;CRITICAL_SECTION;EXIT_PROTOCOL;")
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

// ...existing code...

func (p *Process) run(printer *Printer, wg *sync.WaitGroup) {
	defer wg.Done()

	nrOfSteps := MinSteps + int(float64(MaxSteps-MinSteps)*p.rng.Float64())
	p.storeTrace()

	for step := 0; step < nrOfSteps/4; step++ {
		// LOCAL SECTION
		delay := MinDelay + time.Duration(float64(MaxDelay-MinDelay)*p.rng.Float64())
		time.Sleep(delay)
		p.changeState(LocalSection)

		flagsLock.Lock()
		flags[p.ID] = 1
		flagsLock.Unlock()
		p.changeState(EntryProtocol1)

		for {
			flagsLock.Lock()
			ok := true
			for i := 0; i < NrOfProcesses; i++ {
				if flags[i] != 0 && flags[i] != 1 && flags[i] != 2 {
					ok = false
					break
				}
			}
			flagsLock.Unlock()
			if ok {
				break
			}
			time.Sleep(1 * time.Millisecond)
		}

		flagsLock.Lock()
		flags[p.ID] = 3
		flagsLock.Unlock()
		p.changeState(EntryProtocol3)

		need2 := false
		flagsLock.Lock()
		for i := 0; i < NrOfProcesses; i++ {
			if i != p.ID && flags[i] == 1 {
				need2 = true
				break
			}
		}
		flagsLock.Unlock()
		if need2 {
			flagsLock.Lock()
			flags[p.ID] = 2
			flagsLock.Unlock()
			p.changeState(EntryProtocol2)

			for {
				flagsLock.Lock()
				found4 := false
				for i := 0; i < NrOfProcesses; i++ {
					if flags[i] == 4 {
						found4 = true
						break
					}
				}
				flagsLock.Unlock()
				if found4 {
					break
				}
				time.Sleep(1 * time.Millisecond)
			}
		}

		flagsLock.Lock()
		flags[p.ID] = 4
		flagsLock.Unlock()
		p.changeState(EntryProtocol4)

		for {
			flagsLock.Lock()
			ok := true
			for i := 0; i < p.ID; i++ {
				if flags[i] != 0 && flags[i] != 1 {
					ok = false
					break
				}
			}
			flagsLock.Unlock()
			if ok {
				break
			}
			time.Sleep(1 * time.Millisecond)
		}

		// --- CRITICAL SECTION ---
		p.changeState(CriticalSection)
		delay = MinDelay + time.Duration(float64(MaxDelay-MinDelay)*p.rng.Float64())
		time.Sleep(delay)

		// --- EXIT protocol ---
		for {
			flagsLock.Lock()
			ok := true
			for i := p.ID + 1; i < NrOfProcesses; i++ {
				if flags[i] != 0 && flags[i] != 1 && flags[i] != 4 {
					ok = false
					break
				}
			}
			flagsLock.Unlock()
			if ok {
				break
			}
			time.Sleep(1 * time.Millisecond)
		}

		flagsLock.Lock()
		flags[p.ID] = 0
		flagsLock.Unlock()
		p.changeState(ExitProtocol)
	}

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
