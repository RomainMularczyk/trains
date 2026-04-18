package io

import (
	"fmt"
	"trains/src/core/io/formats/json"
	"trains/src/core/io/formats/yaml"
)

type FileFormat int

const (
	JSON FileFormat = iota
	YAML
)

type FileReader interface {
	Read(path string, content any) error
	Supports(path string) bool
}

type ReaderContext struct {
	reader FileReader
}

func (r *ReaderContext) SetReader(reader FileReader) {
	r.reader = reader
}

func (r *ReaderContext) Read(path string, value any) error {
	return r.reader.Read(path, value)
}

func Reader(format FileFormat) (FileReader, error) {
	switch format {
	case JSON:
		return &json.JSONReader{}, nil
	case YAML:
		return &yaml.YAMLReader{}, nil
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
