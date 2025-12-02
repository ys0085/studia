package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// Node reprezentuje węzeł w drzewie Splay
type Node struct {
	Key   int
	Count int // liczba wystąpień klucza
	Left  *Node
	Right *Node
}

// SplayTree reprezentuje drzewo splay
type SplayTree struct {
	Root *Node
}

// NewSplayTree tworzy nowe puste drzewo Splay
func NewSplayTree() *SplayTree {
	return &SplayTree{Root: nil}
}

// rotateRight wykonuje rotację w prawo
func (st *SplayTree) rotateRight(node *Node) *Node {
	newRoot := node.Left
	node.Left = newRoot.Right
	newRoot.Right = node
	return newRoot
}

// rotateLeft wykonuje rotację w lewo
func (st *SplayTree) rotateLeft(node *Node) *Node {
	newRoot := node.Right
	node.Right = newRoot.Left
	newRoot.Left = node
	return newRoot
}

// splay wykonuje operację splay, przesuwając węzeł z danym kluczem na szczyt
func (st *SplayTree) splay(root *Node, key int) *Node {
	if root == nil || root.Key == key {
		return root
	}

	// Klucz jest w lewym poddrzewie
	if key < root.Key {
		if root.Left == nil {
			return root
		}

		// Zig-Zig (Left Left)
		if key < root.Left.Key {
			root.Left.Left = st.splay(root.Left.Left, key)
			root = st.rotateRight(root)
		}
		// Zig-Zag (Left Right)
		if root.Left != nil && key > root.Left.Key {
			root.Left.Right = st.splay(root.Left.Right, key)
			if root.Left.Right != nil {
				root.Left = st.rotateLeft(root.Left)
			}
		}

		if root.Left == nil {
			return root
		} else {
			return st.rotateRight(root)
		}
	} else {
		// Klucz jest w prawym poddrzewie
		if root.Right == nil {
			return root
		}

		// Zig-Zag (Right Left)
		if key < root.Right.Key {
			root.Right.Left = st.splay(root.Right.Left, key)
			if root.Right.Left != nil {
				root.Right = st.rotateRight(root.Right)
			}
		}
		// Zig-Zig (Right Right)
		if key > root.Right.Key {
			root.Right.Right = st.splay(root.Right.Right, key)
			root = st.rotateLeft(root)
		}

		if root.Right == nil {
			return root
		} else {
			return st.rotateLeft(root)
		}
	}
}

// Insert wstawia klucz do drzewa (obsługuje duplikaty przez licznik)
func (st *SplayTree) Insert(key int) {
	if st.Root == nil {
		st.Root = &Node{Key: key, Count: 1, Left: nil, Right: nil}
		return
	}

	st.Root = st.splay(st.Root, key)

	// Jeśli klucz już istnieje, zwiększamy licznik
	if st.Root.Key == key {
		st.Root.Count++
		return
	}

	// Tworzymy nowy węzeł
	newNode := &Node{Key: key, Count: 1, Left: nil, Right: nil}

	// Jeśli nowy klucz jest mniejszy od korzenia
	if key < st.Root.Key {
		newNode.Right = st.Root
		newNode.Left = st.Root.Left
		st.Root.Left = nil
	} else {
		// Jeśli nowy klucz jest większy od korzenia
		newNode.Left = st.Root
		newNode.Right = st.Root.Right
		st.Root.Right = nil
	}

	st.Root = newNode
}

// Delete usuwa jedno wystąpienie klucza z drzewa
func (st *SplayTree) Delete(key int) bool {
	if st.Root == nil {
		return false
	}

	st.Root = st.splay(st.Root, key)

	// Jeśli klucz nie został znaleziony
	if st.Root.Key != key {
		return false
	}

	// Jeśli jest więcej niż jedno wystąpienie, zmniejszamy licznik
	if st.Root.Count > 1 {
		st.Root.Count--
		return true
	}

	// Usuwamy węzeł
	if st.Root.Left == nil {
		st.Root = st.Root.Right
	} else if st.Root.Right == nil {
		st.Root = st.Root.Left
	} else {
		// Węzeł ma dwoje dzieci
		leftSubtree := st.Root.Left
		st.Root = st.Root.Right

		// Splay maksymalny element w lewym poddrzewie
		leftSubtree = st.splay(leftSubtree, key)
		leftSubtree.Right = st.Root
		st.Root = leftSubtree
	}

	return true
}

