package translationValidator

import (
	"encoding/json"
	"fmt"

	trainsError "trains/src/core/errors"
	"trains/src/core/reader/parser"

	"github.com/go-playground/validator/v10"
)

type BatchValidator struct{}

/*
Validates the structure of the translation result.
*/
func (b *BatchValidator) structure(
	jsonTranslation []parser.TranslationEngineEntry,
) (*[]parser.TranslationEngineEntry, *trainsError.TrainsError) {
	validate := validator.New()

	var translationEntries []parser.TranslationEngineEntry
	for _, translationEntry := range jsonTranslation {
		if err := validate.Struct(translationEntry); err != nil {
			return nil, &trainsError.TrainsError{
				Code: trainsError.TranslationError,
				Message: fmt.Sprintf(
					"Failed to validate translation result. %s",
					err,
				),
				Err: err,
			}
		}
		translationEntries = append(translationEntries, translationEntry)
	}

	return &translationEntries, nil
}

/*
Validates the length of the translation result.
*/
func (b *BatchValidator) length(
	translationEntries []parser.TranslationEngineEntry,
	translationResult parser.TranslationResult,
) *trainsError.TrainsError {
	if len(translationEntries) != len(translationResult.Batch.Units) {
		return &trainsError.TrainsError{
			Code: trainsError.TranslationError,
			Message: fmt.Sprintf(
				"Failed to validate translation result. The number of units does not match the number of entries.",
			),
			Err: nil,
		}
	}

	return nil
}

func (b *BatchValidator) Keys() {
}

func (b *BatchValidator) Placeholders() {
}

func (b *BatchValidator) Values() {
}

/*
Validates the translation result returned by the engine.
*/
func (b *BatchValidator) Validate(
	translations <-chan parser.TranslationResult,
	validatedTranslations chan<- parser.TranslationResult,
) {
	for translation := range translations {
		var rawTranslationOutput []parser.TranslationEngineEntry
		parsingError := json.Unmarshal([]byte(translation.Result), &rawTranslationOutput)
		translation.Error = &trainsError.TrainsError{
			Code: trainsError.TranslationError,
			Message: fmt.Sprintf(
				"Failed to parse the translation result. %s",
				parsingError,
			),
			Err: parsingError,
		}

		validatedTranslation, err := b.structure(rawTranslationOutput)
		translation.Error = err
		translation.Validated = validatedTranslation

		err = b.length(rawTranslationOutput, translation)
		translation.Error = err

		validatedTranslations <- translation
	}
}
