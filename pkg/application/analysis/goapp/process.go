package goapp

import (
	"bytes"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/cocago"
	"strings"
)

func ProcessPackage(path string, debug bool) []*core_domain.CodeFile {
	var GoFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".go")
	}

	files := cocafile.GetFilesWithFilter(path, GoFileFilter)
	filesData := make([]*core_domain.CodeFile, len(files))
	parser := cocago.NewCocagoParser()
	if debug {
		buf := new(bytes.Buffer)
		parser.SetOutput(buf)
	}
	for i, file := range files {
		processFile := parser.ProcessFile(file)
		filesData[i] = &processFile
	}

	return filesData
}
