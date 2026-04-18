package translation

import "encoding/json"

type TranslationBatch struct {
	Units       []TranslationUnit
	Size        int
	Context     string
	ContextSize int
}

type Batch struct{}

func (b TranslationBatch) String() string {
	translationBatch, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(translationBatch)
}

func (b *Batch) Create(translationUnits []TranslationUnit, tokenLimit int) TranslationBatch {
	var batch TranslationBatch
	for _, unit := range translationUnits {
		if batch.Size+batch.ContextSize < tokenLimit {
			batch.Units = append(batch.Units, unit)
		}
	}
	return batch
}
