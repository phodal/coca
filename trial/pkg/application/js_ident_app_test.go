package application

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(JavaScriptApiApp)
	app.Analysis("console.log('hello, world')");

	g.Expect(true).To(Equal(true))
}