package main

import (
	"flag"
	"fmt"
	"lzw/bitreadwrite"
	"os"
)

var table = make(map[int][]byte)

func initTable() {
	for i := 0; i < 256; i++ {
		table[i] = []byte{byte(i)}
	}
}

func LZWDecode(codes []int, bw *bitreadwrite.BitWriter) error {
	initTable()

	if len(codes) == 0 {
		return nil
	}

	old := codes[0]
	s := make([]byte, len(table[old]))
	copy(s, table[old])
	bw.WriteBitsByteArray(s)

	count := 256

	for i := 1; i < len(codes); i++ {
		new := codes[i]
		var entry []byte

		if val, ok := table[new]; ok {
			entry = make([]byte, len(val))
			copy(entry, val)
		} else if new == count {
			entry = make([]byte, len(table[old])+1)
			copy(entry, table[old])
			entry[len(table[old])] = table[old][0]
		} else {
			return fmt.Errorf("invalid code: %d", new)
		}

		bw.WriteBitsByteArray(entry)

		newEntry := make([]byte, len(table[old])+1)
		copy(newEntry, table[old])
		newEntry[len(table[old])] = entry[0]
		table[count] = newEntry
		count++

		old = new
	}

	return nil
}

func readFullElias(br *bitreadwrite.BitReader) ([]int, error) {
	codes := []int{}
	for {
		code, err := readCoding(br)
		//fmt.Printf("%d ", code)
		if err != nil {
			break
		}
		codes = append(codes, code-1)
	}
	return codes, nil
}

var readCoding = func(br *bitreadwrite.BitReader) (int, error) {
	fmt.Println("Oops")
	return -1, nil
}

func readEliasOmega(br *bitreadwrite.BitReader) (int, error) {
	n := 1
	for {
		bit, err := br.ReadBit()
		if err != nil {
			return -1, err
		}
		if bit == 0 {
			break
		}
		b, err := br.ReadBits(n)
		if err != nil {
			return -1, err
		}

		n = int(b + (1 << n))
	}
	return n, nil
}

func readEliasGamma(br *bitreadwrite.BitReader) (int, error) {
	zeroes := 0
	for {
		bit, err := br.ReadBit()
		if err != nil {
			return -1, err
		}
		if bit == 1 {
			break
		}
		zeroes++
	}
	if zeroes == 0 {
		return 1, nil
	}
	rest, err := br.ReadBits(zeroes)
	if err != nil {
		return -1, err
	}
	return int(rest + (1 << zeroes)), nil
}

func readEliasDelta(br *bitreadwrite.BitReader) (int, error) {
	k, err := readEliasGamma(br)
	if err != nil {
		return -1, err
	}
	if k == 1 {
		return 1, nil
	}

	rem, err := br.ReadBits(k - 1)
	if err != nil {
		return -1, err
	}
	return int((1 << (k - 1)) + rem), nil
}

func readFibonacci(br *bitreadwrite.BitReader) (int, error) {
	bits := []bool{}
	for {
		bit, err := br.ReadBit()
		if err != nil {
			return -1, err
		}
		if len(bits) > 0 {
			if bit == 1 && bits[len(bits)-1] {
				break
			}
		}
		bits = append(bits, bit == 1)
	}
	if len(bits) == 0 {
		return 0, nil
	}
	fibs := []int{1, 2}
	for range len(bits) {
		fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2])
	}
	result := 0
	for i, b := range bits {
		if b {
			result += fibs[i]
		}
	}
	return result, nil
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
		readCoding = readEliasOmega
	case "gamma":
		readCoding = readEliasGamma
	case "delta":
		readCoding = readEliasDelta
	case "fibonacci", "fib":
		readCoding = readFibonacci
	default:
		break
	}

	codes, _ := readFullElias(br)
	fmt.Println(codes)
	err = LZWDecode(codes, bw)
	if err != nil {
		fmt.Fprintln(os.Stderr, "encoding error:", err)
		os.Exit(1)
	}

}
