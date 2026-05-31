package configTypes

import (
	lockTypes "trains/src/core/lock/types"
)

type RuntimeConfig struct {
	Batching    Batching
	Config      Config
	IO          IO
	Lock        Lock
	Logger      *StageLogger
	Prompt      Prompt
	Provider    Provider
	Translation Translation
	Locks       map[string]lockTypes.LockFileEntry
}

type ConfigFile struct {
	Batching         Batching
	Config           Config
	IO               IO
	Lock             Lock
	Logging          Logging
	Prompt           Prompt
	Providers        Providers
	SelectedProvider ProviderName
	Translation      Translation
}

type ConfigFileOverrides struct {
	Batching         *BatchingOverrides    `json:"batching,omitempty,dive"`
	IO               *IOOverrides          `json:"io,omitempty,dive"`
	Lock             *LockOverrides        `json:"lock,omitempty,dive"`
	Logging          *LoggingOverrides     `json:"logging,omitempty,dive"`
	Prompt           *PromptOverrides      `json:"prompt,omitempty,dive"`
	Providers        *ProvidersOverrides   `json:"providers,omitempty,dive"`
	SelectedProvider *ProviderName         `json:"selected_provider,omitempty,dive"`
	Translation      *TranslationOverrides `json:"translation,omitempty,dive"`
}

type Config struct {
	Path string
}
