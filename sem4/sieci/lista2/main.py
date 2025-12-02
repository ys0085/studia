import numpy as np
import networkx as nx
import matplotlib.pyplot as plt
import random
from itertools import combinations
from tqdm import tqdm

class NetworkSimulation:
    def __init__(self, num_vertices=20, max_edges=29, p=0.95, T_max=0.1, m=1000):
        """
        Inicjalizacja modelu sieci
        
        Parametry:
        - num_vertices: liczba wierzchołków w grafie (|V|)
        - max_edges: maksymalna liczba krawędzi w grafie (|E|)
        - p: prawdopodobieństwo nieuszkodzenia krawędzi
        - T_max: maksymalne dopuszczalne opóźnienie
        - m: średnia wielkość pakietu w bitach
        """
        self.num_vertices = num_vertices
        self.max_edges = max_edges
        self.p = p
        self.T_max = T_max
        self.m = m
        self.G = None  # Graf
        self.N = None  # Macierz natężeń
        
        # Stworzenie grafu
        self.create_graph()
        
        # Stworzenie macierzy natężeń
        self.create_intensity_matrix()
        
        # Przypisanie przepustowości i przepływów
        self.assign_capacities_and_flows()
        
    def create_graph(self):
        """Tworzy graf o określonej liczbie wierzchołków i krawędzi"""
        # Tworzenie grafu nieskierowanego
        self.G = nx.Graph()
        
        # Dodawanie wierzchołków
        self.G.add_nodes_from(range(self.num_vertices))
        
        # Dodawanie krawędzi zapewniających spójność (minimalne drzewo rozpinające)
        for i in range(1, self.num_vertices):
            self.G.add_edge(i-1, i)
        
        # Dodawanie losowych krawędzi do osiągnięcia pożądanej liczby
        possible_edges = list(combinations(range(self.num_vertices), 2))
        existing_edges = list(self.G.edges())
        possible_edges = [e for e in possible_edges if e not in existing_edges]
        
        # Wylosuj dodatkowe krawędzie
        additional_edges = random.sample(possible_edges, 
                                         min(self.max_edges - (self.num_vertices - 1), 
                                             len(possible_edges)))
        self.G.add_edges_from(additional_edges)
        
        print(f"Utworzono graf o {self.G.number_of_nodes()} wierzchołkach i {self.G.number_of_edges()} krawędziach")
    
    def create_intensity_matrix(self):
        """Tworzy macierz natężeń strumienia pakietów"""
        self.N = np.zeros((self.num_vertices, self.num_vertices))
        
        # Wypełnienie macierzy N losowymi wartościami - niższymi niż poprzednio
        for i in range(self.num_vertices):
            for j in range(self.num_vertices):
                if i != j:  # Nie przesyłamy pakietów do samego siebie
                    # Losowa wartość natężenia od 0.001 do 0.03 pakietów/s (10x mniejsze)
                    self.N[i, j] = round(random.uniform(0.001, 0.03), 3)
        
        print(f"Utworzono macierz natężeń o sumie {np.sum(self.N):.3f} pakietów/s")
    
    def assign_capacities_and_flows(self):
        """Przypisuje przepustowości i przepływy do krawędzi grafu"""
        # Obliczenie przepływów na podstawie najkrótszych ścieżek
        self.compute_flows()
        
        # Przypisanie przepustowości na podstawie przepływów
        for u, v in self.G.edges():
            flow = self.G[u][v]['flow']
            # Przepustowość musi być znacząco większa od przepływu
            capacity = max(1000, int(flow * self.m * 5))  # W bitach/s
            self.G[u][v]['capacity'] = capacity
        
        # Sprawdzenie warunku c(e) > a(e)*m dla każdej krawędzi
        for u, v in self.G.edges():
            flow = self.G[u][v]['flow']
            capacity = self.G[u][v]['capacity']
            if capacity / self.m <= flow:
                print(f"Ostrzeżenie: Warunek c(e)/m > a(e) nie jest spełniony dla krawędzi ({u}, {v})")
                # Zwiększenie przepustowości, aby warunek był spełniony
                self.G[u][v]['capacity'] = int(flow * self.m * 10)
                print(f"Zwiększono przepustowość do {self.G[u][v]['capacity']} dla krawędzi ({u}, {v})")
    
    def compute_flows(self):
        """
        Oblicza przepływy na krawędziach na podstawie macierzy natężeń
        i najkrótszych ścieżek
        """
        # Inicjalizacja przepływów na krawędziach
        for u, v in self.G.edges():
            self.G[u][v]['flow'] = 0.0
        
        # Dla każdej pary (źródło, ujście)
        for i in range(self.num_vertices):
            for j in range(self.num_vertices):
                if i != j and self.N[i, j] > 0:
                    try:
                        # Znajdź najkrótszą ścieżkę
                        path = nx.shortest_path(self.G, source=i, target=j)
                        
                        # Dodaj przepływ do każdej krawędzi na ścieżce
                        for idx in range(len(path) - 1):
                            u, v = path[idx], path[idx + 1]
                            if self.G.has_edge(u, v):
                                self.G[u][v]['flow'] += self.N[i, j]
                            else:
                                self.G[v][u]['flow'] += self.N[i, j]
                    except nx.NetworkXNoPath:
                        print(f"Nie znaleziono ścieżki z {i} do {j}")
    
    def calculate_average_delay(self, graph=None):
        """
        Oblicza średnie opóźnienie pakietu w sieci według wzoru:
        T = 1/G * sum_e(a(e)/(c(e)/m - a(e)))
        
        Parametry:
        - graph: graf dla którego obliczamy opóźnienie (domyślnie self.G)
        """
        if graph is None:
            graph = self.G
            
        G_sum = np.sum(self.N)  # Suma wszystkich elementów macierzy natężeń
        
        delay_sum = 0.0
        for u, v in graph.edges():
            flow = graph[u][v]['flow']
            capacity = graph[u][v]['capacity']
            
            # Sprawdzenie, czy kanał nie jest przeciążony
            if capacity / self.m <= flow:
                return float('inf')  # Nieskończone opóźnienie
                
            if flow > 0:  # Unikamy dzielenia przez zero
                delay_sum += flow / (capacity / self.m - flow)
        
        if G_sum > 0:
            return delay_sum / G_sum
        else:
            return 0.0
    
    def is_network_connected(self, available_edges):
        """Sprawdza, czy sieć jest spójna dla danego zestawu dostępnych krawędzi"""
        # Tworzenie tymczasowego grafu z dostępnymi krawędziami
        temp_graph = nx.Graph()
        temp_graph.add_nodes_from(range(self.num_vertices))
        temp_graph.add_edges_from(available_edges)
        
        # Sprawdzenie, czy graf jest spójny
        return nx.is_connected(temp_graph)
    
    def simulate_network_state(self):
        """
        Symuluje jeden stan sieci z losowo uszkodzonymi krawędziami
        i zwraca informacje czy sieć spełnia warunek T < T_max
        """
        # Losowe uszkodzenie krawędzi
        available_edges = []
        for u, v in self.G.edges():
            if random.random() < self.p:  # Krawędź jest nieuszkodzona
                available_edges.append((u, v))
        
        # Sprawdzenie, czy sieć jest spójna
        if not self.is_network_connected(available_edges):
            # Sieć jest rozspójniona, nie liczymy tego przypadku
            return None
        
        # Tworzenie tymczasowego grafu z dostępnymi krawędziami
        temp_graph = nx.Graph()
        temp_graph.add_nodes_from(range(self.num_vertices))
        
        for u, v in available_edges:
            temp_graph.add_edge(u, v)
            temp_graph[u][v]['flow'] = self.G[u][v]['flow']
            temp_graph[u][v]['capacity'] = self.G[u][v]['capacity']
        
        # Aktualizacja przepływów w temp_graph
        # Musimy przekierować przepływy na dostępne ścieżki
        self.update_flows_for_temp_graph(temp_graph)
        
        # Obliczenie opóźnienia
        delay = self.calculate_average_delay(temp_graph)
        
        # Sprawdzenie warunku T < T_max
        return delay < self.T_max
    
    def update_flows_for_temp_graph(self, temp_graph):
        """Aktualizuje przepływy w tymczasowym grafie po uszkodzeniu niektórych krawędzi"""
        # Resetowanie przepływów
        for u, v in temp_graph.edges():
            temp_graph[u][v]['flow'] = 0.0
        
        # Dla każdej pary (źródło, ujście)
        for i in range(self.num_vertices):
            for j in range(self.num_vertices):
                if i != j and self.N[i, j] > 0:
                    try:
                        # Znajdź najkrótszą ścieżkę w tymczasowym grafie
                        path = nx.shortest_path(temp_graph, source=i, target=j)
                        
                        # Dodaj przepływ do każdej krawędzi na ścieżce
                        for idx in range(len(path) - 1):
                            u, v = path[idx], path[idx + 1]
                            if temp_graph.has_edge(u, v):
                                temp_graph[u][v]['flow'] += self.N[i, j]
                            else:
                                temp_graph[v][u]['flow'] += self.N[i, j]
                    except nx.NetworkXNoPath:
                        # Ignorujemy, bo sieć jest spójna, więc zawsze powinna być ścieżka
                        pass
    
    def estimate_reliability(self, num_trials=1000):
        """
        Szacuje niezawodność sieci jako prawdopodobieństwo, że nierozspójniona sieć
        zachowuje T < T_max przy losowym uszkodzeniu krawędzi z prawdopodobieństwem 1-p
        """
        success_count = 0
        valid_trials = 0
        
        for _ in tqdm(range(num_trials), desc="Symulacja niezawodności"):
            result = self.simulate_network_state()
            
            if result is not None:  # Jeśli sieć jest spójna
                valid_trials += 1
                if result:  # Jeśli T < T_max
                    success_count += 1
        
        # Zabezpieczenie przed dzieleniem przez zero
        if valid_trials == 0:
            print("Ostrzeżenie: Nie było żadnych ważnych prób (sieć zawsze była rozspójniona)")
            return 0
        
        reliability = success_count / valid_trials
        print(f"Trials: {valid_trials}, Success: {success_count}, Ratio: {reliability:.4f}")
        return reliability
    
    def simulate_increasing_intensities(self, min_factor=1.0, max_factor=3.0, steps=10, trials=500):
        """
        Symuluje wpływ zwiększania natężeń na niezawodność sieci
        """
        results = []
        factors = np.linspace(min_factor, max_factor, steps)
        
        original_N = self.N.copy()
        
        for factor in factors:
            # Skalowanie macierzy natężeń
            self.N = original_N * factor
            
            # Aktualizacja przepływów
            self.compute_flows()
            
            # Obliczenie niezawodności
            reliability = self.estimate_reliability(num_trials=trials)
            results.append((factor, reliability))
            
            print(f"Współczynnik natężeń: {factor:.2f}, Niezawodność: {reliability:.4f}")
        
        # Przywrócenie oryginalnej macierzy natężeń
        self.N = original_N
        self.compute_flows()
        
        return results
    
    def simulate_increasing_capacities(self, min_factor=1.0, max_factor=3.0, steps=10, trials=500):
        """
        Symuluje wpływ zwiększania przepustowości na niezawodność sieci
        """
        results = []
        factors = np.linspace(min_factor, max_factor, steps)
        
        # Zapisz oryginalne przepustowości
        original_capacities = {}
        for u, v in self.G.edges():
            original_capacities[(u, v)] = self.G[u][v]['capacity']
        
        for factor in factors:
            # Skalowanie przepustowości
            for u, v in self.G.edges():
                self.G[u][v]['capacity'] = int(original_capacities[(u, v)] * factor)
            
            # Obliczenie niezawodności
            reliability = self.estimate_reliability(num_trials=trials)
            results.append((factor, reliability))
            
            print(f"Współczynnik przepustowości: {factor:.2f}, Niezawodność: {reliability:.4f}")
        
        # Przywrócenie oryginalnych przepustowości
        for u, v in self.G.edges():
            self.G[u][v]['capacity'] = original_capacities[(u, v)]
        
        return results
    
    def simulate_adding_edges(self, max_additional_edges=10, trials=500):
        """
        Symuluje wpływ dodawania nowych krawędzi na niezawodność sieci
        """
        results = []
        
        # Oblicz średnią przepustowość w sieci początkowej
        total_capacity = sum(self.G[u][v]['capacity'] for u, v in self.G.edges())
        avg_capacity = total_capacity / self.G.number_of_edges()
        
        # Utwórz kopię grafu
        original_graph = self.G.copy()
        
        # Znajdź możliwe dodatkowe krawędzie
        possible_edges = list(combinations(range(self.num_vertices), 2))
        possible_edges = [e for e in possible_edges if not original_graph.has_edge(*e)]
        
        # Ograniczenie liczby dodatkowych krawędzi
        max_additional_edges = min(max_additional_edges, len(possible_edges))
        
        # Początkowa niezawodność
        initial_reliability = self.estimate_reliability(num_trials=trials)
        results.append((0, initial_reliability))
        print(f"Początkowa niezawodność: {initial_reliability:.4f}")
        
        # Dodawanie krawędzi po kolei
        for i in range(1, max_additional_edges + 1):
            # Dodaj nową krawędź
            if len(possible_edges) > 0:
                new_edge = random.choice(possible_edges)
                possible_edges.remove(new_edge)
                
                u, v = new_edge
                self.G.add_edge(u, v)
                self.G[u][v]['capacity'] = int(avg_capacity)
                self.G[u][v]['flow'] = 0.0  # Początkowo bez przepływu
                
                # Aktualizacja przepływów
                self.compute_flows()
                
                # Obliczenie niezawodności
                reliability = self.estimate_reliability(num_trials=trials)
                results.append((i, reliability))
                
                print(f"Dodano {i} krawędzi, Niezawodność: {reliability:.4f}")
        
        # Przywróć oryginalny graf
        self.G = original_graph
        
        return results
    
    def visualize_graph(self, title="Topologia sieci"):
        """Wizualizuje graf z przepustowościami i przepływami"""
        plt.figure(figsize=(12, 8))
        
        # Pozyskaj pozycje wierzchołków dla rysowania
        pos = nx.spring_layout(self.G, seed=42)
        
        # Rysuj wierzchołki
        nx.draw_networkx_nodes(self.G, pos, node_size=500, node_color='lightblue')
        
        # Rysuj krawędzie z różnymi szerokościami bazując na przepływach
        edge_width = [max(self.G[u][v]['flow'] * 5 + 0.5, 1) for u, v in self.G.edges()]
        nx.draw_networkx_edges(self.G, pos, width=edge_width, edge_color='gray')
        
        # Rysuj etykiety wierzchołków
        nx.draw_networkx_labels(self.G, pos, font_size=10)
        
        # Przygotuj etykiety krawędzi
        edge_labels = {(u, v): f"c={self.G[u][v]['capacity']/1000:.1f}k\na={self.G[u][v]['flow']:.3f}" 
                      for u, v in self.G.edges()}
        nx.draw_networkx_edge_labels(self.G, pos, edge_labels=edge_labels, font_size=8)
        
        plt.title(title)
        plt.axis('off')
        plt.tight_layout()
        plt.show()
    
    def plot_results(self, results, x_label, y_label="Niezawodność", title="Wyniki symulacji"):
        """Rysuje wykres wyników symulacji"""
        plt.figure(figsize=(10, 6))
        x_values, y_values = zip(*results)
        plt.plot(x_values, y_values, 'o-', linewidth=2)
        plt.xlabel(x_label)
        plt.ylabel(y_label)
        plt.title(title)
        plt.grid(True)
        plt.tight_layout()
        plt.show()

