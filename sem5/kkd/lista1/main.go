package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	counts, transitions := count_file(os.Args[1])
	fmt.Println("Entropia:", entropy(counts))
	fmt.Println("Entropia warunkowa:", conditional_entropy(counts, transitions))
}

// H() = -sum_over_x(p(x)*log(p(x)))
func entropy(variable []int) float64 {
	var variable_sum int
	for _, n := range variable {
		variable_sum += n
	}

	var H float64 = 0.0
	for _, x := range variable {
		if x == 0 {
			continue
		}
		var p_x float64 = float64(x) / float64(variable_sum)
		H -= p_x * math.Log2(p_x)
	}
	return H
}

func conditional_entropy(counts []int, transitions [][]int) float64 {
	var total int
	for _, c := range counts {
		total += c
	}

	var H float64
	for y, t := range transitions {
		if counts[y] == 0 {
			continue
		}
		var p_y float64 = float64(counts[y]) / float64(total)
		for _, c := range t {
			// c == transitions[y][x]
			if c == 0 {
				continue
			}
			p_x_y := float64(c) / float64(counts[y])
			H -= p_y * (p_x_y * math.Log2(p_x_y))
		}
	}
	return H
}
