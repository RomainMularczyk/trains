package errors

import (
	"errors"
	"fmt"
)

type Code string

const (
	InvalidConfigError   Code = "INVALID_CONFIG"
	InvalidLockError     Code = "INVALID_LOCK"
	InvalidProviderError Code = "INVALID_PROVIDER"
	InvalidFileError     Code = "INVALID_FILE"

	FileNotAccessible Code = "FILE_NOT_ACCESSIBLE"
	UnexpectedError   Code = "UNEXPECTED_ERROR"

	TranslationError Code = "TRANSLATION_ERROR"
)

type TrainsError struct {
	Code    Code
	Message string
	Err     error
	Context map[string]string
}

func (e *TrainsError) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

func (e *TrainsError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

/*
Formats the error in a simple way.
*/
func FormatError(e error, verbose bool) string {
	var trainsError *TrainsError

	if errors.As(e, &trainsError) {
		if verbose {
			return VerboseFormatError(trainsError)
		} else {
			return fmt.Sprintf("✖ %s\n", trainsError.Message)
		}
	} else {
		return fmt.Sprintf("✖ %s\n", e.Error())
	}
}

/*
Formats the error by adding the root cause.
*/
func VerboseFormatError(e *TrainsError) string {
	message := fmt.Sprintf(
		"✖ %s\n%s\n",
		e.Message,
		e.Err,
	)
	return message
}
