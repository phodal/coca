package evaluate

type Evaluator interface {
	IsMatch() bool
	Evaluate()
}

type Evaluation struct {
	Evaluator Evaluator
}

func (o *Evaluation) Evaluate() {
	 if o.Evaluator.IsMatch() {
		 o.Evaluator.Evaluate()
	 }
}