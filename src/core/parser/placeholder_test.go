package parser

import "testing"

func TestDetectPlaceholdersWithUnderscores(t *testing.T) {
	input := "This is a {{a_placeholder}}"
	result := DetectPlaceholders(input)
	if len(result) != 1 {
		t.Errorf("Expected 1 placeholder, got %d", len(result))
	}
	if result[0].NameIndices[0] != 12 && result[0].NameIndices[1] != 25 {
		t.Errorf("Expected placeholder name indices to be '[12, 25]', got %d", result[0].NameIndices)
	}
	if result[0].Name != "a_placeholder" {
		t.Errorf("Expected placeholder name to be 'a_placeholder', got %s", result[0].Name)
	}
	if result[0].PatternIndices[0] != 10 && result[0].PatternIndices[1] != 27 {
		t.Errorf("Expected placeholder PatternIndices to be '[10, 27]', got %d", result[0].PatternIndices)
	}
	if result[0].Pattern != "{{a_placeholder}}" {
		t.Errorf("Expected placeholder pattern to be '{{a_placeholder}}', got %s", result[0].Pattern)
	}
}

func TestStartWithPlaceholder(t *testing.T) {
	input := "{{placeholder}} and the rest"
	result := DetectPlaceholders(input)
	if len(result) != 1 {
		t.Errorf("Expected 1 placeholder, got %d", len(result))
	}
	if result[0].NameIndices[0] != 2 && result[0].NameIndices[1] != 13 {
		t.Errorf("Expected placeholder NameIndices to be '[2, 13]', got %d", result[0].NameIndices)
	}
	if result[0].Name != "placeholder" {
		t.Errorf("Expected placeholder Name to be 'placeholder', got %s", result[0].Name)
	}
	if result[0].PatternIndices[0] != 0 && result[0].PatternIndices[1] != 15 {
		t.Errorf("Expected placeholder PatternIndices to be '[0, 15]', got %d", result[0].PatternIndices)
	}
	if result[0].Pattern != "{{placeholder}}" {
		t.Errorf("Expected placeholder Pattern to be '{{placeholder}}', got %s", result[0].Pattern)
	}
}

func TestEndWithPlaceholder(t *testing.T) {
	input := "This is a {{placeholder}}"
	result := DetectPlaceholders(input)
	if len(result) != 1 {
		t.Errorf("Expected 1 placeholder, got %d", len(result))
	}
	if result[0].NameIndices[0] != 12 && result[0].NameIndices[1] != 25 {
		t.Errorf("Expected placeholder NameIndices to be '[12, 25]', got %s", result[0].Name)
	}
	if result[0].Name != "placeholder" {
		t.Errorf("Expected placeholder Name to be 'placeholder', got %s", result[0].Name)
	}
	if result[0].PatternIndices[0] != 10 && result[0].PatternIndices[1] != 27 {
		t.Errorf("Expected placeholder PatternIndices to be '[10, 27]', got %s", result[0].Pattern)
	}
	if result[0].Pattern != "{{placeholder}}" {
		t.Errorf("Expected placeholder Pattern to be '{{placeholder}}', got %s", result[0].Pattern)
	}
}

func TestDetectTwoPlaceholders(t *testing.T) {
	input := "This is a {{a_placeholder}} and {{another_placeholder}}"
	result := DetectPlaceholders(input)
	if len(result) != 2 {
		t.Errorf("Expected 2 placeholders, got %d", len(result))
	}
	if result[0].Name != "a_placeholder" {
		t.Errorf("Expected placeholder name to be 'a_placeholder', got %s", result[0].Name)
	}
	if result[0].NameIndices[0] != 12 && result[0].NameIndices[1] != 25 {
		t.Errorf("Expected placeholder NameIndices to be '[12, 25]', got %s", result[0].Name)
	}
	if result[0].Pattern != "{{a_placeholder}}" {
		t.Errorf("Expected placeholder pattern to be '{{a_placeholder}}', got %s", result[0].Pattern)
	}
	if result[0].PatternIndices[0] != 10 && result[0].PatternIndices[1] != 27 {
		t.Errorf("Expected placeholder PatternIndices to be '[10, 27]', got %s", result[0].Pattern)
	}

	if result[1].Name != "another_placeholder" {
		t.Errorf("Expected placeholder name to be 'another_placeholder', got %s", result[1].Name)
	}
	if result[1].NameIndices[0] != 34 && result[1].NameIndices[1] != 53 {
		t.Errorf("Expected placeholder NameIndices to be '[34, 53]', got %d", result[1].NameIndices)
	}
	if result[1].Pattern != "{{another_placeholder}}" {
		t.Errorf("Expected placeholder pattern to be '{{another_placeholder}}', got %s", result[1].Pattern)
	}
	if result[1].PatternIndices[0] != 32 && result[1].PatternIndices[1] != 55 {
		t.Errorf("Expected placeholder PatternIndices to be '[32, 55]', got %d", result[1].PatternIndices)
	}
}

