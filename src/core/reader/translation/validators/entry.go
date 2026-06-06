package translationValidator

import (
	"fmt"
	trainsError "trains/src/core/errors"
	"trains/src/core/reader/parser"
)

type PlaceholderMatch struct {
	Result  bool
	Missing []string
	Extra   []string
}

type PlaceholderSet map[string]struct{}

/*
Compares two PlaceholderSets and returns a detail of the differences.
- Result: true if the two sets are equal, false otherwise
- Missing: keys in the reference set that are not in the other set
- Extra: keys in the other set that are not in the reference set
*/
func (p PlaceholderSet) equal(other PlaceholderSet) PlaceholderMatch {
	var missing []string
	var extra []string

	for key := range p {
		if _, ok := other[key]; !ok {
			missing = append(missing, key)
		}
	}

	for key := range other {
		if _, ok := p[key]; !ok {
			extra = append(extra, key)
		}
	}

	return PlaceholderMatch{
		Result:  len(missing) == 0 && len(extra) == 0,
		Missing: missing,
		Extra:   extra,
	}
}

/*
Creates a new PlaceholderSet from a slice of Placeholders.
*/
func NewPlaceholderSet(placeholders []parser.Placeholder) PlaceholderSet {
	set := make(PlaceholderSet, len(placeholders))
	for _, placeholder := range placeholders {
		set[placeholder.Name] = struct{}{}
	}
	return set
}

/*
Verifies that the target set does not contain any extra placeholders.
*/
func isPlaceholderInTarget(
	targetSet map[string]PlaceholderSet,
	sourceSet map[string]PlaceholderSet,
) []*trainsError.TrainsError {
	var entryErrors []*trainsError.TrainsError

	for key, target := range targetSet {
		placeholderInSourceFile := sourceSet[key]
		match := placeholderInSourceFile.equal(target)

		if !match.Result {
			entryErrors = append(
				entryErrors,
				&trainsError.TrainsError{
					Code: trainsError.TranslationEntryError,
					Message: fmt.Sprintf(
						`Failed to validate translation result. 
						The placeholders in the target file does not exist in the source file.
						Extra placeholders: %s`,
						match.Extra,
					),
					Err: nil,
				},
			)
		}
	}

	return entryErrors
}

/*
Verifies that the target set does not miss any placeholders.
*/
func isPlaceholderInSource(
	targetSet map[string]PlaceholderSet,
	sourceSet map[string]PlaceholderSet,
) []*trainsError.TrainsError {
	var entryErrors []*trainsError.TrainsError

	for key, source := range sourceSet {
		placeholdersInTargetFile := targetSet[key]
		match := source.equal(placeholdersInTargetFile)

		if !match.Result {
			entryErrors = append(
				entryErrors,
				&trainsError.TrainsError{
					Code: trainsError.TranslationEntryError,
					Message: fmt.Sprintf(
						`Failed to validate translation result. 
						The placeholders in the source file is missing in the target file.
						Missing placeholders: %s`,
						match.Missing,
					),
					Err: nil,
				},
			)
		}
	}

	return entryErrors
}

/*
Validates the placeholders of the translation result.
*/
func (b *BatchValidator) placeholders(
	translationEntries []parser.TranslationEngineEntry,
	translationResult parser.TranslationResult,
) []*trainsError.TrainsError {
	sourcePlaceholder := make(map[string]PlaceholderSet)
	targetPlaceholder := make(map[string]PlaceholderSet)

	for _, translationUnit := range translationResult.Batch.Units {
		placeholdersInSourceFile := translationUnit.GetPlaceholders()
		sourcePlaceholder[translationUnit.Fullkey] = NewPlaceholderSet(*placeholdersInSourceFile)
	}

	for _, translationEntry := range translationEntries {
		placeholders := parser.DetectPlaceholders(translationEntry.Target)
		targetPlaceholder[translationEntry.Key] = NewPlaceholderSet(placeholders)
	}

	var entryErrors []*trainsError.TrainsError
	entryErrors = append(
		entryErrors,
		isPlaceholderInSource(targetPlaceholder, sourcePlaceholder)...,
	)
	entryErrors = append(
		entryErrors,
		isPlaceholderInTarget(targetPlaceholder, sourcePlaceholder)...,
	)

	return entryErrors
}
