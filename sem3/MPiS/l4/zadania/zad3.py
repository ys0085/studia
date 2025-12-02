import numpy as np
import matplotlib.pyplot as plt
from scipy.stats import norm

def calculate(n, k=5000):
    result = []
    for _ in range(k):
        i_above_axis = 0
        val = 0
        for i in range(n):
            next_val = val + np.random.choice([-1, 1])
            if next_val >= 0 and val >= 0: 
                i_above_axis += 1
            val = next_val
        result.append(i_above_axis)
        print(i_above_axis)
    return result
n = [100, 1000, 10000]
for N in n: 
    
    samples = calculate(N)
    plt.hist(samples, bins=20, density=True, alpha=0.6)

    
    plt.title(f"Histogram P_N dla N={N}")
    plt.xlabel("Wartości P_N")
    plt.ylabel("Prawdopodobieństwo")
    plt.savefig(f"./zadania/img/zad3_{N}.png")
    plt.clf()


