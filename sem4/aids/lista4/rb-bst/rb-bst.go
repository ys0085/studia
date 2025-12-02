package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// Color represents the color of a Red-Black tree node
type Color int

const (
	RED   Color = 0
	BLACK Color = 1
)

// Node reprezentuje węzeł w drzewie Red-Black
type Node struct {
	Key    int
	Count  int // liczba wystąpień klucza
	Color  Color
	Left   *Node
	Right  *Node
	Parent *Node
}

// RBTree reprezentuje drzewo Red-Black
type RBTree struct {
	Root *Node
	NIL  *Node // sentinel node
}

// NewRBTree tworzy nowe puste drzewo Red-Black
func NewRBTree() *RBTree {
	nil_node := &Node{Color: BLACK}
	return &RBTree{
		Root: nil_node,
		NIL:  nil_node,
	}
}

// Insert wstawia klucz do drzewa Red-Black
func (rbt *RBTree) Insert(key int) {
	// Sprawdzenie czy klucz już istnieje
	existing := rbt.search(key)
	if existing != rbt.NIL {
		existing.Count++
		return
	}

	// Tworzenie nowego węzła
	newNode := &Node{
		Key:    key,
		Count:  1,
		Color:  RED,
		Left:   rbt.NIL,
		Right:  rbt.NIL,
		Parent: rbt.NIL,
	}

	// Standardowe wstawienie BST
	y := rbt.NIL
	x := rbt.Root

	for x != rbt.NIL {
		y = x
		if newNode.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	newNode.Parent = y

	if y == rbt.NIL {
		rbt.Root = newNode
	} else if newNode.Key < y.Key {
		y.Left = newNode
	} else {
		y.Right = newNode
	}

	// Naprawa właściwości Red-Black
	rbt.insertFixup(newNode)
}

// insertFixup naprawia właściwości Red-Black po wstawieniu
func (rbt *RBTree) insertFixup(z *Node) {
	for z.Parent.Color == RED {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y.Color == RED {
				// Przypadek 1: wuj jest czerwony
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					// Przypadek 2: wuj jest czarny i z jest prawym dzieckiem
					z = z.Parent
					rbt.leftRotate(z)
				}
				// Przypadek 3: wuj jest czarny i z jest lewym dzieckiem
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				rbt.rightRotate(z.Parent.Parent)
			}
		} else {
			y := z.Parent.Parent.Left
			if y.Color == RED {
				// Przypadek 1: wuj jest czerwony
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					// Przypadek 2: wuj jest czarny i z jest lewym dzieckiem
					z = z.Parent
					rbt.rightRotate(z)
				}
				// Przypadek 3: wuj jest czarny i z jest prawym dzieckiem
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				rbt.leftRotate(z.Parent.Parent)
			}
		}
	}
	rbt.Root.Color = BLACK
}

// Delete usuwa jedno wystąpienie klucza z drzewa
func (rbt *RBTree) Delete(key int) bool {
	z := rbt.search(key)
	if z == rbt.NIL {
		return false
	}

	// Jeśli jest więcej niż jedno wystąpienie, zmniejszamy licznik
	if z.Count > 1 {
		z.Count--
		return true
	}

	rbt.deleteNode(z)
	return true
}

// deleteNode usuwa węzeł z drzewa Red-Black
func (rbt *RBTree) deleteNode(z *Node) {
	y := z
	yOriginalColor := y.Color
	var x *Node

	if z.Left == rbt.NIL {
		x = z.Right
		rbt.transplant(z, z.Right)
	} else if z.Right == rbt.NIL {
		x = z.Left
		rbt.transplant(z, z.Left)
	} else {
		y = rbt.minimum(z.Right)
		yOriginalColor = y.Color
		x = y.Right

		if y.Parent == z {
			x.Parent = y
		} else {
			rbt.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}

		rbt.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}

	if yOriginalColor == BLACK {
		rbt.deleteFixup(x)
	}
}

