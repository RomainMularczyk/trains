package configOptions

import (
	"os"
	cmdTypes "trains/src/cli/types"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the prompt configuration from command line options.
*/
func PromptFromOptions(
	cliConfigOptions cmdTypes.CLIConfigOptions,
) (*configTypes.Prompt, *errors.TrainsError) {
	config := configTypes.Prompt{
		Context: cliConfigOptions.Prompt.Context,
	}
	return &config, nil
}

/*
Merges the prompt configuration with the given overrides.
*/
func MergePromptConfig(dest *configTypes.Prompt, overrides *configTypes.PromptOverrides) {
	if overrides.Context != nil {
		dest.Context = *overrides.Context
	}
}

/*
Resolves the prompt configuration from environment variables.
*/
func PromptFromEnv() (*configTypes.PromptOverrides, *errors.TrainsError) {
	config := configTypes.PromptOverrides{}

	// Prompt
	if v := os.Getenv("TRAINS_PROMPT_CONTEXT"); v != "" {
		config.Context = &v
	}

	return &config, nil
}
