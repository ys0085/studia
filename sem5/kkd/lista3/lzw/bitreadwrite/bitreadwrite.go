package bitreadwrite

import (
	"errors"
	"io"
	"os"
)

type BitReader struct {
	r        io.Reader
	buf      byte  // current byte buffer
	bitsLeft uint8 // number of unread bits in buf (0..8)
	err      error
}

// NewBitReader creates a BitReader that reads from r.
func NewBitReader(r io.Reader) *BitReader {
	return &BitReader{r: r, bitsLeft: 0}
}

// NewBitReaderFile opens path and returns a BitReader and the opened *os.File.
// Caller should close the returned *os.File when finished.
func NewBitReaderFile(path string) (*BitReader, *os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	return NewBitReader(f), f, nil
}

// fillByte reads the next byte into the internal buffer.
func (br *BitReader) fillByte() error {
	if br.err != nil {
		return br.err
	}
	var b [1]byte
	n, err := br.r.Read(b[:])
	if n == 0 && err == nil {
		// treat as EOF
		err = io.EOF
	}
	if err != nil {
		br.err = err
		return err
	}
	br.buf = b[0]
	br.bitsLeft = 8
	return nil
}

// ReadBit reads a single bit and returns 0 or 1.
// Bits are returned MSB first (bit 7 down to bit 0).
func (br *BitReader) ReadBit() (int, error) {
	if br.bitsLeft == 0 {
		if err := br.fillByte(); err != nil {
			return 0, err
		}
	}
	bit := int((br.buf >> (br.bitsLeft - 1)) & 1)
	br.bitsLeft--
	if br.bitsLeft == 0 {
		br.buf = 0
	}
	return bit, nil
}

// ReadBits reads n bits (1..64) and returns them packed MSB-first into the low bits of a uint64.
// For example ReadBits(3) might return 0b101 -> 5.
func (br *BitReader) ReadBits(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("ReadBits: n must be between 1 and 64")
	}
	var v uint64
	for i := 0; i < n; i++ {
		b, err := br.ReadBit()
		if err != nil {
			return 0, err
		}
		v = (v << 1) | uint64(b)
	}

	return v, nil
}

// Align discards any remaining bits in the current byte so the next ReadByteAligned
// will start at the next byte boundary.
func (br *BitReader) Align() {
	br.bitsLeft = 0
	br.buf = 0
}

// ReadByteAligned reads a full byte from the underlying reader, but first aligns to the next byte boundary.
func (br *BitReader) ReadByteAligned() (byte, error) {
	br.Align()
	var b [1]byte
	n, err := br.r.Read(b[:])
	if n == 0 && err == nil {
		err = io.EOF
	}
	return b[0], err
}

type BitWriter struct {
	w          io.Writer
	buf        byte  // current byte being filled (bits written from MSB down)
	bitsFilled uint8 // number of bits written into buf (0..8)
	err        error
}

func (bw *BitWriter) GetBitsFilled() uint8 {
	return bw.bitsFilled
}

// NewBitWriter creates a BitWriter that writes to w.
func NewBitWriter(w io.Writer) *BitWriter {
	return &BitWriter{w: w}
}

// NewBitWriterFile creates/opens path for writing and returns a BitWriter and the opened *os.File.
// Caller should close the returned *os.File when finished.
func NewBitWriterFile(path string) (*BitWriter, *os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, nil, err
	}
	return NewBitWriter(f), f, nil
}

func (bw *BitWriter) WriteBitBool(b bool) error {
	if b {
		return bw.WriteBit(1)
	}
	return bw.WriteBit(0)
}

func (bw *BitWriter) WriteBitsBools(bools []bool) error {
	for _, b := range bools {
		if err := bw.WriteBitBool(b); err != nil {
			return err
		}
	}
	return nil
}

func (bw *BitWriter) WriteBitsByteArray(bytes []byte) error {
	for _, b := range bytes {
		if err := bw.WriteBits(8, uint64(b)); err != nil {
			return err
		}
	}
	return nil
}

// WriteBit writes a single bit (0 or 1). Bits are written MSB first into bytes.
func (bw *BitWriter) WriteBit(bit int) error {
	if bit != 0 && bit != 1 {
		return errors.New("WriteBit: bit must be 0 or 1")
	}
	if bw.err != nil {
		return bw.err
	}
	// place bit at position (7 - bitsFilled)
	if bit == 1 {
		bw.buf |= 1 << (7 - bw.bitsFilled)
	}
	bw.bitsFilled++
	if bw.bitsFilled == 8 {
		n, err := bw.w.Write([]byte{bw.buf})
		if err == nil && n != 1 {
			err = io.ErrShortWrite
		}
		if err != nil {
			bw.err = err
			return err
		}
		bw.buf = 0
		bw.bitsFilled = 0
	}
	return nil
}

// WriteBits writes n bits (1..64) from v. Bits are taken MSB-first from the low n bits of v.
// For example WriteBits(3, 0b101) will write bits 1,0,1 in that order.
func (bw *BitWriter) WriteBits(n int, v uint64) error {
	if n <= 0 || n > 64 {
		return errors.New("WriteBits: n must be between 1 and 64")
	}
	for i := n - 1; i >= 0; i-- {
		bit := int((v >> uint(i)) & 1)
		if err := bw.WriteBit(bit); err != nil {
			return err
		}
	}
	return nil
}

// Align pads the current partial byte with zeros and writes it so subsequent writes are byte-aligned.
func (bw *BitWriter) Align() error {
	if bw.err != nil {
		return bw.err
	}
	if bw.bitsFilled == 0 {
		return nil
	}
	n, err := bw.w.Write([]byte{bw.buf})
	if err == nil && n != 1 {
		err = io.ErrShortWrite
	}
	if err != nil {
		bw.err = err
		return err
	}
	bw.buf = 0
	bw.bitsFilled = 0
	return nil
}

// Flush is an alias for Align.
func (bw *BitWriter) Flush() error {
	return bw.Align()
}

// WriteByteAligned aligns first (padding current byte with zeros if needed) then writes the full byte.
func (bw *BitWriter) WriteByteAligned(b byte) error {
	if err := bw.Align(); err != nil {
		return err
	}
	n, err := bw.w.Write([]byte{b})
	if err == nil && n != 1 {
		err = io.ErrShortWrite
	}
	if err != nil {
		bw.err = err
		return err
	}
	return nil
}
