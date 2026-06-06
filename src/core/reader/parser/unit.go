package parser

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	configTypes "trains/src/core/config/types"
)

var TokenPerTextSegmentHeuristic = 4.0

type TranslationUnit struct {
	Fullkey  string
	Path     []string
	Source   string
	Target   string
	Segments []Segment
}

/*
Formats the TranslationUnit as a JSON string.
*/
func (t TranslationUnit) String() string {
	translationSet, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(translationSet)
}

/*
Retrieves all the placeholders from the TranslationUnit.
*/
func (t *TranslationUnit) GetPlaceholders() *[]Placeholder {
	var placeholders []Placeholder

	for _, segment := range t.Segments {
		if segment.Type == PlaceholderSegment {
			placeholders = append(placeholders, *segment.Placeholder)
		}
	}

	return &placeholders
}

/*
Estimates the number of tokens in the TranslationUnit.
*/
func (t *TranslationUnit) EstimateTokenNumber() int {
	estimateNumberOfTokens := 0
	for _, segment := range t.Segments {
		if segment.Type == TextSegment {
			ratio := float64(len(segment.Value)) / TokenPerTextSegmentHeuristic
			estimateNumberOfTokens += int(math.Ceil(ratio))
		}
	}

	return estimateNumberOfTokens
}

func (t *TranslationUnit) ShouldTranslate(config configTypes.RuntimeConfig) {
}

/*
Builds the TranslationEntry from the TranslationUnit.
*/
func (t *TranslationUnit) GetTranslationEntry() string {
	var entry string
	for _, segment := range t.Segments {
		if segment.Type == TextSegment {
			entry += segment.Value
		}
		if segment.Type == PlaceholderSegment {
			entry += segment.Placeholder.Pattern
		}
	}
	return entry
}

/*
Builds the TranslationEntry hash.
*/
func (t *TranslationUnit) GetHash() string {
	hash := sha256.New()
	hash.Write([]byte(t.GetTranslationEntry()))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
