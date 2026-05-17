Dockyard — Architecture
Auth
None required. Dockyard is a local tool with no network dependencies beyond Starship's preset system and optional font download links.
Stack

Language: Go
TUI: Bubble Tea + Lip Gloss + Bubbles
Config target: ~/.config/starship.toml (cross-platform aware)
Distribution: single binary via GitHub Actions (Windows / Mac / Linux)
No database, no backend, no runtime dependencies

Config Path Handling
Starship's config lives at:

Mac/Linux: ~/.config/starship.toml
Windows: %USERPROFILE%\.config\starship.toml

Resolved at runtime using Go's os.UserHomeDir(). No hardcoded paths.
Backup Strategy
Before any write to starship.toml, the tool creates starship.toml.bak in the same directory. Single level only — one backup at a time. If a backup already exists it is overwritten. No versioned history in v1.
Nerd Font Detection
A known Nerd Font glyph is rendered in the terminal and its cell width is checked. If it renders correctly, the check passes silently. If not, the user is guided to install a compatible font. Result is cached locally after first run so the check does not repeat on every launch.
Font Installation by Platform

Mac: automatic install via Homebrew Cask
Linux: automatic download to ~/.local/share/fonts/, run fc-cache
Windows: detection only — if no Nerd Font is found, show a clear step-by-step in-TUI guide with direct download URL and exact installation instructions. No silent install attempted on Windows in v1.

Preset Source
Presets are sourced via starship preset list — Dockyard calls the Starship binary directly rather than maintaining its own copy of preset data. This keeps Dockyard current with any presets Starship adds without maintenance overhead.
Simulated Prompt Examples
When a preset is applied, Dockyard renders static string examples using the actual characters and symbols from the preset's TOML. No live shell execution. Examples shown:

Plain directory (~/projects/dockyard)
Git repo on a clean branch
Git repo with uncommitted changes
Directory with a recognised language runtime (e.g. Node, Python)

Architectural Exclusions (V1)

No custom module editing
No theme import/export
No network calls beyond optional font download
No persistent app state beyond font detection cache

