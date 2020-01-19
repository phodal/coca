package goapp

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/cocago"
)

type GoIdentApp struct {
	Extensions interface{}
}

func (g *GoIdentApp) AnalysisPackageManager(path string) core_domain.CodePackageManagerInfo  {
	return core_domain.CodePackageManagerInfo{}
}

func (g *GoIdentApp) Analysis(code string, fileName string) core_domain.CodeFile {
	parser := cocago.NewCocagoParser()
	var codeMembers []core_domain.CodeMember
	if g.Extensions != nil {
		codeMembers = g.Extensions.([]core_domain.CodeMember)
	}
	return 	*parser.ProcessString(code, fileName, codeMembers)
}

func (g *GoIdentApp) IdentAnalysis(code string, fileName string) []core_domain.CodeMember {
	parser := cocago.NewCocagoParser()
	codeFile := parser.IdentAnalysis(code, fileName)
	return codeFile.Members
}

func (g *GoIdentApp) SetExtensions(extension interface{})  {
	g.Extensions = extension
}
