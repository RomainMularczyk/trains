package translation

import (
	configTypes "trains/src/core/config/types"
	"trains/src/core/reader/parser"
	"trains/src/core/reader/translation/providers"
)

type LLM interface {
	Name() string
	Translate(
		translationBatches <-chan parser.TranslationBatch,
		translations chan<- parser.TranslationResult,
		config configTypes.RuntimeConfig,
	)
}

func NewLLM(config configTypes.RuntimeConfig) LLM {
	switch config.Provider.Name {
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
