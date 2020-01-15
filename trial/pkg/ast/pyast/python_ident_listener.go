package pyast

import (
	"fmt"
	parser "github.com/phodal/coca/languages/python"
	"github.com/phodal/coca/pkg/domain/trial"
	"reflect"
)

type PythonIdentListener struct {
	parser.BasePythonParserListener
}

var currentCodeFile *trial.CodeFile

func NewPythonIdentListener(fileName string) *PythonIdentListener {
	currentCodeFile = &trial.CodeFile{}
	currentCodeFile.FullName = fileName
	return &PythonIdentListener{}
}

func (s *PythonIdentListener) EnterClassdef(ctx *parser.ClassdefContext) {
	dataStruct := trial.CodeDataStruct{
		Name:       ctx.Name().GetText(),
		ID:         "",
		MemberIds:  nil,
		Properties: nil,
	}

	switch x := ctx.GetParent().GetChild(0).(type) {
	case *parser.DecoratorContext:

	default:
		fmt.Println(reflect.TypeOf(x))
	}

	currentCodeFile.DataStructures = append(currentCodeFile.DataStructures, dataStruct)
}

func (s *PythonIdentListener) GetCodeFileInfo() *trial.CodeFile {
	return currentCodeFile
}
