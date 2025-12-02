package main

import (
	"errors"
	"flag"
	"fmt"
	"lzw/bitreadwrite"
	"math"
	"os"
)

var table = make(map[string]int)

func initTable() {
	for i := 0; i < 256; i++ {
		table[string([]byte{byte(i)})] = i
	}
}

func LZWEncode(br *bitreadwrite.BitReader, bw *bitreadwrite.BitWriter) ([]int, error) {
	initTable()
	codes := []int{}

	b, err := br.ReadBits(8)
	if err != nil {
		return []int{}, err
	}

	p := []byte{byte(b)}
	var count int = 256

	for {
		b, err = br.ReadBits(8)
		if err != nil {
			break
		}
		c := byte(b)

		pc := make([]byte, len(p)+1)
		copy(pc, p)
		pc[len(p)] = c

		if _, ok := table[string(pc)]; ok {
			p = pc
		} else {
			codes = append(codes, table[string(p)])
			table[string(pc)] = count
			count++
			p = []byte{c}
		}
	}
	codes = append(codes, table[string(p)])
	return codes, nil
}

var writeCoding = func(bw *bitreadwrite.BitWriter, n int) error {

	return errors.ErrUnsupported
}

func intToBits(n int) []bool {
	bits := []bool{}
	if n == 0 {
		bits = append(bits, false)
		return bits
	}
	for n > 0 {
		bits = append([]bool{n%2 == 1}, bits...)
		n /= 2
	}
	return bits
}

func writeEliasOmega(bw *bitreadwrite.BitWriter, n int) error {
	code := []bool{false}
	//fmt.Println("Writing ", n)
	for {
		if n == 1 {
			break
		}
		bits := intToBits(n)
		code = append(bits, code...)
		n = len(bits) - 1
	}

	bw.WriteBitsBools(code)
	//fmt.Println(code)
	return nil
}

func writeEliasGamma(bw *bitreadwrite.BitWriter, num int) error {
	bits := intToBits(num)
	n := len(bits)

	for range n - 1 {
		err := bw.WriteBitBool(false)
		if err != nil {
			return err
		}
	}

	bw.WriteBitsBools(bits)
	return nil
}

func writeEliasDelta(bw *bitreadwrite.BitWriter, num int) error {
	bits := intToBits(num)
	n := len(bits)

	if err := writeEliasGamma(bw, n); err != nil {
		return err
	}

	if n > 1 {
		if err := bw.WriteBitsBools(bits[1:]); err != nil {
			return err
		}
	}
	return nil
}

func writeFibonacci(bw *bitreadwrite.BitWriter, n int) error {

	fibs := []int{1, 2}
	for fibs[len(fibs)-1] <= n {
		fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2])
	}

	result := []bool{}
	i := len(fibs) - 1
	for i >= 0 {
		if fibs[i] <= n {
			n -= fibs[i]
			result = append([]bool{true}, result...)
		} else {
			if len(result) > 0 {
				result = append([]bool{false}, result...)
			}
		}
		i--
	}

	result = append(result, true) //terminating bit

	bw.WriteBitsBools(result)
	return nil
}

var encType string
var inputFile string
var outputFile string

func init() {
	flag.StringVar(&encType, "enc", "omega", "encoding type (e.g. omega, delta, gamma, fibonacci)")
	flag.StringVar(&inputFile, "in", "", "input filename (required)")
	flag.StringVar(&outputFile, "out", "", "output filename (required)")
	flag.Parse()
}

func main() {
	if inputFile == "" || outputFile == "" {
		fmt.Fprintln(os.Stderr, "usage: -in INPUT -out OUTPUT [-enc omega]")
		os.Exit(2)
	}

	br, fin, err := bitreadwrite.NewBitReaderFile(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to open input:", err)
		os.Exit(1)
	}
	bw, fout, err := bitreadwrite.NewBitWriterFile(outputFile)
	if err != nil {
		fin.Close()
		fmt.Fprintln(os.Stderr, "failed to open output:", err)
		os.Exit(1)
	}

	defer fin.Close()
	defer fout.Close()

	switch encType {
	case "omega":
		writeCoding = writeEliasOmega
	case "gamma":
		writeCoding = writeEliasGamma
	case "delta":
		writeCoding = writeEliasDelta
	case "fibonacci", "fib":
		writeCoding = writeFibonacci
	default:
		fmt.Printf("unknown encoding type: %s", encType)
		os.Exit(2)
	}

	codes, err := LZWEncode(br, bw)

	fmt.Println("Debug: ", codes[len(codes)-10:len(codes)-1])

	for _, code := range codes {
		writeCoding(bw, code+1) // ENSURE ENCODING POSITIVE INTEGER! VALUE OF CODE SHOULD BE LOWERED BY 1 ON DECODING
	}

	switch encType { // pad the last byte with different data depending on encoding used - '0' in elias omega encoding is 1
	case "omega":
		bw.WriteBits(int(8-bw.GetBitsFilled()), 0xFF)
	case "fibonacci", "fib", "delta", "gamma":
		bw.WriteBits(int(8-bw.GetBitsFilled()), 0)
	}

	bw.Flush()

	if err != nil {
		fmt.Fprintln(os.Stderr, "encoding error:", err)
		os.Exit(1)
	}

	// compute sizes
	if infoIn, err := os.Stat(inputFile); err == nil {
		infoOut, err2 := os.Stat(outputFile)
		if err2 == nil {
			origSize := infoIn.Size()
			encSize := infoOut.Size()
			var comp float64
			if origSize > 0 {
				comp = (1.0 - float64(encSize)/float64(origSize)) * 100.0
			}
			fmt.Printf("Original size: %d bytes\n", origSize)
			fmt.Printf("Encoded size:  %d bytes\n", encSize)
			fmt.Printf("Compression:   %.2f%%\n", comp)
		}
	}

	calcEntropy := func(data []byte) float64 {
		n := len(data)
		if n == 0 {
			return 0.0
		}
		var counts [256]int
		for _, b := range data {
			counts[int(b)]++
		}
		ln2 := math.Ln2
		entropy := 0.0
		for _, c := range counts {
			if c == 0 {
				continue
			}
			p := float64(c) / float64(n)
			lnp := math.Log(p)
			log2p := lnp / ln2
			entropy -= p * log2p
		}
		return entropy
	}

	if origData, err := os.ReadFile(inputFile); err == nil {
		if encData, err2 := os.ReadFile(outputFile); err2 == nil {
			entropyOrig := calcEntropy(origData)
			entropyEnc := calcEntropy(encData)
			fmt.Printf("Entropy (original): %.6f bits/symbol\n", entropyOrig)
			fmt.Printf("Entropy (encoded):  %.6f bits/symbol\n", entropyEnc)
		}
	}

}
