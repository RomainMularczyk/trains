package lockTypes

import "encoding/json"

type OriginName string

const (
	LLM   OriginName = "llm"
	Human OriginName = "human"
)

type LockFile struct {
	Version int                      `json:"_version" validate:"required"`
	Entries map[string]LockFileEntry `json:"entries" validate:"required"`
}

type LockFileEntry struct {
	SourceHash      string     `json:"source_hash" validate:"required,len=32"`
	TranslationHash string     `json:"translation_hash" validate:"required,len=32"`
	Origin          OriginName `json:"origin" validate:"required,oneof=llm human"`
}

/*
Format the lock file as a JSON string.
*/
func (c LockFile) String() string {
	config, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(config)
}
