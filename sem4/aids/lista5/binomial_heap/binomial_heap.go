package main

import (
	"fmt"
	"math/rand"
)

type BinomialHeapNode struct {
	p       *BinomialHeapNode
	key     int
	degree  int
	child   *BinomialHeapNode
	sibling *BinomialHeapNode
}

type BinomialHeap struct {
	root *BinomialHeapNode
}

var cmp_count int

func cmp(b bool) bool {
	cmp_count++
	return b
}

func MakeHeap() *BinomialHeap {
	return &BinomialHeap{root: nil}
}

func (h *BinomialHeap) HeapMinimum() *BinomialHeapNode {
	if h.root == nil {
		return nil
	}

	minNode := h.root
	current := h.root.sibling

	for current != nil {
		if cmp(current.key < minNode.key) {
			minNode = current
		}
		current = current.sibling
	}

	return minNode
}

func HeapUnion(h1, h2 *BinomialHeap) *BinomialHeap {
	h := MakeHeap()
	h.root = HeapMerge(h1.root, h2.root)
	if h.root == nil {
		return h
	}
	var prev *BinomialHeapNode = nil
	var current *BinomialHeapNode = h.root
	var next *BinomialHeapNode = current.sibling
	for next != nil {
		if cmp(current.degree != next.degree) || (next.sibling != nil && cmp(next.sibling.degree == current.degree)) {
			prev = current
			current = next
		} else {
			if cmp(current.key <= next.key) {
				current.sibling = next.sibling
				HeapLink(next, current)
			} else {
				if prev == nil {
					h.root = next
				} else {
					prev.sibling = next
				}
				HeapLink(current, next)
				current = next
			}
		}
		next = current.sibling
	}
	return h
}

func HeapLink(h1, h2 *BinomialHeapNode) {
	h1.p = h2
	h1.sibling = h2.child
	h2.child = h1
	h2.degree++
}

func HeapMerge(h1, h2 *BinomialHeapNode) *BinomialHeapNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	var head, tail *BinomialHeapNode

	if cmp(h1.degree <= h2.degree) {
		head = h1
		h1 = h1.sibling
	} else {
		head = h2
		h2 = h2.sibling
	}
	tail = head

	for h1 != nil && h2 != nil {
		var next *BinomialHeapNode
		if cmp(h1.degree <= h2.degree) {
			next = h1
			h1 = h1.sibling
		} else {
			next = h2
			h2 = h2.sibling
		}
		tail.sibling = next
		tail = next
	}

	if h1 != nil {
		tail.sibling = h1
	} else {
		tail.sibling = h2
	}

	return head
}

func (h *BinomialHeap) Insert(key int) {
	newNode := &BinomialHeapNode{key: key, degree: 0, child: nil, sibling: nil}
	newHeap := MakeHeap()
	newHeap.root = newNode

	h.root = HeapUnion(h, newHeap).root
}

func reverseChildren(node *BinomialHeapNode) *BinomialHeapNode {
	var prev *BinomialHeapNode = nil
	for node != nil {
		next := node.sibling
		node.sibling = prev
		node.p = nil
		prev = node
		node = next
	}
	return prev
}

func (h *BinomialHeap) ExtractMin() *BinomialHeapNode {
	minNode := h.HeapMinimum()
	if minNode == nil {
		return nil
	}

	if cmp(h.root == minNode) {
		h.root = minNode.sibling
	} else {
		current := h.root
		for current.sibling != nil && cmp(current.sibling != minNode) {
			current = current.sibling
		}
		if cmp(current.sibling == minNode) {
			current.sibling = minNode.sibling
		}
	}

	childHeap := MakeHeap()
	childHeap.root = reverseChildren(minNode.child)

	h.root = HeapUnion(h, childHeap).root

	return minNode
}

func BinomialHeapExperiment(n int) ([]int, bool, bool) {
	H1 := MakeHeap()
	H2 := MakeHeap()
	seq1 := make([]int, n)
	seq2 := make([]int, n)
	for i := 0; i < n; i++ {
		seq1[i] = rand.Int()
		seq2[i] = rand.Int()
	}

	insertComparisons := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		cmp_count = 0
		H1.Insert(seq1[i])
		insertComparisons = append(insertComparisons, cmp_count)
	}
	for i := 0; i < n; i++ {
		cmp_count = 0
		H2.Insert(seq2[i])
		insertComparisons = append(insertComparisons, cmp_count)
	}

	H := HeapUnion(H1, H2)
	extracts := make([]int, 0, 2*n)
	extractComparisons := make([]int, 0, 2*n)
	for i := 0; i < 2*n; i++ {
		cmp_count = 0
		node := H.ExtractMin()
		extractComparisons = append(extractComparisons, cmp_count)
		if node != nil {
			extracts = append(extracts, node.key)
		}
	}

	sorted := true
	for i := 1; i < len(extracts); i++ {
		if extracts[i] < extracts[i-1] {
			sorted = false
			break
		}
	}
	empty := H.root == nil

	// Concatenate insert and extract comparisons
	allComparisons := append(insertComparisons, extractComparisons...)
	return allComparisons, sorted, empty
}

func BinomialHeapExperimentsSeries() {
	for n := 100; n <= 10000; n += 100 {
		total := 0
		for range 5 {
			comps, _, _ := BinomialHeapExperiment(n)
			sum := 0
			for _, c := range comps {
				sum += c
			}
			total += sum
		}
		avg := float64(total) / float64(n*5)
		fmt.Printf("%d %.4f\n", n, avg)
	}
}

// 5 eksperymentów historycznych dla n=500
func BinomialHeapHistoryExperiments500() {
	for rep := 0; rep < 5; rep++ {
		comps, sorted, empty := BinomialHeapExperiment(500)
		fmt.Printf("Eksperyment %d: posortowane=%v, pusty=%v\n", rep+1, sorted, empty)
		for i, c := range comps {
			fmt.Printf("%d %d\n", i+1, c)
		}
		fmt.Println("---")
	}
}

func main() {
	BinomialHeapHistoryExperiments500()
	BinomialHeapExperimentsSeries()
}
