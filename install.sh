#!/bin/bash
bin=ts-edu
install_path=/usr/local/bin/$bin

if [ ! -f "$install_path" ]; then
    echo "instaling $bin em /usr/local/bin..."
    sudo cp "$bin" "$install_path"
    sudo chmod 755 "$install_path"
else
	echo "$bin already exists in /usr/local/bin, updating..."
    sudo cp "$bin" "$install_path"
    sudo chmod 755 "$install_path"
fi

rm -rf "$FILE"
echo "✔ Installation complete!"
