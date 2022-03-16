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
	g.Expect(len(results.DataStructures)).To(Equal(1))
}

func Test_IdentDSMember(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	code, _ := ioutil.ReadFile("../../../../pkg/domain/core_domain/code_data_struct.go")
	app := &GoIdentApp{}
	results := app.IdentAnalysis(string(code), "file")
	g.Expect(results[0].ID).To(Equal("core_domain::CodeDataStruct"))
}

func Test_ShouldGetModuleNameFromModFile(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := &GoIdentApp{}
	packageManager := app.AnalysisPackageManager("../../../..")

	g.Expect(packageManager.ProjectName).To(Equal("github.com/modernizing/coca"))
}
