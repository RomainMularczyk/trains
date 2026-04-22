package resolvers

import (
	"trains/src/core/config/options"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the configuration from environment variables.
*/
func FromEnv(config configTypes.ConfigFile) (*configTypes.ConfigFile, *errors.TrainsError) {
	batchingConfig, err := configOptions.BatchingFromEnv()
	// TODO: should I let silently fail if env vars are not parsable ?
	if err != nil {
		return nil, err
	}
	ioConfig, err := configOptions.IOFromEnv()
	if err != nil {
		return nil, err
	}
	promptConfig, err := configOptions.PromptFromEnv()
	if err != nil {
		return nil, err
	}
	providerConfig := configOptions.ProviderFromEnv()
	if err != nil {
		return nil, err
	}
	translationConfig := configOptions.TranslationFromEnv(config)

	return configTypes.ConfigFile{
		Batching:    batchingConfig,
		IO:          ioConfig,
		Lock:        config.Lock,
		Prompt:      promptConfig,
		Provider:    configTypes.Providers{},
		Translation: translationConfig,
	}
}
