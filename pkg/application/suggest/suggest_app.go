package suggest

import (
	"github.com/phodal/coca/pkg/domain/jdomain"
	"github.com/phodal/coca/pkg/domain/support_domain"
)

type SuggestApp struct {
}

func NewSuggestApp() SuggestApp {
	return SuggestApp{}
}

func (a SuggestApp) AnalysisPath(deps []jdomain.JClassNode) []api_domain.Suggest {
	var suggests []api_domain.Suggest
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

func factorySuggest(clz jdomain.JClassNode, suggests []api_domain.Suggest) []api_domain.Suggest {
	var constructorCount = 0
	var longestParaConstructorMethod = clz.Methods[0]

	var currentSuggestList []api_domain.Suggest = nil
	for _, method := range clz.Methods {
		if method.IsConstructor {
			constructorCount++

			if len(method.Parameters) >= len(longestParaConstructorMethod.Parameters) {
				longestParaConstructorMethod = method
			}

			declLineNum := method.StopLine - method.StartLine
			// 参数过多，且在构造函数里调用过多
			PARAMETER_LINE_OFFSET := 3
			PARAMETER_METHOD_CALL_OFFSET := 3
			if declLineNum > len(method.Parameters)-PARAMETER_LINE_OFFSET && (len(method.MethodCalls) > len(method.Parameters)+PARAMETER_METHOD_CALL_OFFSET) {
				suggest := api_domain.NewSuggest(clz, "factory", "complex constructor")
				suggest.Line = method.StartLine
				currentSuggestList = append(currentSuggestList, suggest)
			}
		}
	}

	// TODO 合并 suggest
	if constructorCount >= 3 {
		suggest := api_domain.NewSuggest(clz, "factory", "too many constructor")
		suggest.Size = constructorCount
		currentSuggestList = append(currentSuggestList, suggest)
	}

	if len(longestParaConstructorMethod.Parameters) >= 5 {
		suggest := api_domain.NewSuggest(clz, "builder", "too many parameters")
		suggest.Size = len(longestParaConstructorMethod.Parameters)
		currentSuggestList = append(currentSuggestList, suggest)
	}

	suggest := api_domain.MergeSuggest(clz, currentSuggestList)

	if suggest.Pattern != "" {
		suggests = append(suggests, suggest)
	}

	return suggests
}
