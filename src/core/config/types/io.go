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

/*
Applies the given IO overrides to the IO configuration.
*/
func (i *IO) Apply(o *IOOverrides) {
	if o == nil {
		return
	}
	if o.InputFormat != nil {
		i.InputFormat = *o.InputFormat
	}
	if o.OutputFormat != nil {
		i.OutputFormat = *o.OutputFormat
	}
	if o.SourcePath != nil {
		i.SourcePath = *o.SourcePath
	}
	if o.TargetPath != nil {
		i.TargetPath = *o.TargetPath
	}
}
