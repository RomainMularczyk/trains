package providers

import (
	"context"
	"fmt"
	"os"
	"trains/src/core/config/types"
	"trains/src/core/reader/parser"

	"github.com/zendev-sh/goai"
	"github.com/zendev-sh/goai/provider/mistral"
)

type Mistral struct{}

func (m *Mistral) Name() string {
	return "Mistral"
}

func (m *Mistral) Translate(
	translationBatches <-chan parser.TranslationBatch,
	translations chan<- parser.TranslationResult,
	config configTypes.RuntimeConfig,
) {
	if config.Provider.ApiKey != "" {
		os.Setenv("MISTRAL_API_KEY", config.Provider.ApiKey)
	}
	model := mistral.Chat(config.Provider.Model)

	for translationBatch := range translationBatches {
		prompt := fmt.Sprintf(
			configTypes.SYSTEM_PROMPT,
			translationBatch.Context,
			"Target language: French",
			translationBatch.Units,
		)
		result, err := goai.GenerateText(
			context.Background(),
			model,
			goai.WithPrompt(prompt),
		)
		if err != nil {
			continue
		}

		translations <- parser.TranslationResult{
			Batch:  translationBatch,
			Result: result.Text,
		}
	}
}
