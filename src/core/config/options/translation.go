package configOptions

import (
	"fmt"
	"os"
	"trains/src/core/config/types"
)

func TranslationFromEnv(config *types.ConfigFile) {
	if v := os.Getenv("TRAINS_TRANSLATION_SOURCE_LANGUAGE"); v != "" {
		srcLang, err := types.GetLanguage(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.Translation.SourceLanguage = srcLang
	}
	if v := os.Getenv("TRAINS_TRANSLATION_TARGET_LANGUAGE"); v != "" {
		targetLang, err := types.GetLanguage(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.Translation.TargetLanguage = targetLang
	}
}
