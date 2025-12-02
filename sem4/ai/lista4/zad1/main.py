import numpy as np
import matplotlib.pyplot as plt
from tensorflow.keras.datasets import mnist
from collections import Counter

def load_mnist_data(sample_size=10000):
    (x_train, y_train), _ = mnist.load_data()
    x = x_train[:sample_size]
    y = y_train[:sample_size]
    x = x.reshape(sample_size, -1).astype(np.float32) / 255.0
    return x, y

def initialize_centroids_kmeans(X, k):
    n_samples, n_features = X.shape
    centroids = np.zeros((k, n_features))
    centroids[0] = X[np.random.randint(0, n_samples)]
    
    for i in range(1, k):
        distances = np.min(np.linalg.norm(X[:, np.newaxis] - centroids[:i], axis=2)**2, axis=1)
        probs = distances / np.sum(distances)
        chosen_idx = np.random.choice(n_samples, p=probs)
        centroids[i] = X[chosen_idx]
    
    return centroids

def kmeans(X, k, max_iter=100, n_trials=5):
    best_inertia = np.inf
    best_centroids, best_labels = None, None
    
    for trial in range(n_trials):
        centroids = initialize_centroids_kmeans(X, k)
        for _ in range(max_iter):
            dists = np.linalg.norm(X[:, np.newaxis] - centroids, axis=2)
            labels = np.argmin(dists, axis=1)
            new_centroids = np.array([X[labels == i].mean(axis=0) if len(X[labels == i]) > 0 else centroids[i]
                                      for i in range(k)])
            if np.allclose(new_centroids, centroids):
                break
            centroids = new_centroids
        
        inertia = np.sum((X - centroids[labels]) ** 2)
        if inertia < best_inertia:
            best_inertia = inertia
            best_centroids = centroids
            best_labels = labels
    
    return best_centroids, best_labels, best_inertia

def plot_cluster_distribution(y_true, labels, k):
    matrix = np.zeros((10, k), dtype=int)
    for digit in range(10):
        idx = y_true == digit
        counts = Counter(labels[idx])
        for cluster in counts:
            matrix[digit, cluster] = counts[cluster]

    row_sums = matrix.sum(axis=1, keepdims=True)
    matrix = (matrix / row_sums * 100).round(1)
    
    fig, ax = plt.subplots(figsize=(12, 8))
    im = ax.imshow(matrix, cmap='Blues')

    for i in range(10):
        for j in range(k):
            ax.text(j, i, f'{matrix[i, j]}%', ha='center', va='center', color='black')
    
    ax.set_xlabel("Klastry")
    ax.set_ylabel("Cyfry rzeczywiste")
    ax.set_title(f"Rozkład klastra (%) dla k={k}")
    plt.colorbar(im)
    plt.show()

def plot_centroids(centroids):
    side = int(np.sqrt(centroids.shape[1]))
    k = centroids.shape[0]
    cols = 10
    rows = int(np.ceil(k / cols))
    
    fig, axs = plt.subplots(rows, cols, figsize=(15, rows * 1.5))
    axs = axs.flatten()

    for i in range(k):
        axs[i].imshow(centroids[i].reshape(side, side), cmap='gray')
        axs[i].axis('off')
        axs[i].set_title(f'Centroid {i}')
    
    for j in range(k, len(axs)):
        axs[j].axis('off')

    plt.suptitle("Obrazy centroidów")
    plt.tight_layout()
    plt.show()

def main():
    X, y = load_mnist_data(sample_size=10000)
    for k in [10, 15, 20, 30]:
        print(f"\nClustering dla k={k}")
        centroids, labels, inertia = kmeans(X, k, max_iter=100, n_trials=5)
        print(f"Inercja: {inertia:.2f}")
        
        plot_cluster_distribution(y, labels, k)
        plot_centroids(centroids)

if __name__ == "__main__":
    main()
