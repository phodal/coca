package processor

import (
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/trial"
	"github.com/phodal/coca/trial/cocago"
	"strings"
	"sync"
)

func ProcessPackage(path string) []*trial.CodeFile {
	var wg sync.WaitGroup

	var GoFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".go")
	}

	files := cocafile.GetFilesWithFilter(path ,GoFileFilter)
	filesData := make([]*trial.CodeFile, len(files))
	parser := cocago.NewCocagoParser()
	for i, file := range files {
		wg.Add(1)
		go func(i int, file string) {
			defer wg.Done()
			processFile := parser.ProcessFile(file)
			filesData[i] = &processFile
		}(i, file)
	}
	wg.Wait()

	return filesData
}
