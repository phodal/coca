package suggest

import (
	"github.com/phodal/coca/core/models"
)

type SuggestApp struct {
}

func NewSuggestApp() SuggestApp {
	return *&SuggestApp{}
}

func (a SuggestApp) AnalysisPath(deps []models.JClassNode) []Suggest {
	var suggests []Suggest
	// TODO: DSL => class constructor.len > 3
	for _, clz := range deps {
		var constructorCount = 0
		for _, method := range clz.Methods {
			if method.IsConstructor {
				constructorCount++
			}
		}

		if constructorCount >= 3 {
			suggests = append(suggests, NewSuggest(clz, "factory", "to many constructor"))
		}
	}

	return suggests
}
