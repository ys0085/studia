import matplotlib.pyplot as plt

def read_data(filename):
    x, y = [], []
    with open(filename, encoding='utf-8') as f:
        for line in f:
            line = line.strip()
            if not line or line.startswith('#') or line.startswith('//'):
                continue
            parts = line.split()
            if len(parts) == 2:
                try:
                    x_val = float(parts[0])
                    y_val = float(parts[1])
                    x.append(x_val)
                    y.append(y_val)
                except ValueError:
                    continue
    return x, y

if __name__ == "__main__":
    x, y = read_data('results_2')
    plt.figure(figsize=(10, 6))
    plt.plot(x, y, marker='o')
    plt.title('n vs Average Comparisons per Operation')
    plt.xlabel('n')
    plt.ylabel('Average comparisons per operation')
    plt.grid(True)
    plt.tight_layout()
    plt.savefig('results_2_plot.png')
    plt.show()