package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Node reprezentuje węzeł w drzewie BST
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// BST reprezentuje drzewo wyszukiwań binarnych
type BST struct {
	Root *Node
}

// Liczniki operacji
type Counters struct {
	Comparisons   int
	PointerReads  int
	PointerWrites int
	Height        int
}

func NewBST() *BST {
	return &BST{Root: nil}
}

func (bst *BST) Insert(key int, c *Counters) {
	bst.Root = bst.insertNode(bst.Root, key, c)
}

func (bst *BST) insertNode(node *Node, key int, c *Counters) *Node {
	c.PointerReads++
	if node == nil {
		c.PointerWrites++
		return &Node{Key: key}
	}
	c.Comparisons++
	if key == node.Key {
		// Klucz już istnieje, nie wstawiamy ponownie
		return node
	}
	c.Comparisons++
	if key < node.Key {
		node.Left = bst.insertNode(node.Left, key, c)
		c.PointerWrites++
	} else {
		node.Right = bst.insertNode(node.Right, key, c)
		c.PointerWrites++
	}
	return node
}

func (bst *BST) Delete(key int, c *Counters) bool {
	var deleted bool
	bst.Root, deleted = bst.deleteNode(bst.Root, key, c)
	return deleted
}

func (bst *BST) deleteNode(node *Node, key int, c *Counters) (*Node, bool) {
	c.PointerReads++
	if node == nil {
		return nil, false
	}
	var deleted bool
	c.Comparisons++
	if key < node.Key {
		node.Left, deleted = bst.deleteNode(node.Left, key, c)
		c.PointerWrites++
	} else {
		c.Comparisons++
		if key > node.Key {
			node.Right, deleted = bst.deleteNode(node.Right, key, c)
			c.PointerWrites++
		} else {
			deleted = true
			if node.Left == nil {
				c.PointerReads++
				return node.Right, deleted
			}
			if node.Right == nil {
				c.PointerReads++
				return node.Left, deleted
			}
			successor := bst.findMin(node.Right, c)
			node.Key = successor.Key
			node.Right, _ = bst.deleteNode(node.Right, successor.Key, c)
			c.PointerWrites++
		}
	}
	return node, deleted
}

func (bst *BST) findMin(node *Node, c *Counters) *Node {
	for node.Left != nil {
		c.PointerReads++
		node = node.Left
	}
	return node
}

func (bst *BST) Height() int {
	return bst.getHeight(bst.Root)
}

func (bst *BST) getHeight(node *Node) int {
	if node == nil {
		return -1
	}
	leftHeight := bst.getHeight(node.Left)
	rightHeight := bst.getHeight(node.Right)
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

func randomPermutation(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i + 1
	}
	rand.Shuffle(n, func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func runExperiment(n int, scenario int) (avg, max Counters) {
	const trials = 20
	var sum Counters
	var maxOp Counters

	for t := 0; t < trials; t++ {
		bst := NewBST()
		var insertOrder, deleteOrder []int
		if scenario == 0 {
			// Wstawianie rosnącego ciągu, usuwanie losowej permutacji
			insertOrder = make([]int, n)
			for i := 0; i < n; i++ {
				insertOrder[i] = i + 1
			}
			deleteOrder = randomPermutation(n)
		} else {
			// Wstawianie losowej permutacji, usuwanie losowej permutacji
			insertOrder = randomPermutation(n)
			deleteOrder = randomPermutation(n)
		}

		// Insert
		for _, key := range insertOrder {
			c := Counters{}
			bst.Insert(key, &c)
			c.Height = bst.Height()
			sum.Comparisons += c.Comparisons
			sum.PointerReads += c.PointerReads
			sum.PointerWrites += c.PointerWrites
			sum.Height += c.Height
			if c.Comparisons > maxOp.Comparisons {
				maxOp.Comparisons = c.Comparisons
			}
			if c.PointerReads > maxOp.PointerReads {
				maxOp.PointerReads = c.PointerReads
			}
			if c.PointerWrites > maxOp.PointerWrites {
				maxOp.PointerWrites = c.PointerWrites
			}
			if c.Height > maxOp.Height {
				maxOp.Height = c.Height
			}
		}
		// Delete
		for _, key := range deleteOrder {
			c := Counters{}
			bst.Delete(key, &c)
			c.Height = bst.Height()
			sum.Comparisons += c.Comparisons
			sum.PointerReads += c.PointerReads
			sum.PointerWrites += c.PointerWrites
			sum.Height += c.Height
			if c.Comparisons > maxOp.Comparisons {
				maxOp.Comparisons = c.Comparisons
			}
			if c.PointerReads > maxOp.PointerReads {
				maxOp.PointerReads = c.PointerReads
			}
			if c.PointerWrites > maxOp.PointerWrites {
				maxOp.PointerWrites = c.PointerWrites
			}
			if c.Height > maxOp.Height {
				maxOp.Height = c.Height
			}
		}
	}
	ops := 2 * n * trials
	avg.Comparisons = sum.Comparisons / ops
	avg.PointerReads = sum.PointerReads / ops
	avg.PointerWrites = sum.PointerWrites / ops
	avg.Height = sum.Height / ops
	return avg, maxOp
}

func main() {
	fmt.Printf("n,scenario,avg_cmp,avg_read,avg_write,avg_height,max_cmp,max_read,max_write,max_height\n")
	for n := 10000; n <= 100000; n += 10000 {
		for scenario := 0; scenario < 2; scenario++ {
			avg, max := runExperiment(n, scenario)
			fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n",
				n, scenario,
				avg.Comparisons, avg.PointerReads, avg.PointerWrites, avg.Height,
				max.Comparisons, max.PointerReads, max.PointerWrites, max.Height)
		}
	}
}
