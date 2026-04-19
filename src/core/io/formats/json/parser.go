package json

import (
	"strconv"
	"strings"
	"trains/src/core/parser"
	"trains/src/core/translation"
)

type JSONParser struct{}

func (p *JSONParser) Parse(
	data <-chan any,
	translationUnit chan<- translation.TranslationUnit,
) {
	fileContent := <-data
	p.walk(fileContent, []string{}, translationUnit)
}

func (p *JSONParser) walk(
	node any,
	path []string,
	translationUnit chan<- translation.TranslationUnit,
) {
	switch node.(type) {
	// object case
	case map[string]any:
		for key, value := range node.(map[string]any) {
			p.walk(value, append(path, key), translationUnit)
		}
	// array case
	case []any:
		for index, value := range node.([]any) {
			i := strconv.Itoa(index)
			p.walk(value, append(path, i), translationUnit)
		}
	// when a leaf node is reached (i.e. a string)
	case string:
		fullKey := strings.Join(path, ".")

		unit := translation.TranslationUnit{
			Fullkey:  fullKey,
			Path:     append([]string{}, path...),
			Source:   node.(string),
			Segments: parser.Segmentize(node.(string)),
		}
		translationUnit <- unit
	}
}
