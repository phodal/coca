package tsapp

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"testing"
)

func Test_Regression(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/regressions/import_comma_issue.ts")

	results := app.Analysis(string(code), "")

	g.Expect(len(results.Imports)).To(Equal(3))
}

// todo: ignore this test for fast CI
func Test_ProcessErrorGrammar(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)
	code, _ := ioutil.ReadFile("../../../../../_fixtures/ts/regressions/callback_hell.ts")

	results := app.Analysis(string(code), "")

	g.Expect(len(results.Members)).To(Equal(0))
}
