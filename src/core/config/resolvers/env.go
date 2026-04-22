package resolvers

import (
	configOptions "trains/src/core/config/options"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the configuration from environment variables.
*/
func FromEnv() (*configTypes.ConfigFileOverrides, *errors.TrainsError) {
	batchingConfig, err := configOptions.BatchingFromEnv()
	// TODO: should I let silently fail if env vars are not parsable ?
	if err != nil {
		return nil, err
	}
	ioConfig, err := configOptions.IOFromEnv()
	if err != nil {
		return nil, err
	}
	lockConfig, err := configOptions.LockFromEnv()
	if err != nil {
		return nil, err
	}
	promptConfig, err := configOptions.PromptFromEnv()
	if err != nil {
		return nil, err
	}
	providerConfig, err := configOptions.ProviderFromEnv()
	if err != nil {
		return nil, err
	}
	translationConfig, err := configOptions.TranslationFromEnv()
	if err != nil {
		return nil, err
	}

	// Create a new config with all env vars applied
	newConfig := &configTypes.ConfigFileOverrides{
		Batching:    batchingConfig,
		IO:          ioConfig,
		Lock:        lockConfig,
		Prompt:      promptConfig,
		Provider:    providerConfig,
		Translation: translationConfig,
	}

	return newConfig, nil
}
