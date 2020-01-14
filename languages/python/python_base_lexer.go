package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/infrastructure/container"
)

var TabSize = 8
var indents *container.Stack

func init() {
	indents = container.NewStack()
}

type PythonBaseLexer struct {
	*antlr.BaseLexer

	scopeStrictModes []bool
	stackLength      int
	stackIx          int

	lastToken        antlr.Token
	useStrictDefault bool
	useStrictCurrent bool

	_opened int
	buffer  antlr.Token
}

func (l *PythonBaseLexer) Emit() antlr.Token {
	emit := l.BaseLexer.Emit()
	return emit
}

func (l *PythonBaseLexer) IncIndentLevel() {
	l._opened++
}

func (l *PythonBaseLexer) DecIndentLevel() {
	if l._opened > 0 {
		l._opened--
	}
}

func (l *PythonBaseLexer) HandleNewLine() {

}

func (l *PythonBaseLexer) NextToken() antlr.Token {
	if l.GetInputStream().LA(1) == antlr.TokenEOF && indents.Len() > 0 {

	}

	next := l.BaseLexer.NextToken() // Get next token
	if next.GetChannel() == antlr.TokenDefaultChannel {
		// Keep track of the last token on default channel
		l.lastToken = next
	}
	return next
}

func (l *PythonBaseLexer) HandleSpaces() {

}
