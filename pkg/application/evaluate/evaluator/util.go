package evaluator

import (
	"github.com/phodal/coca/pkg/domain/jdomain"
)

type Util struct {
}

func (Util) Evaluate(result *EvaluateModel, node jdomain.JClassNode) {

}

func (s Util) EvaluateList(evaluateModel *EvaluateModel, nodes []jdomain.JClassNode, nodeMap map[string]jdomain.JClassNode, identifiers []jdomain.JIdentifier) {

}
