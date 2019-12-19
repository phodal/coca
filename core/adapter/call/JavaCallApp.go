package call

import (
	"coca/core/models"
	"coca/core/support"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
)

var nodeInfos []models.JClassNode

type JavaCallApp struct {
}

func (j *JavaCallApp) AnalysisPath(codeDir string, classes []string, identNodes []models.JIdentifier) []models.JClassNode {
	nodeInfos = nil
	files := support.GetJavaFiles(codeDir)
	for index := range files {
		nodeInfo := models.NewClassNode()
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := support.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaCallListener(identNodes)
		listener.appendClasses(classes)

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodeInfo = listener.getNodeInfo()
		nodeInfo.Path = file
		nodeInfos = append(nodeInfos, *nodeInfo)
	}

	return nodeInfos
}
