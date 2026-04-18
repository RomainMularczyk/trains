package yaml

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type YAMLReader struct{}

func (r *YAMLReader) Read(path string, content any) error {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(fileContent, content)
}

func (r *YAMLReader) Supports(path string) bool {
	return strings.HasSuffix(path, ".yaml")
}
