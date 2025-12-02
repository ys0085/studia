import matplotlib.pyplot as plt

def read_blocks(filename):
    blocks = []
    current = []
    with open(filename, encoding='utf-8') as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            if line.startswith('#'):
                continue
            parts = line.split()
            if len(parts) == 2:
                try:
                    x, y = float(parts[0]), float(parts[1])
                    current.append((x, y))
                except ValueError:
                    continue
            else:
                continue
            # Heuristic: new block starts when x==1 and current is not empty
            if x == 1 and current and len(current) > 1:
                blocks.append(current[:-1])
                current = [current[-1]]
        if current:
            blocks.append(current)
    return blocks

def plot_history(blocks):
    for i in range(5):
        x, y = zip(*blocks[i])
        plt.figure(figsize=(10,4))
        plt.plot(x, y, marker='.')
        plt.title(f'History plot {i+1}: Operation vs Comparisons')
        plt.xlabel('Operation number')
        plt.ylabel('Comparisons')
        plt.grid(True)
        plt.tight_layout()
        plt.savefig(f'history_{i+1}.png')
        plt.close()

if __name__ == "__main__":
    blocks = read_blocks('results')
    # First 5 blocks: history plots
    plot_history(blocks[:5])
    # Last block: n vs avg
    print("Saved 6 plots: history_1.png ... history_5.png, n_vs_avg.png")