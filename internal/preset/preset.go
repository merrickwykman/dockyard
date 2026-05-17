package preset

// Sep controls how the simulated examples are rendered:
//
//	"powerline" – uses  arrow separators between segments
//	"bracket"   – wraps segments in [ ]
//	"pure"      – minimal, no icons, bare branch name
//	"nofont"    – plain ASCII, no Nerd Font glyphs
//	"plain"     – default Nerd Font icon style
type Preset struct {
	Name    string
	Desc    string
	Sep     string
	ShowVer bool // whether to show runtime version numbers
}

func List() []Preset {
	return []Preset{
		{"bracketed-segments", "Segments wrapped in square brackets", "bracket", true},
		{"catppuccin-powerline", "Catppuccin palette with Powerline arrows", "powerline", true},
		{"gruvbox-rainbow", "Warm Gruvbox colors with rainbow segments", "powerline", true},
		{"jetpack", "Fast and clean with Nerd Font icons", "plain", true},
		{"nerd-font-symbols", "Rich icon set using Nerd Font glyphs", "plain", true},
		{"no-empty-icons", "Icons hidden when segment has no content", "plain", true},
		{"no-nerd-font", "Clean layout without Nerd Font icons", "nofont", true},
		{"no-runtime-versions", "Hides language runtime version numbers", "plain", false},
		{"pastel-powerline", "Soft pastel colors with Powerline separators", "powerline", true},
		{"plain-text-symbols", "Text-only, no special characters required", "nofont", true},
		{"pure-preset", "Minimal single-line prompt inspired by Pure", "pure", true},
		{"tokyo-night", "Dark theme with Tokyo Night color palette", "plain", true},
	}
}
