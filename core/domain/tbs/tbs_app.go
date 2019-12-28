package tbs

import (
	"github.com/phodal/coca/core/models"
)

type TbsApp struct {
}

func NewTbsApp() *TbsApp {
	return &TbsApp{}
}

type TestBadSmell struct {
	FileName    string
	Type        string
	Description string
	Line        int
}

func (a TbsApp) AnalysisPath(deps []models.JClassNode, identifiersMap map[string]models.JIdentifier) []TestBadSmell {
	var results []TestBadSmell = nil
	var identMethodMap = make(map[string]models.JMethod)
	for key, clzMap := range identifiersMap {
		for _, method := range clzMap.Methods {
			identMethodMap[key + "." + method.Name] = method
		}
	}

	for _, clz := range deps {
		// TODO refactoring identify & annotation
		for _, method := range clz.Methods {
			fullName := clz.Package + "." + clz.Class + "." + method.Name
			checkIgnoreTest(clz.Path, identMethodMap[fullName], &results)
			checkEmptyTest(clz.Path, identMethodMap[fullName], method, &results)
		}
	}

	return results
}

func checkEmptyTest(path string, iMethod models.JMethod, cMethod models.JMethod, results *[]TestBadSmell) {
	for _, annotation := range iMethod.Annotations {
		if annotation.QualifiedName == "Test" {
			if len(cMethod.MethodCalls) <= 1 {
				tbs := *&TestBadSmell{
					FileName:    path,
					Type:        "EmptyTest",
					Description: "",
					Line:        0,
				}

				*results = append(*results, tbs)
			}
		}
	}

}

func checkIgnoreTest(clzPath string, method models.JMethod, results *[]TestBadSmell) {
	for _, annotation := range method.Annotations {
		if annotation.QualifiedName == "Ignore" {
			tbs := *&TestBadSmell{
				FileName:    clzPath,
				Type:        "IgnoreTest",
				Description: "",
				Line:        0,
			}

			*results = append(*results, tbs)
		}
	}
}
