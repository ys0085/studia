package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
)

type Edge struct {
	From   int
	To     int
	Weight float64
}

type Graph struct {
	Vertices int
	Edges    []Edge
}

type UnionFind struct {
	parent []int
	rank   []int
}

type TreeNode struct {
	ID       int
	Children []*TreeNode
	Parent   *TreeNode
	Height   int
	Depth    int
}

type Tree struct {
	Root  *TreeNode
	Nodes map[int]*TreeNode
}

func GenerateCompleteGraph(n int) *Graph {
	graph := &Graph{
		Vertices: n,
		Edges:    make([]Edge, 0, n*(n-1)/2),
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			weight := rand.Float64()
			graph.Edges = append(graph.Edges, Edge{
				From:   i,
				To:     j,
				Weight: weight,
			})
		}
	}

	return graph
}

func PrimMST(graph *Graph) ([]Edge, float64) {
	if graph.Vertices <= 1 {
		return []Edge{}, 0.0
	}

	mst := make([]Edge, 0, graph.Vertices-1)
	inMST := make([]bool, graph.Vertices)
	minWeight := make([]float64, graph.Vertices)
	parent := make([]int, graph.Vertices)

	for i := 0; i < graph.Vertices; i++ {
		minWeight[i] = math.Inf(1)
		parent[i] = -1
	}

	minWeight[0] = 0
	totalWeight := 0.0

	adjList := make([][]Edge, graph.Vertices)
	for _, edge := range graph.Edges {
		adjList[edge.From] = append(adjList[edge.From], edge)
		adjList[edge.To] = append(adjList[edge.To], Edge{
			From:   edge.To,
			To:     edge.From,
			Weight: edge.Weight,
		})
	}

	for count := 0; count < graph.Vertices; count++ {
		u := -1
		for v := 0; v < graph.Vertices; v++ {
			if !inMST[v] && (u == -1 || minWeight[v] < minWeight[u]) {
				u = v
			}
		}

		if u == -1 {
			break
		}

		inMST[u] = true

		if parent[u] != -1 {
			mst = append(mst, Edge{
				From:   parent[u],
				To:     u,
				Weight: minWeight[u],
			})
			totalWeight += minWeight[u]
		}

		for _, edge := range adjList[u] {
			v := edge.To
			if !inMST[v] && edge.Weight < minWeight[v] {
				minWeight[v] = edge.Weight
				parent[v] = u
			}
		}
	}

	return mst, totalWeight
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}

// VerifyMST weryfikuje poprawność MST
func VerifyMST(graph *Graph, mst []Edge) bool {
	if len(mst) != graph.Vertices-1 {
		return false
	}

	// Sprawdź czy MST jest drzewem (graf spójny bez cykli)
	uf := NewUnionFind(graph.Vertices)
	for _, edge := range mst {
		if !uf.Union(edge.From, edge.To) {
			return false // Cykl znaleziony
		}
	}

	return true
}

// Funkcja pomocnicza do budowania drzewa z MST
func BuildTreeFromMST(mst []Edge, vertices int, rootID int) *Tree {
	nodes := make(map[int]*TreeNode)

	// Inicjalizuj węzły
	for i := 0; i < vertices; i++ {
		nodes[i] = &TreeNode{
			ID:       i,
			Children: make([]*TreeNode, 0),
			Parent:   nil,
			Height:   0,
			Depth:    0,
		}
	}

	// Buduj listę sąsiedztwa
	adjList := make([][]int, vertices)
	for _, edge := range mst {
		adjList[edge.From] = append(adjList[edge.From], edge.To)
		adjList[edge.To] = append(adjList[edge.To], edge.From)
	}

	// BFS do budowania drzewa z korzeniem rootID
	visited := make([]bool, vertices)
	queue := []int{rootID}
	visited[rootID] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range adjList[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				nodes[neighbor].Parent = nodes[current]
				nodes[current].Children = append(nodes[current].Children, nodes[neighbor])
				nodes[neighbor].Depth = nodes[current].Depth + 1
				queue = append(queue, neighbor)
			}
		}
	}

	tree := &Tree{
		Root:  nodes[rootID],
		Nodes: nodes,
	}

	// Oblicz wysokości węzłów
	calculateHeights(tree.Root)

	return tree
}

// Oblicza wysokość każdego węzła (maksymalna odległość do liścia)
func calculateHeights(node *TreeNode) int {
	if len(node.Children) == 0 {
		node.Height = 0
		return 0
	}

	maxHeight := 0
	for _, child := range node.Children {
		childHeight := calculateHeights(child)
		if childHeight > maxHeight {
			maxHeight = childHeight
		}
	}

	node.Height = maxHeight + 1
	return node.Height
}

