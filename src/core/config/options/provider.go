package configOptions

import (
	"os"
	"strconv"
	cmdTypes "trains/src/cli/types"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the providers configuration from command line options.
*/
func ProvidersFromOptions(
	cliConfigOptions cmdTypes.CLIConfigOptions,
) (*configTypes.ProvidersOverrides, *errors.TrainsError) {
	config := make(configTypes.ProvidersOverrides)

	for name, options := range cliConfigOptions.Providers {
		providerName, err := configTypes.FlagToProvider(options.Name)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid provider name",
				Err:     err,
			}
		}

		config[name].Name = &providerName
		config[name].ApiKey = &options.ApiKey
		config[name].Model = &options.Model
		config[name].BaseUrl = &options.BaseUrl
		config[name].Timeout = &options.Timeout
	}

	return &config, nil
}

/*
Resolves the provider configuration from environment variables.
*/
func ProvidersFromEnv() (*configTypes.ProvidersOverrides, *errors.TrainsError) {
	providers := make(configTypes.ProvidersOverrides)

	providers[configTypes.OpenAI] = OpenAIFromEnv()
	providers[configTypes.Anthropic] = AnthropicFromEnv()
	providers[configTypes.Google] = GoogleFromEnv()
	providers[configTypes.Mistral] = MistralFromEnv()
	providers[configTypes.XAI] = XAIFromEnv()
	providers[configTypes.DeepSeek] = DeepSeekFromEnv()
	providers[configTypes.Cohere] = CohereFromEnv()
	providers[configTypes.Perplexity] = PerplexityFromEnv()
	providers[configTypes.OpenRouter] = OpenRouterFromEnv()
	providers[configTypes.MiniMax] = MiniMaxFromEnv()

	return &providers, nil
}

func providerFromEnv(prefix string) *configTypes.ProviderOverrides {
	apiKey := os.Getenv(prefix + "_API_KEY")
	if apiKey == "" {
		return nil
	}

	provider := configTypes.ProviderOverrides{
		ApiKey: &apiKey,
	}

	if v := os.Getenv(prefix + "_MODEL"); v != "" {
		provider.Model = &v
	}
	if v := os.Getenv(prefix + "_BASE_URL"); v != "" {
		provider.BaseUrl = &v
	}
	if v := os.Getenv(prefix + "_TIMEOUT"); v != "" {
		if timeout, err := strconv.Atoi(v); err == nil {
			provider.Timeout = &timeout
		}
	}

	return &provider
}

func OpenAIFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_OPENAI")
	providerName := configTypes.OpenAI
	if provider != nil {
		provider.Name = &providerName
	}
	return provider
}

func AnthropicFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_ANTHROPIC")
	providerName := configTypes.Anthropic
	if provider != nil {
		provider.Name = &providerName
	}
	return provider
}

func GoogleFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_GOOGLE")
	providerName := configTypes.Google
	if provider != nil {
		provider.Name = &providerName
	}
	return provider
}

func MistralFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_MISTRAL")
	providerName := configTypes.Mistral
	if provider != nil {
		provider.Name = &providerName
	}
	return provider
}

func XAIFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_XAI")
	providerName := configTypes.XAI
	if provider != nil {
		provider.Name = &providerName
	}
	return provider
}

func DeepSeekFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_DEEPSEEK")
	providerName := configTypes.DeepSeek
	if provider != nil {
		provider.Name = &providerName
	}
	return provider
}

func CohereFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_COHERE")
	providerName := configTypes.Cohere
	if provider != nil {
		provider.Name = &providerName
	}
	return provider
}

func PerplexityFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_PERPLEXITY")
	providerName := configTypes.Perplexity
	if provider != nil {
		provider.Name = &providerName
	}
	return provider
}

func OpenRouterFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_OPENROUTER")
	providerName := configTypes.OpenRouter
	if provider != nil {
		provider.Name = &providerName
	}
	return provider
}

func MiniMaxFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_MINIMAX")
	providerName := configTypes.MiniMax
	if provider != nil {
		provider.Name = &providerName
	}
	return provider
}

/*
Merges the provider configuration with the given overrides.
*/
func MergeProviderConfig(dest *configTypes.Provider, overrides *configTypes.ProviderOverrides) {
	if overrides.Name != nil {
		dest.Name = *overrides.Name
	}
	if overrides.ApiKey != nil {
		dest.ApiKey = *overrides.ApiKey
	}
	if overrides.Model != nil {
		dest.Model = *overrides.Model
	}
	if overrides.BaseUrl != nil {
		dest.BaseUrl = *overrides.BaseUrl
	}
	if overrides.Timeout != nil {
		dest.Timeout = *overrides.Timeout
	}
}

/*
Merges the providers configuration with the given overrides.
*/
func MergeProvidersConfig(
	dest *configTypes.Providers,
	overrides *configTypes.ProvidersOverrides,
) {
	for name, override := range *overrides {
		if override == nil {
			continue
		}
		if destProvider, ok := (*dest)[name]; ok && destProvider != nil {
			MergeProviderConfig(destProvider, override)
		}
	}
}
