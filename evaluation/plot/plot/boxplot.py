import matplotlib.pyplot as plt
import pandas as pd

distro = pd.read_csv(
    "../../tests/random700mb/RESULTS/pullTimes/50mbps/concurrent/distromash.csv"
)
docker = pd.read_csv(
    "../../tests/random700mb/RESULTS/pullTimes/50mbps/concurrent/dockerregistry.csv"
)
spegel = pd.read_csv(
    "../../tests/random700mb/RESULTS/pullTimes/50mbps/concurrent/spegel.csv"
)

font = {
    "weight": "semibold",
}


def getData(d):
    node1 = d[d["Node"] == 1]
    node2 = d[d["Node"] == 2]
    node3 = d[d["Node"] == 3]
    node1_pull_times = list(node1["Time"].values)
    node2_pull_times = list(node2["Time"].values)
    node3_pull_times = list(node3["Time"].values)
    return [node1_pull_times, node2_pull_times, node3_pull_times]


def set_box_color(bp, color):
    plt.setp(bp["boxes"], color=color)
    plt.setp(bp["whiskers"], color=color)
    plt.setp(bp["caps"], color=color)
    plt.setp(bp["medians"], color=color)


# Data for the nodes
nodes = [
    "1",
    "2",
    "3",
]

distro_data = getData(distro)
dockerregistry_data = getData(docker)
spegel_data = getData(spegel)

plt.figure(figsize=(9.6, 4.8))
ax = plt.subplot(111)

bpl = ax.boxplot(
    distro_data,
    positions=[0.8, 3.4, 6],
    sym="+",
    notch=False,
    whis=30,
    widths=0.6,
)
bpr = ax.boxplot(
    dockerregistry_data,
    positions=[1.6, 4.2, 6.8],
    sym="+",
    notch=False,
    whis=30,
    widths=0.6,
)
bpm = ax.boxplot(
    spegel_data,
    positions=[2.4, 5, 7.6],
    sym="+",
    notch=False,
    whis=30,
    widths=0.6,
)
set_box_color(bpl, "blue")
set_box_color(bpr, "orange")
set_box_color(bpm, "green")

plt.xticks([1.6, 4.2, 6.8], nodes, fontdict=font)
plt.yticks(fontweight="bold")
plt.ylim(250, 400)

plt.xlabel("Node Number", fontdict=font)
plt.ylabel("Pull Time (s)", fontdict=font)
# plt.title("50 Mbps Sequential Starting of Containers", fontdict=font)

# Add lines between the nodes
plt.vlines(
    x=[2.9, 5.5],
    ymin=250,
    ymax=400,
    colors="#e8e9eb",
    linestyles="dashed",
)

# draw temporary red and blue lines and use them to create a legend
plt.plot([], c="blue", label="DistroMash")
plt.plot([], c="orange", label="Official Docker Registry")
plt.plot([], c="green", label="Spegel")
leg = ax.legend(
    loc="upper center",
    bbox_to_anchor=(0.5, 1.15),
    ncol=3,
    fancybox=True,
    edgecolor="black",
    prop=font,
)
# set the linewidth of each legend object
for legobj in leg.legend_handles:
    legobj.set_linewidth(5.0)

# plt.show()
plt.savefig("50mbpsConcurrent.png")