// OptimalBroadcastOrder wyznacza optymalną kolejność informowania dzieci
func OptimalBroadcastOrder(tree *Tree) map[int][]int {
	order := make(map[int][]int)

	// Dla każdego węzła sortuj dzieci według wysokości malejąco
	for _, node := range tree.Nodes {
		if len(node.Children) > 0 {
			// Utwórz kopię listy dzieci
			children := make([]*TreeNode, len(node.Children))
			copy(children, node.Children)

			// Sortuj dzieci według wysokości malejąco
			sort.Slice(children, func(i, j int) bool {
				return children[i].Height > children[j].Height
			})

			// Zapisz kolejność ID dzieci
			childOrder := make([]int, len(children))
			for i, child := range children {
				childOrder[i] = child.ID
			}
			order[node.ID] = childOrder
		}
	}

	return order
}

// CalculateBroadcastRounds oblicza liczbę rund potrzebną do rozprzestrzenienia informacji
func CalculateBroadcastRounds(tree *Tree, order map[int][]int) int {
	roundsToReach := make(map[int]int)
	roundsToReach[tree.Root.ID] = 0

	// BFS z uwzględnieniem kolejności informowania
	queue := []*TreeNode{tree.Root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if childOrder, exists := order[current.ID]; exists {
			currentRound := roundsToReach[current.ID]

			for i, childID := range childOrder {
				child := tree.Nodes[childID]
				roundsToReach[childID] = currentRound + i + 1
				queue = append(queue, child)
			}
		}
	}

	// Znajdź maksymalną liczbę rund
	maxRounds := 0
	for _, rounds := range roundsToReach {
		if rounds > maxRounds {
			maxRounds = rounds
		}
	}

	return maxRounds
}

// SimulateBroadcast symuluje proces rozprzestrzeniania informacji
func SimulateBroadcast(tree *Tree, order map[int][]int) {
	fmt.Printf("=== Symulacja rozprzestrzeniania informacji ===\n")
	fmt.Printf("Korzeń drzewa: %d\n", tree.Root.ID)

	informed := make(map[int]bool)
	informed[tree.Root.ID] = true
	round := 0

	fmt.Printf("Runda %d: Węzeł %d posiada informację\n", round, tree.Root.ID)

	for {
		round++
		newlyInformed := make([]int, 0)

		// Dla każdego poinformowanego węzła
		for nodeID := range informed {
			if childOrder, exists := order[nodeID]; exists {
				// Sprawdź czy może poinformować kolejne dziecko
				for _, childID := range childOrder {
					if !informed[childID] {
						// Informuj pierwsze nieopoinformowane dziecko
						informed[childID] = true
						newlyInformed = append(newlyInformed, childID)
						break
					}
				}
			}
		}

		if len(newlyInformed) == 0 {
			break
		}

		fmt.Printf("Runda %d: Poinformowano węzły %v\n", round, newlyInformed)
	}

	fmt.Printf("Całkowita liczba rund: %d\n", round-1)
}

// ExperimentResult przechowuje wyniki pojedynczego eksperymentu
type ExperimentResult struct {
	N                 int
	Repetition        int
	MSTWeight         float64
	OptimalRounds     int
	OptimalRoot       int
	AverageRounds     float64
	TestedRoots       int
	MinRoundsAllRoots int
	MaxRoundsAllRoots int
}

// AggregatedResults przechowuje zagregowane wyniki dla danego n
type AggregatedResults struct {
	N                   int
	Repetitions         int
	AvgOptimalRounds    float64
	StdDevOptimalRounds float64
	MinOptimalRounds    int
	MaxOptimalRounds    int
	AvgMSTWeight        float64
	AvgTestedRoots      float64
	TotalExperiments    int
}

// RunExperiments wykonuje systematyczne eksperymenty average case analysis
func RunExperiments(nMin, nMax, step, rep int) []AggregatedResults {
	fmt.Printf("=== AVERAGE CASE ANALYSIS ===\n")
	fmt.Printf("Parametry: n ∈ [%d, %d], krok = %d, powtórzeń = %d\n", nMin, nMax, step, rep)
	fmt.Printf("%-8s %-10s %-8s %-8s %-8s %-10s %-10s %-8s\n",
		"n", "Średnia", "Odch.Std", "Min", "Max", "Śr.Waga", "Śr.Korzeni", "Eksper.")
	fmt.Println(strings.Repeat("-", 80))

	var allResults []AggregatedResults

	for n := nMin; n <= nMax; n += step {
		var experiments []ExperimentResult

		// Wykonaj rep eksperymentów dla danego n
		for r := 0; r < rep; r++ {
			result := runSingleExperiment(n, r)
			experiments = append(experiments, result)
		}

		// Oblicz statystyki zagregowane
		aggregated := aggregateResults(experiments)
		allResults = append(allResults, aggregated)

		// Wyświetl wyniki
		fmt.Printf("%-8d %-10.2f %-8.2f %-8d %-8d %-10.3f %-10.1f %-8d\n",
			aggregated.N,
			aggregated.AvgOptimalRounds,
			aggregated.StdDevOptimalRounds,
			aggregated.MinOptimalRounds,
			aggregated.MaxOptimalRounds,
			aggregated.AvgMSTWeight,
			aggregated.AvgTestedRoots,
			aggregated.TotalExperiments)
	}

	return allResults
}

