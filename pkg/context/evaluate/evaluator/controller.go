package evaluator

import (
	"fmt"
	"github.com/phodal/coca/pkg/domain"
)

type Controller struct {

}

func (Controller) Evaluate(node domain.JClassNode) {
	fmt.Println("controller")
}
