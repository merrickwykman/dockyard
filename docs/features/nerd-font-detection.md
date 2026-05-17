# Feature: Nerd Font Detection
 
## What it does
On first launch, Dockyard tests whether a Nerd Font is active in the current terminal by rendering a known glyph and checking its cell width. The result is cached so the check only runs once.
 
## User-facing behaviour
- If a Nerd Font is detected: proceeds silently, no interruption
- If no Nerd Font is detected: displays a clear explanation of why icons will appear broken, with platform-specific guidance

## Install behaviour by platform
- Mac: offers to run `brew install --cask font-jetbrains-mono-nerd-font` automatically
- Linux: offers to download font to `~/.local/share/fonts/` and run `fc-cache` automatically
- Windows: displays a step-by-step guide in the TUI with a direct download link and exact instructions for installing the font — no silent install attempted

## Additional behaviour
- User can choose to proceed past the font warning if they want to continue anyway
- Cache is stored in the user config directory
- Cache is invalidated if the user changes terminal or requests a recheck
## What it does not do
- Block the app permanently if no Nerd Font is found
- Attempt silent font installation on Windows
- Re-run the detection check on every launch
