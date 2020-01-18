package goapp

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/cocago"
)

type GoIdentApp struct {
	Extensions interface{}
}

func (g *GoIdentApp) Analysis(code string, fileName string) core_domain.CodeFile {
	parser := cocago.NewCocagoParser()
	var codeMembers []core_domain.CodeMember
	if g.Extensions != nil {
		codeMembers = g.Extensions.([]core_domain.CodeMember)
	}
	return 	*parser.ProcessString(code, fileName, codeMembers)
}

func (g *GoIdentApp) IdentAnalysis(s string, file string) []core_domain.CodeMember {
	//parser := cocago.NewCocagoParser()
	//return 	parser.ProcessImports(code, fileName)
	return nil
}

func (g *GoIdentApp) SetExtensions(extension interface{})  {
	g.Extensions = extension
}
