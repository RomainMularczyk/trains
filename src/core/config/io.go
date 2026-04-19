package config

import "trains/src/core/io"

type IO struct {
	InputFormat  io.FileFormat `validate:"required"`
	OutputFormat io.FileFormat `validate:"required"`
	SourcePath   string        `validate:"required"`
	TargetPath   string        `validate:"required"`
}

type IOOverrides struct {
	InputFormat  *io.FileFormat
	OutputFormat *io.FileFormat
	SourcePath   *string
	TargetPath   *string
}
