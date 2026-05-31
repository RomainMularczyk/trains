package validators

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	configTypes "trains/src/core/config/types"
	trainsErrors "trains/src/core/errors"
	"trains/src/utils"

	"github.com/go-playground/validator/v10"
)

func ValidateConfigFile(
	fileContent []byte,
	bootstrapLogger *configTypes.StageLogger,
) (*configTypes.ConfigFileOverrides, *trainsErrors.TrainsError) {
	validate := validator.New()

	var configFile configTypes.ConfigFileOverrides
	decoder := json.NewDecoder(bytes.NewReader(fileContent))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&configFile)
	if decoder.More() {
		return nil, &trainsErrors.TrainsError{
			Code:    trainsErrors.InvalidConfigError,
			Message: "Two configuration files were provided",
			Err:     err,
		}
	}

	if err := bootstrapLogger.Add("parse", err); err != nil {
		return nil, &trainsErrors.TrainsError{
			Code:    trainsErrors.InvalidConfigError,
			Message: "Failed to validate config file structure",
			Err:     err,
		}
	}
	bootstrapLogger.Debug("Configuration file parsed successfully")

	err = validate.Struct(configFile)
	if err := bootstrapLogger.Add("validate", err, utils.FormatJSON(configFile)); err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return nil, &trainsErrors.TrainsError{
				Code:    trainsErrors.InvalidConfigError,
				Message: "Failed to validate config file",
				Err:     err,
			}
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
			return nil, &trainsErrors.TrainsError{
				Code:    trainsErrors.InvalidConfigError,
				Message: "Failed to validate config file",
				Err: fmt.Errorf(
					"Configuration validation failed:\n%s",
					strings.Join(errorMessages, "\n"),
				),
			}
		}

		return nil, &trainsErrors.TrainsError{
			Code:    trainsErrors.InvalidConfigError,
			Message: "Failed to validate config file",
			Err:     fmt.Errorf("Validation config: %w", err),
		}
	}
	bootstrapLogger.Debug("Configuration file validated successfully")
	bootstrapLogger.Info("Configuration file loaded successfully")

	return &configFile, nil
}
