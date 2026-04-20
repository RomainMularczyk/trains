package parser

import (
	"encoding/json"
	"trains/src/core/config/types"
)

type TranslationBatch struct {
	Units       []TranslationUnit
	Size        int
	Context     string
	ContextSize int
}

func (b TranslationBatch) String() string {
	translationBatch, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(translationBatch)
}

func CreateBatch(
	translationUnit <-chan TranslationUnit,
	translationBatches chan<- TranslationBatch,
	tokenLimit int,
	config types.Config,
) {
	batch := TranslationBatch{}

	for unit := range translationUnit {
		// if the batch is full, we send it to the output channel and create a new one
		if batch.Size+batch.ContextSize > tokenLimit && len(batch.Units) > 0 {
			translationBatches <- batch
			batch = TranslationBatch{}
		}
		batch.Size += unit.EstimateTokenNumber()
		batch.Units = append(batch.Units, unit)
	}

	if len(batch.Units) > 0 {
		translationBatches <- batch
	}
}
