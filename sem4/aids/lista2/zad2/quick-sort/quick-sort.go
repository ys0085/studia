package main

import (
	"flag"
	"fmt"
)

var debug_mode bool
var keys []int

func main() {
	var count int
	fmt.Scan(&count)

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

	quick_sort(keys)

	debug_print()

	fmt.Print(cmp_counter, swp_counter)
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

func quick_sort(arr []int) {
	if len(arr) > 1 {
		partitionIndex := partition(arr)

		quick_sort(arr[:partitionIndex])
		quick_sort(arr[partitionIndex+1:])
	}
}

func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	index := -1
	for j, val := range arr {
		if cmp(val < pivot) {
			index++
			swap(arr, index, j)
		}
	}
	swap(arr, index+1, len(arr)-1)
	debug_print()
	return index + 1
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	swp_counter++
}

var cmp_counter int = 0

func cmp(a bool) bool {
	cmp_counter++
	return a
}

var swp_counter int = 0

func init() {
	flag.Parse()

}
