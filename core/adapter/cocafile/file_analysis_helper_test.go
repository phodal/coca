package cocafile

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/infrastructure/ast/identifier"
	"testing"
)

func TestJavaCallApp_ProcessStringWorks(t *testing.T) {
	g := NewGomegaWithT(t)
	parser := ProcessString(`
package com.phodal.coca.analysis.identifier.model;

public class DataClass {
    private String date;

    public String getDate() {
        return date;
    }
}

`)

	context := parser.CompilationUnit()
	listener := identifier.NewJavaIdentifierListener()

	antlr.NewParseTreeWalker().Walk(listener, context)

	identifiers := listener.GetNodes()
	g.Expect(identifiers[0].ClassName).To(Equal("DataClass"))
}
