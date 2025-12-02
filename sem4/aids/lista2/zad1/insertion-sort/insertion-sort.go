package main

import (
	"fmt"
)

var keys []int
var debug_mode bool

func main() {
	var count int
	fmt.Scan(&count)
	if count >= 40 {
		debug_mode = false
	} else {
		debug_mode = true
	}

	keys = make([]int, count)
	for i, _ := range keys {
		fmt.Scan(&(keys[i]))
	}

	debug_print()

	insertion_sort(keys)

	debug_print()

	correct := check(keys)

	if correct {
		fmt.Println("Sorted!")
	} else {
		fmt.Println("Sorting failed.")
	}

	fmt.Printf("Comparisons: %d\n", cmp_counter)
	fmt.Printf("Key swaps: %d\n", swp_counter)
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

func check(arr []int) bool {
	for i, key := range arr[1:] {
		if arr[i] > key {
			return false
		}
	}
	return true
}

func insertion_sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i, key := range arr[1:] {
		for i >= 0 && cmp(arr[i] > key) {
			arr[i+1] = arr[i]
			swp_counter++
			i = i - 1
		}
		arr[i+1] = key
		swp_counter++
		debug_print()
	}
	return arr
}

var cmp_counter int = 0

func cmp(a bool) bool {
	cmp_counter++
	return a
}

var swp_counter int = 0
