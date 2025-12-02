import numpy as np

def plot_clusters(data, labels, title='DBSCAN Clustering Results'):
    import matplotlib.pyplot as plt

    unique_labels = set(labels)
    colors = [plt.cm.Spectral(each) for each in np.linspace(0, 1, len(unique_labels))]

    for k, col in zip(unique_labels, colors):
        if k == -1:
            col = [0, 0, 0, 1]  # Black for noise

        class_member_mask = (labels == k)

        xy = data[class_member_mask]
        plt.scatter(xy[:, 0], xy[:, 1], color=col, s=30, label=f'Cluster {k}')

    plt.title(title)
    plt.legend()
    plt.show()


def calculate_metrics(labels_true, labels_pred, data):
    from sklearn.metrics import silhouette_score, adjusted_rand_score

    # Silhouette only if more than 1 cluster
    if len(set(labels_pred)) > 1:
        silhouette_avg = silhouette_score(data, labels_pred)
    else:
        silhouette_avg = -1
    ari = adjusted_rand_score(labels_true, labels_pred)

    return silhouette_avg, ari


def clustering_report(true_labels, cluster_labels):
    from collections import Counter

    n_points = len(true_labels)
    noise = np.sum(cluster_labels == -1)
    noise_percent = 100 * noise / n_points

    # Map each cluster to the most common true label
    clusters = set(cluster_labels)
    clusters.discard(-1)
    correct = 0
    total_in_clusters = 0
    wrong = 0

    for c in clusters:
        idx = np.where(cluster_labels == c)[0]
        total_in_clusters += len(idx)
        true = true_labels[idx]
        most_common = Counter(true).most_common(1)[0][0]
        correct += np.sum(true == most_common)
        wrong += np.sum(true != most_common)

    accuracy = 100 * correct / total_in_clusters if total_in_clusters > 0 else 0
    wrong_percent = 100 * wrong / total_in_clusters if total_in_clusters > 0 else 0

    print(f"Liczba klastrów (bez szumu): {len(clusters)}")
    print(f"Procent szumu: {noise_percent:.2f}%")
    print(f"Dokładność klasteryzacji (najczęstsza cyfra w klastrze): {accuracy:.2f}%")
    print(f"Procent błędnych klasyfikacji w klastrach: {wrong_percent:.2f}%")