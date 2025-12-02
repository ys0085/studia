package main

import "fmt"

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

	dual_pivot_qs(keys, 0, count-1)

	debug_print()

	fmt.Print(cmp_counter, swp_counter)
}

var cmp_counter int = 0

func cmp(a bool) bool {
	cmp_counter++
	return a
}

var swp_counter int = 0

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

func dual_pivot_qs(arr []int, low, high int) {
	if cmp(low < high) {
		p, q := partition(arr, low, high)
		dual_pivot_qs(arr, low, p-1)
		dual_pivot_qs(arr, p+1, q-1)
		dual_pivot_qs(arr, q+1, high)
	}
}

func partition(arr []int, low, high int) (int, int) {
	if cmp(arr[low] > arr[high]) {
		swap(arr, low, high)
	}
	p, q := arr[low], arr[high]

	i, j, k := low+1, low+1, high-1

	for cmp(k >= j) {
		if cmp(arr[j] < p) {
			swap(arr, i, j)
			i++
		} else if cmp(arr[j] > q) {
			for cmp(k > j) && cmp(arr[k] > q) {
				k--
			}
			swap(arr, j, k)
			k--
			if cmp(arr[j] < p) {
				swap(arr, i, j)
				i++
			}
		}
		j++
	}
	i--
	k++
	swap(arr, low, i)
	swap(arr, high, k)

	return i, k
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	swp_counter++
}
