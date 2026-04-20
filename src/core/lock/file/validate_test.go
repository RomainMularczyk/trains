package lockFile

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
	lockTypes "trains/src/core/lock/types"
)

func SeedLockFile(numEntries int, chanceOfHumanEdits float64) lockTypes.LockFile {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	lock := lockTypes.LockFile{
		Version: 1,
		Entries: map[string]lockTypes.LockFileEntry{},
	}

	for i := range make([]int, numEntries) {
		key := fmt.Sprintf("feature.item_%d", i)
		source := fmt.Sprintf("Source string to translate %d", i)
		translation := fmt.Sprintf("Translated string %d", i)
		srcHash := fmt.Sprintf("%x", sha256.Sum256([]byte(source)))
		translationHash := fmt.Sprintf("%x", sha256.Sum256([]byte(translation)))

		// we simulate some human edits
		if r.Float64() < chanceOfHumanEdits {
			translation = fmt.Sprintf("Human edited string %d", i)
			lock.Entries[key] = lockTypes.LockFileEntry{
				SourceHash:      srcHash,
				TranslationHash: translationHash,
				Origin:          "human",
			}
		} else {
			lock.Entries[key] = lockTypes.LockFileEntry{
				SourceHash:      srcHash,
				TranslationHash: translationHash,
				Origin:          "llm",
			}
		}
	}

	return lock
}

func TestValidate(t *testing.T) {
	seed := SeedLockFile(10, 0.1)
	err := Validate(seed)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestValidate_InvalidOrigin(t *testing.T) {
	lock := lockTypes.LockFile{
		Version: 1,
		Entries: map[string]lockTypes.LockFileEntry{
			"test.key": {
				SourceHash:      strings.Repeat("a", 64),
				TranslationHash: strings.Repeat("b", 64),
				Origin:          "invalid",
			},
		},
	}

	err := Validate(lock)
	if err == nil {
		t.Error("Expected validation error for invalid origin, got nil")
	}
}

func TestValidate_MissingSourceHash(t *testing.T) {
	lock := lockTypes.LockFile{
		Version: 1,
		Entries: map[string]lockTypes.LockFileEntry{
			"test.key": {
				SourceHash:      "",
				TranslationHash: strings.Repeat("b", 64),
				Origin:          "llm",
			},
		},
	}

	err := Validate(lock)
	if err == nil {
		t.Error("Expected validation error for missing source_hash, got nil")
	}
}

func TestValidate_WrongLengthSourceHash(t *testing.T) {
	lock := lockTypes.LockFile{
		Version: 1,
		Entries: map[string]lockTypes.LockFileEntry{
			"test.key": {
				SourceHash:      "tooshort",
				TranslationHash: strings.Repeat("b", 64),
				Origin:          "llm",
			},
		},
	}

	err := Validate(lock)
	if err == nil {
		t.Error("Expected validation error for wrong length source_hash, got nil")
	}
}
