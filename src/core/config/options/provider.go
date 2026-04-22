package configOptions

import (
	"os"
	"strconv"
	cmdTypes "trains/src/cli/types"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the provider configuration from command line options.
*/
func ProviderFromOptions(
	cliConfigOptions cmdTypes.CLIConfigOptions,
) (*configTypes.ProviderOverrides, *errors.TrainsError) {
	config := configTypes.ProviderOverrides{}

	if cliConfigOptions.Provider.Name != "" {
		providerName, err := configTypes.FlagToProvider(cliConfigOptions.Provider.Name)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid provider name",
				Err:     err,
			}
		}
		config.Name = &providerName
	}

	config.ApiKey = &cliConfigOptions.Provider.ApiKey
	config.Model = &cliConfigOptions.Provider.Model
	config.BaseUrl = &cliConfigOptions.Provider.BaseUrl
	config.Timeout = &cliConfigOptions.Provider.Timeout

	return &config, nil
}

/*
Resolves the provider configuration from environment variables.
*/
func ProviderFromEnv() (*configTypes.ProvidersOverrides, *errors.TrainsError) {
	providers := configTypes.ProvidersOverrides{}

	if openAI := OpenAIFromEnv(); openAI != nil {
		providers.OpenAI = openAI
	}
	if anthropic := AnthropicFromEnv(); anthropic != nil {
		providers.Anthropic = anthropic
	}
	if google := GoogleFromEnv(); google != nil {
		providers.Google = google
	}
	if mistral := MistralFromEnv(); mistral != nil {
		providers.Mistral = mistral
	}
	if xAI := XAIFromEnv(); xAI != nil {
		providers.XAI = xAI
	}
	if deepSeeker := DeepSeekerFromEnv(); deepSeeker != nil {
		providers.DeepSeeker = deepSeeker
	}
	if cohere := CohereFromEnv(); cohere != nil {
		providers.Cohere = cohere
	}
	if perplexity := PerplexityFromEnv(); perplexity != nil {
		providers.Perplexity = perplexity
	}
	if openRouter := OpenRouterFromEnv(); openRouter != nil {
		providers.OpenRouter = openRouter
	}
	if miniMax := MiniMaxFromEnv(); miniMax != nil {
		providers.MiniMax = miniMax
	}

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

func DeepSeekerFromEnv() *configTypes.ProviderOverrides {
	provider := providerFromEnv("TRAINS_PROVIDER_DEEPSEEKER")
	providerName := configTypes.DeepSeeker
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
func MergeProvidersConfig(dest *configTypes.Providers, overrides *configTypes.ProvidersOverrides) {
	if overrides.OpenAI != nil && dest.OpenAI != nil {
		MergeProviderConfig(dest.OpenAI, overrides.OpenAI)
	}
	if overrides.Anthropic != nil && dest.Anthropic != nil {
		MergeProviderConfig(dest.Anthropic, overrides.Anthropic)
	}
	if overrides.Google != nil && dest.Google != nil {
		MergeProviderConfig(dest.Google, overrides.Google)
	}
	if overrides.Mistral != nil && dest.Mistral != nil {
		MergeProviderConfig(dest.Mistral, overrides.Mistral)
	}
	if overrides.XAI != nil && dest.XAI != nil {
		MergeProviderConfig(dest.XAI, overrides.XAI)
	}
	if overrides.DeepSeeker != nil && dest.DeepSeeker != nil {
		MergeProviderConfig(dest.DeepSeeker, overrides.DeepSeeker)
	}
	if overrides.Cohere != nil && dest.Cohere != nil {
		MergeProviderConfig(dest.Cohere, overrides.Cohere)
	}
	if overrides.Perplexity != nil && dest.Perplexity != nil {
		MergeProviderConfig(dest.Perplexity, overrides.Perplexity)
	}
	if overrides.OpenRouter != nil && dest.OpenRouter != nil {
		MergeProviderConfig(dest.OpenRouter, overrides.OpenRouter)
	}
	if overrides.MiniMax != nil && dest.MiniMax != nil {
		MergeProviderConfig(dest.MiniMax, overrides.MiniMax)
	}
}
