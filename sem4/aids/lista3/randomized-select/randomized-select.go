package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

var cmp_counter int
var swp_counter int

func cmp(b bool) bool {
	cmp_counter++
	return b
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if cmp(arr[j] <= pivot) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			swp_counter++
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	swp_counter++
	return i + 1
}

func randomizedPartition(arr []int, low, high int) int {
	randomIndex := rand.IntN(high-low+1) + low
	arr[randomIndex], arr[high] = arr[high], arr[randomIndex]
	swp_counter++
	return partition(arr, low, high)
}

func randomizedSelect(arr []int, low, high, k int) int {
	if cmp(low == high) {
		return arr[low]
	}
	q := randomizedPartition(arr, low, high)
	i := q - low + 1
	if cmp(k == i) {
		return arr[q]
	} else if cmp(k < i) {
		return randomizedSelect(arr, low, q-1, k)
	} else {
		return randomizedSelect(arr, q+1, high, k-i)
	}
}

var debug_mode bool
var keys []int

func main() {
	var count int
	var position int
	fmt.Scan(&count)
	fmt.Scan(&position)

	if count < 40 {
		debug_mode = true
	} else {
		debug_mode = false
	}

	keys = make([]int, count)
	for i := range keys {
		fmt.Scan(&(keys[i]))
	}

	debug_print()

	stat := randomizedSelect(keys, 0, count-1, position)

	debug_print()

	fmt.Println(cmp_counter, swp_counter)

	if debug_mode {
		sort.Ints(keys)
		print_full(keys)
		fmt.Println(stat)
	}
}

func debug_print() {
	if debug_mode {
		print_full(keys)
	}
}

func print_full(keys []int) {
	fmt.Print("[")
	for i, key := range keys {
		if i != len(keys)-1 {
			fmt.Printf("%02d ", key)
		} else {
			fmt.Printf("%02d]\n", key)
		}
	}
}
