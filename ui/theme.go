package ui

import (
	"sort"

	"github.com/CuriousFurBytes/diny/ui/themes"
)

var currentTheme *themes.Theme

// customThemes holds themes loaded from user-provided JSON/TOML files.
var customThemes = make(map[string]*themes.Theme)

func init() {
	currentTheme = themes.Catppuccin()
}

// RegisterCustomThemes loads custom themes from the config directory and from
// explicit path mappings (custom_themes config field). Call this after config is loaded.
func RegisterCustomThemes(pathMap map[string]string) {
	// Load themes from ~/.config/diny/themes/ directory
	dirThemes := themes.LoadAllCustomThemes()
	for key, theme := range dirThemes {
		customThemes[key] = theme
	}

	// Load themes from explicit paths in config
	for key, path := range pathMap {
		theme, err := themes.LoadCustomTheme(path)
		if err != nil {
			continue
		}
		customThemes[key] = theme
	}
}

func SetTheme(name string) bool {
	var theme *themes.Theme
	switch name {
	case "catppuccin":
		theme = themes.Catppuccin()
	case "tokyo":
		theme = themes.Tokyo()
	case "nord":
		theme = themes.Nord()
	case "dracula":
		theme = themes.Dracula()
	case "gruvbox-dark":
		theme = themes.GruvboxDark()
	case "onedark":
		theme = themes.OneDark()
	case "monokai":
		theme = themes.Monokai()
	case "solarized-dark":
		theme = themes.SolarizedDark()
	case "solarized-light":
		theme = themes.SolarizedLight()
	case "github-light":
		theme = themes.GithubLight()
	case "gruvbox-light":
		theme = themes.GruvboxLight()
	case "everforest-dark":
		theme = &themes.EverforestDark
	case "flexoki-dark":
		theme = themes.FlexokiDark()
	case "flexoki-light":
		theme = themes.FlexokiLight()
	default:
		// Check custom themes
		if ct, ok := customThemes[name]; ok {
			theme = ct
		} else {
			return false
		}
	}
	currentTheme = theme
	return true
}

func GetCurrentTheme() *themes.Theme {
	return currentTheme
}

func GetAvailableThemes() []string {
	builtIn := []string{
		"catppuccin",
		"tokyo",
		"nord",
		"dracula",
		"gruvbox-dark",
		"onedark",
		"monokai",
		"solarized-dark",
		"everforest-dark",
		"flexoki-dark",
		"solarized-light",
		"github-light",
		"gruvbox-light",
		"flexoki-light",
	}

	custom := GetCustomThemeKeys()
	return append(builtIn, custom...)
}

func GetDarkThemes() []string {
	dark := []string{
		"catppuccin",
		"tokyo",
		"nord",
		"dracula",
		"gruvbox-dark",
		"onedark",
		"monokai",
		"solarized-dark",
		"everforest-dark",
		"flexoki-dark",
	}

	for key, theme := range customThemes {
		if theme.IsDark {
			dark = append(dark, key)
		}
	}

	return dark
}

func GetLightThemes() []string {
	light := []string{
		"solarized-light",
		"github-light",
		"gruvbox-light",
		"flexoki-light",
	}

	for key, theme := range customThemes {
		if !theme.IsDark {
			light = append(light, key)
		}
	}

	return light
}

// GetCustomThemeKeys returns sorted keys of all registered custom themes.
func GetCustomThemeKeys() []string {
	keys := make([]string, 0, len(customThemes))
	for k := range customThemes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
