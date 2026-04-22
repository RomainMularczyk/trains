package configTypes

import (
	"encoding/json"
	lockTypes "trains/src/core/lock/types"
)

type RuntimeConfig struct {
	SelectedProvider Provider
	Locks            []lockTypes.LockFileEntry
}

type ConfigFile struct {
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
Creates a slice of lock file entries from the given lock file.
*/
func CreateLockFileEntries(lockFile lockTypes.LockFile) []lockTypes.LockFileEntry {
	locks := make([]lockTypes.LockFileEntry, 0, len(lockFile.Entries))

	for _, entry := range lockFile.Entries {
		locks = append(locks, entry)
	}

	return locks
}

/*
Format the configuration as a JSON string.
*/
func (c ConfigFile) String() string {
	config, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(config)
}
