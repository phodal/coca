package identifier

import (
	"coca/core/support"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var nodeInfos []JIdentifier = nil

type JavaIdentifierApp struct {
}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) []JIdentifier {
	nodeInfos = nil
	files := support.GetJavaFiles(codeDir)
	for index := range files {
		file := files[index]

		parser := support.ProcessFile(file)
		context := parser.CompilationUnit()

		clzInfo := NewJIdentifier()
		listener := new(JavaIdentifierListener)
		listener.InitNode(clzInfo)

		antlr.NewParseTreeWalker().Walk(listener, context)

		if clzInfo.Name != "" {
			clzInfo.Methods = clzInfo.GetMethods()
			nodeInfos = append(nodeInfos, *clzInfo)
		}
	}

	return nodeInfos
}
