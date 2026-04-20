package json

import (
	"reflect"
	"sort"
	"testing"
	"trains/src/core/parser"
)

func collectTranslations(translationUnit <-chan parser.TranslationUnit) []parser.TranslationUnit {
	result := make([]parser.TranslationUnit, 0)
	for unit := range translationUnit {
		result = append(result, unit)
	}
	return result
}

func sortByFullkey(units []parser.TranslationUnit) []parser.TranslationUnit {
	sorted := make([]parser.TranslationUnit, len(units))
	copy(sorted, units)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Fullkey < sorted[j].Fullkey
	})
	return sorted
}

func TestParseSimpleNestedKeysJson(t *testing.T) {
	input := map[string]any{
		"user": map[string]any{
			"name": "Jovan",
		},
	}

	dataChan := make(chan any, 1)
	dataChan <- input
	close(dataChan)

	translationChan := make(chan parser.TranslationUnit, 10)

	jsonParser := JSONParser{}
	go jsonParser.Parse(dataChan, translationChan)

	result := collectTranslations(translationChan)

	expected := []parser.TranslationUnit{
		{
			Fullkey:  "user.name",
			Path:     []string{"user", "name"},
			Source:   "Jovan",
			Target:   "",
			Segments: []parser.Segment{{Type: parser.TextSegment, Value: "Jovan"}},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected translation set to be %v, got %v", expected, result)
	}
}

func TestParseSimpleFlatJson(t *testing.T) {
	input := map[string]any{
		"user": "Jovan",
	}

	dataChan := make(chan any, 1)
	dataChan <- input
	close(dataChan)

	translationChan := make(chan parser.TranslationUnit, 10)

	jsonParser := JSONParser{}
	go jsonParser.Parse(dataChan, translationChan)

	result := collectTranslations(translationChan)

	expected := []parser.TranslationUnit{
		{
			Fullkey:  "user",
			Path:     []string{"user"},
			Source:   "Jovan",
			Target:   "",
			Segments: []parser.Segment{{Type: parser.TextSegment, Value: "Jovan"}},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected translation set to be %v, got %v", expected, result)
	}
}

func TestParseSimpleNestedJsonWithPlaceholder(t *testing.T) {
	input := map[string]any{
		"user": map[string]any{
			"name": "Jovan",
		},
		"grade": map[string]any{
			"description": "You achieved a grade of {{grade}}",
		},
	}

	dataChan := make(chan any, 1)
	dataChan <- input
	close(dataChan)

	translationChan := make(chan parser.TranslationUnit, 10)

	jsonParser := JSONParser{}
	go jsonParser.Parse(dataChan, translationChan)

	result := sortByFullkey(collectTranslations(translationChan))

	expected := []parser.TranslationUnit{
		{
			Fullkey: "grade.description",
			Path:    []string{"grade", "description"},
			Source:  "You achieved a grade of {{grade}}",
			Target:  "",
			Segments: []parser.Segment{
				{Type: parser.TextSegment, Value: "You achieved a grade of "},
				{Type: parser.PlaceholderSegment, Value: "grade"},
			},
		},
		{
			Fullkey:  "user.name",
			Path:     []string{"user", "name"},
			Source:   "Jovan",
			Target:   "",
			Segments: []parser.Segment{{Type: parser.TextSegment, Value: "Jovan"}},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected translation set to be %v, got %v", expected, result)
	}
}

func TestParseEmptyJson(t *testing.T) {
	input := map[string]any{}

	dataChan := make(chan any, 1)
	dataChan <- input
	close(dataChan)

	translationChan := make(chan parser.TranslationUnit, 10)

	jsonParser := JSONParser{}
	go jsonParser.Parse(dataChan, translationChan)

	result := collectTranslations(translationChan)

	expected := []parser.TranslationUnit{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected translation set to be %v, got %v", expected, result)
	}
}
