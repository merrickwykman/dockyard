# Changelog

## [1.0.1] — 2026-05-17

- Per-preset simulated prompt examples in the detail pane — each preset shows its separator style (Powerline arrows, brackets, plain, pure, no-font)
- `--version` flag added to the binary
- Homebrew tap and PowerShell one-liner installer
- GitHub Actions release workflow

## [1.0.0] — 2026-05-17

Initial release.

- Detects Starship installation on launch; shows platform-specific install guidance (brew / winget / curl) if not found
- Detects Nerd Font availability; offers automatic install on Mac/Linux and a download guide on Windows; user can proceed with a warning
- Reads, backs up, and writes `starship.toml` safely — every write creates a `.bak` before overwriting
- Two-pane preset browser with navigable list and detail pane
- Enter applies the selected preset; R reverts to previous config
- Q exits cleanly from all states
