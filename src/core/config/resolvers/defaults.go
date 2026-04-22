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
	// Config defaults
	DefaultConfigPath = "./trains.json"
	// Lock defaults
	DefaultLockVersion = 1
	DefaultLockPath    = "./trains-lock.json"
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

	config := configTypes.Config{
		Path: DefaultConfigPath,
	}

	io := configTypes.IO{
		InputFormat:  DefaultInputFormat,
		OutputFormat: DefaultOutputFormat,
		SourcePath:   DefaultSourcePath,
		TargetPath:   DefaultTargetPath,
	}

	language := configTypes.Translation{
		SourceLanguage: DefaultSourceLanguage,
		TargetLanguage: DefaultTargetLanguage,
	}

	lock := configTypes.Lock{
		Version: DefaultLockVersion,
		Path:    DefaultLockPath,
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

	configFile := configTypes.ConfigFile{
		Batching: batching,
		Config:   config,
		IO:       io,
		Lock:     lock,
		Prompt:   prompt,
		Provider: configTypes.Providers{
			OpenAI: &provider,
		},
		Translation: language,
	}

	return &configFile
}
