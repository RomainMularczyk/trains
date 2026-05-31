package translationValidator

import (
	"encoding/json"
	"os"
	"testing"
	"trains/src/core/reader/parser"
	testdataBatch "trains/testdata/sources/batch"
)

func TestSimpleBatchValidation(t *testing.T) {
	inputFile, _ := os.ReadFile("../../../../../testdata/outputs/translation/simple.json")
	var input []parser.TranslationEngineEntry
	_ = json.Unmarshal(inputFile, &input)

	testChan := make(chan parser.TranslationResult)
	resultChan := make(chan parser.TranslationResult)

	BatchValidator := BatchValidator{}
	go func() {
		defer close(resultChan)
		BatchValidator.Validate(testChan, resultChan)
	}()

	simpleBatch := testdataBatch.SimpleBatch()
	testChan <- parser.TranslationResult{
		Batch:  simpleBatch,
		Result: string(inputFile),
	}

	close(testChan)

	for result := range resultChan {
		if result.Error != nil {
			t.Error("Test failed")
		}
	}
}

func TestInvalidLengthBatchValidation(t *testing.T) {
}
