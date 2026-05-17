# Plan

 
# Dockyard — Build Plan
 
Each task is designed to be completable in a single Claude Code session. Complete and commit each task before starting the next.
 
---
 
## Task 1: Project Scaffold
 
**Context**
Empty repo with boilerplate docs. Go and Bubble Tea are the chosen stack.
 
**Objective**
Get a working Go project with a running Bubble Tea TUI displaying a placeholder list view.
 
**Files to create**
- go.mod
- go.sum
- main.go
- internal/ui/app.go
- internal/ui/list.go
**Requirements**
- Run `go mod init github.com/MerrickWykman/dockyard`
- Install dependencies: bubbletea, lipgloss, bubbles
- Render a placeholder list with 3-5 dummy items
- App launches, displays TUI, and exits cleanly on Q
**Do not**
- Implement any real logic
- Connect to Starship or read any config files
- Add any styling beyond confirming it renders
**Acceptance checks**
- `go run .` launches without errors
- TUI renders a visible list
- Q exits cleanly
---
 
## Task 2: Starship Detection
 
**Context**
Scaffold is in place. TUI renders. Now add the first real logic layer.
 
**Objective**
On launch, detect whether Starship is installed and accessible. If not, show platform-specific install guidance.
 
**Files to create or edit**
- internal/detect/starship.go
- internal/ui/app.go (integrate detection into launch flow)
**Requirements**
- Check for `starship` in PATH
- If found, retrieve version and proceed
- If not found, display a clear message in the TUI with platform-specific install instructions:
  - Mac: `brew install starship`
  - Windows: `winget install starship`
  - Linux: curl install script
- Detection runs before the main UI loads
**Do not**
- Auto-install Starship silently
- Proceed to main UI if Starship is not found
**Acceptance checks**
- With Starship installed: app proceeds to main UI
- With Starship removed from PATH: app shows guidance screen
- Platform detection works on all three OS targets
---
 
## Task 3: Config Management
 
**Context**
Starship detection is working. Now add the ability to read and write starship.toml safely.
 
**Objective**
Locate, read, back up, and write starship.toml cross-platform.
 
**Files to create**
- internal/config/config.go
**Requirements**
- Resolve config path using `os.UserHomeDir()` — no hardcoded paths
- Read existing starship.toml if present
- Before any write, create starship.toml.bak in the same directory
- Write new config content to starship.toml
- Expose a Revert() function that restores from .bak
- If no .bak exists, Revert() returns a clear error
**Do not**
- Delete or overwrite .bak without first creating a new one
- Implement any UI for this task — logic only
**Acceptance checks**
- Config path resolves correctly on Windows, Mac, and Linux
- Backup is created before every write
- Revert restores previous config correctly
- Revert returns appropriate error when no backup exists
---
 
## Task 4: Nerd Font Detection
 
**Context**
Config management is in place. Now add font detection before the main UI loads.
 
**Objective**
Detect whether a Nerd Font is active. Guide the user to install one if not.
 
**Files to create or edit**
- internal/detect/font.go
- internal/ui/app.go (integrate into launch flow)
**Requirements**
- Render a known Nerd Font glyph and check cell width
- If detection passes, cache result and proceed silently
- Cache stored in user config directory — do not re-run on every launch
- If detection fails:
  - Mac: offer automatic install via Homebrew Cask
  - Linux: offer automatic download to ~/.local/share/fonts/ + fc-cache
  - Windows: display clear step-by-step guide with direct download URL and install instructions
- Detection runs after Starship check, before main UI
**Do not**
- Attempt silent font install on Windows
- Block app permanently if font detection fails — allow user to proceed with a warning
**Acceptance checks**
- Glyph test runs correctly in supported terminals
- Cache prevents re-running on subsequent launches
- Mac/Linux auto-install flow executes without errors
- Windows guidance screen is clear and complete
- User can proceed past font warning if they choose
---
 
## Task 5: Preset Browser
 
**Context**
Detection flows are complete. Now build the core UI — the preset list.
 
**Objective**
Display all available Starship presets in a navigable TUI list with names and descriptions.
 
