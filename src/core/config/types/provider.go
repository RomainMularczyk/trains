package configTypes

import (
	"fmt"
	"trains/src/core/errors"
)

type ProviderName string

const (
	OpenAI     ProviderName = "openai"
	Anthropic  ProviderName = "anthropic"
	Google     ProviderName = "google"
	Mistral    ProviderName = "mistral"
	XAI        ProviderName = "xai"
	DeepSeek   ProviderName = "deepseek"
	Cohere     ProviderName = "cohere"
	Perplexity ProviderName = "perplexity"
	OpenRouter ProviderName = "openrouter"
	MiniMax    ProviderName = "minimax"
)

type Providers map[ProviderName]*Provider

type Provider struct {
	Name    ProviderName `validate:"required,oneof=openai anthropic google mistral xai deepseeker cohere perplexity openrouter minimax"`
	ApiKey  string       `validate:"required"`
	Model   string       `validate:"required"`
	BaseUrl string       `validate:"required"`
	Timeout int          `validate:"required"`
}

type ProvidersOverrides map[ProviderName]*ProviderOverrides

type ProviderOverrides struct {
	Name    *ProviderName
	ApiKey  *string
	Model   *string
	BaseUrl *string
	Timeout *int
}

/*
Applies the given provider overrides to the provider configuration.
*/
func (p *Provider) Apply(o *ProviderOverrides) {
	if o == nil {
		return
	}
	if o.ApiKey != nil {
		p.ApiKey = *o.ApiKey
	}
	if o.Model != nil {
		p.Model = *o.Model
	}
	if o.BaseUrl != nil {
		p.BaseUrl = *o.BaseUrl
	}
	if o.Timeout != nil {
		p.Timeout = *o.Timeout
	}
}

/*
Applies the given providers overrides to the providers configuration.
*/
func (p *Providers) Apply(o *ProvidersOverrides) {
	for name, override := range *o {
		if override == nil {
			continue
		}
		if provider, ok := (*p)[name]; ok && provider != nil {
			provider.Apply(override)
		}
	}
}

func FlagToProvider(flag string) (ProviderName, *errors.TrainsError) {
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
		return XAI, nil
	case "deepseeker":
		return DeepSeek, nil
	case "cohere":
		return Cohere, nil
	case "perplexity":
		return Perplexity, nil
	case "openrouter":
		return OpenRouter, nil
	case "minimax":
		return MiniMax, nil
	default:
		return "", &errors.TrainsError{
			Code:    errors.InvalidConfigError,
			Message: "Unsupported provider",
			Err:     fmt.Errorf("Unsupported provider: %s", flag),
		}
	}
}

/*
Returns the provider config for the given provider name.
*/
func (p Providers) Get(name ProviderName) (*Provider, *errors.TrainsError) {
	provider, ok := p[name]
	if !ok || provider == nil {
		return nil, &errors.TrainsError{
			Code:    errors.InvalidConfigError,
			Message: "Invalid provider name",
			Err:     fmt.Errorf("Invalid provider name: %s", name),
		}
	}
	return provider, nil
}
