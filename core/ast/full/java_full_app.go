package full

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure"
	"path/filepath"
)

type JavaFullApp struct {
}

func NewJavaFullApp() JavaFullApp {
	return *&JavaFullApp{}
}

func (j *JavaFullApp) AnalysisPath(codeDir string, classes []string, identNodes []domain.JIdentifier) []domain.JClassNode {
	files := infrastructure.GetJavaFiles(codeDir)
	return j.AnalysisFiles(identNodes, files, classes)
}

func (j *JavaFullApp) AnalysisFiles(identNodes []domain.JIdentifier, files []string, classes []string) []domain.JClassNode {
	var nodeInfos []domain.JClassNode

	var identMap = make(map[string]domain.JIdentifier)
	for _, ident := range identNodes {
		identMap[ident.GetClassFullName()] = ident
	}

	for _, file := range files {
		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := infrastructure.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaFullListener(identMap, file)
		listener.appendClasses(classes)

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodes := listener.getNodeInfo()
		nodeInfos = append(nodeInfos, nodes...)
	}

	return nodeInfos
}
