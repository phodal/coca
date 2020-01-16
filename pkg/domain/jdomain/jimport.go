package jdomain

import "github.com/phodal/coca/pkg/domain/core_domain"

func NewJImport(str string) core_domain.CodeImport {
	return core_domain.CodeImport{
		Source: str,
	}
}
