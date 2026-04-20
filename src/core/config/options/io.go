package configOptions

import (
	"fmt"
	"os"
	"trains/src/core/config/types"
)

func IOFromEnv(config *types.Config) {
	// IO
	if v := os.Getenv("TRAINS_IO_INPUT_FORMAT"); v != "" {
		fileFormat, err := types.FlagToFileFormat(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.IO.InputFormat = types.FileFormat(fileFormat)
	}
	if v := os.Getenv("TRAINS_IO_OUTPUT_FORMAT"); v != "" {
		fileFormat, err := types.FlagToFileFormat(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.IO.OutputFormat = types.FileFormat(fileFormat)
	}
	if v := os.Getenv("TRAINS_IO_SOURCE_PATH"); v != "" {
		config.IO.SourcePath = v
	}
	if v := os.Getenv("TRAINS_IO_TARGET_PATH"); v != "" {
		config.IO.TargetPath = v
	}
}
