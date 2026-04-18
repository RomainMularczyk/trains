package parser

import (
	"regexp"
)

type Placeholder struct {
	Name           string
	NameIndices    []int
	Pattern        string
	PatternIndices []int
}

func DetectPlaceholders(input string) []Placeholder {
	pattern := regexp.MustCompile(`\{\{\s*([a-zA-Z0-9_]+)\s*\}\}`)
	indices := pattern.FindAllStringSubmatchIndex(input, -1)

	placeholders := make([]Placeholder, 0, len(indices))

	for _, matchIndices := range indices {
		patternStart, patternEnd := matchIndices[0], matchIndices[1]
		nameStart, nameEnd := matchIndices[2], matchIndices[3]
		placeholder := Placeholder{
			Pattern:        input[patternStart:patternEnd],
			PatternIndices: []int{patternStart, patternEnd},
			Name:           input[nameStart:nameEnd],
			NameIndices:    []int{nameStart, nameEnd},
		}
		placeholders = append(placeholders, placeholder)
	}

	return placeholders
}
