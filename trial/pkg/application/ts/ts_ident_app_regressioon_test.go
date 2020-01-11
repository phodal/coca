package ts

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"testing"
)

func Test_Regression(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/regressions/catch_comment_issue.ts")

	results := app.Analysis(string(code), "")

	g.Expect(len(results.Imports)).To(Equal(3))
}
//
//func Test_ProcessErrorGrammar(t *testing.T) {
//	g := NewGomegaWithT(t)
//
//	app := new(TypeScriptApiApp)
//	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/ts-node-starter/src/config/passport.ts")
//
//	results := app.Analysis(string(code), "")
//
//	g.Expect(len(results.Imports)).To(Equal(1))
//}
