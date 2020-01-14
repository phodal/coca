package pyast

import (
	"fmt"
	parser "github.com/phodal/coca/languages/python"
)

type PythonIdentListener struct {
	parser.BasePythonParserListener
}

func NewPythonIdentListener(fileName string) *PythonIdentListener {
	return &PythonIdentListener{}
}

func (s *PythonIdentListener) EnterRoot(ctx *parser.RootContext) {
	fmt.Println(ctx)
}

func (s *PythonIdentListener) EnterSingle_input(ctx *parser.Single_inputContext) {

}


