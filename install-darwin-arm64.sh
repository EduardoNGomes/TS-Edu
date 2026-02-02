#!/bin/bash
set -e

bin=ts-edu-darwin-arm64
target_name=ts-edu
install_path=/usr/local/bin/$target_name
repo_url="https://github.com/edu-gomes/ts-edu"

echo "────────────────────────────────────────────"
echo "🚀 Installing ts-edu"
echo
echo "ℹ️  ts-edu is an open-source CLI tool."
echo "ℹ️  Source code and documentation:"
echo "🔗 $repo_url"
echo
echo "ℹ️  This binary was built from the official"
echo "ℹ️  GitHub repository and distributed via"
echo "ℹ️  GitHub Releases."
echo
echo "ℹ️  macOS may block binaries downloaded from"
echo "ℹ️  the internet. To ensure proper execution,"
echo "ℹ️  we will remove the quarantine attribute."
echo "────────────────────────────────────────────"
echo

if [ ! -f "./$bin" ]; then
    echo "❌ Error: Binary '$bin' not found in the current directory."
    echo "👉 Please download it from:"
    echo "   $repo_url/releases"
    exit 1
fi

if [ ! -f "$install_path" ]; then
    echo "📦 Installing ts-edu to /usr/local/bin..."
else
    echo "♻️  Updating existing ts-edu binary..."
fi

sudo cp "./$bin" "$install_path"

echo "🔓 Removing macOS quarantine attribute..."
sudo xattr -dr com.apple.quarantine "$install_path" || true

sudo chmod 755 "$install_path"

echo
echo "✅ Installation complete!"
echo "👉 Run 'ts-edu' to get started"
echo "📘 Docs & source: $repo_url"
