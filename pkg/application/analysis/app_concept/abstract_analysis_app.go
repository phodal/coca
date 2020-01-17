package app_concept

import "github.com/phodal/coca/pkg/domain/core_domain"

type AbstractAnalysisApp interface {
	Analysis(code string, path string) core_domain.CodeFile
}
