import subprocess
import numpy as np
import matplotlib.pyplot as plt

n_values = list(range(10, 51, 10))  
n_values_large = list(range(1000, 50001, 1000))  
k_values = [1, 10, 100]  

algorithms = {
    "Select 3": "./select/select -n 3",
    "Select 5": "./select/select -n 5",
    "Select 7": "./select/select -n 7",
    "Select 9": "./select/select -n 9"
}

algorithms_large = algorithms

def run_command(command, input_data):
    process = subprocess.Popen(command.split(), stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    stdout, _ = process.communicate(input_data)
    return stdout.strip()

def run_experiments(n_values, m, algorithms):
    results = {algo: {"c": [], "s": []} for algo in algorithms}

    for n in n_values:
        c_sums, s_sums = {algo: 0 for algo in algorithms}, {algo: 0 for algo in algorithms}

        for _ in range(m):
            args = run_command(f"./generator {n}", "")

            for algo, cmd in algorithms.items():
                k = np.random.randint(0, n)
                output = run_command(cmd, args)
                c, s = map(int, output.split())
                if c == 0 or s == 0: raise ValueError
                c_sums[algo] += c
                s_sums[algo] += s

        for algo in algorithms:
            results[algo]["c"].append(c_sums[algo] / m)
            results[algo]["s"].append(s_sums[algo] / m)
        print("finished " + str(n))
    return results
 
results_large = run_experiments(n_values_large, 10, algorithms)  

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


plot_results(results_large, n_values_large, "Large n", "selectvalues.png")