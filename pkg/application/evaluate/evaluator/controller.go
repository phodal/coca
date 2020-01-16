package evaluator

import (
	"fmt"
	"github.com/phodal/coca/pkg/domain/jdomain"
)

type Controller struct {

}

func (Controller) Evaluate(node jdomain.JClassNode) {
	fmt.Println("controller")
}
