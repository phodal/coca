package bs

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestBadSmellApp_ComplexCondition(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := NewBadSmellApp()
	codePath := "../../../_fixtures/bs/ComplexIf.java"
	codePath = filepath.FromSlash(codePath)

	bs := bsApp.AnalysisPath(codePath)
	bsList := bsApp.IdentifyBadSmell(bs, nil)

	g.Expect(len(bsList)).To(Equal(1))
	g.Expect(bsList[0].Bs).To(Equal("complexCondition"))
}
func TestBadSmellApp_DataClass(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := NewBadSmellApp()
	codePath := "../../../_fixtures/bs/DataClass.java"
	codePath = filepath.FromSlash(codePath)

	bs := bsApp.AnalysisPath(codePath)
	bsList := bsApp.IdentifyBadSmell(bs, nil)

	g.Expect(len(bsList)).To(Equal(1))
	g.Expect(bsList[0].Bs).To(Equal("dataClass"))
}

func TestBadSmellApp_LongMethod(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := NewBadSmellApp()
	codePath := "../../../_fixtures/bs/LongMethod.java"
	codePath = filepath.FromSlash(codePath)

	bs := bsApp.AnalysisPath(codePath)
	bsList := bsApp.IdentifyBadSmell(bs, nil)

	g.Expect(len(bsList)).To(Equal(2))
	g.Expect(bsList[0].Bs).To(Equal("longMethod"))
	g.Expect(bsList[1].Bs).To(Equal("refusedBequest"))
}

func TestBadSmellApp_LazyElement(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := NewBadSmellApp()
	codePath := "../../../_fixtures/bs/LazyClass.java"
	codePath = filepath.FromSlash(codePath)

	bs := bsApp.AnalysisPath(codePath)
	bsList := bsApp.IdentifyBadSmell(bs, nil)

	g.Expect(len(bsList)).To(Equal(1))
	g.Expect(bsList[0].Bs).To(Equal("lazyElement"))
}

func TestBadSmellApp_LongParameters(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := NewBadSmellApp()
	codePath := "../../../_fixtures/bs/LongParameter.java"
	codePath = filepath.FromSlash(codePath)

	bs := bsApp.AnalysisPath(codePath)
	bsList := bsApp.IdentifyBadSmell(bs, nil)

	g.Expect(len(bsList)).To(Equal(1))
	g.Expect(bsList[0].Bs).To(Equal("longParameterList"))
}

func TestBadSmellApp_MultipleIf(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := NewBadSmellApp()
	codePath := "../../../_fixtures/bs/MultipleIf.java"
	codePath = filepath.FromSlash(codePath)

	bs := bsApp.AnalysisPath(codePath)
	bsList := bsApp.IdentifyBadSmell(bs, nil)

	g.Expect(len(bsList)).To(Equal(1))
	g.Expect(bsList[0].Bs).To(Equal("repeatedSwitches"))
}

func TestBadSmellApp_LargeClass(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := NewBadSmellApp()
	codePath := "../../../_fixtures/bs/LargeClass.java"
	codePath = filepath.FromSlash(codePath)

	bs := bsApp.AnalysisPath(codePath)
	bsList := bsApp.IdentifyBadSmell(bs, nil)

	g.Expect(len(bsList)).To(Equal(1))
	g.Expect(bsList[0].Bs).To(Equal("largeClass"))
}

func Test_ShouldJobConcurrencyManager(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := NewBadSmellApp()
	codePath := "../../../_fixtures/bs/JobConcurrencyManager.java"
	codePath = filepath.FromSlash(codePath)

	bs := bsApp.AnalysisPath(codePath)
	bsList := bsApp.IdentifyBadSmell(bs, nil)

	g.Expect(len(bsList)).To(Equal(0))
}

func TestBadSmellApp_GraphCall(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := NewBadSmellApp()
	codePath := "../../../_fixtures/bs/graphcall"
	codePath = filepath.FromSlash(codePath)

	bs := bsApp.AnalysisPath(codePath)
	bsList := bsApp.IdentifyBadSmell(bs, nil)

	g.Expect(len(bsList)).To(Equal(1))
	g.Expect(bsList[0].Bs).To(Equal("graphConnectedCall"))
	g.Expect(bsList[0].Description).To(Equal("graphcall.GraphCallA->graphcall.GraphCallB->graphcall.GraphCallC;graphcall.GraphCallA->graphcall.GraphCallC"))
}
