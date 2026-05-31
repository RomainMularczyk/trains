package resolvers

import (
	cmdTypes "trains/src/cli/types"
	configOptions "trains/src/core/config/options"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the configuration from command line options.
*/
func FromOptions(
	cliConfigOptions cmdTypes.CLIConfigOptions,
	bootstrapLogger *configTypes.StageLogger,
) (*configTypes.ConfigFileOverrides, *errors.TrainsError) {
	log := bootstrapLogger.Log.With("component", "config.resolvers.options")

	batchingConfig, err := configOptions.BatchingFromOptions(cliConfigOptions)
	if err != nil {
		return nil, err
	}
	log.Debug("Batching configuration resolved successfully.")

	ioConfig, err := configOptions.IOFromOptions(cliConfigOptions)
	if err != nil {
		return nil, err
	}
	log.Debug("IO configuration resolved successfully.")

	promptConfig, err := configOptions.PromptFromOptions(cliConfigOptions)
	if err != nil {
		return nil, err
	}
	log.Debug("Prompt configuration resolved successfully.")

	providersConfig, err := configOptions.ProvidersFromOptions(cliConfigOptions)
	if err != nil {
		return nil, err
	}

	selectedProviderConfig, err := configOptions.SelectedProviderFromOptions(cliConfigOptions)
	if err != nil {
		return nil, err
	}

	translationConfig, err := configOptions.TranslationFromOptions(cliConfigOptions)
	if err != nil {
		return nil, err
	}

	overrides := &configTypes.ConfigFileOverrides{
		Batching: &configTypes.BatchingOverrides{
			TokenLimit: &batchingConfig.TokenLimit,
			UnitLimit:  &batchingConfig.UnitLimit,
		},
		IO: &configTypes.IOOverrides{
			InputFormat:  &ioConfig.InputFormat,
			OutputFormat: &ioConfig.OutputFormat,
			SourcePath:   &ioConfig.SourcePath,
			TargetPath:   &ioConfig.TargetPath,
		},
		Prompt: &configTypes.PromptOverrides{
			Context: &promptConfig.Context,
		},
		Providers:        providersConfig,
		SelectedProvider: selectedProviderConfig,
		Translation: &configTypes.TranslationOverrides{
			SourceLanguage: &translationConfig.SourceLanguage,
			TargetLanguage: &translationConfig.TargetLanguage,
		},
	}

	return overrides, nil
}
