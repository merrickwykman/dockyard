# Changelog

<!--
Update after meaningful completed tasks — not every minor commit.
Record what changed, why it changed, and any decisions made.
Prevents relitigating decisions in future sessions.
-->

## [Unreleased]
<!--
Staging area. Move to a dated entry on each push to production.
-->

## [1.0.0] — 2026-05-17

Initial release.

- Detects Starship installation on launch; shows platform-specific install guidance (brew / winget / curl) if not found
- Detects Nerd Font availability; offers automatic install on Mac/Linux and a download guide on Windows; user can proceed with a warning
- Reads, backs up, and writes `starship.toml` safely — every write creates a `.bak` before overwriting
- Two-pane preset browser: all presets from `starship preset list` in a navigable list with name and description
- Detail pane shows per-preset simulated prompt examples reflecting each preset's separator style (Powerline arrows, brackets, plain, pure, no-font) and icon set
- Enter applies the selected preset (runs `starship preset <name>`, writes result to `starship.toml`)
- R reverts to the previous config from `.bak`; greyed out with explanation when no backup exists
- Q exits cleanly from all states

## Format

<!--
[DATE] — [feature or change]
- What changed
- Why it changed
- Decisions made and reasoning
- Anything tried that didn't work

Example:
[2026-05-07] — habit completion logic
- Created entries table in Supabase
- Derived streak count on read rather than storing it — storing 
  creates sync issues if entries are deleted
- Tried optimistic UI update — reverted due to flicker on 
  slow connections
-->
