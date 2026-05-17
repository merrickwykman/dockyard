# Dockyard — TUI Design
 
## Design Philosophy
Neutral and minimal. Dockyard's interface is chrome — it exists to surface the presets, not to compete with them. No strong brand colours. No decorative elements. The preset examples are the visual centrepiece.
 
## Layout
 
```
┌─────────────────────┬──────────────────────────────────────┐
│ Presets             │ Tokyo Night                          │
│                     │                                      │
│ > Tokyo Night       │ A clean dark theme inspired by the  │
│   Pastel Powerline  │ Tokyo Night colour palette.          │
│   Gruvbox Rainbow   │                                      │
│   Bracketed Segments│ Preview                              │
│   Nerd Font Symbols │                                      │
│   Pure Preset       │ ~/projects/dockyard                  │
│   ...               │ ~/projects/dockyard on  main        │
│                     │ ~/projects/dockyard on  main [+]    │
│                     │ ~/projects/dockyard via  node v18   │
│                     │                                      │
│                     │ ─────────────────────────────────── │
│                     │ Enter apply  R revert  ? help  Q quit│
└─────────────────────┴──────────────────────────────────────┘
```
 
## Colour Usage
- Background: terminal default (no override)
- List selection highlight: subtle — single colour accent, no bold background
- Detail pane text: terminal default
- Simulated prompt examples: render using actual preset symbol characters
- Status/key hints bar: muted, low contrast
- Error and confirmation messages: brief inline text, no modal dialogs
## Interaction Model
- Arrow keys navigate the preset list
- Detail pane updates in real time as selection changes
- Enter applies immediately — no confirmation dialog
- R reverts immediately if backup exists — no confirmation dialog
- ? shows a brief help overlay
- Q exits cleanly at any time
## Terminal Compatibility
- Minimum width: 80 columns
- Minimum height: 24 rows
- Graceful degradation below minimum — message shown rather than broken layout
- No mouse support required in v1
- Works in Windows Terminal, iTerm2, Alacritty, GNOME Terminal, and most modern emulators