package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
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

// Find znajduje korzeń zbioru z kompresją ścieżki
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union łączy dwa zbiory
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

// GenerateCompleteGraph generuje graf pełny o n wierzchołkach z losowymi wagami
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

// PrimMST implementuje algorytm Prima
func PrimMST(graph *Graph) ([]Edge, float64) {
	if graph.Vertices <= 1 {
		return []Edge{}, 0.0
	}

	mst := make([]Edge, 0, graph.Vertices-1)
	inMST := make([]bool, graph.Vertices)
	minWeight := make([]float64, graph.Vertices)
	parent := make([]int, graph.Vertices)

	// Inicjalizacja
	for i := 0; i < graph.Vertices; i++ {
		minWeight[i] = math.Inf(1)
		parent[i] = -1
	}

	// Rozpocznij od wierzchołka 0
	minWeight[0] = 0
	totalWeight := 0.0

	// Buduj listę sąsiedztwa
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
		// Znajdź wierzchołek o minimalnej wadze
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

		// Dodaj krawędź do MST (pomijając pierwszy wierzchołek)
		if parent[u] != -1 {
			mst = append(mst, Edge{
				From:   parent[u],
				To:     u,
				Weight: minWeight[u],
			})
			totalWeight += minWeight[u]
		}

		// Aktualizuj wagi sąsiadów
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

// KruskalMST implementuje algorytm Kruskala
func KruskalMST(graph *Graph) ([]Edge, float64) {
	if graph.Vertices <= 1 {
		return []Edge{}, 0.0
	}

	uf := NewUnionFind(graph.Vertices)

	S := make([]Edge, len(graph.Edges))
	copy(S, graph.Edges)

	mst := make([]Edge, 0, graph.Vertices-1)
	totalWeight := 0.0

	sort.Slice(S, func(i, j int) bool {
		return S[i].Weight < S[j].Weight
	})

	for len(S) > 0 && len(mst) < graph.Vertices-1 {
		edge := S[0]
		S = S[1:]
		if uf.Union(edge.From, edge.To) {
			mst = append(mst, edge)
			totalWeight += edge.Weight
		}
	}
	return mst, totalWeight
}

// MeasureTime mierzy czas wykonania algorytmu
func MeasureTime(algorithm func(*Graph) ([]Edge, float64), graph *Graph) time.Duration {
	start := time.Now()
	algorithm(graph)
	return time.Since(start)
}

// TestResult przechowuje wyniki testów
type TestResult struct {
	N           int
	PrimTime    time.Duration
	KruskalTime time.Duration
}

func RunExperiments(nMin, nMax, step, rep int) []TestResult {
	results := make([]TestResult, 0)

	fmt.Printf("Rozpoczynam eksperymenty...\n")
	fmt.Printf("nMin=%d, nMax=%d, step=%d, rep=%d\n\n", nMin, nMax, step, rep)

	for n := nMin; n <= nMax; n += step {
		fmt.Printf("Testowanie dla n=%d...\n", n)

		var totalPrimTime, totalKruskalTime time.Duration

		for i := 0; i < rep; i++ {
			graph := GenerateCompleteGraph(n)

			primTime := MeasureTime(PrimMST, graph)
			kruskalTime := MeasureTime(KruskalMST, graph)

			totalPrimTime += primTime
			totalKruskalTime += kruskalTime
		}

		avgPrimTime := totalPrimTime / time.Duration(rep)
		avgKruskalTime := totalKruskalTime / time.Duration(rep)

		results = append(results, TestResult{
			N:           n,
			PrimTime:    avgPrimTime,
			KruskalTime: avgKruskalTime,
		})

		fmt.Printf("  Prim: %v, Kruskal: %v\n", avgPrimTime, avgKruskalTime)
	}

	return results
}

// SaveResultsToCSV zapisuje wyniki do pliku CSV
func SaveResultsToCSV(results []TestResult, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "N,PrimTime(ns),KruskalTime(ns)\n")
	for _, result := range results {
		fmt.Fprintf(file, "%d,%d,%d\n",
			result.N,
			result.PrimTime.Nanoseconds(),
			result.KruskalTime.Nanoseconds())
	}

	return nil
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

func main() {
	nMin := 100
	nMax := 5000
	step := 100
	rep := 20

	fmt.Println("=== PROGRAM PORÓWNANIA ALGORYTMÓW MST ===")
	fmt.Println()

	// Przykład dla małego grafu
	fmt.Println("Przykład dla grafu o 5 wierzchołkach:")
	testGraph := GenerateCompleteGraph(5)

	primMST, primWeight := PrimMST(testGraph)
	kruskalMST, kruskalWeight := KruskalMST(testGraph)

	fmt.Printf("Algorytm Prima:\n")
	fmt.Printf("  Waga MST: %.6f\n", primWeight)
	fmt.Printf("  Poprawność: %t\n", VerifyMST(testGraph, primMST))

	fmt.Printf("Algorytm Kruskala:\n")
	fmt.Printf("  Waga MST: %.6f\n", kruskalWeight)
	fmt.Printf("  Poprawność: %t\n", VerifyMST(testGraph, kruskalMST))

	fmt.Printf("Różnica wag: %.10f\n", math.Abs(primWeight-kruskalWeight))
	fmt.Println()

	// Uruchom eksperymenty
	results := RunExperiments(nMin, nMax, step, rep)

	// Zapisz wyniki
	err := SaveResultsToCSV(results, "results.csv")
	if err != nil {
		fmt.Printf("Błąd przy zapisie wyników: %v\n", err)
	} else {
		fmt.Println("\nWyniki zapisane do pliku 'results.csv'")
	}

	// Wyświetl podsumowanie
	fmt.Println("\n=== PODSUMOWANIE WYNIKÓW ===")
	fmt.Printf("%-5s %-15s %-15s %-10s\n", "N", "Prim (μs)", "Kruskal (μs)", "Stosunek")
	fmt.Println(strings.Repeat("-", 50))

	for _, result := range results {
		primMicros := float64(result.PrimTime.Nanoseconds()) / 1000.0
		kruskalMicros := float64(result.KruskalTime.Nanoseconds()) / 1000.0
		ratio := primMicros / kruskalMicros

		fmt.Printf("%-5d %-15.2f %-15.2f %-10.2f\n",
			result.N, primMicros, kruskalMicros, ratio)
	}
}
