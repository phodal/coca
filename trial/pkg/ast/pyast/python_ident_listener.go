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
	//fmt.Println(reflect.TypeOf(ctx.GetChild(0)))
}


func (s *PythonIdentListener) EnterClassdef(ctx *parser.ClassdefContext) {
	fmt.Println(ctx.Name().GetText())
}