// deleteFixup naprawia właściwości Red-Black po usunięciu
func (rbt *RBTree) deleteFixup(x *Node) {
	for x != rbt.Root && x.Color == BLACK {
		if x == x.Parent.Left {
			w := x.Parent.Right
			if w.Color == RED {
				// Przypadek 1: brat jest czerwony
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.leftRotate(x.Parent)
				w = x.Parent.Right
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				// Przypadek 2: brat jest czarny i oba jego dzieci są czarne
				w.Color = RED
				x = x.Parent
			} else {
				if w.Right.Color == BLACK {
					// Przypadek 3: brat jest czarny, lewe dziecko czerwone, prawe czarne
					w.Left.Color = BLACK
					w.Color = RED
					rbt.rightRotate(w)
					w = x.Parent.Right
				}
				// Przypadek 4: brat jest czarny i prawe dziecko jest czerwone
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rbt.leftRotate(x.Parent)
				x = rbt.Root
			}
		} else {
			w := x.Parent.Left
			if w.Color == RED {
				// Przypadek 1: brat jest czerwony
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.rightRotate(x.Parent)
				w = x.Parent.Left
			}
			if w.Right.Color == BLACK && w.Left.Color == BLACK {
				// Przypadek 2: brat jest czarny i oba jego dzieci są czarne
				w.Color = RED
				x = x.Parent
			} else {
				if w.Left.Color == BLACK {
					// Przypadek 3: brat jest czarny, prawe dziecko czerwone, lewe czarne
					w.Right.Color = BLACK
					w.Color = RED
					rbt.leftRotate(w)
					w = x.Parent.Left
				}
				// Przypadek 4: brat jest czarny i lewe dziecko jest czerwone
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Left.Color = BLACK
				rbt.rightRotate(x.Parent)
				x = rbt.Root
			}
		}
	}
	x.Color = BLACK
}

