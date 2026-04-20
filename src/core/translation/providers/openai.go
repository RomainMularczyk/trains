package providers

import (
	"context"
	"os"
	"trains/src/core/config/types"
	"trains/src/core/parser"

	"github.com/zendev-sh/goai"
	"github.com/zendev-sh/goai/provider/openai"
)

type OpenAI struct{}

func (o *OpenAI) Name() string {
	return "OpenAI"
}

func (o *OpenAI) Translate(
	config types.Config,
	translationBatches <-chan parser.TranslationBatch,
	translations chan<- parser.TranslationUnit,
) {
	if config.Provider.OpenAI.ApiKey != "" {
		os.Setenv("OPENAI_API_KEY", config.Provider.OpenAI.ApiKey)
	}
	model := openai.Chat(config.Provider.OpenAI.Model)

	for translationBatch := range translationBatches {
		result, err := goai.GenerateText(
			context.Background(),
			model,
			goai.WithPrompt(translationBatch.Context),
		)
		if err != nil {
			// TODO: add error handling
			continue
		}

		for _, unit := range translationBatch.Units {
			translations <- parser.TranslationUnit{
				Fullkey:  unit.Fullkey,
				Path:     unit.Path,
				Source:   unit.Source,
				Target:   result.Text,
				Segments: unit.Segments,
			}
		}
	}
}
