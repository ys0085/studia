package main

import "fmt"

func merge_sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := merge_sort(arr[:mid])
	right := merge_sort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	for cmp(i < len(left)) && cmp(j < len(right)) {
		if cmp(left[i] <= right[j]) {
			result[k] = left[i]
			swp_counter++
			i++
		} else {
			result[k] = right[j]
			swp_counter++
			j++
		}
		k++
	}

	for cmp(i < len(left)) {
		result[k] = left[i]
		swp_counter++
		i++
		k++
	}

	for cmp(j < len(right)) {
		result[k] = right[j]
		swp_counter++
		j++
		k++
	}

	return result
}

var keys []int
var debug_mode bool

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

	merge_sort(keys)

	debug_print()

	fmt.Println(cmp_counter, swp_counter)
}

func debug_print() {
	if debug_mode {
		print_full(keys)
	}
}

var cmp_counter int = 0

func cmp(a bool) bool {
	cmp_counter++
	return a
}

var swp_counter int = 0

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
