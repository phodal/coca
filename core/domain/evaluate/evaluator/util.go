package evaluator

import (
	"fmt"
	"github.com/phodal/coca/core/models"
)

type Util struct {

}

func (Util) Evaluate(models.JClassNode) {
	fmt.Println("util")
}