// findMax znajduje maksymalny klucz w poddrzewie
func (st *SplayTree) findMax(node *Node) int {
	for node.Right != nil {
		node = node.Right
	}
	return node.Key
}

// Height zwraca wysokość drzewa
func (st *SplayTree) Height() int {
	return st.getHeight(st.Root)
}

func (st *SplayTree) getHeight(node *Node) int {
	if node == nil {
		return -1 // Wysokość pustego drzewa to -1
	}

	leftHeight := st.getHeight(node.Left)
	rightHeight := st.getHeight(node.Right)

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

// PrintTree wyświetla drzewo w czytelnej formie
func (st *SplayTree) PrintTree() {
	if st.Root == nil {
		fmt.Println("Drzewo jest puste")
		return
	}

	fmt.Println("Struktura drzewa:")
	st.printNode(st.Root, "", true)
	fmt.Printf("Wysokość drzewa: %d\n", st.Height())
	fmt.Println()
}

func (st *SplayTree) printNode(node *Node, prefix string, isLast bool) {
	if node == nil {
		return
	}

	// Wyświetl węzeł z odpowiednim prefixem
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	countStr := ""
	if node.Count > 1 {
		countStr = fmt.Sprintf(" (x%d)", node.Count)
	}

	fmt.Printf("%s%s%d%s\n", prefix, connector, node.Key, countStr)

	// Przygotuj prefix dla dzieci
	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	// Wyświetl dzieci (najpierw prawe, potem lewe, żeby drzewo wyglądało naturalnie)
	hasLeft := node.Left != nil
	hasRight := node.Right != nil

	if hasRight {
		st.printNode(node.Right, childPrefix, !hasLeft)
	}
	if hasLeft {
		st.printNode(node.Left, childPrefix, true)
	}
}

// generatePermutation generuje losową permutację liczb od 1 do n
func generatePermutation(n int) []int {
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i + 1
	}

	// Algorytm Fisher-Yates shuffle
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

	st := NewSplayTree()

	// Wstawianie rosnącego ciągu
	fmt.Println("--- FAZA WSTAWIANIA (rosnący ciąg) ---")
	for i := 1; i <= n; i++ {
		fmt.Printf("insert %d\n", i)
		st.Insert(i)
		st.PrintTree()
	}

	// Usuwanie losowej permutacji
	fmt.Println("--- FAZA USUWANIA (losowa permutacja) ---")
	deleteOrder := generatePermutation(n)
	fmt.Printf("Kolejność usuwania: %v\n\n", deleteOrder)

	for _, key := range deleteOrder {
		fmt.Printf("delete %d\n", key)
		deleted := st.Delete(key)
		if !deleted {
			fmt.Printf("Klucz %d nie został znaleziony!\n", key)
		}
		st.PrintTree()
	}
}

// demonstrateCase2 demonstruje przypadek 2: wstawianie losowej permutacji, usuwanie losowej permutacji
func demonstrateCase2(n int) {
	fmt.Println("=== PRZYPADEK 2: Wstawianie losowej permutacji, usuwanie losowej permutacji ===")
	fmt.Printf("n = %d\n\n", n)

	st := NewSplayTree()

	// Wstawianie losowej permutacji
	fmt.Println("--- FAZA WSTAWIANIA (losowa permutacja) ---")
	insertOrder := generatePermutation(n)
	fmt.Printf("Kolejność wstawiania: %v\n\n", insertOrder)

	for _, key := range insertOrder {
		fmt.Printf("insert %d\n", key)
		st.Insert(key)
		st.PrintTree()
	}

	// Usuwanie losowej permutacji
	fmt.Println("--- FAZA USUWANIA (losowa permutacja) ---")
	deleteOrder := generatePermutation(n)
	fmt.Printf("Kolejność usuwania: %v\n\n", deleteOrder)

	for _, key := range deleteOrder {
		fmt.Printf("delete %d\n", key)
		deleted := st.Delete(key)
		if !deleted {
			fmt.Printf("Klucz %d nie został znaleziony!\n", key)
		}
		st.PrintTree()
	}
}

func main() {
	n := 30

	// Demonstracja przypadku 1
	demonstrateCase1(n)

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

	// Demonstracja przypadku 2
	demonstrateCase2(n)

	fmt.Println("=== KONIEC DEMONSTRACJI ===")
}
