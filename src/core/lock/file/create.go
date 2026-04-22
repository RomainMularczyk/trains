package lockFile

import (
	"encoding/json"
	"os"
	"trains/src/core/errors"
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
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidLockError,
				Message: "Could not read the lock file",
				Err:     err,
			}
		}
		err = json.Unmarshal(file, &fileContent)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidLockError,
				Message: "Could not parse the lock file",
				Err:     err,
			}
		}
		err = Validate(fileContent)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidLockError,
				Message: "Invalid lock file",
				Err:     err,
			}
		}

		return &fileContent, nil
	}

	if !exists {
		_, err = os.Create(path)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidLockError,
				Message: "Could not create the lock file",
				Err:     err,
			}
		}

		return &lockTypes.LockFile{
			Version: 1,
			Entries: map[string]lockTypes.LockFileEntry{},
		}, nil
	}

	return nil, &errors.TrainsError{
		Code:    errors.InvalidLockError,
		Message: "Could not access the lock file",
		Err:     err,
	}
}
