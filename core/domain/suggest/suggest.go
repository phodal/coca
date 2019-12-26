package suggest

import (
	"github.com/phodal/coca/core/models"
)

type SuggestApp struct {
}

func NewSuggestApp() SuggestApp {
	return *&SuggestApp{}
}

func (a SuggestApp) AnalysisPath(deps []models.JClassNode) []Suggest {
	var suggests []Suggest
	for _, clz := range deps {
		if clz.Type == "Class" {
			// TODO: DSL => class constructor.len > 3
			if len(clz.Methods) > 0 {
				suggests = factorySuggest(clz, suggests)
			}
		}
		// TODO: long parameters in constructor builder
	}

	return suggests
}

func factorySuggest(clz models.JClassNode, suggests []Suggest) []Suggest {
	var constructorCount = 0
	var longestParaConstructorMethod = clz.Methods[0]
	for _, method := range clz.Methods {
		if method.IsConstructor {
			constructorCount++

			if len(method.Parameters) >= len(longestParaConstructorMethod.Parameters) {
				longestParaConstructorMethod = method
			}

			declLineNum := method.StopLine - method.StartLine
			// 参数过多，且在构造函数里调用过多
			if declLineNum > len(method.Parameters)-3 && len(method.MethodCalls) > 3 {
				suggest := NewSuggest(clz, "factory", "complex constructor")
				suggest.Line = method.StartLine
				suggests = append(suggests, suggest)
			}
		}
	}

	// TODO 合并 suggest
	if constructorCount >= 3 {
		suggest := NewSuggest(clz, "factory", "too many constructor")
		suggest.Size = constructorCount
		suggests = append(suggests, suggest)
	}

	if len(longestParaConstructorMethod.Parameters) >= 5 {
		suggest := NewSuggest(clz, "builder", "too many parameters")
		suggest.Size = len(longestParaConstructorMethod.Parameters)
		suggests = append(suggests, suggest)
	}

	return suggests
}
