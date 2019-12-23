package evaluator

import (
	"fmt"
	"github.com/phodal/coca/core/models"
)

type Service struct {
}

func (Service) Evaluate(models.JClassNode) {
	fmt.Println("Service")
}
