package evaluator

type Nullable struct {
	Items []string
}

type ServiceIssues struct {
	LifecycleMap map[string]string
}

type EvaluateModel struct {
	Nullable      Nullable
	ServiceIssues ServiceIssues
}

func NewEvaluateModel() EvaluateModel {
	return *&EvaluateModel{Nullable: Nullable{Items: nil}}
}
