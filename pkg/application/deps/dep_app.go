package deps

import (
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain"
	"path/filepath"
	"strings"
)

type DepAnalysisApp struct {
}

func NewDepApp() *DepAnalysisApp {
	return &DepAnalysisApp{}
}

func (d *DepAnalysisApp) BuildImportMap(deps []domain.JClassNode) map[string]domain.JImport {
	var impMap = make(map[string]domain.JImport)
	for _, clz := range deps {
		for _, imp := range clz.Imports {
			impMap[imp.Source] = imp
		}
	}

	return impMap
}

func (d *DepAnalysisApp) AnalysisPath(path string, nodes []domain.JClassNode) []domain.JDependency {
	path, _ = filepath.Abs(path)
	pomXmls := cocafile.GetFilesWithFilter(path, cocafile.PomXmlFilter)
	gradleFiles := cocafile.GetFilesWithFilter(path, cocafile.BuildGradleFilter)

	var mavenDeps []domain.JDependency = nil
	for _, pomFile := range pomXmls {
		currentMavenDeps := AnalysisMaven(pomFile)
		mavenDeps = append(mavenDeps, currentMavenDeps...)
	}
	for _, gradleFile := range gradleFiles {
		dependencies := AnalysisGradleFile(gradleFile)
		mavenDeps = append(mavenDeps, dependencies...)
	}

	importMap := d.BuildImportMap(nodes)

	var needRemoveMap = make(map[int]int)
	for depIndex, dep := range mavenDeps {
		for key := range importMap {
			if strings.Contains(key, dep.GroupId) {
				needRemoveMap[depIndex] = depIndex
				continue
			}
		}
	}

	var results []domain.JDependency = nil
	for index, dep := range mavenDeps {
		if _, ok := needRemoveMap[index]; !ok {
			results = append(results, dep)
		}
	}

	return results
}

var DepApp DepAnalysisApp // export for Plugins