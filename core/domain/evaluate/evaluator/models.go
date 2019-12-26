package evaluator

type Nullable struct {
	Items []string
}

type ServiceIssues struct {
	LifecycleMap  map[string][]string
	ReturnTypeMap map[string][]string
	RelatedMethod []string
}

type EvaluateModel struct {
	Nullable      Nullable
	ServiceIssues ServiceIssues
}

func NewEvaluateModel() EvaluateModel {
	return *&EvaluateModel{Nullable: Nullable{Items: nil}}
}
