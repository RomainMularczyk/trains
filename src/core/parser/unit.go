package parser

import (
	"encoding/json"
	"math"
)

var TOKEN_PER_TEXT_SEGMENT_HEURISTIC = 4.0

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
Estimates the number of tokens in the TranslationUnit.
*/
func (t *TranslationUnit) EstimateTokenNumber() int {
	estimateNumberOfTokens := 0
	for _, segment := range t.Segments {
		if segment.Type == TextSegment {
			ratio := float64(len(segment.Value)) / TOKEN_PER_TEXT_SEGMENT_HEURISTIC
			estimateNumberOfTokens += int(math.Ceil(ratio))
		}
	}

	return estimateNumberOfTokens
}
