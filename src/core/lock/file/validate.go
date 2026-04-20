package lockFile

import (
	"errors"
	"fmt"
	"strings"
	lockTypes "trains/src/core/lock/types"

	"github.com/go-playground/validator/v10"
)

/*
Validates the lock file.
*/
func Validate(fileContent lockTypes.LockFile) error {
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
				return fmt.Errorf(
					"Configuration validation failed:\n%s",
					strings.Join(errorMessages, "\n"),
				)
			}

			return fmt.Errorf("An unexpected validation error occurred on key '%s': '%w'", key, err)
		}
	}

	if len(errorMessages) > 0 {
		return fmt.Errorf(
			"Configuration validation failed:\n%s",
			strings.Join(errorMessages, "\n"),
		)
	}

	return nil
}
