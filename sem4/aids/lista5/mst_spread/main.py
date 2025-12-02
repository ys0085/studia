import matplotlib.pyplot as plt
import numpy as np

# Data from the table
n_values = [100, 200, 300, 400, 500, 600, 700, 800, 900, 1000,
           1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000,
           2100, 2200, 2300, 2400, 2500, 2600, 2700, 2800, 2900, 3000,
           3100, 3200, 3300, 3400, 3500, 3600, 3700, 3800, 3900, 4000,
           4100, 4200, 4300, 4400, 4500, 4600, 4700, 4800, 4900, 5000]

average_values = [13.80, 19.55, 23.00, 26.30, 28.05, 31.50, 31.70, 34.80, 36.50, 37.65,
                 41.65, 40.70, 42.35, 44.45, 44.75, 47.05, 47.50, 46.40, 49.70, 51.10,
                 51.45, 53.30, 51.55, 55.45, 56.00, 55.05, 56.50, 57.40, 53.80, 59.20,
                 59.95, 62.20, 60.25, 61.10, 64.45, 60.55, 64.45, 63.75, 64.60, 63.35,
                 66.70, 66.30, 66.35, 70.30, 64.75, 68.00, 69.95, 68.55, 73.70, 74.60]

# Create the plot
plt.figure(figsize=(12, 8))
plt.plot(n_values, average_values, 'b-', linewidth=2, marker='o', markersize=4)
plt.grid(True, alpha=0.3)
plt.xlabel('N', fontsize=12)
plt.ylabel('Average (Średnia)', fontsize=12)
plt.title('Average vs N', fontsize=14, fontweight='bold')

# Add some styling
plt.tight_layout()
plt.xlim(0, 5200)
plt.ylim(0, max(average_values) * 1.1)


plt.savefig('average_vs_n.png', dpi=300, bbox_inches='tight')

print(f"Data points plotted: {len(n_values)}")
print(f"N range: {min(n_values)} to {max(n_values)}")
print(f"Average range: {min(average_values):.2f} to {max(average_values):.2f}")