# Dockyard

A terminal UI for browsing and applying [Starship](https://starship.rs) prompt presets.

Navigate the full preset list, preview how each one will look, apply it with a single keypress, and revert instantly if you change your mind, pretty much without leaving the terminal. I made this as a tool for myself to easily switch profiles without faffing around.

---

## Requirements

- [Starship](https://starship.rs) installed and in your PATH
- A [Nerd Font](https://www.nerdfonts.com) configured in your terminal (recommended — Dockyard will guide you through installing one if not found)

---

## Install

**macOS / Linux**
```bash
brew install merrickwykman/tap/dockyard
```

**Windows**
```powershell
irm https://raw.githubusercontent.com/merrickwykman/dockyard/main/install.ps1 | iex
```

**Direct download**

Grab the binary for your platform from the [latest release](https://github.com/merrickwykman/dockyard/releases/latest), make it executable, and move it somewhere on your PATH.

```bash
# macOS (Apple Silicon example)
chmod +x dockyard-mac-arm64
mv dockyard-mac-arm64 /usr/local/bin/dockyard
```

**Build from source**
```bash
git clone https://github.com/merrickwykman/dockyard.git
cd dockyard
go build -o dockyard .
```

---

## Usage

```bash
dockyard
```

On launch, Dockyard checks for Starship and a Nerd Font. If either is missing it will guide you through setup before proceeding.

### Key bindings

| Key | Action |
|-----|--------|
| `↑` / `↓` or `k` / `j` | Navigate presets |
| `Enter` | Apply selected preset |
| `R` | Revert to previous config |
| `Q` | Quit |

---

## How it works

- Ships with the full list of official Starship presets and their descriptions
- On apply: backs up your current `starship.toml` to `starship.toml.bak`, then writes the selected preset via `starship preset <name>`
- On revert: restores `starship.toml` from the backup
- Config path resolved via `os.UserHomeDir()` — no hardcoded paths

---

## License

MIT
