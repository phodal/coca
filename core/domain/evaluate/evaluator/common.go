package evaluator

import (
	"fmt"
	"github.com/phodal/coca/core/models"
)

type Common struct {

}

func (Common) Evaluate(models.JClassNode) {
	fmt.Println("common")
}

