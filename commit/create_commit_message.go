package commit

import (
	"github.com/CuriousFurBytes/diny/ai"
	"github.com/CuriousFurBytes/diny/config"
)

func CreateCommitMessage(gitDiff string, cfg *config.Config) (string, error) {
	return ai.GenerateCommitMessage(gitDiff, cfg)
}
