package orchestration

import (
	"fmt"
	"sync"
	"trains/src/core/config/types"
	"trains/src/core/io"
	"trains/src/core/parser"
	"trains/src/core/translation"
)

type Pipeline struct{}

/*
Run the translation pipeline.
*/
func (p *Pipeline) Run(
	root string,
	format types.FileFormat,
	config types.RuntimeConfig,
) {
	// Create channels
	filePaths := make(chan string)
	fileContents := make(chan any)
	translationUnits := make(chan parser.TranslationUnit)
	batches := make(chan parser.TranslationBatch)
	translations := make(chan string)

	processor, err := io.Processor(format)
	llm := translation.NewLLM(config)
	fmt.Println(llm)
	if err != nil {
		// TODO: handle error
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(5)

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
		defer close(batches)
		parser.CreateBatch(translationUnits, batches, 100, config)
	}()

	go func() {
		defer wg.Done()
		defer close(translations)
		llm.Translate(config, batches, translations)
	}()

	go func() {
		for translation := range translations {
			fmt.Println(translation)
		}
	}()

	wg.Wait()
}
