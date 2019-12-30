package evaluator

import (
	"fmt"
	"github.com/phodal/coca/core/domain"
)

type Controller struct {

}

func (Controller) Evaluate(node domain.JClassNode) {
	fmt.Println("controller")
}