// leftRotate wykonuje lewą rotację
func (rbt *RBTree) leftRotate(x *Node) {
	y := x.Right
	x.Right = y.Left

	if y.Left != rbt.NIL {
		y.Left.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == rbt.NIL {
		rbt.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

// rightRotate wykonuje prawą rotację
func (rbt *RBTree) rightRotate(x *Node) {
	y := x.Left
	x.Left = y.Right

	if y.Right != rbt.NIL {
		y.Right.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == rbt.NIL {
		rbt.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}

	y.Right = x
	x.Parent = y
}

// transplant zastępuje jedno poddrzewo drugim
func (rbt *RBTree) transplant(u, v *Node) {
	if u.Parent == rbt.NIL {
		rbt.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}

// search szuka węzła z danym kluczem
func (rbt *RBTree) search(key int) *Node {
	x := rbt.Root
	for x != rbt.NIL && key != x.Key {
		if key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	return x
}

// minimum znajduje węzeł z najmniejszym kluczem w poddrzewie
func (rbt *RBTree) minimum(x *Node) *Node {
	for x.Left != rbt.NIL {
		x = x.Left
	}
	return x
}

// Height zwraca wysokość drzewa
func (rbt *RBTree) Height() int {
	return rbt.getHeight(rbt.Root)
}

func (rbt *RBTree) getHeight(node *Node) int {
	if node == rbt.NIL {
		return -1
	}

	leftHeight := rbt.getHeight(node.Left)
	rightHeight := rbt.getHeight(node.Right)

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

// PrintTree wyświetla drzewo w czytelnej formie
func (rbt *RBTree) PrintTree() {
	if rbt.Root == rbt.NIL {
		fmt.Println("Drzewo jest puste")
		return
	}

	fmt.Println("Struktura drzewa Red-Black:")
	rbt.printNode(rbt.Root, "", true)
	fmt.Printf("Wysokość drzewa: %d\n", rbt.Height())
	fmt.Println()
}

func (rbt *RBTree) printNode(node *Node, prefix string, isLast bool) {
	if node == rbt.NIL {
		return
	}

	// Wyświetl węzeł z odpowiednim prefixem
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	colorStr := "R"
	if node.Color == BLACK {
		colorStr = "B"
	}

	countStr := ""
	if node.Count > 1 {
		countStr = fmt.Sprintf(" (x%d)", node.Count)
	}

	fmt.Printf("%s%s%d%s [%s]\n", prefix, connector, node.Key, countStr, colorStr)

	// Przygotuj prefix dla dzieci
	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	// Wyświetl dzieci (najpierw prawe, potem lewe, żeby drzewo wyglądało naturalnie)
	hasLeft := node.Left != rbt.NIL
	hasRight := node.Right != rbt.NIL

	if hasRight {
		rbt.printNode(node.Right, childPrefix, !hasLeft)
	}
	if hasLeft {
		rbt.printNode(node.Left, childPrefix, true)
	}
}

// generatePermutation generuje losową permutację liczb od 1 do n
func generatePermutation(n int) []int {
	perm := make([]int, n)
	for i := range n {
		perm[i] = i + 1
	}
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		perm[i], perm[j] = perm[j], perm[i]
	}

	return perm
}

// demonstrateCase1 demonstruje przypadek 1: wstawianie rosnącego ciągu, usuwanie losowej permutacji
func demonstrateCase1(n int) {
	fmt.Println("=== PRZYPADEK 1: Wstawianie rosnącego ciągu 1,2,...,n, usuwanie losowej permutacji ===")
	fmt.Printf("n = %d\n\n", n)

	rbt := NewRBTree()

	// Wstawianie rosnącego ciągu
	fmt.Println("--- FAZA WSTAWIANIA (rosnący ciąg) ---")
	for i := 1; i <= n; i++ {
		fmt.Printf("insert %d\n", i)
		rbt.Insert(i)
		rbt.PrintTree()
	}

	// Usuwanie losowej permutacji
	fmt.Println("--- FAZA USUWANIA (losowa permutacja) ---")
	deleteOrder := generatePermutation(n)
	fmt.Printf("Kolejność usuwania: %v\n\n", deleteOrder)

	for _, key := range deleteOrder {
		fmt.Printf("delete %d\n", key)
		deleted := rbt.Delete(key)
		if !deleted {
			fmt.Printf("Klucz %d nie został znaleziony!\n", key)
		}
		rbt.PrintTree()
	}
}

// demonstrateCase2 demonstruje przypadek 2: wstawianie losowej permutacji, usuwanie losowej permutacji
func demonstrateCase2(n int) {
	fmt.Println("=== PRZYPADEK 2: Wstawianie losowej permutacji, usuwanie losowej permutacji ===")
	fmt.Printf("n = %d\n\n", n)

	rbt := NewRBTree()

	// Wstawianie losowej permutacji
	fmt.Println("--- FAZA WSTAWIANIA (losowa permutacja) ---")
	insertOrder := generatePermutation(n)
	fmt.Printf("Kolejność wstawiania: %v\n\n", insertOrder)

	for _, key := range insertOrder {
		fmt.Printf("insert %d\n", key)
		rbt.Insert(key)
		rbt.PrintTree()
	}

	// Usuwanie losowej permutacji
	fmt.Println("--- FAZA USUWANIA (losowa permutacja) ---")
	deleteOrder := generatePermutation(n)
	fmt.Printf("Kolejność usuwania: %v\n\n", deleteOrder)

	for _, key := range deleteOrder {
		fmt.Printf("delete %d\n", key)
		deleted := rbt.Delete(key)
		if !deleted {
			fmt.Printf("Klucz %d nie został znaleziony!\n", key)
		}
		rbt.PrintTree()
	}
}

func main() {
	n := 15 // Zmniejszone dla lepszej czytelności Red-Black tree

	// Demonstracja przypadku 1
	demonstrateCase1(n)

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

	// Demonstracja przypadku 2
	demonstrateCase2(n)

	fmt.Println("=== KONIEC DEMONSTRACJI ===")
}
