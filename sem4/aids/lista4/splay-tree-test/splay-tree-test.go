package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Counters for measuring complexity
type Counters struct {
	Comparisons   int
	PointerReads  int
	PointerWrites int
	Height        int
}

// Instrumented NodeSplay for SplayTree
type NodeSplay struct {
	Key   int
	Count int
	Left  *NodeSplay
	Right *NodeSplay
}

// Instrumented SplayTree for complexity measurement
type SplayTree struct {
	Root *NodeSplay
}

func NewSplayTreeInstr() *SplayTree {
	return &SplayTree{Root: nil}
}

// rotateRight with counters
func (st *SplayTree) rotateRight(node *NodeSplay, c *Counters) *NodeSplay {
	newRoot := node.Left
	c.PointerReads++
	node.Left = newRoot.Right
	c.PointerWrites++
	newRoot.Right = node
	c.PointerWrites++
	return newRoot
}

// rotateLeft with counters
func (st *SplayTree) rotateLeft(node *NodeSplay, c *Counters) *NodeSplay {
	newRoot := node.Right
	c.PointerReads++
	node.Right = newRoot.Left
	c.PointerWrites++
	newRoot.Left = node
	c.PointerWrites++
	return newRoot
}

// splay with counters
func (st *SplayTree) splay(root *NodeSplay, key int, c *Counters) *NodeSplay {
	if root == nil || root.Key == key {
		c.PointerReads++
		if root != nil {
			c.Comparisons++
		}
		return root
	}

	c.Comparisons++
	if key < root.Key {
		c.PointerReads++
		if root.Left == nil {
			return root
		}
		c.Comparisons++
		// Zig-Zig (Left Left)
		if key < root.Left.Key {
			root.Left.Left = st.splay(root.Left.Left, key, c)
			root = st.rotateRight(root, c)
		}
		c.PointerReads++
		if root.Left != nil && key > root.Left.Key {
			root.Left.Right = st.splay(root.Left.Right, key, c)
			if root.Left.Right != nil {
				root.Left = st.rotateLeft(root.Left, c)
			}
		}
		c.PointerReads++
		if root.Left == nil {
			return root
		} else {
			return st.rotateRight(root, c)
		}
	} else {
		c.PointerReads++
		if root.Right == nil {
			return root
		}
		c.Comparisons++
		// Zig-Zag (Right Left)
		if key < root.Right.Key {
			root.Right.Left = st.splay(root.Right.Left, key, c)
			if root.Right.Left != nil {
				root.Right = st.rotateRight(root.Right, c)
			}
		}
		c.Comparisons++
		// Zig-Zig (Right Right)
		if root.Right != nil && key > root.Right.Key {
			root.Right.Right = st.splay(root.Right.Right, key, c)
			root = st.rotateLeft(root, c)
		}
		c.PointerReads++
		if root.Right == nil {
			return root
		} else {
			return st.rotateLeft(root, c)
		}
	}
}

// Insert with counters
func (st *SplayTree) Insert(key int, c *Counters) {
	if st.Root == nil {
		st.Root = &NodeSplay{Key: key, Count: 1, Left: nil, Right: nil}
		c.PointerWrites++
		return
	}

	st.Root = st.splay(st.Root, key, c)

	c.Comparisons++
	if st.Root.Key == key {
		st.Root.Count++
		return
	}

	newNode := &NodeSplay{Key: key, Count: 1, Left: nil, Right: nil}
	c.PointerWrites++

	if key < st.Root.Key {
		newNode.Right = st.Root
		newNode.Left = st.Root.Left
		st.Root.Left = nil
		c.PointerWrites += 3
	} else {
		newNode.Left = st.Root
		newNode.Right = st.Root.Right
		st.Root.Right = nil
		c.PointerWrites += 3
	}

	st.Root = newNode
	c.PointerWrites++
}

// Delete with counters
func (st *SplayTree) Delete(key int, c *Counters) bool {
	if st.Root == nil {
		return false
	}

	st.Root = st.splay(st.Root, key, c)

	c.Comparisons++
	if st.Root.Key != key {
		return false
	}

	if st.Root.Count > 1 {
		st.Root.Count--
		return true
	}

	if st.Root.Left == nil {
		st.Root = st.Root.Right
		c.PointerWrites++
	} else if st.Root.Right == nil {
		st.Root = st.Root.Left
		c.PointerWrites++
	} else {
		leftSubtree := st.Root.Left
		st.Root = st.Root.Right
		c.PointerWrites++

		leftSubtree = st.splay(leftSubtree, key, c)
		leftSubtree.Right = st.Root
		st.Root = leftSubtree
		c.PointerWrites += 2
	}

	return true
}

// Height returns the height of the tree
func (st *SplayTree) Height() int {
	return st.getHeight(st.Root)
}

func (st *SplayTree) getHeight(node *NodeSplay) int {
	if node == nil {
		return -1
	}
	leftHeight := st.getHeight(node.Left)
	rightHeight := st.getHeight(node.Right)
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

// --- Experiment logic ---

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
		st := NewSplayTreeInstr()
		var insertOrder, deleteOrder []int
		if scenario == 0 {
			insertOrder = make([]int, n)
			for i := 0; i < n; i++ {
				insertOrder[i] = i + 1
			}
			deleteOrder = randomPermutation(n)
		} else {
			insertOrder = randomPermutation(n)
			deleteOrder = randomPermutation(n)
		}

		// Insert
		for _, key := range insertOrder {
			c := Counters{}
			st.Insert(key, &c)
			c.Height = st.Height()
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
			st.Delete(key, &c)
			c.Height = st.Height()
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
	rand.Seed(time.Now().UnixNano())
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
