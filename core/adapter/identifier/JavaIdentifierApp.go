package identifier

import (
	"coca/core/models"
	"coca/core/support"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var nodeInfos []models.JIdentifier = nil

type JavaIdentifierApp struct {
}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) []models.JIdentifier {
	nodeInfos = nil
	files := support.GetJavaFiles(codeDir)
	for index := range files {
		file := files[index]

		parser := support.ProcessFile(file)
		context := parser.CompilationUnit()

		clzInfo := models.NewJIdentifier()
		listener := new(JavaIdentifierListener)
		listener.InitNode(clzInfo)

		antlr.NewParseTreeWalker().Walk(listener, context)

		if clzInfo.ClassName != "" {
			clzInfo.Methods = clzInfo.GetMethods()
			nodeInfos = append(nodeInfos, *clzInfo)
		}
	}

	return nodeInfos
}
