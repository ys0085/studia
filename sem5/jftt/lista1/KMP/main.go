package main

import (
	"bufio"
	"fmt"
	"os"
)

func computePrefixFunction(P string) []int {
	m := len(P)
	pi := make([]int, m)
	k := 0
	for q := 1; q < m; q++ {
		for k > 0 && P[k] != P[q] {
			k = pi[k-1]
		}
		if P[k] == P[q] {
			k++
		}
		pi[q] = k
	}
	return pi
}

func kmpMatcher(T, P string) []int {
	n := len(T)
	m := len(P)
	if m == 0 {
		res := make([]int, n+1)
		for i := range res {
			res[i] = i
		}
		return res
	}
	pi := computePrefixFunction(P)
	q := 0
	var occ []int
	for i := 0; i < n; i++ {
		for q > 0 && P[q] != T[i] {
			q = pi[q-1]
		}
		if P[q] == T[i] {
			q++
		}
		if q == m {
			occ = append(occ, i-m+1)
			q = pi[q-1]
		}
	}
	return occ
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Użycie: KMP <wzorzec> <nazwa_pliku>")
		os.Exit(1)
	}
	pattern := os.Args[1]
	filename := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Błąd otwarcia pliku:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	var text []rune
	for scanner.Scan() {
		text = append(text, []rune(scanner.Text())[0])
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Błąd odczytu:", err)
		os.Exit(1)
	}

	occ := kmpMatcher(string(text), pattern)
	for _, pos := range occ {
		fmt.Println(pos)
	}
}
