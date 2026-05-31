package orchestration

import (
	"fmt"
	"sync"
	"trains/src/core/config/types"
	"trains/src/core/io"
	"trains/src/core/reader/parser"
	"trains/src/core/reader/translation"
)

type Pipeline struct{}

/*
Run the translation pipeline.
*/
func (p *Pipeline) Run(
	config configTypes.RuntimeConfig,
) {
	// Create channels
	filePaths := make(chan string)
	fileContents := make(chan any)
	translationUnits := make(chan parser.TranslationUnit)
	batches := make(chan parser.TranslationBatch)
	translations := make(chan string)

	processor, err := io.Processor(config.IO.InputFormat)
	llm := translation.NewLLM(config)
	if err != nil {
		// TODO: handle error
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(5)

	// READ BLOCK
	go func() {
		defer wg.Done()
		defer close(filePaths)
		io.FileExplorerWorker(config.IO.SourcePath, processor.Reader, filePaths)
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
		defer close(batches)
		parser.CreateBatch(translationUnits, batches, config)
	}()

	go func() {
		defer wg.Done()
		defer close(translations)
		llm.Translate(config, batches, translations)
	}()

	// WRITE BLOCK

	wg.Wait()
}
