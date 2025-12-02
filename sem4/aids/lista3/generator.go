package main

import (
	"fmt"
	"math/rand/v2" //This generator is better, but requires at least golang version 1.20
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	k := rand.IntN(n)
	fmt.Print(n, " ")
	fmt.Print(k, " ")
	for range n {
		fmt.Print(rand.IntN(2*n-1), " ")
	}
}
