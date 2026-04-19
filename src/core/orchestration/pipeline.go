package orchestration

import (
	"fmt"
	"sync"
	"trains/src/core/io"
	"trains/src/core/translation"
)

type Pipeline struct{}

func (p *Pipeline) Run(root string, format io.FileFormat) {
	filePaths := make(chan string)
	fileContents := make(chan any)
	translationUnits := make(chan translation.TranslationUnit)
	batches := make([]translation.TranslationBatch, 0)

	processor, err := io.Processor(format)
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		defer close(filePaths)
		io.FileExplorerWorker(root, processor.Reader, filePaths)
	}()

	go func() {
		defer wg.Done()
		defer close(fileContents)
		processor.Reader.Read(filePaths, fileContents)
	}()

	go func() {
		defer wg.Done()
		defer close(translationUnits)
		processor.Parser.Parse(fileContents, translationUnits)
	}()

	go func() {
		defer wg.Done()
		translation.CreateBatch(translationUnits, &batches, 100)
	}()

	wg.Wait()
}
