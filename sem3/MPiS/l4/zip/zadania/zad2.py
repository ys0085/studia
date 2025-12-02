import numpy as np
import matplotlib.pyplot as plt
from scipy.stats import norm

def generate(N, num_samples=100000, num_bins=50):
    samples = np.sum(np.random.choice([-1, 1], size=(num_samples, N)), axis=1)

    hist, bin_edges = np.histogram(samples, bins=num_bins, range=(-N, N), density=True)

    cdf_empirical = np.cumsum(hist) * (bin_edges[1] - bin_edges[0])

    return samples, bin_edges, cdf_empirical

def plot_results(N, samples, bin_edges, cdf_empirical):

    x_vals = np.linspace(-N, N, 1000)
    pdf_normal = norm.pdf(x_vals, 0, np.sqrt(N))
    cdf_normal = norm.cdf(x_vals, 0, np.sqrt(N))
    
    # Histogram
    plt.figure(figsize=(14, 6))
    plt.subplot(1, 2, 1)
    plt.hist(samples, bins=50, range=(-N, N), density=True, alpha=0.6, label="Empiryczne")
    plt.plot(x_vals, pdf_normal, 'r-', label="Rozkład normalny (PDF)")
    plt.title(f"Histogram S_N dla N={N}")
    plt.xlabel("Wartości S_N")
    plt.ylabel("Gęstość prawdopodobieństwa")
    plt.legend()
    
    # Dystrybuanta
    plt.subplot(1, 2, 2)
    plt.step(bin_edges[:-1], cdf_empirical, where='post', label="Dystrybuanta empiryczna", color="blue")
    plt.plot(x_vals, cdf_normal, 'r-', label="Rozkład normalny (CDF)")
    plt.title(f"Dystrybuanta S_N dla N={N}")
    plt.xlabel("Wartości S_N")
    plt.ylabel("P(S_N ≤ x)")
    plt.legend()
    
    plt.tight_layout()
    plt.show()


N_values = [5, 10, 15, 20, 25, 30, 100]

for N in N_values:
    samples, bin_edges, cdf_empirical = generate(N)
    plot_results(N, samples, bin_edges, cdf_empirical)