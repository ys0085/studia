import subprocess
import numpy as np
import matplotlib.pyplot as plt

n_values = list(range(10, 51, 10))  
n_values_large = list(range(1000, 50001, 1000))  
k_values = [1, 10, 100]  

algorithms = {
    "Insertion Sort": "./insertion-sort/insertion-sort",
    "Quick Sort": "./quick-sort/quick-sort",
    "Hybrid Sort (t=10)": "./hybrid-sort/hybrid-sort -t 10",
    "My sort": "./my-sort/my-sort", 
    "Double pivot QS": "./double-pivot-qs/double-pivot-qs",
    "Merge Sort": "./merge-sort/merge-sort"
}

algorithms_large = {
    "Quick Sort": "./quick-sort/quick-sort", 
    "Hybrid Sort (t=10)": "./hybrid-sort/hybrid-sort -t 10", 
    "My sort": "./my-sort/my-sort", 
    "Double pivot QS": "./double-pivot-qs/double-pivot-qs",
    "Merge Sort": "./merge-sort/merge-sort"
}

def run_command(command, input_data):
    process = subprocess.Popen(command.split(), stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    stdout, _ = process.communicate(input_data)
    return stdout.strip()

def run_experiments(n_values, k, algorithms):
    results = {algo: {"c": [], "s": []} for algo in algorithms}

    for n in n_values:
        c_sums, s_sums = {algo: 0 for algo in algorithms}, {algo: 0 for algo in algorithms}

        for _ in range(k):
            numbers = run_command(f"./generator {n}", "").split(" ")

            for algo, cmd in algorithms.items():
                output = run_command(cmd, " ".join(numbers))
                c, s = map(int, output.split()) 

                c_sums[algo] += c
                s_sums[algo] += s

        for algo in algorithms:
            results[algo]["c"].append(c_sums[algo] / k)
            results[algo]["s"].append(s_sums[algo] / k)

    return results

results_small = run_experiments(n_values, 10, algorithms)  
results_large = run_experiments(n_values_large, 10, algorithms_large)  

def plot_results(results, n_values, title, filename):
    plt.figure(figsize=(10, 6))
    for algo in results:
        plt.plot(n_values, results[algo]["c"], label=f"{algo} (Comparisons)")
        plt.plot(n_values, results[algo]["s"], '--', label=f"{algo} (Swaps)")

    plt.xlabel("n (Array Size)")
    plt.ylabel("Operations Count")
    plt.title(title)
    plt.legend()
    plt.grid()
    plt.savefig(filename)


plot_results(results_small, n_values, "Sorting Algorithm Comparisons & Swaps (Small n)", "results1.png")

plot_results(results_large, n_values_large, "Insertion Sort Complexity for Large n", "results2.png")