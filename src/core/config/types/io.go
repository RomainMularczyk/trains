package configTypes

type IO struct {
	InputFormat  FileFormat `validate:"required"`
	OutputFormat FileFormat `validate:"required"`
	SourcePath   string     `validate:"required"`
	TargetPath   string     `validate:"required"`
}

type IOOverrides struct {
	InputFormat  *FileFormat
	OutputFormat *FileFormat
	SourcePath   *string
	TargetPath   *string
}
