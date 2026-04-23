package providers

import (
	"context"
	"fmt"
	"os"
	"trains/src/core/config/types"
	"trains/src/core/parser"

	"github.com/zendev-sh/goai"
	"github.com/zendev-sh/goai/provider/anthropic"
)

type Anthropic struct{}

func (a *Anthropic) Name() string {
	return "Anthropic"
}

func (a *Anthropic) Translate(
	config configTypes.RuntimeConfig,
	translationBatches <-chan parser.TranslationBatch,
	translations chan<- string,
) {
	if config.SelectedProvider.ApiKey != "" {
		os.Setenv("ANTHROPIC_API_KEY", config.SelectedProvider.ApiKey)
	}
	model := anthropic.Chat(config.SelectedProvider.Model)

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
		fmt.Println(result.Text)
		if err != nil {
			continue
		}

		translations <- result.Text
	}
}
