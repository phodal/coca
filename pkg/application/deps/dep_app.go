package deps

import (
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"path/filepath"
	"strings"
)

type DepAnalysisApp struct {
}

func NewDepApp() *DepAnalysisApp {
	return &DepAnalysisApp{}
}

func (d *DepAnalysisApp) BuildImportMap(deps []core_domain.CodeDataStruct) map[string]core_domain.CodeImport {
	var impMap = make(map[string]core_domain.CodeImport)
	for _, clz := range deps {
		for _, imp := range clz.Imports {
			impMap[imp.Source] = imp
		}
	}

	return impMap
}

func (d *DepAnalysisApp) AnalysisPath(path string, nodes []core_domain.CodeDataStruct) []core_domain.CodeDependency {
	path, _ = filepath.Abs(path)
	pomXmls := cocafile.GetFilesWithFilter(path, cocafile.PomXmlFilter)
	gradleFiles := cocafile.GetFilesWithFilter(path, cocafile.BuildGradleFilter)

	var mavenDeps []core_domain.CodeDependency = nil
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

	var results []core_domain.CodeDependency = nil
	for index, dep := range mavenDeps {
		if _, ok := needRemoveMap[index]; !ok {
			results = append(results, dep)
		}
	}

	return results
}

var DepApp DepAnalysisApp // export for Plugins