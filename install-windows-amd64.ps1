$bin = "ts-edu-windows-amd64.exe"
$target_name = "ts-edu.exe"
$repo_url = "https://github.com/edu-gomes/ts-edu"

Write-Host "────────────────────────────────────────────"
Write-Host "🚀 Installing ts-edu"
Write-Host ""
Write-Host "ℹ️  ts-edu is an open-source CLI tool."
Write-Host "ℹ️  Source code and documentation:"
Write-Host "🔗 $repo_url"
Write-Host ""
Write-Host "ℹ️  This binary was built from the official"
Write-Host "ℹ️  GitHub repository and distributed via"
Write-Host "ℹ️  GitHub Releases."
Write-Host ""
Write-Host "ℹ️  The binary will be installed for the"
Write-Host "ℹ️  current user and added to your PATH."
Write-Host "────────────────────────────────────────────"
Write-Host ""

if (-not (Test-Path ".\$bin")) {
    Write-Error "❌ Binary '$bin' not found in the current directory."
    Write-Host "👉 Download it from:"
    Write-Host "   $repo_url/releases"
    exit 1
}

# Safer per-user install directory
$install_dir = "$env:LOCALAPPDATA\ts-edu\bin"

if (-not (Test-Path $install_dir)) {
    Write-Host "📁 Creating install directory: $install_dir"
    New-Item -ItemType Directory -Force -Path $install_dir | Out-Null
}

$dest = Join-Path $install_dir $target_name

Write-Host "📦 Installing ts-edu..."
Copy-Item -Path ".\$bin" -Destination $dest -Force

# Add directory to PATH if not present
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notlike "*$install_dir*") {
    [Environment]::SetEnvironmentVariable(
        "Path",
        "$userPath;$install_dir",
        "User"
    )
    Write-Host "🔧 Added ts-edu to your PATH"
    Write-Host "👉 Restart your terminal to use 'ts-edu'"
}

Write-Host ""
Write-Host "✅ Installation complete!"
Write-Host "👉 Run 'ts-edu' to get started"
Write-Host "📘 Docs & source: $repo_url"
