package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFromFile_ValidConfig(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")

	validJSON := `{
		"Provider": {
			"ApiKey": "test-key",
			"Model": "gpt-4",
			"BaseUrl": "https://api.example.com",
			"Timeout": 30
		},
		"Translation": {
			"SourceLanguage": "en",
			"TargetLanguage": "fr"
		},
		"Batching": {
			"TokenLimit": 1000,
			"UnitLimit": 10
		},
		"Prompt": {
			"Context": "Translate the following text"
		}
	}`

	err := os.WriteFile(configPath, []byte(validJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	expected := Config{
		Provider: Provider{
			ApiKey:  "test-key",
			Model:   "gpt-4",
			BaseUrl: "https://api.example.com",
			Timeout: 30,
		},
		Translation: Translation{
			SourceLanguage: "en",
			TargetLanguage: "fr",
		},
		Batching: Batching{
			TokenLimit: 1000,
			UnitLimit:  10,
		},
		Prompt: Prompt{
			Context: "Translate the following text",
		},
	}

	config, err := fromFile(configPath)
	if !reflect.DeepEqual(*config, expected) {
		t.Errorf("Expected config to be %v, got %v", expected, *config)
	}
}

func TestFromFile_InvalidJSON(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")

	// because of training comas
	invalidJSON := `{
		"Provider": {
			"ApiKey": "test-key",
			"Model": "gpt-4",
			"BaseUrl": "https://api.example.com",
			"Timeout": 30,
		},
	}`

	err := os.WriteFile(configPath, []byte(invalidJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	_, err = fromFile(configPath)
	if err == nil {
		t.Errorf("Expected error for invalid JSON, got nil")
	}
}

func TestFromFile_MissingRequiredFields(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")

	incompleteJSON := `{
		"Provider": {
			"ApiKey": "test-key",
			"Model": "gpt-4"
		},
		"Translation": {
			"SourceLanguage": "en"
		}
	}`

	err := os.WriteFile(configPath, []byte(incompleteJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	_, err = fromFile(configPath)
	if err == nil {
		t.Errorf("Expected error for missing required fields, got nil")
	}
}

func TestFromFile_NonExistentFile(t *testing.T) {
	configPath := "/nonexistent/path/config.json"

	_, err := fromFile(configPath)
	if err == nil {
		t.Errorf("Expected error for non-existent file, got nil")
	}
}
