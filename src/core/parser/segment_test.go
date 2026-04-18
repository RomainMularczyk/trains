package parser

import (
	"reflect"
	"testing"
)

func TestSegmentizeOnlyText(t *testing.T) {
	input := "This is a test"
	result := Segmentize(input)
	expected := Segment{TextSegment, "This is a test"}
	if !reflect.DeepEqual(result[0], expected) {
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
	if len(result) != 3 {
		t.Errorf("Expected 3 segments, got %d", len(result))
	}
	if result[0].Type != TextSegment {
		t.Errorf("Expected segment type to be TextSegment, got %d", result[0].Type)
	}
	if result[1].Type != PlaceholderSegment {
		t.Errorf("Expected segment type to be PlaceholderSegment, got %d", result[1].Type)
	}
	if result[2].Type != TextSegment {
		t.Errorf("Expected segment type to be TextSegment, got %d", result[2].Type)
	}
}

func TestSegmentizePlaceholderAndTextAndPlaceholder(t *testing.T) {
	input := "{{placeholder}} and this is another text {{another_placeholder}}"
	result := Segmentize(input)
	if len(result) != 3 {
		t.Errorf("Expected 3 segments, got %d", len(result))
	}
	if result[0].Type != PlaceholderSegment {
		t.Errorf("Expected segment type to be PlaceholderSegment, got %d", result[0].Type)
	}
	if result[1].Type != TextSegment {
		t.Errorf("Expected segment type to be TextSegment, got %d", result[1].Type)
	}
	if result[2].Type != PlaceholderSegment {
		t.Errorf("Expected segment type to be PlaceholderSegment, got %d", result[2].Type)
	}
}

func TestSegmentizeOnlyOnePlaceholder(t *testing.T) {
	input := "{{placeholder}}"
	result := Segmentize(input)
	if len(result) != 1 {
		t.Errorf("Expected 1 segment, got %d", len(result))
	}
	if result[0].Type != PlaceholderSegment {
		t.Errorf("Expected segment type to be PlaceholderSegment, got %d", result[0].Type)
	}
}

func TestSegmentizeOnlyTwoAdacentPlaceholders(t *testing.T) {
	input := "{{placeholder}}{{another_placeholder}}"
	result := Segmentize(input)
	if len(result) != 2 {
		t.Errorf("Expected 2 segments, got %d", len(result))
	}
	if result[0].Type != PlaceholderSegment {
		t.Errorf("Expected segment type to be PlaceholderSegment, got %d", result[0].Type)
	}
	if result[1].Type != PlaceholderSegment {
		t.Errorf("Expected segment type to be PlaceholderSegment, got %d", result[1].Type)
	}
}

func TestSegmentizeEmptyString(t *testing.T) {
	input := ""
	result := Segmentize(input)
	if len(result) != 0 {
		t.Errorf("Expected 0 segments, got %d", len(result))
	}
}
