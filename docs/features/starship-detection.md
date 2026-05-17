# Feature: Starship Detection
 
## What it does
On every launch, Dockyard checks whether Starship is installed and accessible in the user's PATH before proceeding to the main UI.
 
## User-facing behaviour
- If Starship is found: app proceeds silently, no interruption
- If Starship is not found: app displays a clear message explaining that Starship is required, with platform-specific one-line install instructions

## Install guidance by platform
- Mac: `brew install starship`
- Windows: `winget install starship`
- Linux: `curl -sS https://starship.rs/install.sh | sh`

## What it does not do
- Auto-install Starship without user action
- Proceed to the main UI if Starship is missing
- Show a generic error — guidance is always platform-specific
