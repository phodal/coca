package application

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(JavaScriptApiApp)
	results := app.Analysis("console.log('hello, world')");

	g.Expect(len(results.MethodCalls)).To(Equal(1))
	g.Expect(results.MethodCalls[0].Class).To(Equal("console"))
}