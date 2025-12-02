package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Color int

const (
	RED   Color = 0
	BLACK Color = 1
)

// Counters for measuring complexity
type Counters struct {
	Comparisons   int
	PointerReads  int
	PointerWrites int
	Height        int
}

// --- Instrumented Red-Black Tree for complexity measurement ---

// NodeRB is a wrapper for the RBTree Node to allow counting
type NodeRB struct {
	Key    int
	Count  int
	Color  Color
	Left   *NodeRB
	Right  *NodeRB
	Parent *NodeRB
}

// RBTree is a Red-Black tree with counters
type RBTree struct {
	Root *NodeRB
	NIL  *NodeRB
}

// NewRBTreeTest creates a new empty Red-Black tree for testing
func NewRBTreeTest() *RBTree {
	nilNode := &NodeRB{Color: BLACK}
	return &RBTree{
		Root: nilNode,
		NIL:  nilNode,
	}
}

// Insert with counters
func (rbt *RBTree) Insert(key int, c *Counters) {
	existing := rbt.search(key, c)
	if existing != rbt.NIL {
		existing.Count++
		return
	}

	newNode := &NodeRB{
		Key:    key,
		Count:  1,
		Color:  RED,
		Left:   rbt.NIL,
		Right:  rbt.NIL,
		Parent: rbt.NIL,
	}

	y := rbt.NIL
	x := rbt.Root

	for x != rbt.NIL {
		c.PointerReads++
		y = x
		c.Comparisons++
		if newNode.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	newNode.Parent = y

	if y == rbt.NIL {
		rbt.Root = newNode
		c.PointerWrites++
	} else if newNode.Key < y.Key {
		y.Left = newNode
		c.PointerWrites++
	} else {
		y.Right = newNode
		c.PointerWrites++
	}

	rbt.insertFixup(newNode, c)
}

// insertFixup with counters
func (rbt *RBTree) insertFixup(z *NodeRB, c *Counters) {
	for z.Parent.Color == RED {
		c.PointerReads++
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			c.PointerReads++
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					rbt.leftRotate(z, c)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				rbt.rightRotate(z.Parent.Parent, c)
			}
		} else {
			y := z.Parent.Parent.Left
			c.PointerReads++
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					rbt.rightRotate(z, c)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				rbt.leftRotate(z.Parent.Parent, c)
			}
		}
	}
	rbt.Root.Color = BLACK
}

// Delete with counters
func (rbt *RBTree) Delete(key int, c *Counters) bool {
	z := rbt.search(key, c)
	if z == rbt.NIL {
		return false
	}
	if z.Count > 1 {
		z.Count--
		return true
	}
	rbt.deleteNode(z, c)
	return true
}

// deleteNode with counters
func (rbt *RBTree) deleteNode(z *NodeRB, c *Counters) {
	y := z
	yOriginalColor := y.Color
	var x *NodeRB

	if z.Left == rbt.NIL {
		x = z.Right
		rbt.transplant(z, z.Right, c)
	} else if z.Right == rbt.NIL {
		x = z.Left
		rbt.transplant(z, z.Left, c)
	} else {
		y = rbt.minimum(z.Right, c)
		yOriginalColor = y.Color
		x = y.Right

		if y.Parent == z {
			x.Parent = y
		} else {
			rbt.transplant(y, y.Right, c)
			y.Right = z.Right
			y.Right.Parent = y
		}

		rbt.transplant(z, y, c)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}

	if yOriginalColor == BLACK {
		rbt.deleteFixup(x, c)
	}
}

// deleteFixup with counters
func (rbt *RBTree) deleteFixup(x *NodeRB, c *Counters) {
	for x != rbt.Root && x.Color == BLACK {
		c.PointerReads++
		if x == x.Parent.Left {
			w := x.Parent.Right
			c.PointerReads++
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.leftRotate(x.Parent, c)
				w = x.Parent.Right
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					rbt.rightRotate(w, c)
					w = x.Parent.Right
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rbt.leftRotate(x.Parent, c)
				x = rbt.Root
			}
		} else {
			w := x.Parent.Left
			c.PointerReads++
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.rightRotate(x.Parent, c)
				w = x.Parent.Left
			}
			if w.Right.Color == BLACK && w.Left.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Left.Color == BLACK {
					w.Right.Color = BLACK
					w.Color = RED
					rbt.leftRotate(w, c)
					w = x.Parent.Left
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Left.Color = BLACK
				rbt.rightRotate(x.Parent, c)
				x = rbt.Root
			}
		}
	}
	x.Color = BLACK
}

// leftRotate with counters
func (rbt *RBTree) leftRotate(x *NodeRB, c *Counters) {
	y := x.Right
	x.Right = y.Left
	c.PointerWrites++
	if y.Left != rbt.NIL {
		y.Left.Parent = x
		c.PointerWrites++
	}
	y.Parent = x.Parent
	c.PointerWrites++
	if x.Parent == rbt.NIL {
		rbt.Root = y
		c.PointerWrites++
	} else if x == x.Parent.Left {
		x.Parent.Left = y
		c.PointerWrites++
	} else {
		x.Parent.Right = y
		c.PointerWrites++
	}
	y.Left = x
	x.Parent = y
	c.PointerWrites += 2
}

// rightRotate with counters
func (rbt *RBTree) rightRotate(x *NodeRB, c *Counters) {
	y := x.Left
	x.Left = y.Right
	c.PointerWrites++
	if y.Right != rbt.NIL {
		y.Right.Parent = x
		c.PointerWrites++
	}
	y.Parent = x.Parent
	c.PointerWrites++
	if x.Parent == rbt.NIL {
		rbt.Root = y
		c.PointerWrites++
	} else if x == x.Parent.Right {
		x.Parent.Right = y
		c.PointerWrites++
	} else {
		x.Parent.Left = y
		c.PointerWrites++
	}
	y.Right = x
	x.Parent = y
	c.PointerWrites += 2
}

// transplant with counters
func (rbt *RBTree) transplant(u, v *NodeRB, c *Counters) {
	if u.Parent == rbt.NIL {
		rbt.Root = v
		c.PointerWrites++
	} else if u == u.Parent.Left {
		u.Parent.Left = v
		c.PointerWrites++
	} else {
		u.Parent.Right = v
		c.PointerWrites++
	}
	v.Parent = u.Parent
	c.PointerWrites++
}

// search with counters
func (rbt *RBTree) search(key int, c *Counters) *NodeRB {
	x := rbt.Root
	for x != rbt.NIL && key != x.Key {
		c.PointerReads++
		c.Comparisons++
		if key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	return x
}

// minimum with counters
func (rbt *RBTree) minimum(x *NodeRB, c *Counters) *NodeRB {
	for x.Left != rbt.NIL {
		c.PointerReads++
		x = x.Left
	}
	return x
}

// Height returns the height of the tree
func (rbt *RBTree) Height() int {
	return rbt.getHeight(rbt.Root)
}

func (rbt *RBTree) getHeight(node *NodeRB) int {
	if node == rbt.NIL {
		return -1
	}
	leftHeight := rbt.getHeight(node.Left)
	rightHeight := rbt.getHeight(node.Right)
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
		rbt := NewRBTreeTest()
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
			rbt.Insert(key, &c)
			c.Height = rbt.Height()
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
			rbt.Delete(key, &c)
			c.Height = rbt.Height()
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
