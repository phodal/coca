package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type PythonBaseLexer struct {
	*antlr.BaseLexer

	scopeStrictModes []bool
	stackLength      int
	stackIx          int

	lastToken        antlr.Token
	useStrictDefault bool
	useStrictCurrent bool
}

func (l *PythonBaseLexer) IncIndentLevel() {

}

func (l *PythonBaseLexer) DecIndentLevel() {

}

func (l *PythonBaseLexer) HandleNewLine() {

}

func (l *PythonBaseLexer) HandleSpaces() {

}
