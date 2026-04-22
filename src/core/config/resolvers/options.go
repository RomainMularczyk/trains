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
) (*configTypes.ConfigFileOverrides, *errors.TrainsError) {
	batchingConfig, err := configOptions.BatchingFromOptions(cliConfigOptions)
	if err != nil {
		return nil, err
	}

	ioConfig, err := configOptions.IOFromOptions(cliConfigOptions)
	if err != nil {
		return nil, err
	}

	promptConfig, err := configOptions.PromptFromOptions(cliConfigOptions)
	if err != nil {
		return nil, err
	}

	providerConfig, err := configOptions.ProviderFromOptions(cliConfigOptions)
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
		Translation: &configTypes.TranslationOverrides{
			SourceLanguage: &translationConfig.SourceLanguage,
			TargetLanguage: &translationConfig.TargetLanguage,
		},
	}

	// Provider overrides need special handling based on provider name
	if providerConfig.Name != "" {
		providerOverrides := &configTypes.ProviderOverrides{
			ApiKey:  &providerConfig.ApiKey,
			Model:   &providerConfig.Model,
			BaseUrl: &providerConfig.BaseUrl,
			Timeout: &providerConfig.Timeout,
		}

		switch providerConfig.Name {
		case configTypes.OpenAI:
			overrides.Provider = &configTypes.ProvidersOverrides{OpenAI: providerOverrides}
		case configTypes.Anthropic:
			overrides.Provider = &configTypes.ProvidersOverrides{Anthropic: providerOverrides}
		case configTypes.Google:
			overrides.Provider = &configTypes.ProvidersOverrides{Google: providerOverrides}
		case configTypes.Mistral:
			overrides.Provider = &configTypes.ProvidersOverrides{Mistral: providerOverrides}
		case configTypes.XAI:
			overrides.Provider = &configTypes.ProvidersOverrides{XAI: providerOverrides}
		case configTypes.DeepSeeker:
			overrides.Provider = &configTypes.ProvidersOverrides{DeepSeeker: providerOverrides}
		case configTypes.Cohere:
			overrides.Provider = &configTypes.ProvidersOverrides{Cohere: providerOverrides}
		case configTypes.Perplexity:
			overrides.Provider = &configTypes.ProvidersOverrides{Perplexity: providerOverrides}
		case configTypes.OpenRouter:
			overrides.Provider = &configTypes.ProvidersOverrides{OpenRouter: providerOverrides}
		case configTypes.MiniMax:
			overrides.Provider = &configTypes.ProvidersOverrides{MiniMax: providerOverrides}
		}
	}

	return overrides, nil
}
