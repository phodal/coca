package domain

import (
	"coca/src/adapter/models"
	"coca/src/domain/support"
	"fmt"
)

type ConceptAnalyser struct {

}

func NewConceptAnalyser() ConceptAnalyser {
	return *&ConceptAnalyser{}
}

func (c ConceptAnalyser) run()  {

}

func (c ConceptAnalyser) Analysis(clzs *[]models.JClassNode) {
	buildMethodsFromDeps(*clzs)
}

func buildMethodsFromDeps(clzs []models.JClassNode) {
	var methodsName []string
	var methodStr string
	for _, clz := range clzs {
		for _, method := range clz.Methods {
			methodName := method.Name
			methodsName = append(methodsName, methodName)
			methodStr = methodStr + " " + methodName
		}
	}

	camelcase := support.SegmentConceptCamelcase(methodsName)
	fmt.Println(camelcase)
}
