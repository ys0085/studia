package main

import (
	"fmt"
	"sync"
	"time"
)

func sort(vals ...int) []int {
	var sorted []int
	var wg sync.WaitGroup
	for _, val := range vals {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			time.Sleep(time.Duration(v) * time.Millisecond)
			sorted = append(sorted, v)
		}(val)
	}
	wg.Wait()
	return sorted
}

func main() {
	start := time.Now()

	nums := []int{8, 9, 10, 7, 1, 1, 1000, 19, 12346}
	fmt.Println(nums)
	fmt.Println(sort(nums...))

	elapsed := time.Since(start)
	fmt.Println("Sorting took ", elapsed)

}
