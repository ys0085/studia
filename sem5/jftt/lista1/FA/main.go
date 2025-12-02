package main

import (
	"bufio"
	"fmt"
	"os"
)

func computeTransitionFunction(P string, alphabet map[rune]bool) []map[rune]int {
	m := len(P)
	delta := make([]map[rune]int, m+1)
	for q := 0; q <= m; q++ {
		delta[q] = make(map[rune]int)
		for a := range alphabet {
			k := min(m, q+1)
			for k > 0 && P[:k] != suffix(P[:q]+string(a), k) {
				k--
			}
			delta[q][a] = k
		}
	}
	return delta
}

func suffix(s string, k int) string {
	if len(s) < k {
		return s
	}
	return s[len(s)-k:]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func finiteAutomatonMatcher(T, P string) []int {
	if len(P) == 0 {
		res := make([]int, len(T)+1)
		for i := range res {
			res[i] = i
		}
		return res
	}
	alphabet := make(map[rune]bool)
	for _, ch := range T + P {
		alphabet[ch] = true
	}
	delta := computeTransitionFunction(P, alphabet)
	q := 0
	m := len(P)
	var occ []int
	for i, ch := range T {
		if next, ok := delta[q][ch]; ok {
			q = next
		} else {
			q = 0
		}
		if q == m {
			occ = append(occ, i-m+1)
		}
	}
	return occ
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Użycie: FA <wzorzec> <nazwa_pliku>")
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

	reader := bufio.NewReader(file)
	content, _ := reader.ReadString(0)

	occ := finiteAutomatonMatcher(content, pattern)
	for _, pos := range occ {
		fmt.Println(pos)
	}
}
