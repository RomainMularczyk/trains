package parser

import (
	"reflect"
	"testing"
)

func TestSegmentizeOnlyText(t *testing.T) {
	input := "This is a test"
	result := Segmentize(input)
	expected := []Segment{{TextSegment, "This is a test", nil}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizeTextAndPlaceholder(t *testing.T) {
	input := "This is a {{placeholder}}"
	result := Segmentize(input)
	expected := []Segment{
		{TextSegment, "This is a ", nil},
		{PlaceholderSegment, "placeholder", &Placeholder{
			Name:           "placeholder",
			NameIndices:    []int{12, 23},
			Pattern:        "{{placeholder}}",
			PatternIndices: []int{10, 25},
		}},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizeTextPlaceholderAndText(t *testing.T) {
	input := "This is a {{placeholder}} and this is another text"
	result := Segmentize(input)
	expected := []Segment{
		{TextSegment, "This is a ", nil},
		{PlaceholderSegment, "placeholder", &Placeholder{
			Name:           "placeholder",
			NameIndices:    []int{12, 23},
			Pattern:        "{{placeholder}}",
			PatternIndices: []int{10, 25},
		}},
		{TextSegment, " and this is another text", nil},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizePlaceholderAndTextAndPlaceholder(t *testing.T) {
	input := "{{placeholder}} and this is another text {{another_placeholder}}"
	result := Segmentize(input)
	expected := []Segment{
		{PlaceholderSegment, "placeholder", &Placeholder{
			Name:           "placeholder",
			NameIndices:    []int{2, 13},
			Pattern:        "{{placeholder}}",
			PatternIndices: []int{0, 15},
		}},
		{TextSegment, " and this is another text ", nil},
		{PlaceholderSegment, "another_placeholder", &Placeholder{
			Name:           "another_placeholder",
			NameIndices:    []int{43, 62},
			Pattern:        "{{another_placeholder}}",
			PatternIndices: []int{41, 64},
		}},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizeOnlyOnePlaceholder(t *testing.T) {
	input := "{{placeholder}}"
	result := Segmentize(input)
	expected := []Segment{
		{PlaceholderSegment, "placeholder", &Placeholder{
			Name:           "placeholder",
			NameIndices:    []int{2, 13},
			Pattern:        "{{placeholder}}",
			PatternIndices: []int{0, 15},
		}},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizeOnlyTwoAdacentPlaceholders(t *testing.T) {
	input := "{{placeholder}}{{another_placeholder}}"
	result := Segmentize(input)
	expected := []Segment{
		{PlaceholderSegment, "placeholder", &Placeholder{
			Name:           "placeholder",
			NameIndices:    []int{2, 13},
			Pattern:        "{{placeholder}}",
			PatternIndices: []int{0, 15},
		}},
		{PlaceholderSegment, "another_placeholder", &Placeholder{
			Name:           "another_placeholder",
			NameIndices:    []int{17, 36},
			Pattern:        "{{another_placeholder}}",
			PatternIndices: []int{15, 38},
		}},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizeEmptyString(t *testing.T) {
	input := ""
	result := Segmentize(input)
	expected := []Segment{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}
