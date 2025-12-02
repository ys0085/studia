import matplotlib.pyplot as plt
import csv

def read_csv(filename):
    N = []
    prim = []
    kruskal = []
    with open(filename, encoding='utf-8') as f:
        reader = csv.reader(f)
        header = next(reader)  # Skip header
        for row in reader:
            if not row or row[0].startswith('//'):
                continue
            try:
                n = int(row[0])
                p = int(row[1])
                k = int(row[2])
                N.append(n)
                prim.append(p/1e6)
                kruskal.append(k/1e6)
            except Exception:
                continue
    return N, prim, kruskal

if __name__ == "__main__":
    N, prim, kruskal = read_csv('results.csv')
    plt.figure(figsize=(10,6))
    plt.plot(N, prim, label="Prim's algorithm", marker='o')
    plt.plot(N, kruskal, label="Kruskal's algorithm", marker='s')
    plt.xlabel('N (number of vertices)')
    plt.ylabel('Time (ms)')
    plt.title("Prim's vs Kruskal's Algorithm Running Time")
    plt.legend()
    plt.grid(True)
    plt.tight_layout()
    plt.savefig('mst_algorithms_times.png')
    plt.show()