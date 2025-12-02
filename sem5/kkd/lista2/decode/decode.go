package main

import (
	"fmt"
	"os"
)

const (
	CodeValueBits = 32
	MaxCode       = (1 << CodeValueBits) - 1
	FirstQtr      = MaxCode/4 + 1
	Half          = 2 * FirstQtr
	ThirdQtr      = 3 * FirstQtr
	EOFSymbol     = 256
	NumSymbols    = 257
	MaxFreq       = 16383
)

type ArithmeticDecoder struct {
	freq    [NumSymbols]int
	cumFreq [NumSymbols + 1]int
	low     int
	high    int
	value   int
	input   []byte
	bytePos int
	bitPos  int
}

func NewArithmeticDecoder(data []byte) *ArithmeticDecoder {
	dec := &ArithmeticDecoder{
		low:     0,
		high:    MaxCode,
		input:   data,
		bytePos: 0,
		bitPos:  7,
	}

	for i := 0; i < NumSymbols; i++ {
		dec.freq[i] = 1
	}
	dec.updateCumFreq()

	dec.value = 0
	for i := 0; i < CodeValueBits; i++ {
		dec.value = (dec.value << 1) | dec.inputBit()
	}

	return dec
}

func (dec *ArithmeticDecoder) updateCumFreq() {
	dec.cumFreq[0] = 0
	for i := 0; i < NumSymbols; i++ {
		dec.cumFreq[i+1] = dec.cumFreq[i] + dec.freq[i]
	}
}

func (dec *ArithmeticDecoder) updateModel(symbol int) {
	dec.freq[symbol]++

	if dec.cumFreq[NumSymbols] >= MaxFreq {
		for i := 0; i < NumSymbols; i++ {
			dec.freq[i] = (dec.freq[i] + 1) / 2
		}
	}
	dec.updateCumFreq()
}

func (dec *ArithmeticDecoder) inputBit() int {
	if dec.bytePos >= len(dec.input) {
		return 0
	}

	bit := 0
	if (dec.input[dec.bytePos] & (1 << dec.bitPos)) != 0 {
		bit = 1
	}

	dec.bitPos--
	if dec.bitPos < 0 {
		dec.bytePos++
		dec.bitPos = 7
	}

	return bit
}

func (dec *ArithmeticDecoder) decodeSymbol() int {
	totalFreq := dec.cumFreq[NumSymbols]
	rangeSize := dec.high - dec.low + 1

	cumTarget := ((dec.value-dec.low+1)*totalFreq - 1) / rangeSize

	symbol := 0
	for symbol = 0; symbol < NumSymbols-1; symbol++ {
		if dec.cumFreq[symbol+1] > cumTarget {
			break
		}
	}

	// Aktualizuj zakres
	dec.high = dec.low + (rangeSize*dec.cumFreq[symbol+1])/totalFreq - 1
	dec.low = dec.low + (rangeSize*dec.cumFreq[symbol])/totalFreq

	// Skalowanie
	for {
		if dec.high < Half {
			// Nic nie rób
		} else if dec.low >= Half {
			dec.value -= Half
			dec.low -= Half
			dec.high -= Half
		} else if dec.low >= FirstQtr && dec.high < ThirdQtr {
			dec.value -= FirstQtr
			dec.low -= FirstQtr
			dec.high -= FirstQtr
		} else {
			break
		}
		dec.low = 2 * dec.low
		dec.high = 2*dec.high + 1
		dec.value = (dec.value << 1) | dec.inputBit()
	}

	dec.updateModel(symbol)

	return symbol
}

func (dec *ArithmeticDecoder) decode() []byte {
	var output []byte

	for {
		symbol := dec.decodeSymbol()
		if symbol == EOFSymbol {
			break
		}
		output = append(output, byte(symbol))
	}

	return output
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Użycie: %s <plik_wejściowy> <plik_wyjściowy>\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Wczytaj skompresowane dane
	compressed, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Błąd odczytu pliku: %v\n", err)
		os.Exit(1)
	}

	if len(compressed) == 0 {
		fmt.Println("Plik wejściowy jest pusty")
		os.WriteFile(outputFile, []byte{}, 0644)
		return
	}

	// Dekodowanie
	decoder := NewArithmeticDecoder(compressed)
	decompressed := decoder.decode()

	// Zapisz wynik
	err = os.WriteFile(outputFile, decompressed, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Błąd zapisu pliku: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Dekompresja zakończona: %d bajtów\n", len(decompressed))
}
