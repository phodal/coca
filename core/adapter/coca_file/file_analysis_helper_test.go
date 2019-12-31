package coca_file

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/infrastructure/ast/identifier"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)
	parser := ProcessString("package com.phodal.coca;")

	context := parser.CompilationUnit()
	listener := identifier.NewJavaIdentifierListener()

	antlr.NewParseTreeWalker().Walk(listener, context)

	identifiers := listener.GetNodes()
	g.Expect(len(identifiers)).To(Equal(0))
}