func TestDetectOnlyPlaceholder(t *testing.T) {
	input := "{{placeholder}}"
	result := DetectPlaceholders(input)
	if len(result) != 1 {
		t.Errorf("Expected 1 placeholder, got %d", len(result))
	}
	if result[0].NameIndices[0] != 2 && result[0].NameIndices[1] != 13 {
		t.Errorf("Expected placeholder NameIndices to be '[2, 13]', got %d", result[0].NameIndices)
	}
	if result[0].Name != "placeholder" {
		t.Errorf("Expected placeholder Name to be 'placeholder', got %s", result[0].Name)
	}
	if result[0].PatternIndices[0] != 0 && result[0].PatternIndices[1] != 15 {
		t.Errorf("Expected placeholder PatternIndices to be '[0, 15]', got %d", result[0].PatternIndices)
	}
	if result[0].Pattern != "{{placeholder}}" {
		t.Errorf("Expected placeholder Pattern to be '{{placeholder}}', got %s", result[0].Pattern)
	}
}

func TestTwoAdjacentPlaceholders(t *testing.T) {
	input := "{{placeholder}}{{another_placeholder}}"
	result := DetectPlaceholders(input)
	if len(result) != 2 {
		t.Errorf("Expected 2 placeholders, got %d", len(result))
	}
	if result[0].NameIndices[0] != 2 && result[0].NameIndices[1] != 13 {
		t.Errorf("Expected placeholder NameIndices to be '[2, 13]', got %d", result[0].NameIndices)
	}
	if result[0].Name != "placeholder" {
		t.Errorf("Expected placeholder Name to be 'placeholder', got %s", result[0].Name)
	}
	if result[0].PatternIndices[0] != 0 && result[0].PatternIndices[1] != 15 {
		t.Errorf("Expected placeholder PatternIndices to be '[0, 15]', got %d", result[0].PatternIndices)
	}
	if result[0].Pattern != "{{placeholder}}" {
		t.Errorf("Expected placeholder Pattern to be '{{placeholder}}', got %s", result[0].Pattern)
	}
	if result[1].NameIndices[0] != 17 && result[1].NameIndices[1] != 36 {
		t.Errorf("Expected placeholder NameIndices to be '[17, 36]', got %d", result[1].NameIndices)
	}
	if result[1].Name != "another_placeholder" {
		t.Errorf("Expected placeholder Name to be 'another_placeholder', got %s", result[1].Name)
	}
	if result[1].PatternIndices[0] != 15 && result[1].PatternIndices[1] != 38 {
		t.Errorf("Expected placeholder PatternIndices to be '[15, 38]', got %d", result[1].PatternIndices)
	}
	if result[1].Pattern != "{{another_placeholder}}" {
		t.Errorf("Expected placeholder pattern to be '{{another_placeholder}}', got %s", result[1].Pattern)
	}
}

func TestDetectNoPlaceholderInEmptyString(t *testing.T) {
	input := ""
	result := DetectPlaceholders(input)
	if len(result) != 0 {
		t.Errorf("Expected 0 placeholders, got %d", len(result))
	}
}

func TestDetectNoPlaceholdersWithSingleBrace(t *testing.T) {
	input := "{not_a_placeholder}"
	result := DetectPlaceholders(input)
	if len(result) != 0 {
		t.Errorf("Expected 0 placeholders, got %d", len(result))
	}
}

func TestDetectMissingLeftBrace(t *testing.T) {
	input := "not_a_placeholder}}"
	result := DetectPlaceholders(input)
	if len(result) != 0 {
		t.Errorf("Expected 0 placeholders, got %d", len(result))
	}
}

func TestDetectMissingRightBrace(t *testing.T) {
	input := "{{not_a_placeholder"
	result := DetectPlaceholders(input)
	if len(result) != 0 {
		t.Errorf("Expected 0 placeholders, got %d", len(result))
	}
}

func TestNestedPlaceholder(t *testing.T) {
	input := "{{{{placeholder}}}}"
	result := DetectPlaceholders(input)
	if len(result) != 1 {
		t.Errorf("Expected 1 placeholder, got %d", len(result))
	}
}

func TestDetectNoPlaceholderWhenSpaceSeperated(t *testing.T) {
	input := "This is not {{ a placeholder }}"
	result := DetectPlaceholders(input)
	if len(result) != 0 {
		t.Errorf("Expected 0 placeholders, got %d", len(result))
	}
}
