package orchestration

import (
	"trains/src/core/io"
)

pathChan := make(chan FileJob, 100)
fileContentChan := make(chan FileContent, 100)
translationUnitChan := make(chan TranslationUnit, 100)
batchChan := make(chan TranslationSet, 100)

go FileExplorerWorker(pathChan)

for i := 0; i < 10; i++ {
	go FileWorker(pathChan, fileContentChan)
}

for i := 0; i < 10; i++ {
	go TranslateWorker(fileContentChan, translationUnitChan)
}

for i := 0; i < 10; i++ {
	go BatchWorker(translationUnitChan, batchChan)
}
