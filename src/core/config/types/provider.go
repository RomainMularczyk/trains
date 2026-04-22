package configTypes

import (
	"fmt"
)

type ProviderName string

const (
	OpenAI     ProviderName = "openai"
	Anthropic  ProviderName = "anthropic"
	Google     ProviderName = "google"
	Mistral    ProviderName = "mistral"
	xAI        ProviderName = "xai"
	DeepSeeker ProviderName = "deepseeker"
	Cohere     ProviderName = "cohere"
	Perplexity ProviderName = "perplexity"
	OpenRouter ProviderName = "openrouter"
	MiniMax    ProviderName = "minimax"
)

type Providers struct {
	OpenAI     *Provider `json:"openai"`
	Anthropic  *Provider `json:"anthropic"`
	Google     *Provider `json:"google"`
	Mistral    *Provider `json:"mistral"`
	xAI        *Provider `json:"xai"`
	DeepSeeker *Provider `json:"deepseeker"`
	Cohere     *Provider `json:"cohere"`
	Perplexity *Provider `json:"perplexity"`
	OpenRouter *Provider `json:"openrouter"`
	MiniMax    *Provider `json:"minimax"`
}

type ProvidersOverrides struct {
	OpenAI     *ProviderOverrides
	Anthropic  *ProviderOverrides
	Google     *ProviderOverrides
	Mistral    *ProviderOverrides
	xAI        *ProviderOverrides
	DeepSeeker *ProviderOverrides
	Cohere     *ProviderOverrides
	Perplexity *ProviderOverrides
	OpenRouter *ProviderOverrides
	MiniMax    *ProviderOverrides
}

type Provider struct {
	Name    ProviderName `validate:"required"`
	ApiKey  string       `validate:"required"`
	Model   string       `validate:"required"`
	BaseUrl string       `validate:"required"`
	Timeout int          `validate:"required"`
}

type ProviderOverrides struct {
	ApiKey  *string
	Model   *string
	BaseUrl *string
	Timeout *int
}

func FlagToProvider(flag string) (ProviderName, error) {
	switch flag {
	case "openai":
		return OpenAI, nil
	case "anthropic":
		return Anthropic, nil
	case "google":
		return Google, nil
	case "mistral":
		return Mistral, nil
	case "xai":
		return xAI, nil
	case "deepseeker":
		return DeepSeeker, nil
	case "cohere":
		return Cohere, nil
	case "perplexity":
		return Perplexity, nil
	case "openrouter":
		return OpenRouter, nil
	case "minimax":
		return MiniMax, nil
	default:
		return "", fmt.Errorf("Unsupported provider: %s", flag)
	}
}

/*
Returns the provider config for the given provider name.
*/
func ProviderNameToProviderConfig(name ProviderName, config ConfigFile) *Provider {
	switch name {
	case OpenAI:
		return config.Provider.OpenAI
	case Anthropic:
		return config.Provider.Anthropic
	case Google:
		return config.Provider.Google
	case Mistral:
		return config.Provider.Mistral
	case xAI:
		return config.Provider.xAI
	case DeepSeeker:
		return config.Provider.DeepSeeker
	case Cohere:
		return config.Provider.Cohere
	case Perplexity:
		return config.Provider.Perplexity
	case OpenRouter:
		return config.Provider.OpenRouter
	case MiniMax:
		return config.Provider.MiniMax
	default:
		return config.Provider.OpenAI
	}
}
