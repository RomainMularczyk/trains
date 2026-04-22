package errors

import (
	"errors"
	"fmt"
)

type Code string

const (
	InvalidConfigError Code = "INVALID_CONFIG"
	InvalidLockError   Code = "INVALID_LOCK"
)

type TrainsError struct {
	Code    Code
	Message string
	Err     error
	Context map[string]string
}

func (e *TrainsError) Error() string {
	return e.Message
}

func (e *TrainsError) Unwrap() error {
	return e.Err
}

func FormatError(e error, verobose bool) string {
	var trainsError *TrainsError

	if errors.As(e, &trainsError) {

		message := fmt.Sprintf("✖ %s\n", trainsError.Message)

		return message
	}

	return fmt.Sprintf("✖ %s\n", e.Error())
}
