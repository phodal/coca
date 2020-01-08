package deps

import (
	"github.com/phodal/coca/core/domain"
)

type DepApp struct {
}

func NewDepApp() DepApp {
	return *&DepApp{}
}

func (d *DepApp) BuildImportMap(deps []domain.JClassNode) map[string]domain.JImport {
	var impMap = make(map[string]domain.JImport)
	for _, clz := range deps {
		for _, imp := range clz.Imports {
			impMap[imp.Source] = imp
		}
	}

	return impMap
}
