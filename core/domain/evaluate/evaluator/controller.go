package evaluator

import (
	"fmt"
	"github.com/phodal/coca/core/models"
)

type Controller struct {

}

func (Controller) Evaluate(node models.JClassNode) {
	fmt.Println("controller")
}
