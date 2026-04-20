package yaml

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type YAMLReader struct{}

func (r *YAMLReader) Read(path <-chan string, content any) error {
	filePath := <-path
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(fileContent, content)
}

func (r *YAMLReader) Supports(path string) bool {
	return strings.HasSuffix(path, ".yaml")
}
