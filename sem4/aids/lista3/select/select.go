package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
)

var cmp_counter int = 0
var swp_counter int = 0

var size int

func init() {
	flag.IntVar(&size, "n", 5, "Group size for median of medians")
	flag.Parse()
}

func Select(arr []int, k int, groupSize int) int {
	cmp := func(a bool) bool {
		cmp_counter++
		return a
	}

	if len(arr) <= groupSize {
		_, c, s := insertion_sort(arr)
		cmp_counter += c
		swp_counter += s
		return arr[k-1]
	}

	medians := []int{}
	for i := 0; i < len(arr); i += groupSize {
		end := int(math.Min(float64(i+groupSize), float64(len(arr))))
		group := make([]int, end-i)
		copy(group, arr[i:end])
		_, c, s := insertion_sort(group)
		cmp_counter += c
		swp_counter += s
		medians = append(medians, group[len(group)/2])
	}

	pivot := Select(medians, (len(medians)+1)/2, groupSize)

	low := []int{}
	high := []int{}
	equal := []int{}
	for _, num := range arr {
		if cmp(num < pivot) {
			low = append(low, num)
			swp_counter++
		} else if cmp(num > pivot) {
			high = append(high, num)
			swp_counter++
		} else {
			equal = append(equal, num)
			swp_counter++
		}
	}

	if k <= len(low) {
		return Select(low, k, groupSize)
	} else if k <= len(low)+len(equal) {
		return pivot
	} else {
		return Select(high, k-len(low)-len(equal), groupSize)
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

	stat := Select(keys, position, size)

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
