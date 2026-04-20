package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"trains/src/core/config/options"
	"trains/src/core/config/types"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

/*
Resolve the configuration applying variable precedence.
Variable precedence order:
 1. Command line arguments
 2. Environment variables
 3. Configuration file
 4. Default values
*/
func ResolveConfig(path string) types.Config {
	config, err := fromFile(path)
	if err != nil {
		fmt.Println(err)
	}

	fromEnv(config)
	fromOptions(config)

	return *config
}

/*
Reads a JSON configuration file and returns a Config object.
*/
func fromFile(path string) (*types.Config, error) {
	validate = validator.New()

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config types.Config
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

func fromEnv(config *types.Config) {
	configOptions.BatchingFromEnv(config)
	configOptions.IOFromEnv(config)
	configOptions.PromptFromEnv(config)
	configOptions.ProviderFromEnv(config)
	configOptions.TranslationFromEnv(config)
}

func fromOptions(config *types.Config) {

}
