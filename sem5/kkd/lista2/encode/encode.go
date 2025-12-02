package main

import (
	"fmt"
	"math"
	"os"
)

const (
	CodeValueBits = 32 //zmiana
	MaxCode       = (1 << CodeValueBits) - 1
	FirstQtr      = MaxCode/4 + 1
	Half          = 2 * FirstQtr
	ThirdQtr      = 3 * FirstQtr
	EOFSymbol     = 256
	NumSymbols    = 257
	MaxFreq       = 16383 // Zapobiega przepełnieniu
)

type ArithmeticEncoder struct {
	freq         [NumSymbols]int
	cumFreq      [NumSymbols + 1]int
	low          int
	high         int
	bitsToFollow int
	buffer       byte
	bitPos       int
	output       []byte
}

func NewArithmeticEncoder() *ArithmeticEncoder {
	enc := &ArithmeticEncoder{
		low:    0,
		high:   MaxCode,
		bitPos: 7,
	}
	// Inicjalizacja modelu - wszystkie symbole z częstością 1
	for i := 0; i < NumSymbols; i++ {
		enc.freq[i] = 1
	}
	enc.updateCumFreq()
	return enc
}

func (enc *ArithmeticEncoder) updateCumFreq() {
	enc.cumFreq[0] = 0
	for i := 0; i < NumSymbols; i++ {
		enc.cumFreq[i+1] = enc.cumFreq[i] + enc.freq[i]
	}
}

func (enc *ArithmeticEncoder) updateModel(symbol int) {
	enc.freq[symbol]++
	if enc.cumFreq[NumSymbols] >= MaxFreq {
		for i := 0; i < NumSymbols; i++ {
			enc.freq[i] = (enc.freq[i] + 1) / 2
		}
	}
	enc.updateCumFreq()
}

func (enc *ArithmeticEncoder) outputBit(bit int) {
	if bit != 0 {
		enc.buffer |= (1 << enc.bitPos)
	}
	enc.bitPos--

	if enc.bitPos < 0 {
		enc.output = append(enc.output, enc.buffer)
		enc.buffer = 0
		enc.bitPos = 7
	}
}

func (enc *ArithmeticEncoder) bitPlusFollow(bit int) {
	enc.outputBit(bit)
	for enc.bitsToFollow > 0 {
		enc.outputBit(1 - bit)
		enc.bitsToFollow--
	}
}

func (enc *ArithmeticEncoder) encodeSymbol(symbol int) {
	totalFreq := enc.cumFreq[NumSymbols]
	rangeSize := enc.high - enc.low + 1

	// Oblicz nowy zakres
	enc.high = enc.low + (rangeSize*enc.cumFreq[symbol+1])/totalFreq - 1
	enc.low = enc.low + (rangeSize*enc.cumFreq[symbol])/totalFreq

	// Skalowanie
	for {
		if enc.high < Half {
			// Lewy zakres [0, Half)
			enc.bitPlusFollow(0)
		} else if enc.low >= Half {
			// Prawy zakres [Half, MaxCode]
			enc.bitPlusFollow(1)
			enc.low -= Half
			enc.high -= Half
		} else if enc.low >= FirstQtr && enc.high < ThirdQtr {
			// Środkowy zakres [FirstQtr, ThirdQtr)
			enc.bitsToFollow++
			enc.low -= FirstQtr
			enc.high -= FirstQtr
		} else {
			break
		}
		enc.low = 2 * enc.low
		enc.high = 2*enc.high + 1
	}

	enc.updateModel(symbol)
}

func (enc *ArithmeticEncoder) finish() {
	enc.bitsToFollow++
	if enc.low < FirstQtr {
		enc.bitPlusFollow(0)
	} else {
		enc.bitPlusFollow(1)
	}

	// Dopełnij ostatni bajt
	if enc.bitPos < 7 {
		enc.output = append(enc.output, enc.buffer)
	}
}

func (enc *ArithmeticEncoder) encode(data []byte) []byte {
	for _, b := range data {
		enc.encodeSymbol(int(b))
	}
	enc.encodeSymbol(EOFSymbol)
	enc.finish()
	return enc.output
}

func calculateEntropy(data []byte) float64 {
	if len(data) == 0 {
		return 0
	}

	freq := make(map[byte]int)
	for _, b := range data {
		freq[b]++
	}

	entropy := 0.0
	total := float64(len(data))
	for _, count := range freq {
		if count > 0 {
			p := float64(count) / total
			entropy -= p * math.Log2(p)
		}
	}
	return entropy
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Użycie: %s <plik_wejściowy> <plik_wyjściowy>\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Wczytaj dane
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Błąd odczytu pliku: %v\n", err)
		os.Exit(1)
	}

	if len(data) == 0 {
		fmt.Println("Plik wejściowy jest pusty")
		os.WriteFile(outputFile, []byte{}, 0644)
		return
	}

	// Kodowanie
	encoder := NewArithmeticEncoder()
	compressed := encoder.encode(data)

	// Zapisz wynik
	err = os.WriteFile(outputFile, compressed, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Błąd zapisu pliku: %v\n", err)
		os.Exit(1)
	}

	// Statystyki
	originalSize := len(data)
	compressedSize := len(compressed)

	entropy := calculateEntropy(data)
	avgCodeLength := float64(compressedSize*8) / float64(originalSize)
	compressionRatio := float64(originalSize) / float64(compressedSize)

	fmt.Printf("Entropia: %.4f bitów/symbol\n", entropy)
	fmt.Printf("Średnia długość kodowania: %.4f bitów/symbol\n", avgCodeLength)
	fmt.Printf("Stopień kompresji: %.4f\n", compressionRatio)
	fmt.Printf("Rozmiar oryginalny: %d bajtów\n", originalSize)
	fmt.Printf("Rozmiar skompresowany: %d bajtów\n", compressedSize)
}
