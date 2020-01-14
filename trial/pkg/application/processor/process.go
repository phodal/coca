package processor

import (
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/trial"
	"github.com/phodal/coca/trial/cocago"
	"strings"
)

func ProcessPackage(path string, debug bool) []*trial.CodeFile {
	var GoFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".go")
	}

	files := cocafile.GetFilesWithFilter(path, GoFileFilter)
	filesData := make([]*trial.CodeFile, len(files))
	parser := cocago.NewCocagoParser()
	if debug {
		parser.SetOutput(true)
	}
	for i, file := range files {
		processFile := parser.ProcessFile(file)
		filesData[i] = &processFile
	}

	return filesData
}
