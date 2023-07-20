
#!/bin/bash

# Source and destination directories
source_dir="."
dest_dir="../FusionAutomationService/SandboxAddIn/ACT"

# Source file name
source_file="act.darwin"

# Destination file name (with modified name)
dest_file="act"

# Copy the file to the destination directory
cp "${source_dir}/${source_file}" "${dest_dir}/${dest_file}"

# Modify the file permissions to make it executable
chmod +x "${dest_dir}/${dest_file}"

# Source file name
source_file="act.win64.exe"

# Destination file name (with modified name)
dest_file="act.exe"

# Copy the file to the destination directory
cp "${source_dir}/${source_file}" "${dest_dir}/${dest_file}"

# Modify the file permissions to make it executable
chmod +x "${dest_dir}/${dest_file}"

echo "File copied and modified successfully."
