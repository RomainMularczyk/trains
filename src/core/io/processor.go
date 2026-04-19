package io

import (
	"fmt"
	"trains/src/core/io/formats/json"
	"trains/src/core/translation"
)

type FileFormat int

const (
	JSON FileFormat = iota
	YAML
)

type FormatDriver struct {
	Reader FileReader
	Parser Parser
}

type Parser interface {
	Parse(
		data <-chan any,
		translationUnit chan<- translation.TranslationUnit,
	)
}

type FileReader interface {
	Read(path <-chan string, content chan<- any)
	Supports(path string) bool
}

type ReaderContext struct {
	reader FileReader
}

func (r *ReaderContext) Read(path <-chan string, value chan<- any) {
	r.reader.Read(path, value)
}

func Processor(format FileFormat) (*FormatDriver, error) {
	switch format {
	case JSON:
		return &FormatDriver{
			&json.JSONReader{},
			&json.JSONParser{},
		}, nil
	default:
		return nil, fmt.Errorf("Unsupported format: %s", format)
	}
}

func FlagToFileFormat(flag string) (FileFormat, error) {
	switch flag {
	case "json":
		return JSON, nil
	case "yaml":
		return YAML, nil
	default:
		return 0, fmt.Errorf("Unsupported format: %s", flag)
	}
}
