package evaluate

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cocatest/testhelper"
	"github.com/phodal/coca/pkg/application/evaluate/evaluator"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"path/filepath"
	"testing"
)

func TestAnalyser_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	analyser := NewEvaluateAnalyser()
	codePath := "../../../_fixtures/evaluate/service.json"
	codePath = filepath.FromSlash(codePath)
	file := cmd_util.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	analyser.Analysis(parsedDeps, nil)

	g.Expect(true).To(Equal(true))
}

func Test_Service_LifeCycle(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	analyser := NewEvaluateAnalyser()
	codePath := "../../../_fixtures/evaluate/service_lifecycle.json"
	codePath = filepath.FromSlash(codePath)
	file := cmd_util.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	result := analyser.Analysis(parsedDeps, nil)

	g.Expect(len(result.ServiceSummary.LifecycleMap["do"])).To(Equal(2))
	g.Expect(result.ServiceSummary.LifecycleMap["do"][0]).To(Equal("doSave"))
	g.Expect(result.ServiceSummary.LifecycleMap["do"][1]).To(Equal("doUpdate"))
}

func Test_Service_Same_Return_Type(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	analyser := NewEvaluateAnalyser()
	codePath := "../../../_fixtures/evaluate/service_same_return_type.json"
	codePath = filepath.FromSlash(codePath)
	file := cmd_util.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	results := analyser.Analysis(parsedDeps, nil)

	g.Expect(len(results.ServiceSummary.ReturnTypeMap)).To(Equal(1))
}

func Test_Long_Parameters(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	analyser := NewEvaluateAnalyser()
	codePath := "../../../_fixtures/evaluate/service_long_parameters.json"
	codePath = filepath.FromSlash(codePath)
	file := cmd_util.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	result := analyser.Analysis(parsedDeps, nil)

	g.Expect(result.ServiceSummary.RelatedMethod[0]).To(Equal("address"))
	g.Expect(result.ServiceSummary.RelatedMethod[1]).To(Equal("age"))
	g.Expect(result.ServiceSummary.RelatedMethod[2]).To(Equal("firstname"))
	g.Expect(result.ServiceSummary.RelatedMethod[3]).To(Equal("lastname"))
}

func TestNullPointException(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/evaluate/null"
	result := buildEvaluateResult(codePath)

	g.Expect(len(result.Nullable.Items)).To(Equal(2))
}

func TestStaticUtils(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/evaluate/utils"
	result := buildEvaluateResult(codePath)

	g.Expect(result.Summary.UtilsCount).To(Equal(1))
}

func Test_CheckFornull(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/evaluate/checkfornull"
	result := buildEvaluateResult(codePath)

	g.Expect(len(result.Nullable.Items)).To(Equal(1))
}

func buildEvaluateResult(codePath string) evaluator.EvaluateModel {
	callNodes, _, identifiers := testhelper.BuildAnalysisDeps(codePath)

	analyser := NewEvaluateAnalyser()
	result := analyser.Analysis(callNodes, identifiers)
	return result
}
