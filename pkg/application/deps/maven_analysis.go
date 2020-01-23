package deps

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/xmlparse"
	"os"
)

func AnalysisMaven(xmlPath string) []core_domain.CodeDependency {
	xmlFile, _ := os.Open(xmlPath)
	parseXml := xmlparse.ParseXML(xmlFile)
	for _, element := range parseXml.Elements {
		val := element.Val.(xmlparse.XMLNode)
		if val.Name == "dependencies" {
			return BuildDeps(val)
		}
	}
	return nil
}

func BuildDeps(val xmlparse.XMLNode) []core_domain.CodeDependency {
	var deps []core_domain.CodeDependency = nil
	for _, depElement := range val.Elements {
		depNode := depElement.Val.(xmlparse.XMLNode)
		dependency := core_domain.NewCodeDependency("", "")

		for _, depValue := range depNode.Elements {
			node := depValue.Val.(xmlparse.XMLNode)
			if node.Name == "groupId" {
				for _, textNode := range node.Elements {
					dependency.GroupId = textNode.Val.(string)
				}
			}

			if node.Name == "artifactId" {
				for _, textNode := range node.Elements {
					dependency.ArtifactId = textNode.Val.(string)
				}
			}

			if node.Name == "scope" {
				for _, textNode := range node.Elements {
					dependency.Scope = textNode.Val.(string)
				}
			}
		}

		deps = append(deps, *dependency)
	}

	return deps
}