# Przykładowe użycie
def main():
    random.seed(42)
    np.random.seed(42)
    
    # Inicjalizacja modelu sieci z parametrami
    # - Mniejsze T_max żeby sieć mogła łatwiej nie spełniać warunku
    # - Większe m dla większego obciążenia sieci
    sim = NetworkSimulation(num_vertices=20, max_edges=29, p=0.95, T_max=0.05, m=2000)
    
    # Wizualizacja grafu
    sim.visualize_graph()
    
    # Obliczenie aktualnego opóźnienia
    delay = sim.calculate_average_delay()
    print(f"Średnie opóźnienie w sieci: {delay:.4f} s")
    print(f"T_max: {sim.T_max:.4f} s")
    
    # Oszacowanie niezawodności
    reliability = sim.estimate_reliability(num_trials=500)
    print(f"Oszacowanie niezawodności sieci: {reliability:.4f}")
    
    # Jeśli niezawodność jest nadal 0, ustawmy lepsze parametry
    if reliability < 0.01:
        print("Niezawodność jest nadal bliska zeru. Dostosowywanie parametrów...")
        sim.T_max = delay * 1.5  # Ustawienie T_max jako 150% aktualnego opóźnienia
        print(f"Ustawiono nowy T_max: {sim.T_max:.4f} s")
        reliability = sim.estimate_reliability(num_trials=500)
        print(f"Nowa niezawodność sieci: {reliability:.4f}")
    
    # Symulacja zwiększania natężeń
    intensity_results = sim.simulate_increasing_intensities(min_factor=0.5, max_factor=1.5, steps=5, trials=300)
    sim.plot_results(intensity_results, x_label="Współczynnik natężeń", 
                   title="Wpływ zwiększania natężeń na niezawodność sieci")
    
    # Symulacja zwiększania przepustowości
    capacity_results = sim.simulate_increasing_capacities(min_factor=1.0, max_factor=2.0, steps=5, trials=300)
    sim.plot_results(capacity_results, x_label="Współczynnik przepustowości", 
                   title="Wpływ zwiększania przepustowości na niezawodność sieci")
    
    # Symulacja dodawania krawędzi
    edge_results = sim.simulate_adding_edges(max_additional_edges=5, trials=300)
    sim.plot_results(edge_results, x_label="Liczba dodanych krawędzi", 
                   title="Wpływ dodawania krawędzi na niezawodność sieci")
    
    sim.visualize_graph()

if __name__ == "__main__":
    main()