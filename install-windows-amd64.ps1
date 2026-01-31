$bin = "ts-edu-windows-amd64.exe"
$target_name = "ts-edu.exe"
# Installation to a common path in System32 or similar requires Admin, but often handled differently.
# Here we'll try to find a suitable path or just user's local bin if exists.
# For simplicity matching the 'system install' pattern, we'll assume C:\Windows\System32 or allow user to pick.
# Actually, strict equivalent to /usr/local/bin is hard in Windows without standard paths.
# We will check if a directory is in PATH or suggest one.

# Let's try to install to a folder in Program Files and add to PATH, or just copy to where the script is run?
# The request was "follow the pattern". Pattern is: copy to global path.
# We'll assume usage of a common tools folder or C:\Windows (dangerous/messy).

# Safer approach for Windows script:
Write-Host "Installing $bin..."

if (-not (Test-Path ".\$bin")) {
    Write-Error "Binary $bin not found in current directory."
    exit 1
}

# Determine install path (user profile bin is safer)
$install_dir = "$env:LOCALAPPDATA\ts-edu\bin"
if (-not (Test-Path $install_dir)) {
    New-Item -ItemType Directory -Force -Path $install_dir | Out-Null
}

$dest = Join-Path $install_dir $target_name
Copy-Item -Path ".\$bin" -Destination $dest -Force

# Add to PATH if not present
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notlike "*$install_dir*") {
    [Environment]::SetEnvironmentVariable("Path", "$userPath;$install_dir", "User")
    Write-Host "Added $install_dir to User PATH. You may need to restart your terminal."
}

Write-Host "✔ Installation complete! Installed to $dest"
