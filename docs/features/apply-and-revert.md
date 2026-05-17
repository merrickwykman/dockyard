 
# Feature: Apply and Revert
 
## What it does
Applies a selected Starship preset to the user's starship.toml with a single keypress, and reverts to the previous config with a single keypress. The backup is handled automatically — the user never needs to manage files manually.
 
## Apply behaviour
1. User presses Enter on a selected preset
2. Existing starship.toml is backed up to starship.toml.bak in the same directory
3. Selected preset config is written to starship.toml
4. Confirmation message shown in TUI
5. Simulated prompt examples rendered in detail pane
6. User opens a new terminal tab to see it live
## Revert behaviour
1. User presses R
2. starship.toml.bak is restored to starship.toml
3. Confirmation message shown in TUI
4. If no backup exists, R is visually disabled with a clear explanation
## Backup strategy
- Single level only — one .bak file at a time
- Each new apply overwrites the previous .bak
- No versioned history in v1
## What it does not do
- Provide multi-level undo
- Keep a config history
- Prompt for confirmation before applying — apply is immediate