**Files to create or edit**
- internal/preset/preset.go
- internal/ui/list.go (replace placeholder with real data)
- internal/ui/detail.go
**Requirements**
- Call `starship preset list` to retrieve available presets
- Parse output into name + description pairs
- Display in a two-pane layout: list on left, detail on right
- Arrow key navigation updates detail pane in real time
- Detail pane shows preset name, description, and key binding hints
- Neutral/minimal colour scheme using Lip Gloss — no strong colours in chrome
**Do not**
- Apply any preset in this task
- Implement simulated examples yet
**Acceptance checks**
- All presets from `starship preset list` appear in the list
- Navigation updates detail pane correctly
- Layout renders cleanly at standard terminal widths (80col minimum)
- No crashes on empty or unexpected preset list output
---
 
## Task 6: Simulated Prompt Examples + Apply & Revert
 
**Context**
Preset browser is working. Now add the apply action and simulated examples.
 
**Objective**
Apply a preset on Enter, show simulated prompt examples, revert on R.
 
**Files to create or edit**
- internal/preset/preview.go
- internal/ui/detail.go (add examples to detail pane)
- internal/ui/app.go (wire up key bindings)
**Requirements**
- On Enter: back up current starship.toml, write selected preset, show confirmation
- Render simulated prompt examples in detail pane using preset TOML symbols:
  - Plain directory: `~/projects/dockyard`
  - Git repo clean: `~/projects/dockyard on main`
  - Git repo dirty: `~/projects/dockyard on main [+]`
  - Language runtime: `~/projects/dockyard via node v18`
- On R: restore starship.toml.bak, show confirmation
- If no backup exists, R is greyed out with explanation
- On Q: exit cleanly
**Do not**
- Execute a live shell to generate previews
- Add any config editing beyond preset apply
**Acceptance checks**
- Preset is correctly written to starship.toml on Enter
- Backup exists after apply
- Simulated examples render correctly for each preset
- Revert restores previous config correctly
- Revert option is correctly disabled when no backup exists
---
 
## Task 7: Distribution
 
**Context**
All features complete and tested. Now set up automated binary distribution.
 
**Objective**
GitHub Actions workflow that builds and releases binaries for Windows, Mac, and Linux on a version tag push.
 
**Files to create**
- .github/workflows/release.yml
**Requirements**
- Trigger on push of a tag matching `v*`
- Build targets:
  - windows/amd64 → dockyard-windows-amd64.exe
  - darwin/amd64 → dockyard-mac-amd64
  - darwin/arm64 → dockyard-mac-arm64
  - linux/amd64 → dockyard-linux-amd64
- Attach all binaries to a GitHub Release automatically
- Release notes pulled from CHANGELOG.md
**Do not**
- Add code signing in v1
- Set up Homebrew tap or winget manifest in v1
**Acceptance checks**
- Pushing a v* tag triggers the workflow
- All four binaries are built without errors
- Binaries are attached to the GitHub Release
- Release notes are populated from CHANGELOG.md
---



## Active task

### Task 2: Starship Detection

**Context**
Scaffold is in place. TUI renders. Now add the first real logic layer.

**Objective**
On launch, detect whether Starship is installed and accessible. If not, show platform-specific install guidance.

**Files to create or edit**
- internal/detect/starship.go — create
- internal/ui/app.go — edit (integrate detection into launch flow)

**Requirements**
- Check for `starship` in PATH
- If found, retrieve version and proceed
- If not found, display a clear message in the TUI with platform-specific install instructions:
  - Mac: `brew install starship`
  - Windows: `winget install starship`
  - Linux: curl install script
- Detection runs before the main UI loads

**Do not**
- Auto-install Starship silently
- Proceed to main UI if Starship is not found

**Acceptance checks**
- [ ] With Starship installed: app proceeds to main UI
- [ ] With Starship removed from PATH: app shows guidance screen
- [ ] Platform detection works on all three OS targets

---

## Backlog

Tasks 3–7 as defined above.

---

## Completed

### Task 1: Project Scaffold ✓
Go module initialised at `github.com/MerrickWykman/dockyard`. Bubble Tea TUI renders a placeholder list of 5 dummy preset items using `bubbles/list`. App launches in alt-screen mode, exits cleanly on Q. No real logic or Starship integration — scaffold only.
