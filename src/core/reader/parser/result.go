package parser

import trainsError "trains/src/core/errors"

type TranslationEngineEntry struct {
	Key    string `json:"Key" validate:"required,min=1"`
	Target string `json:"Target" validate:"required,min=1"`
}

type TranslationResult struct {
	Batch            TranslationBatch
	Result           string
	Validated        []TranslationEngineEntry
	ValidationErrors *TranslationValiationErrors
}

type TranslationValiationErrors struct {
	Response []*trainsError.TrainsError
	Batch    []*trainsError.TrainsError
	Entry    []*trainsError.TrainsError
}
