package configOptions

import (
	"os"
	"trains/src/core/config/types"
)

func PromptFromEnv(config *types.Config) {
	// Prompt
	if v := os.Getenv("TRAINS_PROMPT_CONTEXT"); v != "" {
		config.Prompt.Context = v
	}
}
