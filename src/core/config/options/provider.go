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
) (*configTypes.Provider, *errors.TrainsError) {
	config := configTypes.Provider{}

	if cliConfigOptions.Provider.Name != "" {
		providerName, err := configTypes.FlagToProvider(cliConfigOptions.Provider.Name)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid provider name",
				Err:     err,
			}
		}
		config.Name = providerName
	}

	config.ApiKey = cliConfigOptions.Provider.ApiKey
	config.Model = cliConfigOptions.Provider.Model
	config.BaseUrl = cliConfigOptions.Provider.BaseUrl
	config.Timeout = cliConfigOptions.Provider.Timeout

	return &config, nil
}

/*
Resolves the provider configuration from environment variables.
*/
func ProviderFromEnv() (*configTypes.Providers, *errors.TrainsError) {
	providers := configTypes.Providers{}

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

func providerFromEnv(prefix string) *configTypes.Provider {
	apiKey := os.Getenv(prefix + "_API_KEY")
	if apiKey == "" {
		return nil
	}

	provider := configTypes.Provider{
		Name:   configTypes.ProviderName(""),
		ApiKey: apiKey,
	}

	if v := os.Getenv(prefix + "_MODEL"); v != "" {
		provider.Model = v
	}
	if v := os.Getenv(prefix + "_BASE_URL"); v != "" {
		provider.BaseUrl = v
	}
	if v := os.Getenv(prefix + "_TIMEOUT"); v != "" {
		if timeout, err := strconv.Atoi(v); err == nil {
			provider.Timeout = timeout
		}
	}

	return &provider
}

func OpenAIFromEnv() *configTypes.Provider {
	provider := providerFromEnv("TRAINS_PROVIDER_OPENAI")
	if provider != nil {
		provider.Name = configTypes.OpenAI
	}
	return provider
}

func AnthropicFromEnv() *configTypes.Provider {
	provider := providerFromEnv("TRAINS_PROVIDER_ANTHROPIC")
	if provider != nil {
		provider.Name = configTypes.Anthropic
	}
	return provider
}

func GoogleFromEnv() *configTypes.Provider {
	provider := providerFromEnv("TRAINS_PROVIDER_GOOGLE")
	if provider != nil {
		provider.Name = configTypes.Google
	}
	return provider
}

func MistralFromEnv() *configTypes.Provider {
	provider := providerFromEnv("TRAINS_PROVIDER_MISTRAL")
	if provider != nil {
		provider.Name = configTypes.Mistral
	}
	return provider
}

func XAIFromEnv() *configTypes.Provider {
	provider := providerFromEnv("TRAINS_PROVIDER_XAI")
	if provider != nil {
		provider.Name = configTypes.XAI
	}
	return provider
}

func DeepSeekerFromEnv() *configTypes.Provider {
	provider := providerFromEnv("TRAINS_PROVIDER_DEEPSEEKER")
	if provider != nil {
		provider.Name = configTypes.DeepSeeker
	}
	return provider
}

func CohereFromEnv() *configTypes.Provider {
	provider := providerFromEnv("TRAINS_PROVIDER_COHERE")
	if provider != nil {
		provider.Name = configTypes.Cohere
	}
	return provider
}

func PerplexityFromEnv() *configTypes.Provider {
	provider := providerFromEnv("TRAINS_PROVIDER_PERPLEXITY")
	if provider != nil {
		provider.Name = configTypes.Perplexity
	}
	return provider
}

func OpenRouterFromEnv() *configTypes.Provider {
	provider := providerFromEnv("TRAINS_PROVIDER_OPENROUTER")
	if provider != nil {
		provider.Name = configTypes.OpenRouter
	}
	return provider
}

func MiniMaxFromEnv() *configTypes.Provider {
	provider := providerFromEnv("TRAINS_PROVIDER_MINIMAX")
	if provider != nil {
		provider.Name = configTypes.MiniMax
	}
	return provider
}
