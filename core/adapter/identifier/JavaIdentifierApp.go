package identifier

import (
	parser2 "coca/core/languages/java"
	"coca/core/models"
	"coca/core/support"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var nodeInfos []models.JsonIdentifier = nil

type JavaIdentifierApp struct {
}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) []models.JsonIdentifier {
	nodeInfos = nil
	files := support.GetJavaFiles(codeDir)
	for index := range files {
		file := files[index]
		node := models.NewJsonIdentifier()

		parser := support.ProcessFile(file)
		context := parser.CompilationUnit()

		clzInfo := models.NewJIdentifier()
		listener := new(JavaIdentifierListener)
		listener.InitNode(clzInfo)

		antlr.NewParseTreeWalker().Walk(listener, context)

		if clzInfo.Name != "" {
			node = &models.JsonIdentifier{clzInfo.Pkg, clzInfo.Name, clzInfo.Type, clzInfo.GetMethods()}
			nodeInfos = append(nodeInfos, *node)
		}
	}

	return nodeInfos
}
