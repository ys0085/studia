import numpy as np
from scipy.stats import binom

n_values = [100, 1000, 10000]
prob = 0.5
results = []

for n in n_values:
    mean = n * prob
    var = n * prob * (1 - prob)
    
    threshold_a = 1.2 * mean
    markov_a = mean / threshold_a
    exact_a = 1 - binom.cdf(threshold_a - 1, n, prob)
    
    threshold_b = 0.1 * mean
    chebyshev_b = var / (threshold_b ** 2)
    exact_b = (1 - binom.cdf(mean + threshold_b - 1, n, prob)) + binom.cdf(mean - threshold_b, n, prob)
    
    results.append({
        "n": n,
        "Markov (a)": markov_a,
        "Exact (a)": exact_a,
        "Chebyshev (b)": chebyshev_b,
        "Exact (b)": exact_b,
    })

import pandas as pd
df = pd.DataFrame(results)
print(df)