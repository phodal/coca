package app_concept

import "github.com/phodal/coca/pkg/domain/core_domain"

type AbstractAnalysisApp interface {
	Analysis(code string, path string) core_domain.CodeContainer
	IdentAnalysis(s string, file string) []core_domain.CodeMember
	SetExtensions(extension interface{})
	AnalysisPackageManager(path string) core_domain.CodePackageInfo
}
