package lockFile

import (
	"encoding/json"
	"os"
	"trains/src/core/errors"
	"trains/src/core/lock/types"
	"trains/src/utils"
)

/*
Checks if the lock file exists and is accessible.
*/
func Exists(path string) (bool, bool, *errors.TrainsError) {
	_, err := os.Stat(path)
	if err == nil {
		return true, true, &errors.TrainsError{
			Code:    errors.InvalidLockError,
			Message: "Lock file already exists",
			Err:     err,
		}
	}
	if os.IsNotExist(err) {
		return false, false, nil
	}
	return true, false, &errors.TrainsError{
		Code:    errors.InvalidLockError,
		Message: "Could not check if lock file exists",
		Err:     err,
	}
}

/*
Create the lock file.
*/
func create(path string) (*lockTypes.LockFile, *errors.TrainsError) {
	_, err := os.Create(path)
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

func load(path string) (*lockTypes.LockFile, *errors.TrainsError) {
	var fileContent lockTypes.LockFile
	file, err := os.ReadFile(path)

	if err != nil {
		return nil, &errors.TrainsError{
			Code:    errors.InvalidLockError,
			Message: "Could not read the lock file",
			Err:     err,
		}
	}

	if !utils.IsFileEmpty(file) {
		return &fileContent, nil
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

/*
Loads the lock file from the given path.
If the file does not exist, it will be created.
*/
func LoadOrCreate(path string) (*lockTypes.LockFile, *errors.TrainsError) {
	exists, accessible, err := Exists(path)

	if !accessible {
		return nil, &errors.TrainsError{
			Code:    errors.FileNotAccessible,
			Message: "The file is not accessible",
			Err:     err,
		}
	}

	if !exists {
		return create(path)
	}

	return load(path)
}
