package resolvers

import (
	"os"
	configTypes "trains/src/core/config/types"
	"trains/src/core/config/validators"
	trainsErrors "trains/src/core/errors"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

/*
Reads a JSON configuration file and returns a Config object.
*/
func FromFile(
	path string,
	bootstrapLogger *configTypes.StageLogger,
) (*configTypes.ConfigFileOverrides, *trainsErrors.TrainsError) {
	fileContent, err := os.ReadFile(path)
	if err := bootstrapLogger.Add("read", err); err != nil {
		return nil, &trainsErrors.TrainsError{
			Code:    trainsErrors.InvalidConfigError,
			Message: "Failed to read config file",
			Err:     err,
		}
	}
	bootstrapLogger.Debug("Configuration file read successfully")

	configFile, validError := validators.ValidateConfigFile(fileContent, bootstrapLogger)
	if err != nil {
		return nil, validError
	}

	return configFile, nil
}
