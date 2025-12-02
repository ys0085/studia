package main

import (
	"os"
)

func count_file(path string) ([]int, [][]int) {
	var counts []int = make([]int, 256)
	var transitions [][]int = make([][]int, 256)
	for i := range transitions {
		transitions[i] = make([]int, 256)
	}

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	curr := make([]byte, 1)
	var prev byte = 0

	var e error
	var n int

	for e == nil {
		n, e = file.Read(curr)
		if n > 0 {
			counts[curr[0]]++
			transitions[prev][curr[0]]++
			prev = curr[0]
		}
	}
	return counts, transitions
}
