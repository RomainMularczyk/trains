package parser

import (
	"reflect"
	"testing"
)

func TestSegmentizeOnlyText(t *testing.T) {
	input := "This is a test"
	result := Segmentize(input)
	expected := []Segment{{TextSegment, "This is a test"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizeTextAndPlaceholder(t *testing.T) {
	input := "This is a {{placeholder}}"
	result := Segmentize(input)
	expected := []Segment{{TextSegment, "This is a "}, {PlaceholderSegment, "placeholder"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizeTextPlaceholderAndText(t *testing.T) {
	input := "This is a {{placeholder}} and this is another text"
	result := Segmentize(input)
	expected := []Segment{{TextSegment, "This is a "}, {PlaceholderSegment, "placeholder"}, {TextSegment, " and this is another text"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizePlaceholderAndTextAndPlaceholder(t *testing.T) {
	input := "{{placeholder}} and this is another text {{another_placeholder}}"
	result := Segmentize(input)
	expected := []Segment{{PlaceholderSegment, "placeholder"}, {TextSegment, " and this is another text "}, {PlaceholderSegment, "another_placeholder"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizeOnlyOnePlaceholder(t *testing.T) {
	input := "{{placeholder}}"
	result := Segmentize(input)
	expected := []Segment{{PlaceholderSegment, "placeholder"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected segment to be %v, got %v", expected, result)
	}
}

func TestSegmentizeOnlyTwoAdacentPlaceholders(t *testing.T) {
	input := "{{placeholder}}{{another_placeholder}}"
	result := Segmentize(input)
	expected := []Segment{{PlaceholderSegment, "placeholder"}, {PlaceholderSegment, "another_placeholder"}}
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
