package providers

import (
	"context"
	"fmt"
	"os"
	configTypes "trains/src/core/config/types"
	"trains/src/core/reader/parser"

	"github.com/zendev-sh/goai"
	"github.com/zendev-sh/goai/provider/openai"
)

type OpenAI struct{}

func (o *OpenAI) Name() string {
	return "OpenAI"
}

func (o *OpenAI) Translate(
	config configTypes.RuntimeConfig,
	translationBatches <-chan parser.TranslationBatch,
	translations chan<- string,
) {
	if config.Provider.ApiKey != "" {
		os.Setenv("OPENAI_API_KEY", config.Provider.ApiKey)
	}
	model := openai.Chat(config.Provider.Model)

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

		translations <- result.Text
	}
}
