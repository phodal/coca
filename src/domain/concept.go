package domain

import (
	"coca/src/adapter/models"
	"fmt"
)

type ConceptAnalyser struct {

}

func NewConceptAnalyser() ConceptAnalyser {
	return *&ConceptAnalyser{}
}

func (c ConceptAnalyser) run()  {

}

func (c ConceptAnalyser) Analysis(path string, clzs *[]models.JClassNode) {
	fmt.Println(clzs)
}