// runSingleExperiment wykonuje pojedynczy eksperyment dla danego n
func runSingleExperiment(n, repetition int) ExperimentResult {
	// Generuj losowy graf pełny
	graph := GenerateCompleteGraph(n)

	// Znajdź MST używając algorytmu Prima
	mst, mstWeight := PrimMST(graph)

	if !VerifyMST(graph, mst) {
		fmt.Printf("BŁĄD: Niepoprawne MST dla n=%d, rep=%d\n", n, repetition)
		return ExperimentResult{}
	}

	// Testuj różne korzenie
	maxRootsToTest := n
	if n > 50 { // Ograniczenie dla większych grafów
		maxRootsToTest = min(50, n)
	}

	bestRounds := math.MaxInt32
	bestRoot := 0
	totalRounds := 0
	minRounds := math.MaxInt32
	maxRounds := 0

	for root := 0; root < maxRootsToTest; root++ {
		tree := BuildTreeFromMST(mst, n, root)
		order := OptimalBroadcastOrder(tree)
		rounds := CalculateBroadcastRounds(tree, order)

		totalRounds += rounds
		if rounds < minRounds {
			minRounds = rounds
		}
		if rounds > maxRounds {
			maxRounds = rounds
		}

		if rounds < bestRounds {
			bestRounds = rounds
			bestRoot = root
		}
	}

	avgRounds := float64(totalRounds) / float64(maxRootsToTest)

	return ExperimentResult{
		N:                 n,
		Repetition:        repetition,
		MSTWeight:         mstWeight,
		OptimalRounds:     bestRounds,
		OptimalRoot:       bestRoot,
		AverageRounds:     avgRounds,
		TestedRoots:       maxRootsToTest,
		MinRoundsAllRoots: minRounds,
		MaxRoundsAllRoots: maxRounds,
	}
}

func aggregateResults(experiments []ExperimentResult) AggregatedResults {
	if len(experiments) == 0 {
		return AggregatedResults{}
	}

	n := experiments[0].N
	sumOptimalRounds := 0
	sumMSTWeight := 0.0
	sumTestedRoots := 0
	minOptimal := math.MaxInt32
	maxOptimal := 0

	// Oblicz sumy
	for _, exp := range experiments {
		sumOptimalRounds += exp.OptimalRounds
		sumMSTWeight += exp.MSTWeight
		sumTestedRoots += exp.TestedRoots

		if exp.OptimalRounds < minOptimal {
			minOptimal = exp.OptimalRounds
		}
		if exp.OptimalRounds > maxOptimal {
			maxOptimal = exp.OptimalRounds
		}
	}

	// Oblicz średnie
	count := len(experiments)
	avgOptimal := float64(sumOptimalRounds) / float64(count)
	avgMSTWeight := sumMSTWeight / float64(count)
	avgTestedRoots := float64(sumTestedRoots) / float64(count)

	// Oblicz odchylenie standardowe
	sumSquaredDiffs := 0.0
	for _, exp := range experiments {
		diff := float64(exp.OptimalRounds) - avgOptimal
		sumSquaredDiffs += diff * diff
	}
	stdDev := math.Sqrt(sumSquaredDiffs / float64(count))

	return AggregatedResults{
		N:                   n,
		Repetitions:         count,
		AvgOptimalRounds:    avgOptimal,
		StdDevOptimalRounds: stdDev,
		MinOptimalRounds:    minOptimal,
		MaxOptimalRounds:    maxOptimal,
		AvgMSTWeight:        avgMSTWeight,
		AvgTestedRoots:      avgTestedRoots,
		TotalExperiments:    count,
	}
}

// min helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	// Parametry eksperymentów
	nMin := 100
	nMax := 5000
	step := 100
	rep := 20

	fmt.Println("=== PROGRAM ANALIZY ALGORYTMU ROZPRZESTRZENIANIA INFORMACJI ===")

	results := RunExperiments(nMin, nMax, step, rep)

	fmt.Printf("\n=== DEMONSTRACJA ALGORYTMU ===\n")
	graph := GenerateCompleteGraph(8)
	_, weight := PrimMST(graph)

	fmt.Printf("Graf demonstracyjny: 8 wierzchołków, MST waga: %.3f\n", weight)

	// Opcjonalnie: zapisz wyniki do dalszej analizy
	fmt.Printf("\n=== PODSUMOWANIE EKSPERYMENTÓW ===\n")
	fmt.Printf("Przeprowadzono łącznie %d eksperymentów\n", len(results)*rep)
	fmt.Printf("Zakres n: [%d, %d] z krokiem %d\n", nMin, nMax, step)
	fmt.Printf("Liczba powtórzeń na każde n: %d\n", rep)

	if len(results) > 0 {
		minAvg := results[0].AvgOptimalRounds
		maxAvg := results[len(results)-1].AvgOptimalRounds
		improvementRatio := maxAvg / minAvg

		fmt.Printf("Średnia liczba rund:\n")
		fmt.Printf("  Dla n=%d: %.2f\n", results[0].N, minAvg)
		fmt.Printf("  Dla n=%d: %.2f\n", results[len(results)-1].N, maxAvg)
		fmt.Printf("  Stosunek wzrostu: %.2fx\n", improvementRatio)
	}
}
