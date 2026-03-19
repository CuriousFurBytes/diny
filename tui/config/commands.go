package configtui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/CuriousFurBytes/diny/config"
)

func doSaveConfig(cfg *config.Config, configPath, configType string) tea.Cmd {
	return func() tea.Msg {
		var err error
		if configType == "global" {
			err = config.Save(cfg, configPath)
		} else {
			err = config.SaveLocal(toLocalConfig(cfg), configPath)
		}
		if err != nil {
			return errMsg{err}
		}
		return savedMsg{}
	}
}

func doSaveAndQuit(cfg *config.Config, configPath, configType string) tea.Cmd {
	return func() tea.Msg {
		var err error
		if configType == "global" {
			err = config.Save(cfg, configPath)
		} else {
			err = config.SaveLocal(toLocalConfig(cfg), configPath)
		}
		if err != nil {
			return errMsg{err}
		}
		return saveAndQuitMsg{}
	}
}

func toLocalConfig(cfg *config.Config) *config.LocalConfig {
	conventional := cfg.Commit.Conventional
	emoji := cfg.Commit.Emoji
	hashAfterCommit := cfg.Commit.HashAfterCommit
	aiMode := cfg.AI.Mode

	return &config.LocalConfig{
		Theme:        cfg.Theme,
		CustomThemes: cfg.CustomThemes,
		AI: config.LocalAIConfig{
			Mode:     &aiMode,
			LocalURL: cfg.AI.LocalURL,
			APIURL:   cfg.AI.APIURL,
			APIKey:   cfg.AI.APIKey,
			Model:    cfg.AI.Model,
			Command:  cfg.AI.Command,
		},
		Commit: config.LocalCommitConfig{
			Conventional:       &conventional,
			ConventionalFormat: cfg.Commit.ConventionalFormat,
			Emoji:              &emoji,
			EmojiMap:           cfg.Commit.EmojiMap,
			HashAfterCommit:    &hashAfterCommit,
			Tone:               cfg.Commit.Tone,
			Length:             cfg.Commit.Length,
			CustomInstructions: cfg.Commit.CustomInstructions,
		},
	}
}
