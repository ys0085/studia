package main

import (
	"flag"
	"os"
	"tga/decoder"
)

var inputFilePath string
var outputFilePath string

func init() {
	flag.StringVar(&inputFilePath, "Input File", "", "File in TGA format, to decode")
	flag.StringVar(&outputFilePath, "Output File", "", "File to encode into")
}

func main() {
	flag.Parse()

	f, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = decoder.DecodeTGA(f)
}
