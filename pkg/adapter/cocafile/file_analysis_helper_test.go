package cocafile

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java/java_identify"
	"testing"
)

func TestJavaCallApp_ProcessStringWorks(t *testing.T) {
	g := NewGomegaWithT(t)
	parser := ast_java.ProcessJavaString(`
package com.phodal.coca.analysis.identifier.model;

public class DataClass {
    private String date;

    public String getDate() {
        return date;
    }
}

`)

	context := parser.CompilationUnit()
	listener := java_identify.NewJavaIdentifierListener()

	antlr.NewParseTreeWalker().Walk(listener, context)

	identifiers := listener.GetNodes()
	g.Expect(identifiers[0].NodeName).To(Equal("DataClass"))
}
