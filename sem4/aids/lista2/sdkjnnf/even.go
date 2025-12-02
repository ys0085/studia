package main

import (
	"fmt"
	"time"
)

func isEven(num int) bool {
	if num == 1 {
		return false
	} else if num == 0 {
		return true
	} else {
		return isEven(num - 2)
	}
}

func isEven2(num int) bool {
	return num%2 == 0
}

func main() {
	start := time.Now()
	fmt.Println(isEven(1000000))
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	fmt.Println(isEven2(1000000))
	elapsed = time.Since(start)
	fmt.Println(elapsed)
}
