package preset

type Preset struct {
	Name string
	Desc string
}

func List() []Preset {
	return []Preset{
		{"bracketed-segments", "Segments wrapped in square brackets"},
		{"catppuccin-powerline", "Catppuccin palette with Powerline arrows"},
		{"gruvbox-rainbow", "Warm Gruvbox colors with rainbow segments"},
		{"jetpack", "Fast and clean with Nerd Font icons"},
		{"nerd-font-symbols", "Rich icon set using Nerd Font glyphs"},
		{"no-empty-icons", "Icons hidden when segment has no content"},
		{"no-nerd-font", "Clean layout without Nerd Font icons"},
		{"no-runtime-versions", "Hides language runtime version numbers"},
		{"pastel-powerline", "Soft pastel colors with Powerline separators"},
		{"plain-text-symbols", "Text-only, no special characters required"},
		{"pure-preset", "Minimal single-line prompt inspired by Pure"},
		{"tokyo-night", "Dark theme with Tokyo Night color palette"},
	}
}
