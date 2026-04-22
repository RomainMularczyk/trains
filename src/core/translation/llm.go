package translation

import (
	"trains/src/core/config/types"
	"trains/src/core/parser"
	"trains/src/core/translation/providers"
)

type LLM interface {
	Name() string
	Translate(
		config types.RuntimeConfig,
		translationBatches <-chan parser.TranslationBatch,
		translations chan<- string,
	)
}

func NewLLM(config types.RuntimeConfig) LLM {
	switch config.SelectedProvider.Name {
	case types.OpenAI:
		return &providers.OpenAI{}
	case types.Anthropic:
		return &providers.Anthropic{}
	case types.Mistral:
		return &providers.Mistral{}
	default:
		return &providers.OpenAI{}
	}
}
