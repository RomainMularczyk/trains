package json

import (
	"strconv"
	"strings"
	"trains/src/core/translation"
)

type JSONParser struct{}

func (p *JSONParser) Parse(data any) translation.TranslationSet {
	var translationSet translation.TranslationSet
	p.walk(data, []string{}, &translationSet)
	return translationSet
}

func (p *JSONParser) walk(
	node any,
	path []string,
	translationSet *translation.TranslationSet,
) {
	switch node.(type) {
	// object case
	case map[string]any:
		for key, value := range node.(map[string]any) {
			p.walk(value, append(path, key), translationSet)
		}
	// array case
	case []any:
		for index, value := range node.([]any) {
			i := strconv.Itoa(index)
			p.walk(value, append(path, i), translationSet)
		}
	// when a leaf node is reached (i.e. a string)
	case string:
		fullKey := strings.Join(path, ".")

		unit := translation.TranslationUnit{
			Fullkey: fullKey,
			Path:    append([]string{}, path...),
			Source:  node.(string),
		}
		translationSet.Units = append(translationSet.Units, unit)
	}
}
