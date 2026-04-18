package io

import (
	"io/fs"
	"path/filepath"
)

func FileExplorerWorker(root string, reader FileReader, filePool chan<- string) error {
	defer close(filePool)

	// handle case in which we don't have a base directory
	// defined
	if root == "" {
		root = "."
	}

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && reader.Supports(path) {
			filePool <- path
		}

		return nil
	})
}
