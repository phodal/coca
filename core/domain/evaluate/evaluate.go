package evaluate

type Evaluator interface {
	IsMatch() bool
	Evaluate()
}

type Evaluation struct {
	Evaluator Evaluator
}

func (o *Evaluation) IsMatch() bool {
	return o.Evaluator.IsMatch()
}

func (o *Evaluation) Evaluate() {
	o.Evaluator.Evaluate()
}
