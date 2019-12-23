package evaluator

import (
	"fmt"
	"github.com/phodal/coca/core/models"
)

type Model struct {

}

func (Model) Evaluate(models.JClassNode) {
	fmt.Println("model")
}
