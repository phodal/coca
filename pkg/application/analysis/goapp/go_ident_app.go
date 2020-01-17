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
	var imports []core_domain.CodeImport
	if g.Extensions != nil {
		imports = g.Extensions.([]core_domain.CodeImport)
	}
	return 	*parser.ProcessString(code, fileName, imports)
}

func (g *GoIdentApp) AnalysisImport(code string, fileName string) []core_domain.CodeImport {
	//parser := cocago.NewCocagoParser()
	//return 	parser.ProcessImports(code, fileName)
	return nil
}

func (g *GoIdentApp) SetExtensions(extension interface{})  {
	g.Extensions = extension
}
