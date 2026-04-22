package resolvers

import configTypes "trains/src/core/config/types"

const (
	// Batching defaults
	DefaultTokenLimit           = 10_000
	DefaultTranslationUnitLimit = 30
	// IO defaults
	DefaultInputFormat  = configTypes.JSON
	DefaultOutputFormat = configTypes.JSON
	DefaultSourcePath   = "./source"
	DefaultTargetPath   = "./target"
	// Prompt defaults
	DefaultPromptContext = "Translate the following text"
	// Provider defaults
	DefaultProviderName    = configTypes.OpenAI
	DefaultProviderApiKey  = ""
	DefaultProviderModel   = "gpt-5.2-mini"
	DefaultProviderBaseUrl = "https://api.openai.com"
	DefaultProviderTimeout = 30
	// Translation defaults
	DefaultSourceLanguage = configTypes.English
	DefaultTargetLanguage = configTypes.French
)

/*
Resolves the configuration from default values.
*/
func FromDefaults() *configTypes.ConfigFile {
	batching := configTypes.Batching{
		TokenLimit: DefaultTokenLimit,
		UnitLimit:  DefaultTranslationUnitLimit,
	}

	io := configTypes.IO{
		InputFormat:  DefaultInputFormat,
		OutputFormat: DefaultOutputFormat,
		SourcePath:   DefaultSourcePath,
		TargetPath:   DefaultTargetPath,
	}

	prompt := configTypes.Prompt{
		Context: DefaultPromptContext,
	}

	provider := configTypes.Provider{
		Name:    DefaultProviderName,
		ApiKey:  DefaultProviderApiKey,
		Model:   DefaultProviderModel,
		BaseUrl: DefaultProviderBaseUrl,
		Timeout: DefaultProviderTimeout,
	}

	translation := configTypes.Translation{
		SourceLanguage: DefaultSourceLanguage,
		TargetLanguage: DefaultTargetLanguage,
	}

	config := configTypes.ConfigFile{
		Provider: configTypes.Providers{
			OpenAI: &provider,
		},
		Translation: translation,
		Prompt:      prompt,
		Batching:    batching,
		IO:          io,
	}

	return &config
}
