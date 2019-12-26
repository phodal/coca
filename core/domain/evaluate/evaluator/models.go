package evaluator

type Nullable struct {
	Items []string
}

type ServiceIssues struct {
	LifecycleMap  map[string][]string
	ReturnTypeMap map[string][]string
	RelatedMethod []string
}

type NormalIssues struct {
	StaticMethodCount int
}

type Summary struct {
	ClassCount  int
	MethodCount int
}

type EvaluateModel struct {
	Nullable      Nullable
	ServiceIssues ServiceIssues
	NormalIssues  NormalIssues
	Summary       Summary
}

func NewEvaluateModel() EvaluateModel {
	return *&EvaluateModel{Nullable: Nullable{Items: nil}}
}
