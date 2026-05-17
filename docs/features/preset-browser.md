# Feature: Preset Browser
 
## What it does
The main screen of Dockyard. Displays all built-in Starship presets in a navigable two-pane TUI layout. The user browses presets by name, reads a description, sees simulated prompt examples, and applies or reverts with single keypresses.
 
## Layout
- Left pane: scrollable list of preset names
- Right pane: selected preset detail — name, description, simulated prompt examples, key binding hints
## Preset data source
Presets are retrieved by calling `starship preset list` against the installed Starship binary. Dockyard does not maintain its own copy of preset data.
 
## Simulated prompt examples
When a preset is selected, the detail pane renders static examples using the preset's actual symbols and characters:
- Plain directory: `~/projects/dockyard`
- Git repo on a clean branch: `~/projects/dockyard on main`
- Git repo with uncommitted changes: `~/projects/dockyard on main [+]`
- Directory with a language runtime: `~/projects/dockyard via node v18`
## Key bindings
| Key | Action |
|---|---|
| ↑ / ↓ | Navigate preset list |
| Enter | Apply selected preset |
| R | Revert to previous config |
| Q | Quit |
| ? | Show help |
 
## Visual style
Neutral and minimal. Lip Gloss used for layout and subtle highlights only. The preset examples are the visual focus — Dockyard's chrome stays out of the way.
 
## What it does not do
- Provide a live shell-executed preview
- Allow editing of individual modules or colours
- Import community or custom themes