package resolvers

import (
	"trains/src/core/config/options"
	"trains/src/core/config/types"
)

/*
Resolves the configuration from environment variables.
*/
func FromEnv(config *types.ConfigFile) {
	configOptions.BatchingFromEnv(config)
	configOptions.IOFromEnv(config)
	configOptions.PromptFromEnv(config)
	configOptions.ProviderFromEnv(config)
	configOptions.TranslationFromEnv(config)
}
