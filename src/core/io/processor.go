package io

import (
	"fmt"
	"trains/src/core/config/types"
	"trains/src/core/io/formats/json"
	"trains/src/core/parser"
)

type FormatDriver struct {
	Reader FileReader
	Parser Parser
}

type Parser interface {
	Parse(
		data <-chan any,
		translationUnit chan<- parser.TranslationUnit,
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

func Processor(format types.FileFormat) (*FormatDriver, error) {
	switch format {
	case types.JSON:
		return &FormatDriver{
			&json.JSONReader{},
			&json.JSONParser{},
		}, nil
	default:
		return nil, fmt.Errorf("Unsupported format: %v", format)
	}
}
