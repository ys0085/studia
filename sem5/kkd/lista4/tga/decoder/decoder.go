package decoder

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
	"tga/fcursor"
)

/*

====================================================================================================
										TGA FILE FORMAT
====================================================================================================
01. 	offset=0  		| bytes=1 		| ID Length
02. 	offset=1  		| bytes=1 		| Color Map Type
03. 	offset=2  		| bytes=1 		| Image Type
04. 	offset=3  		| bytes=5 		| Color Map Specification
|-4.1. 	offset=3  		| bytes=2 		| First Entry Index
|-4.2. 	offset=5  		| bytes=2 		| Color Map Length
'-4.3. 	offset=7  		| bytes=1 		| Color Map Entry Size
05. 	offset=8  		| bytes=10 		| Image Specification
|-5.1. 	offset=8  		| bytes=2 		| X-origin
|-5.2. 	offset=10 		| bytes=2 		| Y-origin
|-5.3. 	offset=12 		| bytes=2 		| Image Width
|-5.4. 	offset=14 		| bytes=2 		| Image Height
|-5.5. 	offset=16 		| bytes=1 		| Pixel Depth
'-5.6. 	offset=17 		| bytes=1 		| Image Descriptor
========================================== Image Data ==============================================
06. 	offset=18 		| bytes=def 01	| Image ID
07. 	offset=?  		| bytes=def 04	| Color Map Data
08. 	offset=?  		| bytes=def 05	| Image Data
============================================ Footer ================================================
F01. 	offset=0 		| bytes=4 		| Extension Area Offset (assumed 0 in this exercise)
F02. 	offset=4 		| bytes=4 		| Developer Directory Offset (assumed 0 in this exercise)
F03. 	offset=8 		| bytes=16		| Signature ("TRUEVISION-XFILE")
F04. 	offset=24 		| bytes=1		| Dot (.)
F05. 	offset=25 		| bytes=1		| Null (0x00)
====================================================================================================


*/

var errFileFormatError = errors.New("wrong file format")

type TGAHeader struct {
	IDLength          uint8
	ColorMapType      uint8
	ImageType         uint8
	FirstEntryIndex   uint16
	ColorMapLength    uint16
	ColorMapEntrySize uint8
	XOrigin           uint16
	YOrigin           uint16
	ImageWidth        uint16
	ImageHeight       uint16
	PixelDepth        uint8
	ImageDescriptor   byte
}

// checkFooter reads and validates the TGA footer. Moves the cursor back to the original position after checking.
func checkFooter(fc *fcursor.FileCursor) error {
	startPos, _ := fc.Tell()
	fc.Seek(26, io.SeekEnd)
	defer fc.Seek(startPos, io.SeekStart)

	ExtAreaOffsetBytes, err := fc.Read(4)
	if err != nil {
		return err
	}
	DevDirOffsetBytes, err := fc.Read(4)
	if err != nil {
		return err
	}

	// Ensure that both offsets are zero, external areas and developer directories are not supported in this exercise
	if string(ExtAreaOffsetBytes) != "\x00\x00\x00\x00" || string(DevDirOffsetBytes) != "\x00\x00\x00\x00" {
		return errors.ErrUnsupported
	}
	SignatureBytes, err := fc.Read(16)
	if err != nil {
		return err
	}
	if string(SignatureBytes) != "TRUEVISION-XFILE" {
		return errFileFormatError
	}

	DotByte, err := fc.ReadByte()
	if err != nil {
		return err
	}
	if DotByte != '.' {
		return errFileFormatError
	}
	NullByte, err := fc.ReadByte()
	if err != nil {
		return err
	}
	if NullByte != 0x00 {
		return errFileFormatError
	}
	return nil
}

// decodeHeader reads and decodes the TGA header from the beginning of the file. Moves the cursor back to the original position after reading.
func decodeHeader(fc *fcursor.FileCursor) (TGAHeader, error) {
	startPos, _ := fc.Tell()
	defer fc.Seek(startPos, io.SeekStart)
	// Read ID Length
	IDLengthByte, err := fc.ReadByte()
	if err != nil {
		return TGAHeader{}, err
	}
	// Read Color Map Type
	ColorMapTypeByte, err := fc.ReadByte()
	if err != nil {
		return TGAHeader{}, err
	}
	// Read Image Type
	ImageTypeByte, err := fc.ReadByte()
	if err != nil {
		return TGAHeader{}, err
	}
	// Read Color Map Specification
	ColorMapSpecBytes, err := fc.Read(5)
	if err != nil {
		return TGAHeader{}, err
	}

	// Read Image Specification
	ImageSpecBytes, err := fc.Read(10)
	if err != nil {
		return TGAHeader{}, err
	}

	header := TGAHeader{
		IDLength:          uint8(IDLengthByte),
		ColorMapType:      uint8(ColorMapTypeByte),
		ImageType:         uint8(ImageTypeByte),
		FirstEntryIndex:   binary.LittleEndian.Uint16(ColorMapSpecBytes[0:2]),
		ColorMapLength:    binary.LittleEndian.Uint16(ImageSpecBytes[2:4]),
		ColorMapEntrySize: uint8(ColorMapSpecBytes[4]),
		XOrigin:           binary.LittleEndian.Uint16(ImageSpecBytes[0:2]),
		YOrigin:           binary.LittleEndian.Uint16(ImageSpecBytes[2:4]),
		ImageWidth:        binary.LittleEndian.Uint16(ImageSpecBytes[4:6]),
		ImageHeight:       binary.LittleEndian.Uint16(ImageSpecBytes[6:8]),
		PixelDepth:        uint8(ImageSpecBytes[8]),
		ImageDescriptor:   ImageSpecBytes[9],
	}

	return header, nil
}

func readImageID(fc *fcursor.FileCursor, header TGAHeader) (string, error) {
	if header.IDLength == 0 {
		return "", nil
	}
	startPos, _ := fc.Tell()
	defer fc.Seek(startPos, io.SeekStart)
	fc.Seek(18, io.SeekStart) // Move to the start of Image ID
	IDBytes, err := fc.Read(int(header.IDLength))
	if err != nil {
		return "", err
	}
	return string(IDBytes), nil
}

func readColorMapData(fc *fcursor.FileCursor, header TGAHeader) ([][]byte, error) {
	if header.ColorMapType != 1 {
		return nil, nil
	}
	startPos, _ := fc.Tell()
	defer fc.Seek(startPos, io.SeekStart)
	fc.Seek(18+int64(header.IDLength), io.SeekStart)
	ColorMapSize := int(header.ColorMapLength) * int((header.ColorMapEntrySize+7)/8)
	//ColorMapBytes, err := fc.Read(ColorMapSize)
	if err != nil {
		return nil, err
	}
	//
	return nil, nil
}

func DecodeTGA(f *os.File) error {
	fc := fcursor.NewFileCursorFromFile(f)
	defer fc.Close()

	err := checkFooter(fc)
	if err != nil {
		return err
	}
	header, err := decodeHeader(fc)
	if err != nil {
		return err
	}

}
