package main

import (
	"fmt"
)

func findRuns(arr []int) [][2]int {
	if len(arr) == 0 {
		return [][2]int{}
	}

	var runs [][2]int
	start := 0

	for i := 1; i < len(arr); i++ {
		if cmp(arr[i] < arr[i-1]) {
			runs = append(runs, [2]int{start, i - 1})
			start = i
		}
	}

	runs = append(runs, [2]int{start, len(arr) - 1})
	return runs
}

func gallopingMerge(arr []int, left, mid, right int) {
	leftSize := mid - left + 1
	rightSize := right - mid

	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)

	for i := range leftSize {
		leftArr[i] = arr[left+i]
	}
	for i := range rightSize {
		rightArr[i] = arr[mid+1+i]
	}

	i, j, k := 0, 0, left
	minGallop := 7

	for cmp(i < leftSize) && cmp(j < rightSize) {
		leftWins, rightWins := 0, 0

		for {
			if cmp(leftArr[i] <= rightArr[j]) {
				arr[k] = leftArr[i]
				swp_counter++
				i++
				k++
				leftWins++
				rightWins = 0
				if i == leftSize {
					break
				}
				if cmp(leftWins >= minGallop) {
					break
				}
			} else {
				arr[k] = rightArr[j]
				swp_counter++
				j++
				k++
				rightWins++
				leftWins = 0
				if j == rightSize {
					break
				}
				if cmp(rightWins >= minGallop) {
					break
				}
			}
		}

		if cmp(leftWins >= minGallop) {
			for cmp(i < leftSize) && cmp(leftArr[i] <= rightArr[j]) {
				arr[k] = leftArr[i]
				swp_counter++
				i++
				k++
			}
		}

		if cmp(rightWins >= minGallop) {
			for cmp(j < rightSize) && cmp(rightArr[j] < leftArr[i]) {
				arr[k] = rightArr[j]
				swp_counter++
				j++
				k++
			}
		}
	}

	for cmp(i < leftSize) {
		arr[k] = leftArr[i]
		swp_counter++
		i++
		k++
	}

	for cmp(j < rightSize) {
		arr[k] = rightArr[j]
		swp_counter++
		j++
		k++
	}
}

func advancedMergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	runs := findRuns(arr)

	for len(runs) > 1 {
		minRunsIdx := 0
		minRunsLen := len(arr) + 2

		for i := range len(runs) - 1 {
			currLen := runs[i+1][1] - runs[i][0] + 1
			if currLen < minRunsLen {
				minRunsLen = currLen
				minRunsIdx = i
			}
		}

		gallopingMerge(arr, runs[minRunsIdx][0], runs[minRunsIdx][1], runs[minRunsIdx+1][1])

		newRun := [2]int{runs[minRunsIdx][0], runs[minRunsIdx+1][1]}
		newRuns := make([][2]int, 0, len(runs)-1)

		for i := range minRunsIdx {
			newRuns = append(newRuns, runs[i])
		}

		newRuns = append(newRuns, newRun)

		for i := minRunsIdx + 2; i < len(runs); i++ {
			newRuns = append(newRuns, runs[i])
		}

		runs = newRuns
	}
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

	advancedMergeSort(keys)

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
