package goapp

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"testing"
)

func Test_ProcessPackage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	code, _ := ioutil.ReadFile("../../../../pkg/domain/core_domain/code_data_struct.go")
	app := &GoIdentApp{}
	results := app.Analysis(string(code), "domain")
	g.Expect(len(results.DataStructures)).To(Equal(2))
}