package translationValidator

import (
	"encoding/json"
	"fmt"
	"slices"

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
				Code: trainsError.TranslationBatchError,
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
	numUnits := len(translationResult.Batch.Units)
	numEntries := len(translationEntries)

	if numUnits != numEntries {
		return &trainsError.TrainsError{
			Code: trainsError.TranslationBatchError,
			Message: fmt.Sprintf(
				`Failed to validate translation result. 
				The number of units does not match the number of entries.
				Reveived %d translation entries, expected %d.`,
				numEntries,
				numUnits,
			),
			Err: nil,
		}
	}

	return nil
}

/*
Validates the keys of the translation result.
*/
func (b *BatchValidator) keys(
	translationEntries []parser.TranslationEngineEntry,
	translationResult parser.TranslationResult,
) *trainsError.TrainsError {
	var keysInSourceFile []string
	for _, translationUnit := range translationResult.Batch.Units {
		keysInSourceFile = append(keysInSourceFile, translationUnit.Fullkey)
	}

	for _, translationEntry := range translationEntries {
		if !slices.Contains(keysInSourceFile, translationEntry.Key) {
			return &trainsError.TrainsError{
				Code: trainsError.TranslationEntryError,
				Message: fmt.Sprintf(
					`Failed to validate translation result. 
					The key %s is not present in the source file.`,
					translationEntry.Key,
				),
				Err: nil,
			}
		}
	}

	return nil
}

/*
Verifies that each placeholder detected in the target file is also
present in the source file.
*/
func (b *BatchValidator) isPlaceholderInSource() {
}

/*
Verifies that each placeholder present in the source file is also
detected in the target file.
*/
func (b *BatchValidator) isPlaceholderInTarget() {
}

func (b *BatchValidator) values() {
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
		translation.ValidationErrors.Response = append(
			translation.ValidationErrors.Response,
			&trainsError.TrainsError{
				Code: trainsError.TranslationBatchError,
				Message: fmt.Sprintf(
					"Failed to parse the translation result. %s",
					parsingError,
				),
				Err: parsingError,
			},
		)

		validatedTranslation, err := b.structure(rawTranslationOutput)
		translation.ValidationErrors.Batch = append(translation.ValidationErrors.Batch, err)
		translation.Validated = *validatedTranslation

		err = b.length(rawTranslationOutput, translation)
		translation.ValidationErrors.Batch = append(translation.ValidationErrors.Batch, err)

		err = b.keys(rawTranslationOutput, translation)
		translation.ValidationErrors.Entry = append(translation.ValidationErrors.Batch, err)

		placeholderErrors := b.placeholders(rawTranslationOutput, translation)
		translation.ValidationErrors.Entry = append(translation.ValidationErrors.Entry, placeholderErrors...)

		validatedTranslations <- translation
	}
}
