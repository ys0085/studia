import subprocess
import numpy as np
import matplotlib.pyplot as plt

n_values = list(range(10, 51, 10))  
n_values_large = list(range(1000, 100001, 1000))  
# k_values = [1, 10, 100]  

algorithms = {
    "Binary Search": "./binary-search/binary-search",
}

algorithms_large = algorithms

def run_command(command, input_data):
    process = subprocess.Popen(command.split(), stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    stdout, _ = process.communicate(input_data)
    return stdout.strip()

def run_experiments(n_values, m, algorithms):
    results = {algo: {"c": []} for algo in algorithms}

    for n in n_values:
        c_sums = {algo: 0 for algo in algorithms}

        for _ in range(m):
            arr = np.random.randint(1, 2 * n + 1, size=n)
            target = np.random.randint(1, 2 * n + 1)
            arr.sort()
            args = str(n) + ' ' + str(target) + ' '.join(map(str, arr))

            for algo, cmd in algorithms.items():
                output = run_command(cmd, args)
                c = int(output) 

                c_sums[algo] += c


        for algo in algorithms:
            results[algo]["c"].append(c_sums[algo] / m)
        print("done with " + str(n))
        print(results[algo]["c"][-1])

    return results
 
results_large = run_experiments(n_values_large, 10, algorithms_large)  

def plot_results(results, n_values, title, filename):
    plt.figure(figsize=(10, 6))
    for algo in results:
        plt.plot(n_values, results[algo]["c"], label=f"{algo} (Comparisons)")

    plt.xlabel("n (Array Size)")
    plt.ylabel("Operations Count")
    plt.title(title)
    plt.legend()
    plt.grid()
    plt.savefig(filename)

plot_results(results_large, n_values_large, "Large n", "BinSearch.png")