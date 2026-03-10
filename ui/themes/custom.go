package themes

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/lipgloss"
)

// CustomThemeFile represents the structure of a custom theme file (JSON or TOML).
type CustomThemeFile struct {
	Name              string `json:"name" toml:"name"`
	IsDark            bool   `json:"is_dark" toml:"is_dark"`
	PrimaryForeground string `json:"primary_foreground" toml:"primary_foreground"`
	PrimaryBackground string `json:"primary_background" toml:"primary_background"`
	SuccessForeground string `json:"success_foreground" toml:"success_foreground"`
	SuccessBackground string `json:"success_background" toml:"success_background"`
	ErrorForeground   string `json:"error_foreground" toml:"error_foreground"`
	ErrorBackground   string `json:"error_background" toml:"error_background"`
	WarningForeground string `json:"warning_foreground" toml:"warning_foreground"`
	WarningBackground string `json:"warning_background" toml:"warning_background"`
	MutedForeground   string `json:"muted_foreground" toml:"muted_foreground"`
}

func (f *CustomThemeFile) toTheme() *Theme {
	return &Theme{
		Name:              f.Name,
		IsDark:            f.IsDark,
		PrimaryForeground: lipgloss.Color(f.PrimaryForeground),
		PrimaryBackground: lipgloss.Color(f.PrimaryBackground),
		SuccessForeground: lipgloss.Color(f.SuccessForeground),
		SuccessBackground: lipgloss.Color(f.SuccessBackground),
		ErrorForeground:   lipgloss.Color(f.ErrorForeground),
		ErrorBackground:   lipgloss.Color(f.ErrorBackground),
		WarningForeground: lipgloss.Color(f.WarningForeground),
		WarningBackground: lipgloss.Color(f.WarningBackground),
		MutedForeground:   lipgloss.Color(f.MutedForeground),
	}
}

func (f *CustomThemeFile) validate() error {
	if f.Name == "" {
		return fmt.Errorf("theme name is required")
	}
	colors := map[string]string{
		"primary_foreground": f.PrimaryForeground,
		"primary_background": f.PrimaryBackground,
		"success_foreground": f.SuccessForeground,
		"success_background": f.SuccessBackground,
		"error_foreground":   f.ErrorForeground,
		"error_background":   f.ErrorBackground,
		"warning_foreground": f.WarningForeground,
		"warning_background": f.WarningBackground,
		"muted_foreground":   f.MutedForeground,
	}
	for field, val := range colors {
		if val == "" {
			return fmt.Errorf("%s is required", field)
		}
	}
	return nil
}

// LoadCustomTheme loads a theme from a JSON or TOML file.
func LoadCustomTheme(path string) (*Theme, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read theme file: %w", err)
	}

	var themeFile CustomThemeFile
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &themeFile); err != nil {
			return nil, fmt.Errorf("failed to parse JSON theme: %w", err)
		}
	case ".toml":
		if err := toml.Unmarshal(data, &themeFile); err != nil {
			return nil, fmt.Errorf("failed to parse TOML theme: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported theme file format '%s', use .json or .toml", ext)
	}

	if err := themeFile.validate(); err != nil {
		return nil, fmt.Errorf("invalid theme file: %w", err)
	}

	return themeFile.toTheme(), nil
}

// GetThemesDir returns the path to the custom themes directory.
func GetThemesDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".config", "diny", "themes")
}

// LoadAllCustomThemes scans the themes directory and loads all valid theme files.
// Returns a map of theme key (filename without extension) to Theme.
func LoadAllCustomThemes() map[string]*Theme {
	themes := make(map[string]*Theme)
	dir := GetThemesDir()
	if dir == "" {
		return themes
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return themes
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext != ".json" && ext != ".toml" {
			continue
		}

		key := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		path := filepath.Join(dir, entry.Name())

		theme, err := LoadCustomTheme(path)
		if err != nil {
			continue
		}

		themes[key] = theme
	}

	return themes
}
