package cmd

import (
	"github.com/CuriousFurBytes/diny/changelog"
	"github.com/spf13/cobra"
)

var changelogCmd = &cobra.Command{
	Use:   "changelog",
	Short: "Generate an AI-powered changelog for your repository",
	Run: func(cmd *cobra.Command, args []string) {
		changelog.Main(AppConfig)
	},
}

func init() {
	rootCmd.AddCommand(changelogCmd)
}
