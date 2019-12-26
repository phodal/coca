package call

import (
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
)

var nodeInfos []models.JClassNode

type JavaCallApp struct {
}

func NewJavaCallApp() JavaCallApp {
	return *&JavaCallApp{}
}

func (j *JavaCallApp) AnalysisPath(codeDir string, classes []string, identNodes []models.JIdentifier) []models.JClassNode {
	nodeInfos = nil
	files := support.GetJavaFiles(codeDir)

	var identMap = make(map[string]models.JIdentifier)
	for _, ident := range identNodes {
		identMap[ident.Package + "." + ident.ClassName] = ident
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
