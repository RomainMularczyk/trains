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
