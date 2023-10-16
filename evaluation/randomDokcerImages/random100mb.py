import os

file_size = 100 * 1024 * 1024  # 100MB
file_name = "random_binary_file.bin"

# Generate random binary data
random_data = os.urandom(file_size)

# Write the data to a file
with open(file_name, "wb") as file:
    file.write(random_data)

print(f"Random binary file '{file_name}' with {file_size} bytes has been created.")
