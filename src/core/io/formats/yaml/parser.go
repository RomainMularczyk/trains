package yaml

import (
	"strconv"
	"strings"

	yamlv3 "gopkg.in/yaml.v3"

	"trains/src/core/parser"
	"trains/src/core/translation"
)

type YAMLParser struct{}

func (p *YAMLParser) Parse(
	node <-chan *yamlv3.Node,
	translationUnit chan<- translation.TranslationUnit,
) {
	p.walk(node, []string{}, translationUnit)
}

func (p *YAMLParser) walk(
	node *yamlv3.Node,
	path []string,
	translationUnit chan<- translation.TranslationUnit,
) {
	switch node.Kind {

	// mapping (YAML object)
	case yamlv3.MappingNode:
		// Content is [key, value, key, value, ...]
		for i := 0; i < len(node.Content); i += 2 {
			keyNode := node.Content[i]
			valueNode := node.Content[i+1]

			key := keyNode.Value
			p.walk(valueNode, append(path, key), translationUnit)
		}

	// sequence (YAML array)
	case yamlv3.SequenceNode:
		for i, valueNode := range node.Content {
			index := strconv.Itoa(i)
			p.walk(valueNode, append(path, index), translationUnit)
		}

	// scalar (leaf node)
	case yamlv3.ScalarNode:
		fullKey := strings.Join(path, ".")

		unit := translation.TranslationUnit{
			Fullkey:  fullKey,
			Path:     append([]string{}, path...),
			Source:   node.Value,
			Segments: parser.Segmentize(node.Value),
		}

		translationUnit <- unit
	}
}
