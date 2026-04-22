package configOptions

import (
	"os"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

func PromptFromOptions(config *configTypes.ConfigFile) {
}

/*
Resolves the prompt configuration from environment variables.
*/
func PromptFromEnv() (*configTypes.Prompt, *errors.TrainsError) {
	config := configTypes.Prompt{}

	// Prompt
	if v := os.Getenv("TRAINS_PROMPT_CONTEXT"); v != "" {
		config.Context = v
	}

	return &config, nil
}
