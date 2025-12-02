package main

func insertion_sort(arr []int) ([]int, int, int) {
	// Returns the sorted array, number of comparisons, and number of swaps
	cmp_counter := 0
	swp_counter := 0
	cmp := func(a bool) bool {
		cmp_counter++
		return a
	}

	if len(arr) < 2 {
		return arr, cmp_counter, swp_counter
	}
	for i, key := range arr[1:] {
		for i >= 0 && cmp(arr[i] > key) {
			arr[i+1] = arr[i]
			swp_counter++
			i = i - 1
		}
		arr[i+1] = key
		swp_counter++
	}
	return arr, cmp_counter, swp_counter
}
