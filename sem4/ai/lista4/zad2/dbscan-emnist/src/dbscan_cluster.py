import numpy as np

class DBSCANClusterer:
    def __init__(self, eps=0.5, min_samples=5):
        self.eps = eps
        self.min_samples = min_samples
        self.labels_ = None

    def fit(self, X):
        n = X.shape[0]
        labels = np.full(n, -1, dtype=int)
        cluster_id = 0
        visited = np.zeros(n, dtype=bool)

        def region_query(point_idx):
            dists = np.linalg.norm(X - X[point_idx], axis=1)
            return np.where(dists <= self.eps)[0]

        for i in range(n):
            if visited[i]:
                continue
            visited[i] = True
            neighbors = region_query(i)
            if len(neighbors) < self.min_samples:
                labels[i] = -1  # noise
            else:
                labels[i] = cluster_id
                seeds = set(neighbors)
                seeds.discard(i)
                while seeds:
                    curr = seeds.pop()
                    if not visited[curr]:
                        visited[curr] = True
                        curr_neighbors = region_query(curr)
                        if len(curr_neighbors) >= self.min_samples:
                            seeds.update(curr_neighbors)
                    if labels[curr] == -1:
                        labels[curr] = cluster_id
                cluster_id += 1
        self.labels_ = labels

    def predict(self, X):
        return self.labels_