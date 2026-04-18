package json

import (
	"reflect"
	"testing"
	"trains/src/core/parser"
	"trains/src/core/translation"
)

func TestParseSimpleNestedKeysJson(t *testing.T) {
	input := map[string]any{
		"user": map[string]any{
			"name": "Jovan",
		},
	}
	jsonParser := JSONParser{}
	result := jsonParser.Parse(input)
	expected := []translation.TranslationUnit{
		{
			Fullkey:  "user.name",
			Path:     []string{"user", "name"},
			Source:   "Jovan",
			Target:   "",
			Segments: []parser.Segment{{parser.TextSegment, "Jovan"}},
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
	jsonParser := JSONParser{}
	result := jsonParser.Parse(input)
	expected := []translation.TranslationUnit{
		{
			Fullkey:  "user",
			Path:     []string{"user"},
			Source:   "Jovan",
			Target:   "",
			Segments: []parser.Segment{{parser.TextSegment, "Jovan"}},
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
	jsonParser := JSONParser{}
	result := jsonParser.Parse(input)
	expected := []translation.TranslationUnit{
		{
			Fullkey:  "user.name",
			Path:     []string{"user", "name"},
			Source:   "Jovan",
			Target:   "",
			Segments: []parser.Segment{{parser.TextSegment, "Jovan"}},
		},
		{
			Fullkey:  "grade.description",
			Path:     []string{"grade", "description"},
			Source:   "You achieved a grade of {{grade}}",
			Target:   "",
			Segments: []parser.Segment{{parser.TextSegment, "You achieved a grade of "}, {parser.PlaceholderSegment, "grade"}},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected translation set to be %v, got %v", expected, result)
	}
}

func TestParseEmptyJson(t *testing.T) {
	input := map[string]any{}
	jsonParser := JSONParser{}
	result := jsonParser.Parse(input)
	expected := []translation.TranslationUnit{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected translation set to be %v, got %v", expected, result)
	}
}
