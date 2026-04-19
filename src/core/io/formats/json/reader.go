package json

import (
	"encoding/json"
	"os"
	"strings"
)

type JSONReader struct{}

func (r *JSONReader) Read(path <-chan string, content chan<- any) {
	for filePath := range path {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}
		var data any
		err = json.Unmarshal(fileContent, &data)
		if err != nil {
			continue
		}
		content <- data
	}
}

func (r *JSONReader) Supports(path string) bool {
	return strings.HasSuffix(path, ".json")
}
