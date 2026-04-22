package resolvers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	configTypes "trains/src/core/config/types"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

/*
Reads a JSON configuration file and returns a Config object.
*/
func FromFile(
	path string,
) (*configTypes.ConfigFileOverrides, error) {
	validate = validator.New()

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configFile configTypes.ConfigFileOverrides
	decoder := json.NewDecoder(bytes.NewReader(fileContent))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&configFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileContent, &configFile)
	if err != nil {
		return nil, err
	}

	err = validate.Struct(configFile)
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

	return &configFile, nil
}
