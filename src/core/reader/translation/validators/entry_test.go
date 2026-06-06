package translationValidator

import (
	"slices"
	"testing"
	trainsError "trains/src/core/errors"
	"trains/src/core/reader/parser"
	testdataEntry "trains/testdata/outputs/translation/entry"
	testdataBatch "trains/testdata/sources/batch"
)

// --------------------------------------------------------------
// PlaceholderSet
// --------------------------------------------------------------

func TestPlaceholderSetEqual(t *testing.T) {
	sourceSet := PlaceholderSet{
		"name": struct{}{},
		"age":  struct{}{},
	}
	targetSet := PlaceholderSet{
		"name": struct{}{},
		"age":  struct{}{},
	}

	result := sourceSet.equal(targetSet)

	if !result.Result {
		t.Errorf("Expected true, got false")
	}
}

func TestPlaceholderSetMissingOne(t *testing.T) {
	sourceSet := PlaceholderSet{
		"name": struct{}{},
		"age":  struct{}{},
	}
	targetSet := PlaceholderSet{
		"name": struct{}{},
	}

	result := sourceSet.equal(targetSet)

	if result.Missing == nil {
		t.Errorf("Expected a missing placeholder, got nil")
	}
}

func TestPlaceholderSetExtra(t *testing.T) {
	sourceSet := PlaceholderSet{
		"name": struct{}{},
	}
	targetSet := PlaceholderSet{
		"name": struct{}{},
		"age":  struct{}{},
	}

	result := sourceSet.equal(targetSet)

	if result.Extra == nil {
		t.Errorf("Expected an extra placeholder, got nil")
	}
}

// --------------------------------------------------------------
// PlaceholderSet
// --------------------------------------------------------------

// Verifies that the target set contains the same placeholders as the source set.
func TestPlaceholderValidation(t *testing.T) {
	translationEntries := testdataEntry.TwoEntries()
	translationResult := parser.TranslationResult{
		Batch: testdataBatch.TwoUnitsBatch(),
	}
	batchValidator := &BatchValidator{}
	errors := batchValidator.placeholders(translationEntries, translationResult)

	if len(errors) != 0 {
		t.Errorf("Expected no errors, got %d", len(errors))
	}
}

// Verifies that the target set does not contain any extra placeholders
// that are not present in the source set.
func TestPlaceholderValidationWithExtraPlaceholder(t *testing.T) {
	translationEntries := testdataEntry.TwoEntries()
	translationResult := parser.TranslationResult{
		Batch: testdataBatch.OneUnitBatch(),
	}
	batchValidator := &BatchValidator{}
	errors := batchValidator.placeholders(translationEntries, translationResult)

	if !slices.ContainsFunc(
		errors,
		func(err *trainsError.TrainsError) bool {
			return err.Code == trainsError.TranslationEntryError &&
				err.Message == `Failed to validate translation result.
				The placeholders in the target file does not exist in the source file.
				Extra placeholders: ["user.age"]`
		},
	) {
		t.Errorf("Expected an extra placeholder, got %v", errors)
	}
}

// Verifies that the target set does not miss any placeholders that are
// present in the source set.
func TestPlaceholderValidationWithMissingPlaceholder(t *testing.T) {
	translationEntries := testdataEntry.OneEntry()
	translationResult := parser.TranslationResult{
		Batch: testdataBatch.TwoUnitsBatch(),
	}
	batchValidator := &BatchValidator{}
	errors := batchValidator.placeholders(translationEntries, translationResult)

	if !slices.ContainsFunc(errors, func(err *trainsError.TrainsError) bool {
		return err.Code == trainsError.TranslationEntryError && err.Message == `Failed to validate translation result. 
		The placeholders in the source file is missing in the target file.
		Missing placeholders: ["user.name"]`
	}) {
		t.Errorf("Expected a missing placeholder, got %v", errors)
	}
}
