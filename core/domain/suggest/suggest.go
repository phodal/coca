package suggest

import "github.com/phodal/coca/core/models"

type SuggestApp struct {

}

func NewSuggestApp() SuggestApp {
	return *&SuggestApp{}
}

func (a SuggestApp) AnalysisPath(deps []models.JClassNode) {

}