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
// Placeholder Validation
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
			return err.Code == trainsError.TranslationEntryError
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
		return err.Code == trainsError.TranslationEntryError
	}) {
		t.Errorf("Expected a missing placeholder, got %v", errors)
	}
}

// --------------------------------------------------------------
// Values Validation
// --------------------------------------------------------------

// Verifies that the target does not contain any empty values.
func TestValuesValidation(t *testing.T) {
	translationEntries := testdataEntry.TwoEntries()
	translationResult := parser.TranslationResult{
		Batch: testdataBatch.TwoUnitsBatch(),
	}
	batchValidator := &BatchValidator{}
	errors := batchValidator.values(translationEntries, translationResult)

	if len(errors) != 0 {
		t.Errorf("Expected no errors, got %d", len(errors))
	}
}

// Verifies that the target set does not contain an empty target
// value.
func TestValuesValidationWithEmptyTarget(t *testing.T) {
	translationEntries := testdataEntry.TwoEntries()
	translationEntries[0].Target = ""
	translationResult := parser.TranslationResult{
		Batch: testdataBatch.TwoUnitsBatch(),
	}

	batchValidator := &BatchValidator{}
	errors := batchValidator.values(translationEntries, translationResult)

	if !slices.ContainsFunc(
		errors,
		func(err *trainsError.TrainsError) bool {
			return err.Code == trainsError.TranslationEntryError
		},
	) {
		t.Errorf("Expected an error, got %v", errors)
	}
}

// Verifies that if the target is empty, the source is also empty
// in order to return no error.
func TestValuesValidationWithEmptySourceAndEmptyTarget(t *testing.T) {
	translationEntries := testdataEntry.TwoEntries()
	translationEntries[0].Target = ""
	translationResult := parser.TranslationResult{
		Batch: testdataBatch.TwoUnitsBatch(),
	}
	translationResult.Batch.Units[0].Source = ""

	batchValidator := &BatchValidator{}
	errors := batchValidator.values(translationEntries, translationResult)

	if slices.ContainsFunc(
		errors,
		func(err *trainsError.TrainsError) bool {
			return err.Code == trainsError.TranslationEntryError
		},
	) {
		t.Errorf("Expected no error, got %v", errors)
	}
}
