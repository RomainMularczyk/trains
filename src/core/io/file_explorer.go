package io

import (
	"io/fs"
	"path/filepath"
)

func FileExplorerWorker(root string, reader FileReader, filePaths chan<- string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && reader.Supports(path) {
			filePaths <- path
		}

		return nil
	})
}
