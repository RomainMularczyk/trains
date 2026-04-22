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
	XAI        ProviderName = "xai"
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
	XAI        *Provider `json:"xai"`
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
	XAI        *ProviderOverrides
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
	if o == nil {
		return
	}
	if o.OpenAI != nil && p.OpenAI != nil {
		p.OpenAI.Apply(o.OpenAI)
	}
	if o.Anthropic != nil && p.Anthropic != nil {
		p.Anthropic.Apply(o.Anthropic)
	}
	if o.Google != nil && p.Google != nil {
		p.Google.Apply(o.Google)
	}
	if o.Mistral != nil && p.Mistral != nil {
		p.Mistral.Apply(o.Mistral)
	}
	if o.XAI != nil && p.XAI != nil {
		p.XAI.Apply(o.XAI)
	}
	if o.DeepSeeker != nil && p.DeepSeeker != nil {
		p.DeepSeeker.Apply(o.DeepSeeker)
	}
	if o.Cohere != nil && p.Cohere != nil {
		p.Cohere.Apply(o.Cohere)
	}
	if o.Perplexity != nil && p.Perplexity != nil {
		p.Perplexity.Apply(o.Perplexity)
	}
	if o.OpenRouter != nil && p.OpenRouter != nil {
		p.OpenRouter.Apply(o.OpenRouter)
	}
	if o.MiniMax != nil && p.MiniMax != nil {
		p.MiniMax.Apply(o.MiniMax)
	}
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
		return XAI, nil
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
	case XAI:
		return config.Provider.XAI
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
