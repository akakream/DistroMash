import matplotlib.pyplot as plt

# Data for pull times for each node
node1_pull_times = [63.28, 61.51, 61.50, 62.92, 63.04, 61.59, 62.94]
node2_pull_times = [61.82, 62.09, 61.54, 63.00, 61.74, 61.43, 61.91]
node3_pull_times = [61.82, 62.09, 61.54, 63.00, 61.74, 61.43, 61.91]

# Data for the nodes
nodes = ["Node 1", "Node 2", "Node 3"]

# Combine the data for plotting
data = [node1_pull_times, node2_pull_times, node3_pull_times]

# Create a box plot
plt.boxplot(data, labels=nodes)

# Set y-axis limits from 40 to 70
plt.ylim(40, 70)

# Set axis labels and title
plt.xlabel("Node")
plt.ylabel("Pull Time")
plt.title("Box Plot of Pull Times for Nodes")

# Show the plot
plt.show()
