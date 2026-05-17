# Dockyard installer for Windows
# Usage: irm https://raw.githubusercontent.com/MerrickWykman/dockyard/main/install.ps1 | iex

$ErrorActionPreference = "Stop"

$repo    = "MerrickWykman/dockyard"
$binName = "dockyard.exe"
$installDir = Join-Path $env:USERPROFILE ".local\bin"

Write-Host "Fetching latest Dockyard release..."

$release = Invoke-RestMethod "https://api.github.com/repos/$repo/releases/latest"
$asset   = $release.assets | Where-Object { $_.name -eq "dockyard-windows-amd64.exe" }

if (-not $asset) {
    Write-Error "Could not find Windows binary in the latest release. Please check https://github.com/$repo/releases"
    exit 1
}

New-Item -ItemType Directory -Force -Path $installDir | Out-Null

$dest = Join-Path $installDir $binName
Write-Host "Downloading $($asset.name)..."
Invoke-WebRequest $asset.browser_download_url -OutFile $dest

Write-Host ""
Write-Host "Installed: $dest"
Write-Host ""

# Check whether install dir is already in the user PATH
$userPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($userPath -like "*$installDir*") {
    Write-Host "dockyard is ready. Run: dockyard"
} else {
    Write-Host "To use dockyard from any terminal, add the install directory to your PATH."
    Write-Host ""
    Write-Host "Run this once to update your user PATH permanently:"
    Write-Host ""
    Write-Host "  [Environment]::SetEnvironmentVariable('PATH', `"`$env:PATH;$installDir`", 'User')"
    Write-Host ""
    Write-Host "Then restart your terminal and run: dockyard"
}
