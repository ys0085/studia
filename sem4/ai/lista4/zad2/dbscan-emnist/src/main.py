import numpy as np
from data_loader import load_emnist
from dbscan_cluster import DBSCANClusterer
from utils import plot_clusters, calculate_metrics, clustering_report

def main():
    data, labels = load_emnist()

    from sklearn.decomposition import PCA
    data_pca = PCA(n_components=10).fit_transform(data)

    data_small, labels_small = data_pca, labels

    clusterer = DBSCANClusterer(eps=1.8, min_samples=5)
    clusterer.fit(data_small)
    cluster_labels = clusterer.predict(data_small)

    plot_clusters(data_small, cluster_labels)

    metrics = calculate_metrics(labels_small, cluster_labels, data_small)
    print("Silhouette, ARI:", metrics)

    clustering_report(labels_small, cluster_labels)

if __name__ == "__main__":
    main()