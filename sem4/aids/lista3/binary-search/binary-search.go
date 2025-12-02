package main

import "fmt"

var cmp_counter int

func binary_search(arr []int, target int, low int, high int) bool {
	cmp := func(b bool) bool {
		cmp_counter++
		return b
	}

	if cmp(low > high) {
		return false
	}
	mid := low + (high-low)/2
	if cmp(arr[mid] == target) {
		return true
	} else if cmp(arr[mid] < target) {
		return binary_search(arr, target, mid+1, high)
	} else {
		return binary_search(arr, target, low, mid-1)
	}

}

var count int
var keys []int
var target int

func main() {
	fmt.Scan(&count)
	fmt.Scan(&target)

	keys = make([]int, count)
	for i := range keys {
		fmt.Scan(&(keys[i]))
	}
	found := binary_search(keys, target, 0, count-1)

	if len(keys) < 40 {
		fmt.Println(found)
	}
	fmt.Println(cmp_counter)
}

//
