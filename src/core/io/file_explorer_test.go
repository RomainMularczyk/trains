package io

import (
	"sync"
	"testing"
	configTypes "trains/src/core/config/types"
)

func TestDiscoverFiles(t *testing.T) {
	driver, err := Processor(configTypes.JSON)
	if err != nil {
		t.Errorf("Error creating reader: %v", err)
	}

	filePaths := make(chan string, 10)
	var files []string
	var wg sync.WaitGroup
	wg.Add(1)

	done := make(chan error)
	go func() {
		done <- FileExplorerWorker("../../../tests", driver.Reader, filePaths)
		close(filePaths)
	}()

	go func() {
		defer wg.Done()
		for path := range filePaths {
			files = append(files, path)
		}
	}()

	if err := <-done; err != nil {
		t.Errorf("Error discovering files: %v", err)
	}
	wg.Wait()

	if len(files) != 1 {
		t.Errorf("Expected 1 file, got %d", len(files))
	}
}
