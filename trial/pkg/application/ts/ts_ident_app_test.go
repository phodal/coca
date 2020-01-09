package js_ident

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_TypeScriptConsoleLog(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	results := app.Analysis("console.log('hello, world')");

	g.Expect(len(results.MethodCalls)).To(Equal(1))
	g.Expect(results.MethodCalls[0].Class).To(Equal("console"))
}