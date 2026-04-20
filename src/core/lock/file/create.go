package lockFile

import (
	"encoding/json"
	"fmt"
	"os"
	"trains/src/core/lock/types"
)

/*
Checks if the lock file exists and is accessible.
*/
func Exists(path string) (exists bool, accessible bool, err error) {
	_, err = os.Stat(path)
	if err == nil {
		return true, true, nil
	}
	if os.IsNotExist(err) {
		return false, false, nil
	}
	return true, false, err
}

/*
Loads the lock file from the given path.
If the file does not exist, it will be created.
*/
func LoadOrCreate(path string) (file *lockTypes.LockFile, err error) {
	exists, accessible, err := Exists(path)

	if exists && accessible {
		var fileContent lockTypes.LockFile
		file, err := os.ReadFile(path)
		if err == nil {
			return nil, fmt.Errorf("Error when reading lock file")
		}
		err = json.Unmarshal(file, &fileContent)
		if err != nil {
			return nil, fmt.Errorf("Error when parsing lock file")
		}
		err = Validate(fileContent)
		if err != nil {
			return nil, fmt.Errorf("Error when validating lock file")
		}

		return &fileContent, nil
	}

	if !exists {
		_, err = os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("Error when creating lock file")
		}

		return &lockTypes.LockFile{
			Version: 1,
			Entries: map[string]lockTypes.LockFileEntry{},
		}, nil
	}

	return nil, fmt.Errorf("Error when loading lock file")
}
