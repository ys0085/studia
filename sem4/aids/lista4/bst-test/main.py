import pandas as pd
import matplotlib.pyplot as plt
import io
from textwrap import dedent

# Read the data
df = pd.read_csv("results.csv")

# Define colors for each data type
colors = {
    'cmp': '#1f77b4',     # blue
    'read': '#ff7f0e',    # orange
    'write': '#2ca02c',   # green
    'height': '#d62728'   # red
}

# Create figure with subplots
fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(15, 6))

# Filter data for each scenario
scenario_0 = df[df['scenario'] == 0].sort_values('n')
scenario_1 = df[df['scenario'] == 1].sort_values('n')

# Plot Scenario 0
ax1.set_title('Scenario 0', fontsize=14, fontweight='bold')
ax1.plot(scenario_0['n'], scenario_0['avg_cmp'], color=colors['cmp'], linestyle='-', label='avg_cmp', linewidth=2)
ax1.plot(scenario_0['n'], scenario_0['max_cmp'], color=colors['cmp'], linestyle='--', label='max_cmp', linewidth=2)
ax1.plot(scenario_0['n'], scenario_0['avg_read'], color=colors['read'], linestyle='-', label='avg_read', linewidth=2)
ax1.plot(scenario_0['n'], scenario_0['max_read'], color=colors['read'], linestyle='--', label='max_read', linewidth=2)
ax1.plot(scenario_0['n'], scenario_0['avg_write'], color=colors['write'], linestyle='-', label='avg_write', linewidth=2)
ax1.plot(scenario_0['n'], scenario_0['max_write'], color=colors['write'], linestyle='--', label='max_write', linewidth=2)
ax1.plot(scenario_0['n'], scenario_0['avg_height'], color=colors['height'], linestyle='-', label='avg_height', linewidth=2)
ax1.plot(scenario_0['n'], scenario_0['max_height'], color=colors['height'], linestyle='--', label='max_height', linewidth=2)

ax1.set_xlabel('n')
ax1.set_ylabel('Value')
ax1.legend(bbox_to_anchor=(1.05, 1), loc='upper left')
ax1.grid(True, alpha=0.3)

# Plot Scenario 1
ax2.set_title('Scenario 1', fontsize=14, fontweight='bold')
ax2.plot(scenario_1['n'], scenario_1['avg_cmp'], color=colors['cmp'], linestyle='-', label='avg_cmp', linewidth=2)
ax2.plot(scenario_1['n'], scenario_1['max_cmp'], color=colors['cmp'], linestyle='--', label='max_cmp', linewidth=2)
ax2.plot(scenario_1['n'], scenario_1['avg_read'], color=colors['read'], linestyle='-', label='avg_read', linewidth=2)
ax2.plot(scenario_1['n'], scenario_1['max_read'], color=colors['read'], linestyle='--', label='max_read', linewidth=2)
ax2.plot(scenario_1['n'], scenario_1['avg_write'], color=colors['write'], linestyle='-', label='avg_write', linewidth=2)
ax2.plot(scenario_1['n'], scenario_1['max_write'], color=colors['write'], linestyle='--', label='max_write', linewidth=2)
ax2.plot(scenario_1['n'], scenario_1['avg_height'], color=colors['height'], linestyle='-', label='avg_height', linewidth=2)
ax2.plot(scenario_1['n'], scenario_1['max_height'], color=colors['height'], linestyle='--', label='max_height', linewidth=2)

ax2.set_xlabel('n')
ax2.set_ylabel('Value')
ax2.legend(bbox_to_anchor=(1.05, 1), loc='upper left')
ax2.grid(True, alpha=0.3)

# Adjust layout to prevent legend cutoff
plt.tight_layout()

# Show the plot
plt.show()

# Optional: Save the plot
plt.savefig('bst-test.png', dpi=300, bbox_inches='tight')