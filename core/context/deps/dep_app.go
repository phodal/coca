package deps

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure/ast/groovy"
	"github.com/phodal/coca/core/infrastructure/xmlparse"
	. "github.com/phodal/coca/languages/groovy"
	"os"
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

func AnalysisMaven(xmlPath string) []domain.JDependency {
	xmlFile, _ := os.Open(xmlPath)
	parseXml := xmlparse.ParseXml(xmlFile)
	for _, element := range parseXml.Elements {
		val := element.Val.(xmlparse.XmlNode)
		if val.Name == "dependencies" {
			return BuildDeps(val)
		}
	}
	return nil
}

func BuildDeps(val xmlparse.XmlNode) []domain.JDependency {
	var deps []domain.JDependency = nil
	for _, depElement := range val.Elements {
		depNode := depElement.Val.(xmlparse.XmlNode)
		dependency := domain.NewJDependency("", "")

		for _, depValue := range depNode.Elements {
			node := depValue.Val.(xmlparse.XmlNode)
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

func AnalysisGradle(str string) {
	parser := ProcessGroovyString(str)
	context := parser.CompilationUnit()

	listener := groovy.NewGroovyIdentListener()

	antlr.NewParseTreeWalker().Walk(listener, context)
}

func ProcessGroovyString(code string) *GroovyParser {
	is := antlr.NewInputStream(code)
	lexer := NewGroovyLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewGroovyParser(stream)
	return parser
}
