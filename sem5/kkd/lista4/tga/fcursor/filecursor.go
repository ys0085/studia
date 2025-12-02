package fcursor

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

// FileCursor is a small wrapper around *os.File with helpers for reading
// bytes, seeking, skipping and reading LE/BE integers.
type FileCursor struct {
	f *os.File
}

// NewFileCursor opens path and returns a FileCursor.
func NewFileCursor(path string) (*FileCursor, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &FileCursor{f: f}, nil
}

func NewFileCursorFromFile(f *os.File) *FileCursor {
	return &FileCursor{f: f}
}

// Close closes the underlying file.
func (r *FileCursor) Close() error {
	if r.f == nil {
		return nil
	}
	err := r.f.Close()
	r.f = nil
	return err
}

// Seek sets the offset for the next Read. whence uses io.SeekStart/Current/End.
func (r *FileCursor) Seek(offset int64, whence int) (int64, error) {
	return r.f.Seek(offset, whence)
}

// Tell returns the current offset in the file.
func (r *FileCursor) Tell() (int64, error) {
	return r.f.Seek(0, io.SeekCurrent)
}

// Read reads exactly n bytes; returns io.EOF or io.ErrUnexpectedEOF on failure.
func (r *FileCursor) Read(n int) ([]byte, error) {
	if n < 0 {
		return nil, errors.New("negative read length")
	}
	buf := make([]byte, n)
	_, err := io.ReadFull(r.f, buf)
	return buf, err
}

// ReadByte reads a single byte.
func (r *FileCursor) ReadByte() (byte, error) {
	var b [1]byte
	_, err := io.ReadFull(r.f, b[:])
	return b[0], err
}

// ReadUint8 is an alias for ReadByte.
func (r *FileCursor) ReadUint8() (uint8, error) {
	b, err := r.ReadByte()
	return uint8(b), err
}

// ReadUint16LE reads a little-endian uint16.
func (r *FileCursor) ReadUint16LE() (uint16, error) {
	var v uint16
	err := binary.Read(r.f, binary.LittleEndian, &v)
	return v, err
}

// ReadUint16BE reads a big-endian uint16.
func (r *FileCursor) ReadUint16BE() (uint16, error) {
	var v uint16
	err := binary.Read(r.f, binary.BigEndian, &v)
	return v, err
}

// ReadUint32LE reads a little-endian uint32.
func (r *FileCursor) ReadUint32LE() (uint32, error) {
	var v uint32
	err := binary.Read(r.f, binary.LittleEndian, &v)
	return v, err
}

// ReadUint32BE reads a big-endian uint32.
func (r *FileCursor) ReadUint32BE() (uint32, error) {
	var v uint32
	err := binary.Read(r.f, binary.BigEndian, &v)
	return v, err
}

// ReadAt reads exactly n bytes at a specific offset without changing cursor.
func (r *FileCursor) ReadAt(offset int64, n int) ([]byte, error) {
	if n < 0 {
		return nil, errors.New("negative read length")
	}
	buf := make([]byte, n)
	_, err := r.f.ReadAt(buf, offset)
	return buf, err
}

// Skip advances the file position by n bytes (positive or negative allowed).
// Equivalent to Seek(n, io.SeekCurrent).
func (r *FileCursor) Skip(n int64) (int64, error) {
	return r.Seek(n, io.SeekCurrent)
}

// ReadString reads exactly n bytes and returns them as a string.
func (r *FileCursor) ReadString(n int) (string, error) {
	b, err := r.Read(n)
	return string(b), err
}

// ReadNullTerminatedString reads bytes until a 0x00 is found or maxLen bytes.
// If maxLen is 0, it will read until 0x00 or EOF.
func (r *FileCursor) ReadNullTerminatedString(maxLen int) (string, error) {
	var buf []byte
	for {
		if maxLen > 0 && len(buf) >= maxLen {
			return string(buf), nil
		}
		b, err := r.ReadByte()
		if err != nil {
			// return what we have and the error
			if len(buf) == 0 {
				return "", err
			}
			return string(buf), err
		}
		if b == 0x00 {
			return string(buf), nil
		}
		buf = append(buf, b)
	}
}

// ReadToEnd reads and returns the remaining file bytes from current position.
func (r *FileCursor) ReadToEnd() ([]byte, error) {
	return io.ReadAll(r.f)
}
