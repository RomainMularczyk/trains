package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"trains/src/core/io"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Provider    Provider
	Translation Translation
	Prompt      Prompt
	Batching    Batching
	IO          IO
}

type CLIOptions struct {
	Provider    *ProviderOverrides
	Translation *TranslationOverrides
	Prompt      *PromptOverrides
	Batching    *BatchingOverrides
	IO          *IOOverrides
}

/*
Format the configuration as a JSON string.
*/
func (c Config) String() string {
	config, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(config)
}

var validate *validator.Validate

/*
Resolve the configuration applying variable precedence.
Variable precedence order:
 1. Command line arguments
 2. Environment variables
 3. Configuration file
 4. Default values
*/
func ResolveConfig(path string) Config {
	config, err := fromFile(path)
	if err != nil {
		return Config{}
	}

	fromEnv(config)
	fromOptions(config)

	return *config
}

/*
Reads a JSON configuration file and returns a Config object.
*/
func fromFile(path string) (*Config, error) {
	validate = validator.New()

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}

	err = validate.Struct(config)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return nil, err
		}

		var validateErrors validator.ValidationErrors
		if errors.As(err, &validateErrors) {
			var errorMessages []string
			for _, e := range validateErrors {
				errorMessages = append(
					errorMessages,
					fmt.Sprintf(
						"[%s][%s] The property '%s' is required.",
						e.StructNamespace(),
						e.Type(),
						e.Field()),
				)
			}
			return nil, fmt.Errorf(
				"Configuration validation failed:\n%s",
				strings.Join(errorMessages, "\n"),
			)
		}

		return nil, fmt.Errorf("Validation config: %w", err)
	}

	return &config, nil
}

/*
Reads environment variables and overrides the configuration.
*/
func fromEnv(config *Config) {
	// Provider
	if v := os.Getenv("TRAINS_PROVIDER_API_KEY"); v != "" {
		config.Provider.ApiKey = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_MODEL"); v != "" {
		config.Provider.Model = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_BASE_URL"); v != "" {
		config.Provider.BaseUrl = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_TIMEOUT"); v != "" {
		if timeout, err := strconv.Atoi(v); err == nil {
			config.Provider.Timeout = timeout
		}
	}

	// Translation
	if v := os.Getenv("TRAINS_TRANSLATION_SOURCE_LANGUAGE"); v != "" {
		srcLang, err := GetLanguage(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.Translation.SourceLanguage = srcLang
	}
	if v := os.Getenv("TRAINS_TRANSLATION_TARGET_LANGUAGE"); v != "" {
		targetLang, err := GetLanguage(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.Translation.TargetLanguage = targetLang
	}

	// IO
	if v := os.Getenv("TRAINS_IO_INPUT_FORMAT"); v != "" {
		fileFormat, err := io.FlagToFileFormat(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.IO.InputFormat = io.FileFormat(fileFormat)
	}
	if v := os.Getenv("TRAINS_IO_OUTPUT_FORMAT"); v != "" {
		fileFormat, err := io.FlagToFileFormat(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.IO.OutputFormat = io.FileFormat(fileFormat)
	}
	if v := os.Getenv("TRAINS_IO_SOURCE_PATH"); v != "" {
		config.IO.SourcePath = v
	}
	if v := os.Getenv("TRAINS_IO_TARGET_PATH"); v != "" {
		config.IO.TargetPath = v
	}

	// Batching
	if v := os.Getenv("TRAINS_BATCHING_TOKEN_LIMIT"); v != "" {
		tokenLimit, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.Batching.TokenLimit = tokenLimit
	}

	// Prompt
	if v := os.Getenv("TRAINS_PROMPT_CONTEXT"); v != "" {
		config.Prompt.Context = v
	}
}

func fromOptions(config *Config) {

}
