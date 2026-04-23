package translation

import (
	"trains/src/core/config/types"
	"trains/src/core/parser"
	"trains/src/core/translation/providers"
)

type LLM interface {
	Name() string
	Translate(
		config configTypes.RuntimeConfig,
		translationBatches <-chan parser.TranslationBatch,
		translations chan<- string,
	)
}

func NewLLM(config configTypes.RuntimeConfig) LLM {
	switch config.SelectedProvider.Name {
	case configTypes.OpenAI:
		return &providers.OpenAI{}
	case configTypes.Anthropic:
		return &providers.Anthropic{}
	case configTypes.Mistral:
		return &providers.Mistral{}
	default:
		return &providers.OpenAI{}
	}
}
