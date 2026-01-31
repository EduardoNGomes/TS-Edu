#!/bin/bash
bin=ts-edu-linux-amd64
target_name=ts-edu
install_path=/usr/local/bin/$target_name

if [ ! -f "./$bin" ]; then
    echo "Error: Binary $bin not found in current directory."
    exit 1
fi

if [ ! -f "$install_path" ]; then
    echo "Installing $bin to /usr/local/bin..."
    sudo cp "./$bin" "$install_path"
    sudo chmod 755 "$install_path"
else
    echo "$target_name already exists in /usr/local/bin, updating..."
    sudo cp "./$bin" "$install_path"
    sudo chmod 755 "$install_path"
fi

echo "✔ Installation complete!"
