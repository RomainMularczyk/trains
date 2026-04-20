package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFromFile_ValidConfig(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")

	validJSON := `{
		"Provider": {
			"openai": {
				"ApiKey": "test-key",
				"Model": "gpt-4",
				"BaseUrl": "https://api.example.com",
				"Timeout": 30
			},
			"anthropic": {
				"ApiKey": "test-key",
				"Model": "claude-3",
				"BaseUrl": "https://api.anthropic.com",
				"Timeout": 30
			},
			"google": {
				"ApiKey": "test-key",
				"Model": "gemini-pro",
				"BaseUrl": "https://api.google.com",
				"Timeout": 30
			},
			"mistral": {
				"ApiKey": "test-key",
				"Model": "mistral-large",
				"BaseUrl": "https://api.mistral.com",
				"Timeout": 30
			},
			"xai": {
				"ApiKey": "test-key",
				"Model": "grok",
				"BaseUrl": "https://api.x.ai",
				"Timeout": 30
			},
			"deepseeker": {
				"ApiKey": "test-key",
				"Model": "deepseek-chat",
				"BaseUrl": "https://api.deepseek.com",
				"Timeout": 30
			},
			"cohere": {
				"ApiKey": "test-key",
				"Model": "command",
				"BaseUrl": "https://api.cohere.com",
				"Timeout": 30
			},
			"perplexity": {
				"ApiKey": "test-key",
				"Model": "llama-3",
				"BaseUrl": "https://api.perplexity.com",
				"Timeout": 30
			},
			"openrouter": {
				"ApiKey": "test-key",
				"Model": "gpt-4",
				"BaseUrl": "https://api.openrouter.com",
				"Timeout": 30
			},
			"minimax": {
				"ApiKey": "test-key",
				"Model": "minimax-01",
				"BaseUrl": "https://api.minimax.com",
				"Timeout": 30
			}
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
		},
		"IO": {
			"InputFormat": 1,
			"OutputFormat": 1,
			"SourcePath": "./source",
			"TargetPath": "./target"
		}
	}`

	err := os.WriteFile(configPath, []byte(validJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	config, err := fromFile(configPath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if config == nil {
		t.Fatal("Expected config to be non-nil")
	}

	// Verify specific fields
	if config.Provider.OpenAI.ApiKey != "test-key" {
		t.Errorf("Expected OpenAI ApiKey to be 'test-key', got '%s'", config.Provider.OpenAI.ApiKey)
	}
	if config.Translation.SourceLanguage != "en" {
		t.Errorf("Expected SourceLanguage to be 'en', got '%s'", config.Translation.SourceLanguage)
	}
	if config.Batching.TokenLimit != 1000 {
		t.Errorf("Expected TokenLimit to be 1000, got %d", config.Batching.TokenLimit)
	}
	if config.Prompt.Context != "Translate the following text" {
		t.Errorf("Expected Context to be 'Translate the following text', got '%s'", config.Prompt.Context)
	}
	if config.IO.SourcePath != "./source" {
		t.Errorf("Expected SourcePath to be './source', got '%s'", config.IO.SourcePath)
	}
}

func TestFromFile_InvalidJSON(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")

	// because of trailing commas
	invalidJSON := `{
		"Provider": {
			"openai": {
				"ApiKey": "test-key",
				"Model": "gpt-4",
				"BaseUrl": "https://api.example.com",
				"Timeout": 30,
			},
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
			"openai": {
				"ApiKey": "test-key",
				"Model": "gpt-4"
			}
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
