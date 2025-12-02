package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// Node reprezentuje węzeł w drzewie BST
type Node struct {
	Key   int
	Count int // liczba wystąpień klucza
	Left  *Node
	Right *Node
}

// BST reprezentuje drzewo wyszukiwań binarnych
type BST struct {
	Root *Node
}

// NewBST tworzy nowe puste drzewo BST
func NewBST() *BST {
	return &BST{Root: nil}
}

// Insert wstawia klucz do drzewa (obsługuje duplikaty przez licznik)
func (bst *BST) Insert(key int) {
	bst.Root = bst.insertNode(bst.Root, key)
}

func (bst *BST) insertNode(node *Node, key int) *Node {
	// Jeśli węzeł jest nil, tworzymy nowy węzeł
	if node == nil {
		return &Node{Key: key, Count: 1, Left: nil, Right: nil}
	}

	// Jeśli klucz już istnieje, zwiększamy licznik
	if key == node.Key {
		node.Count++
		return node
	}

	// Rekurencyjnie wstawiamy w odpowiednie poddrzewo
	if key < node.Key {
		node.Left = bst.insertNode(node.Left, key)
	} else {
		node.Right = bst.insertNode(node.Right, key)
	}

	return node
}

// Delete usuwa jedno wystąpienie klucza z drzewa
func (bst *BST) Delete(key int) bool {
	var deleted bool
	bst.Root, deleted = bst.deleteNode(bst.Root, key)
	return deleted
}

func (bst *BST) deleteNode(node *Node, key int) (*Node, bool) {
	if node == nil {
		return nil, false
	}

	var deleted bool

	if key < node.Key {
		node.Left, deleted = bst.deleteNode(node.Left, key)
	} else if key > node.Key {
		node.Right, deleted = bst.deleteNode(node.Right, key)
	} else {
		// Znaleźliśmy węzeł do usunięcia
		deleted = true

		// Jeśli jest więcej niż jedno wystąpienie, zmniejszamy licznik
		if node.Count > 1 {
			node.Count--
			return node, deleted
		}

		// Usuwamy węzeł - przypadki standardowe dla BST
		if node.Left == nil {
			return node.Right, deleted
		}
		if node.Right == nil {
			return node.Left, deleted
		}

		// Węzeł ma dwoje dzieci - znajdź następnik (najmniejszy w prawym poddrzewie)
		successor := bst.findMin(node.Right)
		node.Key = successor.Key
		node.Count = successor.Count

		// Usuń następnik z prawego poddrzewa
		node.Right, _ = bst.deleteNode(node.Right, successor.Key)
	}

	return node, deleted
}

// findMin znajduje węzeł z najmniejszym kluczem w poddrzewie
func (bst *BST) findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// Height zwraca wysokość drzewa
func (bst *BST) Height() int {
	return bst.getHeight(bst.Root)
}

func (bst *BST) getHeight(node *Node) int {
	if node == nil {
		return -1 // Wysokość pustego drzewa to -1
	}

	leftHeight := bst.getHeight(node.Left)
	rightHeight := bst.getHeight(node.Right)

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

// PrintTree wyświetla drzewo w czytelnej formie
func (bst *BST) PrintTree() {
	if bst.Root == nil {
		fmt.Println("Drzewo jest puste")
		return
	}

	fmt.Println("Struktura drzewa:")
	bst.printNode(bst.Root, "", true)
	fmt.Printf("Wysokość drzewa: %d\n", bst.Height())
	fmt.Println()
}

func (bst *BST) printNode(node *Node, prefix string, isLast bool) {
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
		bst.printNode(node.Right, childPrefix, !hasLeft)
	}
	if hasLeft {
		bst.printNode(node.Left, childPrefix, true)
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

	bst := NewBST()

	// Wstawianie rosnącego ciągu
	fmt.Println("--- FAZA WSTAWIANIA (rosnący ciąg) ---")
	for i := 1; i <= n; i++ {
		fmt.Printf("insert %d\n", i)
		bst.Insert(i)
		bst.PrintTree()
	}

	// Usuwanie losowej permutacji
	fmt.Println("--- FAZA USUWANIA (losowa permutacja) ---")
	deleteOrder := generatePermutation(n)
	fmt.Printf("Kolejność usuwania: %v\n\n", deleteOrder)

	for _, key := range deleteOrder {
		fmt.Printf("delete %d\n", key)
		deleted := bst.Delete(key)
		if !deleted {
			fmt.Printf("Klucz %d nie został znaleziony!\n", key)
		}
		bst.PrintTree()
	}
}

// demonstrateCase2 demonstruje przypadek 2: wstawianie losowej permutacji, usuwanie losowej permutacji
func demonstrateCase2(n int) {
	fmt.Println("=== PRZYPADEK 2: Wstawianie losowej permutacji, usuwanie losowej permutacji ===")
	fmt.Printf("n = %d\n\n", n)

	bst := NewBST()

	// Wstawianie losowej permutacji
	fmt.Println("--- FAZA WSTAWIANIA (losowa permutacja) ---")
	insertOrder := generatePermutation(n)
	fmt.Printf("Kolejność wstawiania: %v\n\n", insertOrder)

	for _, key := range insertOrder {
		fmt.Printf("insert %d\n", key)
		bst.Insert(key)
		bst.PrintTree()
	}

	// Usuwanie losowej permutacji
	fmt.Println("--- FAZA USUWANIA (losowa permutacja) ---")
	deleteOrder := generatePermutation(n)
	fmt.Printf("Kolejność usuwania: %v\n\n", deleteOrder)

	for _, key := range deleteOrder {
		fmt.Printf("delete %d\n", key)
		deleted := bst.Delete(key)
		if !deleted {
			fmt.Printf("Klucz %d nie został znaleziony!\n", key)
		}
		bst.PrintTree()
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
