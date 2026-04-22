package configOptions

import (
	"fmt"
	"os"
	configTypes "trains/src/core/config/types"
)

func TranslationFromEnv(config *configTypes.ConfigFile) {
	if v := os.Getenv("TRAINS_TRANSLATION_SOURCE_LANGUAGE"); v != "" {
		srcLang, err := configTypes.GetLanguage(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.Translation.SourceLanguage = srcLang
	}
	if v := os.Getenv("TRAINS_TRANSLATION_TARGET_LANGUAGE"); v != "" {
		targetLang, err := configTypes.GetLanguage(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.Translation.TargetLanguage = targetLang
	}
}
