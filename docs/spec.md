# Dockyard — Product Specification
 
## What is it?
Dockyard is a cross-platform TUI (terminal user interface) tool for browsing, applying, and managing Starship prompt themes. It wraps Starship's config system with a friendly interactive interface, handling Nerd Font detection, preset browsing, simulated previews, and config writing — without the user ever touching a file.
 
## Who is it for?
Developers who are comfortable in a terminal but aren't deep customisers. They want a good-looking prompt without spending an afternoon in dotfiles. They know what Starship is (or are willing to install it), but find manual TOML editing a barrier.
 
## What problem does it solve?
Setting up a Starship theme properly involves installing fonts, editing config files, cross-referencing glyph cheat sheets, and guessing what things will look like. Most people either give up or spend hours on it. Dockyard removes every step of that friction.
 
## V1 Core Features
1. Starship detection — detect if installed, provide guided install if missing
2. Nerd Font detection — detect if a Nerd Font is active, guide install if not
3. Preset browser — navigable list of all built-in Starship presets with descriptions
4. Simulated prompt examples — show what the selected preset looks like in common scenarios
5. Instant apply and revert — apply a preset with one keypress, revert to previous with one keypress
## Success Definition
V1 is successful when it is publicly released, announced in relevant developer and terminal communities (r/unixporn, r/commandline, r/starship), and receives genuine organic traction — stars, shares, and real user feedback.
 
## Explicitly Out of Scope for V1
- Custom module editing (colours, icons, individual segment configuration)
- Non-Starship prompt managers (Oh My Posh, Powerlevel10k, etc.)
- Custom or community theme importing
- GUI version
- Cloud sync or sharing features
- Multi-level undo or config history
## Monetisation
None. Dockyard is open source and free. It is a credibility-building tool released under the Merrick Wykman name.
 
---
