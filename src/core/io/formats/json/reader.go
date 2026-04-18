package json

import (
	"encoding/json"
	"os"
	"strings"
)

type JSONReader struct{}

func (r *JSONReader) Read(path string, content any) error {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileContent, content)
}

func (r *JSONReader) Supports(path string) bool {
	return strings.HasSuffix(path, ".json")
}
