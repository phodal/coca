package evaluator

type Nullable struct {
	Items []string
}

type ServiceSummary struct {
	LifecycleMap  map[string][]string
	ReturnTypeMap map[string][]string
	RelatedMethod []string
}

type NormalIssues struct {
}

type UtilsSummary struct {
}

type Summary struct {
	UtilsCount               int
	ClassCount               int
	MethodCount              int
	NormalMethodCount        int
	TotalMethodLength        int
	StaticMethodCount        int
	MethodLengthStdDeviation float64
	MethodNumStdDeviation    float64
}

type EvaluateModel struct {
	Nullable       Nullable
	ServiceSummary ServiceSummary
	UtilsSummary   UtilsSummary
	Summary        Summary
}

func NewEvaluateModel() EvaluateModel {
	return EvaluateModel{Nullable: Nullable{Items: nil}}
}
