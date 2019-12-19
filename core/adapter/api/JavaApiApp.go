package api

import (
	"coca/core/adapter/identifier"
	"coca/core/models"
	"coca/core/support"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
)

var parsedDeps []models.JClassNode
var allApis []RestApi

type JavaApiApp struct {

}

func (j *JavaApiApp) AnalysisPath(codeDir string, depPath string) []RestApi {
	parsedDeps = nil
	file := support.ReadFile(depPath)
	if file == nil {
		return nil
	}

	_ = json.Unmarshal(file, &parsedDeps)

	identifiers := LoadIdentify(depPath)
	var identifiersMap = make(map[string]models.JIdentifier)

	for _, ident := range identifiers {
		identifiersMap[ident.Package + "." + ident.ClassName] = ident
	}

	files := support.GetJavaFiles(codeDir)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := support.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaApiListener(identifiersMap)
		listener.appendClasses(parsedDeps)

		antlr.NewParseTreeWalker().Walk(listener, context)

		allApis = listener.getClassApis()
	}

	return *&allApis
}


func LoadIdentify(importPath string) []models.JIdentifier {
	var identifiers []models.JIdentifier

	apiContent := support.ReadCocaFile("identify.json")
	if apiContent == nil {
		identifierApp := new(identifier.JavaIdentifierApp)
		ident := identifierApp.AnalysisPath(importPath)

		identModel, _ := json.MarshalIndent(ident, "", "\t")
		support.WriteToCocaFile("identify.json", string(identModel))

		return *&ident
	}
	_ = json.Unmarshal(apiContent, &identifiers)

	return *&identifiers
}