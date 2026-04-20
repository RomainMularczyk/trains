package providers

import (
	"context"
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
	config types.Config,
	translationBatches <-chan parser.TranslationBatch,
	translations chan<- parser.TranslationUnit,
) {
	if config.Provider.Anthropic.ApiKey != "" {
		os.Setenv("ANTHROPIC_API_KEY", config.Provider.Anthropic.ApiKey)
	}
	model := anthropic.Chat(config.Provider.Anthropic.Model)

	for translationBatch := range translationBatches {
		result, err := goai.GenerateText(
			context.Background(),
			model,
			goai.WithPrompt(translationBatch.Context),
		)
		if err != nil {
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
