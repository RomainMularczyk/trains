package lockFile

import (
	"errors"
	"fmt"
	"strings"
	trainsError "trains/src/core/errors"
	lockTypes "trains/src/core/lock/types"

	"github.com/go-playground/validator/v10"
)

/*
Validates the lock file.
*/
func Validate(fileContent lockTypes.LockFile) *trainsError.TrainsError {
	validate := validator.New()

	var errorMessages []string
	for key, entry := range fileContent.Entries {
		err := validate.Struct(entry)
		if err != nil {
			var validateErrors validator.ValidationErrors

			if errors.As(err, &validateErrors) {
				for _, e := range validateErrors {
					errorMessages = append(
						errorMessages,
						fmt.Sprintf(
							"key '%s': field '%s' failed validation tag '%s'",
							key,
							e.Field(),
							e.Tag(),
						),
					)
					continue
				}
				return &trainsError.TrainsError{
					Message: fmt.Sprintf(
						"Configuration validation failed:\n%s",
						strings.Join(errorMessages, "\n"),
					),
					Code: trainsError.InvalidLockError,
				}
			}

			return &trainsError.TrainsError{
				Message: fmt.Sprintf(
					"An unexpected validation error occurred on key '%s': '%w'",
					key,
					err,
				),
				Code: trainsError.UnexpectedError,
				Err:  err,
			}
		}
	}

	if len(errorMessages) > 0 {
		return &trainsError.TrainsError{
			Message: fmt.Sprintf(
				"Configuration validation failed:\n%s",
				strings.Join(errorMessages, "\n"),
			),
			Code: trainsError.InvalidLockError,
		}
	}

	return nil
}
