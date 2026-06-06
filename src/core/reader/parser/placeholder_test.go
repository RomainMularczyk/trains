package parser

import (
	"reflect"
	"testing"
)

// --------------------------------------------------------------
// DetectPlaceholders
// --------------------------------------------------------------

func TestDetectPlaceholdersWithUnderscores(t *testing.T) {
	input := "This is a {{a_placeholder}}"
	result := DetectPlaceholders(input)
	expected := []Placeholder{
		{
			Name:           "a_placeholder",
			NameIndices:    []int{12, 25},
			Pattern:        "{{a_placeholder}}",
			PatternIndices: []int{10, 27},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestStartWithPlaceholder(t *testing.T) {
	input := "{{placeholder}} and the rest"
	result := DetectPlaceholders(input)
	expected := []Placeholder{
		{
			Name:           "placeholder",
			NameIndices:    []int{2, 13},
			Pattern:        "{{placeholder}}",
			PatternIndices: []int{0, 15},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestEndWithPlaceholder(t *testing.T) {
	input := "This is a {{placeholder}}"
	result := DetectPlaceholders(input)
	expected := []Placeholder{
		{
			Name:           "placeholder",
			NameIndices:    []int{12, 23},
			Pattern:        "{{placeholder}}",
			PatternIndices: []int{10, 25},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDetectTwoPlaceholders(t *testing.T) {
	input := "This is a {{a_placeholder}} and {{another_placeholder}}"
	result := DetectPlaceholders(input)
	expected := []Placeholder{
		{
			Name:           "a_placeholder",
			NameIndices:    []int{12, 25},
			Pattern:        "{{a_placeholder}}",
			PatternIndices: []int{10, 27},
		},
		{
			Name:           "another_placeholder",
			NameIndices:    []int{34, 53},
			Pattern:        "{{another_placeholder}}",
			PatternIndices: []int{32, 55},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDetectOnlyPlaceholder(t *testing.T) {
	input := "{{placeholder}}"
	result := DetectPlaceholders(input)
	expected := []Placeholder{
		{
			Name:           "placeholder",
			NameIndices:    []int{2, 13},
			Pattern:        "{{placeholder}}",
			PatternIndices: []int{0, 15},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTwoAdjacentPlaceholders(t *testing.T) {
	input := "{{placeholder}}{{another_placeholder}}"
	result := DetectPlaceholders(input)
	expected := []Placeholder{
		{
			Name:           "placeholder",
			NameIndices:    []int{2, 13},
			Pattern:        "{{placeholder}}",
			PatternIndices: []int{0, 15},
		},
		{
			Name:           "another_placeholder",
			NameIndices:    []int{17, 36},
			Pattern:        "{{another_placeholder}}",
			PatternIndices: []int{15, 38},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDetectNoPlaceholderInEmptyString(t *testing.T) {
	input := ""
	result := DetectPlaceholders(input)
	expected := []Placeholder{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDetectNoPlaceholdersWithSingleBrace(t *testing.T) {
	input := "{not_a_placeholder}"
	result := DetectPlaceholders(input)
	expected := []Placeholder{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDetectMissingLeftBrace(t *testing.T) {
	input := "not_a_placeholder}}"
	result := DetectPlaceholders(input)
	expected := []Placeholder{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDetectMissingRightBrace(t *testing.T) {
	input := "{{not_a_placeholder"
	result := DetectPlaceholders(input)
	expected := []Placeholder{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNestedPlaceholder(t *testing.T) {
	input := "{{{{placeholder}}}}"
	result := DetectPlaceholders(input)
	expected := []Placeholder{
		{
			Name:           "placeholder",
			NameIndices:    []int{4, 15},
			Pattern:        "{{placeholder}}",
			PatternIndices: []int{2, 17},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDetectNoPlaceholderWhenSpaceSeperated(t *testing.T) {
	input := "This is not {{ a placeholder }}"
	result := DetectPlaceholders(input)
	expected := []Placeholder{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// --------------------------------------------------------------
// GetPlaceholders
// --------------------------------------------------------------

func TestRetrievePlaceholders(t *testing.T) {
	input := TranslationUnit{
		Fullkey: "user.name",
		Path:    []string{"user", "name"},
		Source:  "My name is {{name}}",
		Target:  "",
		Segments: []Segment{
			{TextSegment, "My name is ", nil},
			{PlaceholderSegment, "name",
				&Placeholder{
					Name:           "name",
					NameIndices:    []int{19, 24},
					Pattern:        "{{name}}",
					PatternIndices: []int{17, 25}},
			},
		},
	}
	result := input.GetPlaceholders()

	if len(*result) != 1 {
		t.Errorf("Expected 1 placeholder, got %d", len(*result))
	}
}
