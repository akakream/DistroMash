import matplotlib.pyplot as plt
from datetime import datetime
import pandas as pd


def read_csv(file_path):
    # Read the CSV file into a DataFrame.
    df = pd.read_csv(file_path)

    # Access 'Time' and 'Kbps out' columns.
    time_column = df["Time"]
    kbps_out_column = df["Kbps out"]

    # Extract time and Kbps out data
    time = [datetime.strptime(item, "%H:%M:%S") for item in time_column]
    kbps_out = [item / 1000 for item in kbps_out_column]  # Convert Kbps to Mbps

    return time, kbps_out


# Replace 'your_data.csv' with the actual file path to your CSV file.
docker_file_path = "../networkDataDockerRegistry.csv"
distro_file_path = "../networkDataDistroMash.csv"
spegel_file_path = "../networkDataDistroMash.csv"

do_time, do_out = read_csv(docker_file_path)
do_time = [(t - do_time[0]).total_seconds() for t in do_time]
di_time, di_out = read_csv(distro_file_path)
di_time = [(t - di_time[0]).total_seconds() for t in di_time]
sp_time, sp_out = read_csv(spegel_file_path)
sp_time = [(t - sp_time[0]).total_seconds() for t in sp_time]

# Apply a rolling median filter to smooth the data
window_size = 3  # Adjust the window size as needed
do_out = (
    pd.Series(do_out).rolling(window=window_size, center=True, min_periods=1).median()
)
di_out = (
    pd.Series(di_out).rolling(window=window_size, center=True, min_periods=1).median()
)
sp_out = (
    pd.Series(sp_out).rolling(window=window_size, center=True, min_periods=1).median()
)

# Create the plot
plt.figure(figsize=(12, 6))
plt.plot(do_time, do_out, label="Official Docker Registry", color="orange")
plt.plot(di_time, di_out, label="DistroMash", color="blue")
plt.plot(sp_time, sp_out, label="Spegel", color="green")
plt.xlabel("Experiment Duration (s)")
plt.ylabel("Network Utilization (Mbps)")
plt.title("Network Utilization Over Time", fontweight="bold")
plt.xticks(rotation=45)
plt.legend()
plt.tight_layout()

plt.xlim(0, max(do_time))

# Show the plot
plt.show()
