package call

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"path/filepath"
)

type JavaCallApp struct {
}

func NewJavaCallApp() JavaCallApp {
	return *&JavaCallApp{}
}

func (j *JavaCallApp) AnalysisPath(codeDir string, classes []string, identNodes []models.JIdentifier) []models.JClassNode {
	files := support.GetJavaFiles(codeDir)
	return j.AnalysisFiles(identNodes, files, classes)
}

func (j *JavaCallApp) AnalysisFiles(identNodes []models.JIdentifier, files []string, classes []string) []models.JClassNode {
	var nodeInfos []models.JClassNode

	var identMap = make(map[string]models.JIdentifier)
	for _, ident := range identNodes {
		identMap[ident.Package+"."+ident.ClassName] = ident
	}

	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := support.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaCallListener(identMap, file)
		listener.appendClasses(classes)

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodes := listener.getNodeInfo()
		nodeInfos = append(nodeInfos, nodes...)
	}

	return nodeInfos
}
