package types

import "encoding/json"

type Config struct {
	Provider    Providers
	Translation Translation
	Prompt      Prompt
	Batching    Batching
	IO          IO
	Lock        Lock
}

type CLIOptions struct {
	Provider    *ProvidersOverrides
	Translation *TranslationOverrides
	Prompt      *PromptOverrides
	Batching    *BatchingOverrides
	IO          *IOOverrides
	Lock        *LockOverrides
}

/*
Format the configuration as a JSON string.
*/
func (c Config) String() string {
	config, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(config)
}
