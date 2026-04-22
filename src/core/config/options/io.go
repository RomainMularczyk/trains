package configOptions

import (
	"fmt"
	"os"
	configTypes "trains/src/core/config/types"
)

func IOFromEnv(config *configTypes.ConfigFile) {
	// IO
	if v := os.Getenv("TRAINS_IO_INPUT_FORMAT"); v != "" {
		fileFormat, err := configTypes.FlagToFileFormat(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.IO.InputFormat = configTypes.FileFormat(fileFormat)
	}
	if v := os.Getenv("TRAINS_IO_OUTPUT_FORMAT"); v != "" {
		fileFormat, err := configTypes.FlagToFileFormat(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.IO.OutputFormat = configTypes.FileFormat(fileFormat)
	}
	if v := os.Getenv("TRAINS_IO_SOURCE_PATH"); v != "" {
		config.IO.SourcePath = v
	}
	if v := os.Getenv("TRAINS_IO_TARGET_PATH"); v != "" {
		config.IO.TargetPath = v
	}
}
